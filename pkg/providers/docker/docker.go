package docker

import (
	"context"
	_ "embed"
	"fmt"
	"io"

	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	githubTokenEnvVar = "GITHUB_TOKEN"
)

//go:embed config/template-cp.yaml
var defaultCAPIConfigCP string

//go:embed config/template-md.yaml
var defaultCAPIConfigMD string

var eksaDockerResourceType = fmt.Sprintf("dockerdatacenterconfigs.%s", v1alpha1.GroupVersion.Group)

type ProviderClient interface {
	GetDockerLBPort(ctx context.Context, clusterName string) (port string, err error)
}

// Provider implements providers.Provider for the docker cluster-api provider.
type Provider struct {
	docker                ProviderClient
	datacenterConfig      *v1alpha1.DockerDatacenterConfig
	providerKubectlClient ProviderKubectlClient
	templateBuilder       *DockerTemplateBuilder
}

// KubeconfigReader reads the kubeconfig secret from the cluster.
type KubeconfigReader interface {
	GetClusterKubeconfig(ctx context.Context, clusterName, kubeconfigPath string) ([]byte, error)
}

// KubeconfigWriter reads the kubeconfig secret on a docker cluster and copies the contents to a writer.
type KubeconfigWriter struct {
	docker ProviderClient
	reader KubeconfigReader
}

// InstallCustomProviderComponents is a no-op. It implements providers.Provider.
func (p *Provider) InstallCustomProviderComponents(ctx context.Context, kubeconfigFile string) error {
	_ = "STUB: not implemented"
	return nil
}

type ProviderKubectlClient interface {
	GetEksaCluster(ctx context.Context, cluster *types.Cluster, clusterName string) (*v1alpha1.Cluster, error)
	GetMachineDeployment(ctx context.Context, machineDeploymentName string, opts ...executables.KubectlOpt) (*clusterv1beta2.MachineDeployment, error)
	GetKubeadmControlPlane(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...executables.KubectlOpt) (*controlplanev1beta2.KubeadmControlPlane, error)
	GetEtcdadmCluster(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...executables.KubectlOpt) (*etcdv1.EtcdadmCluster, error)
	UpdateAnnotation(ctx context.Context, resourceType, objectName string, annotations map[string]string, opts ...executables.KubectlOpt) error
}

// NewProvider returns a new Provider.
func NewProvider(providerConfig *v1alpha1.DockerDatacenterConfig, docker ProviderClient, providerKubectlClient ProviderKubectlClient, now types.NowFunc) *Provider {
	_ = "STUB: not implemented"
	return nil
}

// BootstrapClusterOpts returns a list of options to be used when creating the bootstrap cluster.
func (p *Provider) BootstrapClusterOpts(_ *cluster.Spec) ([]bootstrapper.BootstrapClusterOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PreCAPIInstallOnBootstrap is a no-op. It implements providers.Provider.
func (p *Provider) PreCAPIInstallOnBootstrap(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"

	// PostBootstrapSetup is a no-op. It implements providers.Provider.
	return nil
}

func (p *Provider) PostBootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	_ = "STUB: not implemented"

	// PostWorkloadInit is a no-op. It implements providers.Provider.
	return nil
}

func (p *Provider) PostWorkloadInit(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"

	// Name returns the name of the provider.
	return nil
}

func (p *Provider) Name() string { _ = "STUB: not implemented"; return "" }

// DatacenterResourceType returns the resource type for the dockerdatacenterconfigs.
func (p *Provider) DatacenterResourceType() string { _ = "STUB: not implemented"; return "" }

// MachineResourceType returns nothing because docker has no machines. It implements providers.Provider.
func (p *Provider) MachineResourceType() string {
	_ = "STUB: not implemented"

	// PostClusterDeleteValidate is a no-op. It implements providers.Provider.
	return ""
}

func (p *Provider) PostClusterDeleteValidate(_ context.Context, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	// No validations
	return nil
}

// PostMoveManagementToBootstrap is a no-op. It implements providers.Provider.
func (p *Provider) PostMoveManagementToBootstrap(_ context.Context, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	// NOOP
	return nil

	// SetupAndValidateCreateCluster validates the cluster spec and sets up any provider-specific resources.
}

