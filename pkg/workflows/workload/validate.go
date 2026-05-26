package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
	"github.com/aws/eks-anywhere/pkg/validations"
)

type (
	setAndValidateUpgradeWorkloadTask struct{}
	setAndValidateCreateWorkloadTask  struct{}
)

// Run setAndValidateCreateWorkloadTask performs actions needed to validate creating the workload cluster.
func (s *setAndValidateCreateWorkloadTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *setAndValidateCreateWorkloadTask) providerValidation(ctx context.Context, commandContext *task.CommandContext) []validations.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (s *setAndValidateCreateWorkloadTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *setAndValidateCreateWorkloadTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *setAndValidateCreateWorkloadTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"

	// Run setAndValidateWorkloadTask performs actions needed to validate the workload cluster.
	return nil
}

func (s *setAndValidateUpgradeWorkloadTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *setAndValidateUpgradeWorkloadTask) providerValidation(ctx context.Context, commandContext *task.CommandContext) []validations.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (s *setAndValidateUpgradeWorkloadTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *setAndValidateUpgradeWorkloadTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *setAndValidateUpgradeWorkloadTask) Checkpoint() *task.CompletedTask {
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
