package task

import (
	"context"
	"time"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces"
)

// Task is a logical unit of work - meant to be implemented by each Task.
type Task interface {
	Run(ctx context.Context, commandContext *CommandContext) Task
	Name() string
	Checkpoint() *CompletedTask
	Restore(ctx context.Context, commandContext *CommandContext, completedTask *CompletedTask) (Task, error)
}

// Command context maintains the mutable and shared entities.
type CommandContext struct {
	ClientFactory         interfaces.ClientFactory
	Bootstrapper          interfaces.Bootstrapper
	Provider              providers.Provider
	ClusterManager        interfaces.ClusterManager
	GitOpsManager         interfaces.GitOpsManager
	Validations           interfaces.Validator
	Writer                filewriter.FileWriter
	EksdInstaller         interfaces.EksdInstaller
	EksaInstaller         interfaces.EksaInstaller
	PackageManager        interfaces.PackageManager
	EksdUpgrader          interfaces.EksdUpgrader
	ClusterUpgrader       interfaces.ClusterUpgrader
	ClusterCreator        interfaces.ClusterCreator
	ClusterDeleter        interfaces.ClusterDeleter
	CAPIManager           interfaces.CAPIManager
	ClusterSpec           *cluster.Spec
	CurrentClusterSpec    *cluster.Spec
	UpgradeChangeDiff     *types.ChangeDiff
	BootstrapCluster      *types.Cluster
	ManagementCluster     *types.Cluster
	WorkloadCluster       *types.Cluster
	Profiler              *Profiler
	OriginalError         error
	BackupClusterStateDir string
	ForceCleanup          bool
	ClusterMover          interfaces.ClusterMover
	IamAuth               interfaces.AwsIamAuth
}

func (c *CommandContext) SetError(err error) { _ = "STUB: not implemented"; return }

type Profiler struct {
	metrics map[string]map[string]time.Duration
	starts  map[string]map[string]time.Time
}

// profiler for a Task.
func (pp *Profiler) SetStartTask(taskName string) { _ = "STUB: not implemented"; return }

// this can be used to profile sub tasks.
func (pp *Profiler) SetStart(taskName string, msg string) { _ = "STUB: not implemented"; return }

// needs to be called after setStart.
func (pp *Profiler) MarkDoneTask(taskName string) { _ = "STUB: not implemented"; return }

// this can be used to profile sub tasks.
func (pp *Profiler) MarkDone(taskName string, msg string) { _ = "STUB: not implemented"; return }

// get Metrics.
func (pp *Profiler) Metrics() map[string]map[string]time.Duration {
	_ = "STUB: not implemented"

	// debug logs for task metric.
	return nil
}

func (pp *Profiler) logProfileSummary(taskName string) { _ = "STUB: not implemented"; return }

// Manages Task execution.
type taskRunner struct {
	task           Task
	writer         filewriter.FileWriter
	withCheckpoint bool
}

type TaskRunnerOpt func(*taskRunner)

func WithCheckpointFile() TaskRunnerOpt { _ = "STUB: not implemented"; return *new(TaskRunnerOpt) }

func (tr *taskRunner) RunTask(ctx context.Context, commandContext *CommandContext) error {
	_ = "STUB: not implemented"
	return nil
}

func taskRunnerFinalBlock(startTime time.Time) { _ = "STUB: not implemented"; return }

func NewTaskRunner(task Task, writer filewriter.FileWriter, opts ...TaskRunnerOpt) *taskRunner {
	_ = "STUB: not implemented"
	return nil
}

func (tr *taskRunner) saveCheckpoint(checkpointInfo CheckpointInfo, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (tr *taskRunner) setupCheckpointInfo(commandContext *CommandContext, checkpointFileName string) (CheckpointInfo, error) {
	_ = "STUB: not implemented"
	return *new(CheckpointInfo), nil
}

type TaskCheckpoint interface{}

type CheckpointInfo struct {
	CompletedTasks map[string]*CompletedTask `json:"completedTasks"`
}

type CompletedTask struct {
	Checkpoint TaskCheckpoint `json:"checkpoint"`
}

func newCheckpointInfo() CheckpointInfo { _ = "STUB: not implemented"; return *new(CheckpointInfo) }

func (c CheckpointInfo) taskCompleted(name string, completedTask *CompletedTask) {
	_ = "STUB: not implemented"
	return
}

func readCheckpointFile(file string) (*CheckpointInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
	UnmarshalTaskCheckpoint marshals the received task checkpoint (type interface{}) then unmarshalls it into the desired type

specified in the Restore() method.
When reading from a yaml file, there isn't a direct way in Go to do a type conversion from interface{} to the desired type.
We use interface{} because the TaskCheckpoint type will vary depending on what's needed for a specific task. The known workaround
for this is to marshal & unmarshal it into the checkpoint type.
*/
func UnmarshalTaskCheckpoint(taskCheckpoint TaskCheckpoint, config TaskCheckpoint) error {
	_ = "STUB: not implemented"
	return nil
}
