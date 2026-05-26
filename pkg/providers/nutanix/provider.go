package nutanix

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/crypto"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
)

//go:embed config/cp-template.yaml
var defaultCAPIConfigCP string

//go:embed config/md-template.yaml
var defaultClusterConfigMD string

//go:embed config/secret-template.yaml
var secretTemplate string

//go:embed config/machine-health-check-template.yaml
var mhcTemplate []byte

var (
	eksaNutanixDatacenterResourceType = fmt.Sprintf("nutanixdatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaNutanixMachineResourceType    = fmt.Sprintf("nutanixmachineconfigs.%s", v1alpha1.GroupVersion.Group)
	// list of env variables required by CAPX to be present and defined beforehand.
	requiredEnvs = []string{nutanixEndpointKey, constants.NutanixUsernameKey, constants.NutanixPasswordKey, expClusterResourceSetKey}
)

// Provider implements the Nutanix Provider.
type Provider struct {
	clusterConfig    *v1alpha1.Cluster
	datacenterConfig *v1alpha1.NutanixDatacenterConfig
	machineConfigs   map[string]*v1alpha1.NutanixMachineConfig
	templateBuilder  *TemplateBuilder
	kubectlClient    ProviderKubectlClient
	validator        *Validator
	writer           filewriter.FileWriter
	ipValidator      IPValidator
	skipIPCheck      bool
}

var _ providers.Provider = &Provider{}

// NewProvider returns a new nutanix provider.
func NewProvider(
	datacenterConfig *v1alpha1.NutanixDatacenterConfig,
	machineConfigs map[string]*v1alpha1.NutanixMachineConfig,
	clusterConfig *v1alpha1.Cluster,
	providerKubectlClient ProviderKubectlClient,
	writer filewriter.FileWriter,
	clientCache *ClientCache,
	ipValidator IPValidator,
	certValidator crypto.TlsValidator,
	httpClient *http.Client,
	now types.NowFunc,
	skipIPCheck bool,
) *Provider {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) BootstrapClusterOpts(_ *cluster.Spec) ([]bootstrapper.BootstrapClusterOption, error) {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil, nil
}

func (p *Provider) BootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

func (p *Provider) PostBootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

func (p *Provider) PostWorkloadInit(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

func (p *Provider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Provider) DatacenterResourceType() string { _ = "STUB: not implemented"; return "" }

func (p *Provider) MachineResourceType() string { _ = "STUB: not implemented"; return "" }

func (p *Provider) generateSSHKeysIfNotSet() error { _ = "STUB: not implemented"; return nil }

// use the same key

func (p *Provider) PostClusterDeleteValidate(ctx context.Context, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

func (p *Provider) SetupAndValidateCreateCluster(ctx context.Context, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) SetupAndValidateDeleteCluster(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupAndValidateUpgradeCluster - Performs necessary setup and validations for upgrade cluster operation.
func (p *Provider) SetupAndValidateUpgradeCluster(ctx context.Context, _ *types.Cluster, clusterSpec *cluster.Spec, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupAndValidateUpgradeManagementComponents performs necessary setup for upgrade management components operation.
func (p *Provider) SetupAndValidateUpgradeManagementComponents(_ context.Context, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): Add validations when this is supported
	return nil
}

func (p *Provider) UpdateSecrets(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	// check if CAPI Secret name and EKS-A Secret name are not the same
	// this is to ensure that the EKS-A Secret that is watched and CAPX Secret that is reconciled are not the same
	return nil
}

// updateEKSASecrets generates and applies the EKSA secret on the cluster.
func (p *Provider) updateEKSASecrets(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) GenerateStorageClass() []byte {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

// GenerateMHC returns MachineHealthCheck for the cluster in yaml format.
func (p *Provider) GenerateMHC(_ *cluster.Spec) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Provider) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

// Version returns the version of the provider.
func (p *Provider) Version(components *cluster.ManagementComponents) string {
	_ = "STUB: not implemented"
	return ""
}

// EnvMap returns the environment variables for the provider.
func (p *Provider) EnvMap(_ *cluster.ManagementComponents, _ *cluster.Spec) (map[string]string, error) {
	_ = "STUB: not implemented"
	// TODO(nutanix): determine if any env vars are needed and add them to requiredEnvs
	return nil, nil
}

func (p *Provider) GetDeployments() map[string][]string { _ = "STUB: not implemented"; return nil }

// GetInfrastructureBundle returns the infrastructure bundle for the provider.
func (p *Provider) GetInfrastructureBundle(components *cluster.ManagementComponents) *types.InfrastructureBundle {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) DatacenterConfig(_ *cluster.Spec) providers.DatacenterConfig {
	_ = "STUB: not implemented"
	return *new(providers.DatacenterConfig)
}

// MachineConfigs returns a MachineConfig slice.
func (p *Provider) MachineConfigs(_ *cluster.Spec) []providers.MachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func configsMapToSlice(c map[string]providers.MachineConfig) []providers.MachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) ValidateNewSpec(_ context.Context, _ *types.Cluster, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

// ChangeDiff returns the component change diff for the provider.
func (p *Provider) ChangeDiff(currentComponents, newComponents *cluster.ManagementComponents) *types.ComponentChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) RunPostControlPlaneUpgrade(ctx context.Context, oldClusterSpec *cluster.Spec, clusterSpec *cluster.Spec, workloadCluster *types.Cluster, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

func (p *Provider) RunPostControlPlaneCreation(ctx context.Context, clusterSpec *cluster.Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

func (p *Provider) MachineDeploymentsToDelete(workloadCluster *types.Cluster, currentSpec, newSpec *cluster.Spec) []string {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) InstallCustomProviderComponents(ctx context.Context, kubeconfigFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) PreCAPIInstallOnBootstrap(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) PostMoveManagementToBootstrap(ctx context.Context, bootstrapCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// TODO(nutanix): figure out if we need something else here
	return nil
}

// PreCoreComponentsUpgrade staisfies the Provider interface.
func (p *Provider) PreCoreComponentsUpgrade(
	ctx context.Context,
	cluster *types.Cluster,
	managementComponents *cluster.ManagementComponents,
	clusterSpec *cluster.Spec,
) error {
	_ = "STUB: not implemented"
	return nil
}
