package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
	"github.com/aws/eks-anywhere/pkg/types"
)

type installEksaComponentsOnBootstrapTask struct{}

func (s *installEksaComponentsOnBootstrapTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installEksaComponentsOnBootstrapTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installEksaComponentsOnBootstrapTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *installEksaComponentsOnBootstrapTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

type installEksaComponentsOnWorkloadTask struct{}

func (s *installEksaComponentsOnWorkloadTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installEksaComponentsOnWorkloadTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installEksaComponentsOnWorkloadTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *installEksaComponentsOnWorkloadTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func installEKSAComponents(ctx context.Context, commandContext *task.CommandContext, targetCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
