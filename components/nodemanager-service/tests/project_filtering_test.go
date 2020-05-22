package manager

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chef/automate/api/interservice/nodemanager/manager"
	"github.com/chef/automate/api/interservice/nodemanager/nodes"
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	authzConstants "github.com/chef/automate/components/authz-service/constants"
	nodesserver "github.com/chef/automate/components/nodemanager-service/server/nodes"
	"github.com/chef/automate/lib/grpc/auth_context"
)

func TestReadProjectFilteringNodes(t *testing.T) {
	timestamp, err := ptypes.TimestampProto(time.Now())
	require.NoError(t, err)

	db, err := createPGDB()
	require.NoError(t, err)

	nodeManager := nodesserver.New(db, nil, "")

	// Adding a manual node
	mgr1 := manager.NodeManager{Name: "mgr1", Type: "aws-ec2"}
	mgrID1, err := db.AddNodeManager(&mgr1, "11111111")
	require.NoError(t, err)
	defer db.DeleteNodeManager(mgrID1)

	node1 := manager.ManagerNode{Id: "i-1111111", Region: "us-west-2", Host: "Node1"}

	instances := []*manager.ManagerNode{&node1}
	manualNodeIds := db.AddManagerNodesToDB(instances, mgrID1, "242403433", []*manager.CredentialsByTags{}, "aws-ec2")
	require.NoError(t, err)
	defer func() {
		for _, node := range manualNodeIds {
			db.DeleteNode(node)
		}
	}()

	assert.Equal(t, 1, len(manualNodeIds))
	manualNodeID := manualNodeIds[0]

	cases := []struct {
		description     string
		requestProjects []string
		nodeProjects    []string
		isError         bool
	}{
		{
			description:     "Node project matching request's projects",
			requestProjects: []string{"target_project"},
			nodeProjects:    []string{"target_project"},
			isError:         false,
		},
		{
			description:     "Node project not matching request's projects",
			requestProjects: []string{"missed_target_project"},
			nodeProjects:    []string{"target_project"},
			isError:         true,
		},
		{
			description:     "Node has no projects; request's project has unassigned project",
			requestProjects: []string{authzConstants.UnassignedProjectID},
			nodeProjects:    []string{},
			isError:         false,
		},
		{
			description:     "Node has a project assigned; request's project has unassigned and the matching project",
			requestProjects: []string{authzConstants.UnassignedProjectID, "target_project"},
			nodeProjects:    []string{"target_project"},
			isError:         false,
		},
		{
			description:     "Node has a project assigned; request's projects has only the unassigned project",
			requestProjects: []string{authzConstants.UnassignedProjectID},
			nodeProjects:    []string{"target_project"},
			isError:         true,
		},
		{
			description:     "Node is assigned a project; request for all projects",
			requestProjects: []string{authzConstants.AllProjectsExternalID},
			nodeProjects:    []string{"target_project"},
			isError:         false,
		},
		{
			description:     "Node has no projects; requested for all projects",
			requestProjects: []string{authzConstants.AllProjectsExternalID},
			nodeProjects:    []string{},
			isError:         false,
		},
		{
			description:     "Node has no projects; requested projects is empty",
			requestProjects: []string{},
			nodeProjects:    []string{},
			isError:         false,
		},
		{
			description:     "Node has one project not matching any of several requested projects",
			requestProjects: []string{"project3", "project4", "project7", "project6"},
			nodeProjects:    []string{"project9"},
			isError:         true,
		},
		{
			description:     "Node has one project matching one of several requested projects",
			requestProjects: []string{"project3", "project4", "project7", "project6"},
			nodeProjects:    []string{"project7"},
			isError:         false,
		},
		{
			description:     "Node with several projects where one matches a single requested project",
			requestProjects: []string{"project3"},
			nodeProjects:    []string{"project3", "project4", "project7", "project6"},
			isError:         false,
		},
		{
			description:     "Node with several projects where one matches one of several requested projects",
			requestProjects: []string{"project3", "project10", "project12", "project13"},
			nodeProjects:    []string{"project3", "project4", "project7", "project6"},
			isError:         false,
		},
		{
			description:     "Node with several projects that do not matche several requested projects",
			requestProjects: []string{"project14", "project10", "project12", "project13"},
			nodeProjects:    []string{"project3", "project4", "project7", "project6"},
			isError:         true,
		},
		{
			description:     "Node with several projects where two match two of several requested projects",
			requestProjects: []string{"project3", "project10", "project12", "project13"},
			nodeProjects:    []string{"project3", "project10", "project7", "project6"},
			isError:         false,
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Project filter: %s", test.description),
			func(t *testing.T) {

				// Create the context with projects added
				ctx := contextWithProjects(test.requestProjects)

				// Create Ingest node
				node := &manager.NodeMetadata{
					Uuid:        "node1",
					Projects:    test.nodeProjects,
					LastContact: timestamp,
					RunData: &nodes.LastContactData{
						Id:      createUUID(),
						EndTime: timestamp,
						Status:  nodes.LastContactData_PASSED,
					},
				}

				// ingest node
				err = db.ProcessIncomingNode(node)
				require.NoError(t, err)

				// Delete created node after the test is complete
				defer db.DeleteNode(node.Uuid)

				// Call Read to get the ingested node with project filtering context.
				nodeResponse, err := nodeManager.Read(ctx, &nodes.Id{
					Id: node.Uuid,
				})

				if test.isError {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, node.Uuid, nodeResponse.Id)
				}

				// Test that we can read the manually added node for all cases.
				manualNodeResponse, err := nodeManager.Read(ctx, &nodes.Id{Id: manualNodeID})
				require.NoError(t, err)
				assert.Equal(t, manualNodeID, manualNodeResponse.Id)
			})
	}
}

