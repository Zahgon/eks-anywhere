package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type moveClusterManagementForDeleteTask struct{}

func (s *moveClusterManagementForDeleteTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *moveClusterManagementForDeleteTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *moveClusterManagementForDeleteTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *moveClusterManagementForDeleteTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
