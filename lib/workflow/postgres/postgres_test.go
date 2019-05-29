// +build integration

package postgres

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/chef/automate/lib/workflow"

	"github.com/stretchr/testify/assert"

	"github.com/chef/automate/lib/platform/pg"
	"github.com/chef/automate/lib/workflow/backend"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

const defaultDatabaseName = "workflow_test"

func defaultConnURIForDatabase(dbname string) string {
	if os.Getenv("PG_URI") != "" {
		return os.Getenv("PG_URI")
	}
	connInfo := pg.A2ConnInfo{
		Host:  "localhost",
		Port:  5432,
		User:  "automate",
		Certs: pg.A2SuperuserCerts,
	}
	return connInfo.ConnURI(dbname)
}

func runResetDB() error {
	dbName := defaultDatabaseName

	db, err := sql.Open("postgres", defaultConnURIForDatabase("template1"))
	if err != nil {
		return errors.Wrap(err, "could not initialize db connection")
	}
	defer db.Close()
	_, err = db.Exec(pg.DropDatabaseQuery(dbName))
	if err != nil {
		return errors.Wrap(err, "could not drop database")
	}
	_, err = db.Exec(pg.CreateDatabaseQuery(dbName))
	if err != nil {
		return errors.Wrap(err, "could not create database")
	}
	return nil
}

func TestNoAvailableTasks(t *testing.T) {
	taskName := "task_name"
	err := runResetDB()
	require.NoError(t, err)
	b1 := NewPostgresBackend(defaultConnURIForDatabase(defaultDatabaseName))
	err = b1.Init()
	require.NoError(t, err)
	defer b1.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, _, err = b1.DequeueTask(ctx, taskName)
	require.Equal(t, workflow.ErrNoTasks, err)
}

func TestMultipleTasksCanDequeueConcurrently(t *testing.T) {
	taskName := "task_name"
	workflowName := "workflow_name"
	err := runResetDB()
	require.NoError(t, err)
	b1 := NewPostgresBackend(defaultConnURIForDatabase(defaultDatabaseName))
	err = b1.Init()
	require.NoError(t, err)
	defer b1.Close()

	b2 := NewPostgresBackend(defaultConnURIForDatabase(defaultDatabaseName))
	err = b2.Init()
	require.NoError(t, err)
	defer b2.Close()

	b3 := NewPostgresBackend(defaultConnURIForDatabase(defaultDatabaseName))
	err = b3.Init()
	require.NoError(t, err)
	defer b3.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = b1.EnqueueWorkflow(ctx, &backend.WorkflowInstance{
		InstanceName: "workflow-instance",
		WorkflowName: workflowName,
		Status:       "running",
	})

	require.NoError(t, err, "failed to enqueue workflow")

	_, completer, err := b1.DequeueWorkflow(ctx, []string{workflowName})
	require.NoError(t, err, "failed to dequeue workflow")

	completer.EnqueueTask(&backend.Task{
		Name: taskName,
	}, backend.TaskEnqueueOpts{TryRemaining: 1})
	completer.EnqueueTask(&backend.Task{
		Name: taskName,
	}, backend.TaskEnqueueOpts{TryRemaining: 1})
	completer.Continue(nil)

	tx1, err := b1.db.BeginTx(ctx, nil)
	require.NoError(t, err)

	tx2, err := b2.db.BeginTx(ctx, nil)
	require.NoError(t, err)

	tx3, err := b3.db.BeginTx(ctx, nil)
	require.NoError(t, err)

	b1TID, _, err := b1.dequeueTask(tx1, taskName)
	require.NoError(t, err)

	b2TID, _, err := b2.dequeueTask(tx2, taskName)
	require.NoError(t, err)

	_, _, err = b3.dequeueTask(tx3, taskName)
	require.Equal(t, workflow.ErrNoTasks, err)

	assert.NotZero(t, b1TID)
	assert.NotZero(t, b2TID)
	assert.NotEqual(t, b1TID, b2TID)

	err = tx1.Commit()
	require.NoError(t, err)

	err = tx2.Commit()
	require.NoError(t, err)
}

func TestTaskCompleteWhileWorkflowIsRunning(t *testing.T) {
	taskName := "task_name"
	workflowName := "workflow_name"
	err := runResetDB()
	require.NoError(t, err)
	b1 := NewPostgresBackend(defaultConnURIForDatabase(defaultDatabaseName))
	err = b1.Init()
	require.NoError(t, err)
	defer b1.Close()

	b2 := NewPostgresBackend(defaultConnURIForDatabase(defaultDatabaseName))
	err = b2.Init()
	require.NoError(t, err)
	defer b2.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = b1.EnqueueWorkflow(ctx, &backend.WorkflowInstance{
		InstanceName: "workflow-instance",
		WorkflowName: workflowName,
		Status:       "running",
	})

	require.NoError(t, err, "failed to enqueue workflow")

	_, completer, err := b1.DequeueWorkflow(ctx, []string{workflowName})
	require.NoError(t, err, "failed to dequeue workflow")

	completer.EnqueueTask(&backend.Task{
		Name: taskName,
	}, backend.TaskEnqueueOpts{TryRemaining: 1})
	completer.Continue(nil)

	_, taskCompleter, err := b1.DequeueTask(ctx, taskName)
	require.NoError(t, err)
	taskCompleter.Succeed(nil)
}
