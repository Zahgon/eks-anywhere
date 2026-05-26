package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type installCAPIComponentsForDeleteTask struct{}

func (s *installCAPIComponentsForDeleteTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installCAPIComponentsForDeleteTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installCAPIComponentsForDeleteTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *installCAPIComponentsForDeleteTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
