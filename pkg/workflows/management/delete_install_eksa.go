package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type installEksaComponentsOnBootstrapForDeleteTask struct{}

func (s *installEksaComponentsOnBootstrapForDeleteTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installEksaComponentsOnBootstrapForDeleteTask) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *installEksaComponentsOnBootstrapForDeleteTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *installEksaComponentsOnBootstrapForDeleteTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
