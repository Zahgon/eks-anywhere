package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type moveClusterManagementTask struct{}

func (s *moveClusterManagementTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *moveClusterManagementTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *moveClusterManagementTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *moveClusterManagementTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
