package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type pauseGitOpsReconcile struct{}

// Run pauseGitOpsReconcile pause GitOps reconciler before management cluster upgrade.
func (s *pauseGitOpsReconcile) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *pauseGitOpsReconcile) Name() string { _ = "STUB: not implemented"; return "" }

func (s *pauseGitOpsReconcile) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *pauseGitOpsReconcile) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

// reconcileGitOps updates all the places that have a cluster definition to follow the cluster config provided to this workflow:
// the cluster config in the git repo if GitOps is enabled. It also resumes the GitOps reconciliations.
type reconcileGitOps struct{}

// Run reconcileGitOps resumes GitOps reconciler and performs other GitOps related tasks after management cluster upgrade.
func (s *reconcileGitOps) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *reconcileGitOps) Name() string { _ = "STUB: not implemented"; return "" }

func (s *reconcileGitOps) Checkpoint() *task.CompletedTask { _ = "STUB: not implemented"; return nil }

func (s *reconcileGitOps) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
