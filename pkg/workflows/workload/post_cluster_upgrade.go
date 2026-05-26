package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type postClusterUpgrade struct{}

// Run postClusterUpgrade implements steps to be performed after the upgrade process.
func (s *postClusterUpgrade) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *postClusterUpgrade) Name() string { _ = "STUB: not implemented"; return "" }

func (s *postClusterUpgrade) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *postClusterUpgrade) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
