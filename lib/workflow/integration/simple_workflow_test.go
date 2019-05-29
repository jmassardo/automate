// +build integration

package integration_test

import (
	"context"
	"sync"
	"time"

	"github.com/chef/automate/lib/workflow"
)

// TestCompleteSimpleWorkflow tests that a workflow the launches a
// single task completes
//
// Workflow:
// - OnStart -> Launch 'simple' task
// - OnTaskComplete -> Done
func (suite *WorkflowTestSuite) TestCompleteSimpleWorkflow() {
	taskName := randName("simple")
	workflowName := randName("simple")
	instanceName := randName("instance")

	// There will be once task that runs, along
	// with the TaskCompleted
	wgTask := sync.WaitGroup{}
	wgTask.Add(2)

	m := suite.newManager(
		WithTaskExecutorF(
			taskName,
			func(context.Context, workflow.Task) (interface{}, error) {
				wgTask.Done()
				return nil, nil
			}),
		WithWorkflowExecutor(
			workflowName,
			&workflowExecutorWrapper{
				onStart: func(w workflow.WorkflowInstance, ev workflow.StartEvent) workflow.Decision {
					err := w.EnqueueTask(taskName, nil)
					suite.Require().NoError(err, "failed to enqueue task")
					return w.Continue(nil)
				},
				onTaskComplete: func(w workflow.WorkflowInstance, ev workflow.TaskCompleteEvent) workflow.Decision {
					suite.Assert().Equal(1, w.TotalCompletedTasks())
					suite.Assert().Equal(1, w.TotalEnqueuedTasks())
					wgTask.Done()
					return w.Complete()
				},
			},
		),
	)
	defer m.Stop()
	err := m.EnqueueWorkflow(context.Background(), workflowName, instanceName, nil)
	suite.Require().NoError(err, "Failed to enqueue workflow")
	wgTask.Wait()
	time.Sleep(20 * time.Millisecond)
	err = m.Stop()
	suite.NoError(err)
}
