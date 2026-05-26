package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type preClusterUpgrade struct{}

// Run preClusterUpgrade implements steps to be performed before workload cluster's upgrade.
func (s *preClusterUpgrade) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	// Take CAPI backup filtering on only current workload cluster.
	return *new(task.Task)
}

func (s *preClusterUpgrade) Name() string { _ = "STUB: not implemented"; return "" }

func (s *preClusterUpgrade) Checkpoint() *task.CompletedTask { _ = "STUB: not implemented"; return nil }

func (s *preClusterUpgrade) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
