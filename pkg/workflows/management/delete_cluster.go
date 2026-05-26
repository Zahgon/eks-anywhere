package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type deleteManagementCluster struct{}

func (s *deleteManagementCluster) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *deleteManagementCluster) Name() string { _ = "STUB: not implemented"; return "" }

func (s *deleteManagementCluster) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *deleteManagementCluster) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

type cleanupGitRepo struct{}

func (s *cleanupGitRepo) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *cleanupGitRepo) Name() string { _ = "STUB: not implemented"; return "" }

func (s *cleanupGitRepo) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *cleanupGitRepo) Checkpoint() *task.CompletedTask { _ = "STUB: not implemented"; return nil }
