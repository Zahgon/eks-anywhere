package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

// createWorkloadClusterTask implementation.
type createWorkloadClusterTask struct{}

func (s *createWorkloadClusterTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *createWorkloadClusterTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *createWorkloadClusterTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *createWorkloadClusterTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
