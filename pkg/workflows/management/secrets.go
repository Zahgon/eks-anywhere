package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type (
	updateSecrets       struct{}
	updateSecretsCreate struct{}
)

// Run updateSecrets updates management cluster's secrets.
func (s *updateSecrets) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *updateSecrets) Name() string { _ = "STUB: not implemented"; return "" }

func (s *updateSecrets) Checkpoint() *task.CompletedTask { _ = "STUB: not implemented"; return nil }

func (s *updateSecrets) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

// Run updateSecrets updates management cluster's secrets.
func (s *updateSecretsCreate) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *updateSecretsCreate) Name() string { _ = "STUB: not implemented"; return "" }

func (s *updateSecretsCreate) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *updateSecretsCreate) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
