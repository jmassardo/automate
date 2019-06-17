// +build integration

package integration_test

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/chef/automate/lib/workflow"
)

// TestWorkflowFail tests that a call to w.Fail ends the workflow
// with an error that can be recovered.
func (suite *WorkflowTestSuite) TestWorkflowFail() {
	workflowName := randName("failing")
	instanceName := randName("instance")

	doneChan := make(chan struct{})
	m := suite.newManager(
		WithWorkflowExecutor(
			workflowName,
			&workflowExecutorWrapper{
				onStart: func(w workflow.WorkflowInstance, ev workflow.StartEvent) workflow.Decision {
					close(doneChan)
					return w.Fail(errors.New("expected test error"))
				},
				onTaskComplete: func(w workflow.WorkflowInstance, ev workflow.TaskCompleteEvent) workflow.Decision {
					return w.Complete()
				},
			},
		),
	)
	defer m.Stop()
	err := m.EnqueueWorkflow(context.Background(), workflowName, instanceName, nil)
	suite.Require().NoError(err, "Failed to enqueue workflow")
	<-doneChan
	time.Sleep(20 * time.Millisecond)
	w, err := m.GetWorkflowInstanceByName(context.Background(), instanceName, workflowName)
	suite.NoError(err)
	suite.Error(w.Err())
	suite.False(w.IsRunning())
	suite.Equal(w.Err().Error(), "expected test error")
	err = m.Stop()
	suite.NoError(err)
}
