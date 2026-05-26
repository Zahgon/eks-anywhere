package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type installCAPIComponentsTask struct{}

func (s *installCAPIComponentsTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installCAPIComponentsTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installCAPIComponentsTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *installCAPIComponentsTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
