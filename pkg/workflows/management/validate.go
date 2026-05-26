package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
	"github.com/aws/eks-anywhere/pkg/validations"
)

type setupAndValidateCreate struct{}

// setupAndValidateCreate implementation

func (s *setupAndValidateCreate) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *setupAndValidateCreate) providerValidation(ctx context.Context, commandContext *task.CommandContext) []validations.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (s *setupAndValidateCreate) Name() string { _ = "STUB: not implemented"; return "" }

func (s *setupAndValidateCreate) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *setupAndValidateCreate) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

type setupAndValidateUpgrade struct{}

// Run setupAndValidate validates management cluster before upgrade process starts.
func (s *setupAndValidateUpgrade) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *setupAndValidateUpgrade) providerValidation(ctx context.Context, commandContext *task.CommandContext) []validations.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (s *setupAndValidateUpgrade) Name() string { _ = "STUB: not implemented"; return "" }

func (s *setupAndValidateUpgrade) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *setupAndValidateUpgrade) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

type setupAndValidateDelete struct{}

func (s *setupAndValidateDelete) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *setupAndValidateDelete) Name() string { _ = "STUB: not implemented"; return "" }

func (s *setupAndValidateDelete) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *setupAndValidateDelete) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
