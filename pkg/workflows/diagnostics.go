package workflows

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/task"
)

type CollectDiagnosticsTask struct {
	*CollectWorkloadClusterDiagnosticsTask
	*CollectMgmtClusterDiagnosticsTask
}

type CollectWorkloadClusterDiagnosticsTask struct{}

type CollectMgmtClusterDiagnosticsTask struct{}

// CollectDiagnosticsTask implementation

func (s *CollectDiagnosticsTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *CollectDiagnosticsTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *CollectDiagnosticsTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *CollectDiagnosticsTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"

	// CollectWorkloadClusterDiagnosticsTask implementation
	return nil
}

// Run starts collecting the logs for workload cluster diagnostics.
func (s *CollectWorkloadClusterDiagnosticsTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

// Name returns the name of CollectWorkloadClusterDiagnosticsTask.
func (s *CollectWorkloadClusterDiagnosticsTask) Name() string { _ = "STUB: not implemented"; return "" }

// Restore restores from CollectWorkloadClusterDiagnosticsTask.
func (s *CollectWorkloadClusterDiagnosticsTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"

	// Checkpoint sets a checkpoint at CollectWorkloadClusterDiagnosticsTask.
	return *new(task.Task), nil
}

func (s *CollectWorkloadClusterDiagnosticsTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"

	// CollectMgmtClusterDiagnosticsTask implementation
	return nil
}

func (s *CollectMgmtClusterDiagnosticsTask) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *CollectMgmtClusterDiagnosticsTask) Name() string { _ = "STUB: not implemented"; return "" }

func (s *CollectMgmtClusterDiagnosticsTask) Restore(ctx context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *CollectMgmtClusterDiagnosticsTask) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}
