package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type upgradeCuratedPackagesTask struct{}

func (s *upgradeCuratedPackagesTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *upgradeCuratedPackagesTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *upgradeCuratedPackagesTask) Restore(_ context.Context, _ *task.CommandContext, _ *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *upgradeCuratedPackagesTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
