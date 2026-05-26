package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
	"github.com/aws/eks-anywhere/pkg/types"
)

type ensureEtcdCAPIComponentsExist struct{}

// Run ensureEtcdCAPIComponentsExist ensures ETCD CAPI providers on the management cluster.
func (s *ensureEtcdCAPIComponentsExist) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *ensureEtcdCAPIComponentsExist) Name() string { _ = "STUB: not implemented"; return "" }

func (s *ensureEtcdCAPIComponentsExist) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *ensureEtcdCAPIComponentsExist) Restore(_ context.Context, _ *task.CommandContext, _ *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

type upgradeCoreComponents struct {
	UpgradeChangeDiff *types.ChangeDiff
}

func runUpgradeCoreComponents(ctx context.Context, commandContext *task.CommandContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Run upgradeCoreComponents upgrades pre cluster upgrade components.
func (s *upgradeCoreComponents) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *upgradeCoreComponents) Name() string { _ = "STUB: not implemented"; return "" }

func (s *upgradeCoreComponents) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *upgradeCoreComponents) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
