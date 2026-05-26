package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type deleteBootstrapClusterTask struct{}

func (s *deleteBootstrapClusterTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *deleteBootstrapClusterTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *deleteBootstrapClusterTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *deleteBootstrapClusterTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
