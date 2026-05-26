package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type createCluster struct{}

// Run createCluster performs actions needed to create the management cluster.
func (c *createCluster) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (c *createCluster) Name() string { _ = "STUB: not implemented"; return "" }

func (c *createCluster) Checkpoint() *task.CompletedTask { _ = "STUB: not implemented"; return nil }

func (c *createCluster) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
