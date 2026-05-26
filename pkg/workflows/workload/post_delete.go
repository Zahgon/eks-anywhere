package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type (
	cleanupGitRepo     struct{}
	postDeleteWorkload struct{}
)

func (s *postDeleteWorkload) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *postDeleteWorkload) Name() string { _ = "STUB: not implemented"; return "" }

func (s *postDeleteWorkload) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *postDeleteWorkload) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
