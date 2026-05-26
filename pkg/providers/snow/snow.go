package snow

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	providerValidator "github.com/aws/eks-anywhere/pkg/providers/validator"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	eksaSnowCredentialsFileKey = "EKSA_AWS_CREDENTIALS_FILE"
	eksaSnowCABundlesFileKey   = "EKSA_AWS_CA_BUNDLES_FILE"
	snowCredentialsKey         = "AWS_B64ENCODED_CREDENTIALS"
	snowCertsKey               = "AWS_B64ENCODED_CA_BUNDLES"
	maxRetries                 = 30
	backOffPeriod              = 5 * time.Second
)

var (
	snowDatacenterResourceType = fmt.Sprintf("snowdatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	snowMachineResourceType    = fmt.Sprintf("snowmachineconfigs.%s", v1alpha1.GroupVersion.Group)
)

type SnowProvider struct {
	kubeUnAuthClient KubeUnAuthClient
	retrier          *retrier.Retrier
	configManager    *ConfigManager
	ipValidator      *providerValidator.IPValidator
	skipIpCheck      bool
	log              logr.Logger
}

type KubeUnAuthClient interface {
	KubeconfigClient(kubeconfig string) kubernetes.Client
	Apply(ctx context.Context, kubeconfig string, obj runtime.Object) error
}

func NewProvider(kubeUnAuthClient KubeUnAuthClient, configManager *ConfigManager, skipIpCheck bool) *SnowProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *SnowProvider) SetupAndValidateCreateCluster(ctx context.Context, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) SetupAndValidateUpgradeCluster(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupAndValidateUpgradeManagementComponents performs necessary setup for upgrade management components operation.
func (p *SnowProvider) SetupAndValidateUpgradeManagementComponents(_ context.Context, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) SetupAndValidateDeleteCluster(ctx context.Context, _ *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) UpdateSecrets(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// CAPIObjects generates the control plane and worker nodes objects for snow provider from clusterSpec.
func CAPIObjects(ctx context.Context, log logr.Logger, clusterSpec *cluster.Spec, kubeClient kubernetes.Client) (controlPlaneSpec, workersSpec []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func kubernetesToRuntimeObjects(objs []kubernetes.Object) []runtime.Object {
	_ = "STUB: not implemented"
	return nil
}

// PreCAPIInstallOnBootstrap runs the steps that are provider specific before CAPI is installed on the bootstrap cluster.
func (p *SnowProvider) PreCAPIInstallOnBootstrap(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) PostBootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) PostWorkloadInit(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) BootstrapClusterOpts(_ *cluster.Spec) ([]bootstrapper.BootstrapClusterOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *SnowProvider) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"

	// Version returns the snow version from the management components.
	return nil
}

func (p *SnowProvider) Version(components *cluster.ManagementComponents) string {
	_ = "STUB: not implemented"
	return ""
}

// EnvMap returns the environment variables for the snow provider.
func (p *SnowProvider) EnvMap(managementComponents *cluster.ManagementComponents, clusterSpec *cluster.Spec) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *SnowProvider) GetDeployments() map[string][]string { _ = "STUB: not implemented"; return nil }

// GetInfrastructureBundle returns the infrastructure bundle from the management components.
func (p *SnowProvider) GetInfrastructureBundle(components *cluster.ManagementComponents) *types.InfrastructureBundle {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) DatacenterConfig(clusterSpec *cluster.Spec) providers.DatacenterConfig {
	_ = "STUB: not implemented"
	return *new(providers.DatacenterConfig)
}

func (p *SnowProvider) DatacenterResourceType() string { _ = "STUB: not implemented"; return "" }

func (p *SnowProvider) MachineResourceType() string { _ = "STUB: not implemented"; return "" }

func (p *SnowProvider) MachineConfigs(clusterSpec *cluster.Spec) []providers.MachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) ValidateNewSpec(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"

	// ChangeDiff returns the change diff from the management components.
	return nil
}

func (p *SnowProvider) ChangeDiff(currentComponents, newComponents *cluster.ManagementComponents) *types.ComponentChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

func (p *SnowProvider) RunPostControlPlaneUpgrade(ctx context.Context, oldClusterSpec *cluster.Spec, clusterSpec *cluster.Spec, workloadCluster *types.Cluster, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func bundleImagesEqual(new, old releasev1alpha1.SnowBundle) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *SnowProvider) machineConfigsChanged(ctx context.Context, cluster *types.Cluster, spec *cluster.Spec) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *SnowProvider) datacenterChanged(ctx context.Context, cluster *types.Cluster, spec *cluster.Spec) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// namespaceOrDefault return the object namespace or default if it's empty.
func namespaceOrDefault(obj client.Object) string { _ = "STUB: not implemented"; return "" }

func (p *SnowProvider) PostClusterDeleteValidate(_ context.Context, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	// No validations
	return nil
}

func (p *SnowProvider) PostMoveManagementToBootstrap(_ context.Context, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	// NOOP
	return nil
}

func (p *SnowProvider) InstallCustomProviderComponents(ctx context.Context, kubeconfigFile string) error {
	_ = "STUB: not implemented"

	// PreCoreComponentsUpgrade staisfies the Provider interface.
	return nil
}

func (p *SnowProvider) PreCoreComponentsUpgrade(
	ctx context.Context,
	cluster *types.Cluster,
	managementComponents *cluster.ManagementComponents,
	clusterSpec *cluster.Spec,
) error {
	_ = "STUB: not implemented"
	return nil
}
