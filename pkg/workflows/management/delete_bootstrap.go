package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type deleteBootstrapClusterForDeleteTask struct{}

func (s *deleteBootstrapClusterForDeleteTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *deleteBootstrapClusterForDeleteTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *deleteBootstrapClusterForDeleteTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *deleteBootstrapClusterForDeleteTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
