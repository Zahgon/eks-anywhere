package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type writeClusterConfig struct{}

// Run writeClusterConfig writes new management cluster's cluster config file to the destination after the create/upgrade process.
func (s *writeClusterConfig) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

// Handle AWS IAM kubeconfig generation/cleanup during cluster operations

// New cluster creation

// Workload cluster upgrade scenarios

// AWS IAM being added during upgrade

// AWS IAM being removed during upgrade - cleanup existing kubeconfig

func (s *writeClusterConfig) Name() string { _ = "STUB: not implemented"; return "" }

func (s *writeClusterConfig) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *writeClusterConfig) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