func TestListProjectFilteringAllNodes(t *testing.T) {
	db, err := createPGDB()
	require.NoError(t, err)

	nodeManager := nodesserver.New(db, nil, "")

	timestamp, err := ptypes.TimestampProto(time.Now())

	// Adding two manual nodes
	node1 := nodes.Node{Name: "test-manual-node-1"}
	manualNodeID1, err := db.AddNode(&node1)
	require.NoError(t, err)

	node2 := nodes.Node{Name: "test-manual-node-2"}
	manualNodeID2, err := db.AddNode(&node2)
	require.NoError(t, err)

	manualNodeIds := []string{manualNodeID1, manualNodeID2}
	logrus.Infof("manual node ids %s %s", manualNodeID1, manualNodeID2)
	defer func() {
		for _, node := range manualNodeIds {
			db.DeleteNode(node)
		}
	}()

	assert.Equal(t, 2, len(manualNodeIds))

	cases := []struct {
		description     string
		ctx             context.Context
		nodes           []*manager.NodeMetadata
		expectedNodeIDs []string
	}{
		{
			description: "Three nodes matching on the same project tag",
			ctx:         contextWithProjects([]string{"target_project"}),
			nodes: []*manager.NodeMetadata{
				{
					Uuid:     "node1",
					Projects: []string{"target_project"},
				},
				{
					Uuid:     "node2",
					Projects: []string{"project8", "target_project"},
				},
				{
					Uuid:     manualNodeID1,
					Projects: []string{"project8", "target_project"},
				},
			},
			// bc of the project filter we only expect the three above to match
			expectedNodeIDs: []string{"node1", "node2", manualNodeID1},
		},
		{
			description: "Three nodes matching with two project tags",
			ctx:         contextWithProjects([]string{"target_project_1", "target_project_2"}),
			nodes: []*manager.NodeMetadata{
				{
					Uuid:     "node1",
					Projects: []string{"target_project_1"},
				},
				{
					Uuid:     "node2",
					Projects: []string{"target_project_2"},
				},
				{
					Uuid:     manualNodeID1,
					Projects: []string{"target_project_2"},
				},
			},
			// bc of the project filter we only expect the three above to match
			expectedNodeIDs: []string{"node1", "node2", manualNodeID1},
		},
		{
			description: "Three nodes with only one with a matching project",
			ctx:         contextWithProjects([]string{"target_project"}),
			nodes: []*manager.NodeMetadata{
				{
					Uuid:     "node1",
					Projects: []string{"target_project"},
				},
				{
					Uuid:     "node2",
					Projects: []string{"project8"},
				},
				{
					Uuid:     manualNodeID1,
					Projects: []string{"project8"},
				},
			},
			// bc of the project filter we only expect the one above to match
			expectedNodeIDs: []string{"node1"},
		},
		{
			description: "Four nodes with different projects and one missing a project where all match " +
				"because the AllProjectsID is requested",
			ctx: contextWithProjects([]string{authzConstants.AllProjectsExternalID}),
			nodes: []*manager.NodeMetadata{
				{
					Uuid:     "node1",
					Projects: []string{"project3"},
				},
				{
					Uuid:     "node2",
					Projects: []string{"project8", "project7"},
				},
				{
					Uuid:     "node3",
					Projects: []string{},
				},
				{
					Uuid:     manualNodeID1,
					Projects: []string{},
				},
			},
			// bc we requested all projects we expect the four nodes above + the manualNodeID2
			expectedNodeIDs: []string{"node1", "node2", "node3", manualNodeID2, manualNodeID1},
		},
		{
			description: "Three nodes one with a project tag and one with none. Matching two unassigned",
			ctx:         contextWithProjects([]string{authzConstants.UnassignedProjectID}),
			nodes: []*manager.NodeMetadata{
				{
					Uuid:     "node1",
					Projects: []string{"project9"},
				},
				{
					Uuid:     "node2",
					Projects: []string{},
				},
				{
					Uuid:     manualNodeID1,
					Projects: []string{},
				},
			},
			// bc of the request for unassigned, we expect above three nodes + manualNodeID2
			expectedNodeIDs: []string{"node2", manualNodeID1, manualNodeID2},
		},
		{
			description: "Two nodes with projects assigned, with unassigned request no matches",
			ctx:         contextWithProjects([]string{authzConstants.UnassignedProjectID}),
			nodes: []*manager.NodeMetadata{
				{
					Uuid:     "node1",
					Projects: []string{"project9"},
				},
				{
					Uuid:     "node2",
					Projects: []string{"project7"},
				},
			},
			//bc of the request for unassigned, we expect only manualNodeID2
			expectedNodeIDs: []string{manualNodeID2},
		},
		{
			description: "Two nodes one unassigned and one with a node, with unassigned and macthing " +
				"project request matching both",
			ctx: contextWithProjects([]string{authzConstants.UnassignedProjectID, "target_project"}),
			nodes: []*manager.NodeMetadata{
				{
					Uuid:     "node1",
					Projects: []string{"target_project"},
				},
				{
					Uuid:     "node2",
					Projects: []string{},
				},
			},
			//bc of the request for unassigned, we expect above two nodes + manualNodeID2
			expectedNodeIDs: []string{"node1", "node2", manualNodeID2},
		},
		{
			description: "Three nodes two unassigned and one with a node, with unassigned and macthing " +
				"project request matching both",
			ctx: contextWithProjects([]string{authzConstants.UnassignedProjectID, "target_project"}),
			nodes: []*manager.NodeMetadata{
				{
					Uuid:     "node1",
					Projects: []string{"target_project"},
				},
				{
					Uuid:     "node2",
					Projects: []string{},
				},
				{
					Uuid:     manualNodeID1,
					Projects: []string{},
				},
			},
			//bc of the request for unassigned, we expect above three nodes + manualNodeID2
			expectedNodeIDs: []string{"node1", "node2", manualNodeID1, manualNodeID2},
		},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Project filter: %s", test.description),
			func(t *testing.T) {
				// Ingest nodes
				for _, node := range test.nodes {
					node.LastContact = timestamp
					node.RunData = &nodes.LastContactData{
						Id:      createUUID(),
						EndTime: timestamp,
						Status:  nodes.LastContactData_PASSED,
					}

					err = db.ProcessIncomingNode(node)
					require.NoError(t, err)
				}
				// Delete nodes after the test is complete
				defer func() {
					for _, node := range test.nodes {
						db.DeleteNode(node.Uuid)
					}
				}()

				// Call List to get all ingested nodes with project filtering context.
				nodesResponse, err := nodeManager.List(test.ctx, &nodes.Query{})
				require.NoError(t, err)

				// Get all the node IDs returned.
				actualNodeIDs := []string{}
				for _, node := range nodesResponse.Nodes {
					actualNodeIDs = append(actualNodeIDs, node.Id)
				}
				logrus.Infof("expected: %v", test.expectedNodeIDs)
				logrus.Infof("actual: %v", actualNodeIDs)

				assert.ElementsMatch(t, actualNodeIDs, test.expectedNodeIDs)
			})
	}
}

func contextWithProjects(projects []string) context.Context {
	ctx := context.Background()
	return auth_context.NewContext(ctx, []string{}, projects, "", "")
}
