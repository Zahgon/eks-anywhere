package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type deleteWorkloadCluster struct{}

func (s *deleteWorkloadCluster) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *deleteWorkloadCluster) Name() string { _ = "STUB: not implemented"; return "" }

func (s *deleteWorkloadCluster) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *deleteWorkloadCluster) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
