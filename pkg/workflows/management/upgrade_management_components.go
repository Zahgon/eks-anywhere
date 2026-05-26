package management

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/task"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/validations"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces"
	v1releasealpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

// UpgradeManagementComponentsWorkflow is a schema for upgrade management components.
type UpgradeManagementComponentsWorkflow struct {
	clientFactory  interfaces.ClientFactory
	provider       providers.Provider
	clusterManager interfaces.ClusterManager
	gitOpsManager  interfaces.GitOpsManager
	writer         filewriter.FileWriter
	capiManager    interfaces.CAPIManager
	eksdInstaller  interfaces.EksdInstaller
	eksdUpgrader   interfaces.EksdUpgrader
}

// NewUpgradeManagementComponentsRunner builds a new UpgradeManagementCommponents construct.
func NewUpgradeManagementComponentsRunner(
	clientFactory interfaces.ClientFactory,
	provider providers.Provider,
	capiManager interfaces.CAPIManager,
	clusterManager interfaces.ClusterManager,
	gitOpsManager interfaces.GitOpsManager,
	writer filewriter.FileWriter,
	eksdUpgrader interfaces.EksdUpgrader,
	eksdInstaller interfaces.EksdInstaller,
) *UpgradeManagementComponentsWorkflow {
	_ = "STUB: not implemented"
	return nil
}

// UMCValidator is a struct that holds a cluster and a kubectl executable.
// It is used to perform preflight validations on the cluster.
type UMCValidator struct {
	cluster         *types.Cluster
	eksaRelease     *v1releasealpha1.EKSARelease
	kubectl         validations.KubectlClient
	skipValidations []string
}

// NewUMCValidator is a constructor function that creates a new instance of UMCValidator.
func NewUMCValidator(cluster *types.Cluster, eksaRelease *v1releasealpha1.EKSARelease, kubectl validations.KubectlClient, skipValidations []string) *UMCValidator {
	_ = "STUB: not implemented"
	return nil
}

// PreflightValidations is a method of the UMCValidator struct.
// It performs preflight validations on the cluster and returns a slice of Validation objects.
func (u *UMCValidator) PreflightValidations(ctx context.Context) []validations.Validation {
	_ = "STUB: not implemented"
	return nil
}

// When upgrading management components, the only preflight check that can be skipped
// is the EKS-A version skew validation. While the skipValidations flag accepts multiple
// values for consistency with the cluster upgrade command, only the EKS-A version skew
// check (validations.EksaVersionSkew) will be honored - all other skip requests will
// be ignored.

// Run Upgrade implements upgrade functionality for management cluster's upgrade operation.
func (umc *UpgradeManagementComponentsWorkflow) Run(ctx context.Context, clusterSpec *cluster.Spec, managementCluster *types.Cluster, validator interfaces.Validator) error {
	_ = "STUB: not implemented"
	return nil
}

type setupAndValidateMC struct{}

// Run setupAndValidate validates management cluster before upgrade process starts.
func (s *setupAndValidateMC) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *setupAndValidateMC) Name() string { _ = "STUB: not implemented"; return "" }

func (s *setupAndValidateMC) Restore(_ context.Context, _ *task.CommandContext, _ *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *setupAndValidateMC) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

// This struct is similar to upgradeCoreComponents, but its returned value is different in Run() function.
type upgradeCoreComponentsMC struct {
	UpgradeChangeDiff *types.ChangeDiff
}

func (s *upgradeCoreComponentsMC) Name() string { _ = "STUB: not implemented"; return "" }

func (s *upgradeCoreComponentsMC) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *upgradeCoreComponentsMC) Restore(_ context.Context, commandContext *task.CommandContext, completedTask *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}

func (s *upgradeCoreComponentsMC) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

// This struct is similar to installNewComponents, but its returned value is different in Run() function.
type installNewComponentsMC struct{}

func (s *installNewComponentsMC) Run(ctx context.Context, commandContext *task.CommandContext) task.Task {
	_ = "STUB: not implemented"
	return *new(task.Task)
}

func (s *installNewComponentsMC) Name() string { _ = "STUB: not implemented"; return "" }

func (s *installNewComponentsMC) Checkpoint() *task.CompletedTask {
	_ = "STUB: not implemented"
	return nil
}

func (s *installNewComponentsMC) Restore(_ context.Context, _ *task.CommandContext, _ *task.CompletedTask) (task.Task, error) {
	_ = "STUB: not implemented"
	return *new(task.Task), nil
}
