package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type upgradeCluster struct{}

// Run upgradeCluster performs actions needed to upgrade the management cluster.
func (s *upgradeCluster) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *upgradeCluster) Name() string { _ = "STUB: not implemented"; return "" }

func (s *upgradeCluster) Checkpoint() *task.CompletedTask { _ = "STUB: not implemented"; return nil }

func (s *upgradeCluster) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