func (p *Provider) SetupAndValidateCreateCluster(ctx context.Context, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupAndValidateDeleteCluster is a no-op. It implements providers.Provider.
func (p *Provider) SetupAndValidateDeleteCluster(ctx context.Context, _ *types.Cluster, _ *cluster.Spec) error {
	_ = "STUB: not implemented"

	// SetupAndValidateUpgradeCluster is a no-op. It implements providers.Provider.
	return nil
}

func (p *Provider) SetupAndValidateUpgradeCluster(ctx context.Context, _ *types.Cluster, _ *cluster.Spec, _ *cluster.Spec) error {
	_ = "STUB: not implemented"

	// SetupAndValidateUpgradeManagementComponents performs necessary setup for upgrade management components operation.
	return nil
}

func (p *Provider) SetupAndValidateUpgradeManagementComponents(_ context.Context, _ *cluster.Spec) error {
	_ = "STUB: not implemented"

	// UpdateSecrets is a no-op. It implements providers.Provider.
	return nil
}

func (p *Provider) UpdateSecrets(ctx context.Context, cluster *types.Cluster, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	// Not implemented
	return nil
}

// NewDockerTemplateBuilder returns a docker template builder object.
func NewDockerTemplateBuilder(now types.NowFunc) *DockerTemplateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// DockerTemplateBuilder builds the docker templates.
type DockerTemplateBuilder struct {
	now types.NowFunc
}

// GenerateCAPISpecControlPlane generates a yaml spec with the CAPI objects representing the control plane.
func (d *DockerTemplateBuilder) GenerateCAPISpecControlPlane(clusterSpec *cluster.Spec, buildOptions ...providers.BuildMapOption) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateCAPISpecWorkers generates a yaml spec with the CAPI objects representing the worker nodes for a particular eks-a cluster.
func (d *DockerTemplateBuilder) GenerateCAPISpecWorkers(clusterSpec *cluster.Spec, workloadTemplateNames, kubeadmconfigTemplateNames map[string]string) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CAPIWorkersSpecWithInitialNames generates a yaml spec with the CAPI objects representing the worker
// nodes for a particular eks-a cluster. It uses default initial names (ended in '-1') for the docker
// machine templates and kubeadm config templates.
func (d *DockerTemplateBuilder) CAPIWorkersSpecWithInitialNames(spec *cluster.Spec) (content []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initialNamesForWorkers(spec *cluster.Spec) (machineTemplateNames, kubeadmConfigTemplateNames map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func kubeletCgroupDriverExtraArgs(kubeVersion v1alpha1.KubernetesVersion) (clusterapi.ExtraArgs, error) {
	_ = "STUB: not implemented"
	return *new(clusterapi.ExtraArgs), nil
}

func buildTemplateMapCP(clusterSpec *cluster.Spec) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fail-cgroupv1 flag was introduced in Kubernetes 1.31

// fail-cgroupv1 flag was introduced in Kubernetes 1.31

func buildTemplateMapMD(clusterSpec *cluster.Spec, workerNodeGroupConfiguration v1alpha1.WorkerNodeGroupConfiguration) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fail-cgroupv1 flag was introduced in Kubernetes 1.31

// fail-cgroupv1 flag was introduced in Kubernetes 1.31

// UpdateKubeConfig updates the kubeconfig secret on a docker cluster.
func (p *Provider) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"
	// The Docker provider is for testing only. We don't want to change the interface just for the test
	return nil
}

// NewKubeconfigWriter creates a KubeconfigWriter.
func NewKubeconfigWriter(docker ProviderClient, reader KubeconfigReader) KubeconfigWriter {
	_ = "STUB: not implemented"
	return *new(KubeconfigWriter)
}

// WriteKubeconfig retrieves the contents of the specified cluster's kubeconfig from a secret and copies it to an io.Writer.
func (kr KubeconfigWriter) WriteKubeconfig(ctx context.Context, clusterName, kubeconfigPath string, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteKubeconfigContent retrieves the contents of the specified cluster's kubeconfig from a secret and copies it to an io.Writer.
func (kr KubeconfigWriter) WriteKubeconfigContent(ctx context.Context, clusterName string, content []byte, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// this is required for docker provider.
func updateKubeconfig(content *[]byte, dockerLbPort string) { _ = "STUB: not implemented"; return }

// Version returns the version of the provider.
func (p *Provider) Version(components *cluster.ManagementComponents) string {
	_ = "STUB: not implemented"
	return ""
}

// EnvMap returns a map of environment variables to be set when running the docker clusterctl command.
func (p *Provider) EnvMap(_ *cluster.ManagementComponents, _ *cluster.Spec) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDeployments returns a map of namespaces to deployments that should be running for the provider.
func (p *Provider) GetDeployments() map[string][]string { _ = "STUB: not implemented"; return nil }

// GetInfrastructureBundle returns the infrastructure bundle for the provider.
func (p *Provider) GetInfrastructureBundle(components *cluster.ManagementComponents) *types.InfrastructureBundle {
	_ = "STUB: not implemented"
	return nil
}

// DatacenterConfig returns the datacenter config for the provider.
func (p *Provider) DatacenterConfig(_ *cluster.Spec) providers.DatacenterConfig {
	_ = "STUB: not implemented"
	return *new(providers.DatacenterConfig)
}

// MachineConfigs is a no-op. It implements providers.Provider.
func (p *Provider) MachineConfigs(_ *cluster.Spec) []providers.MachineConfig {
	_ = "STUB: not implemented"

	// ValidateNewSpec is a no-op. It implements providers.Provider.
	return nil
}

func (p *Provider) ValidateNewSpec(_ context.Context, _ *types.Cluster, _ *cluster.Spec) error {
	_ = "STUB: not implemented"

	// ChangeDiff returns the component change diff for the provider.
	return nil
}

func (p *Provider) ChangeDiff(currentComponents, newComponents *cluster.ManagementComponents) *types.ComponentChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

// RunPostControlPlaneUpgrade is a no-op. It implements providers.Provider.
func (p *Provider) RunPostControlPlaneUpgrade(ctx context.Context, oldClusterSpec *cluster.Spec, clusterSpec *cluster.Spec, workloadCluster *types.Cluster, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"

	// RunPostControlPlaneCreation is a no-op. It implements providers.Provider.
	return nil
}

func (p *Provider) RunPostControlPlaneCreation(ctx context.Context, clusterSpec *cluster.Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func getHAProxyImageRepo(haProxyImage releasev1alpha1.Image) string {
	_ = "STUB: not implemented"
	return ""
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

func populateRegistryMirrorValues(clusterSpec *cluster.Spec, values map[string]interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
