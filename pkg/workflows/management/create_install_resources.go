package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type installProviderSpecificResources struct{}

func (s *installProviderSpecificResources) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installProviderSpecificResources) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installProviderSpecificResources) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *installProviderSpecificResources) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
