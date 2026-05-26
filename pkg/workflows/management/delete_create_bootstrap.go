package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type createBootStrapClusterForDeleteTask struct{}

func (s *createBootStrapClusterForDeleteTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *createBootStrapClusterForDeleteTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *createBootStrapClusterForDeleteTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *createBootStrapClusterForDeleteTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
