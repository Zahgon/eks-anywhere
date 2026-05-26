package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type installNewComponents struct{}

func runInstallNewComponents(ctx context.Context, commandContext *task.CommandContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Run installNewComponents performs actions needed to upgrade the management cluster.
func (s *installNewComponents) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installNewComponents) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installNewComponents) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *installNewComponents) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
