package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/task"
)

type writeUpgradeClusterConfig struct{}

// Run writeClusterConfig writes new management cluster's cluster config file to the destination after the upgrade process.
func (s *writeUpgradeClusterConfig) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

// Handle AWS IAM kubeconfig generation/cleanup during management cluster upgrade

// AWS IAM being added during upgrade

// AWS IAM being removed during upgrade - cleanup existing kubeconfig

func (s *writeUpgradeClusterConfig) Name() string { _ = "STUB: not implemented"; return "" }

func (s *writeUpgradeClusterConfig) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *writeUpgradeClusterConfig) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

type writeCreateClusterConfig struct{}

func (s *writeCreateClusterConfig) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *writeCreateClusterConfig) Name() string { _ = "STUB: not implemented"; return "" }

func (s *writeCreateClusterConfig) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *writeCreateClusterConfig) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func writeClusterConfigToDisk(clusterSpec *cluster.Spec, datacenterConfig providers.DatacenterConfig, machineConfigs []providers.MachineConfig, writer filewriter.FileWriter) error {
	_ = "STUB: not implemented"
	return nil
}
