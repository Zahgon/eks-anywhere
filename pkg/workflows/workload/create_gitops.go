package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type installGitOpsManagerTask struct{}

func (s *installGitOpsManagerTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installGitOpsManagerTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installGitOpsManagerTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *installGitOpsManagerTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
