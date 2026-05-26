package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type createBootStrapClusterTask struct{}

func (s *createBootStrapClusterTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *createBootStrapClusterTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *createBootStrapClusterTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *createBootStrapClusterTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
