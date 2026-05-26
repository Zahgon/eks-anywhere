package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type installCuratedPackagesTask struct{}

func (s *installCuratedPackagesTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installCuratedPackagesTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installCuratedPackagesTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *installCuratedPackagesTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
