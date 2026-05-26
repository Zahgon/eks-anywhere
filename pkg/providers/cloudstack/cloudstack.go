package cloudstack

import (
	"context"
	_ "embed"
	"fmt"

	etcdv1beta1 "github.com/aws/etcdadm-controller/api/v1beta1"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack/decoder"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	eksaLicense         = "EKSA_LICENSE"
	etcdTemplateNameKey = "etcdTemplateName"
	cpTemplateNameKey   = "controlPlaneTemplateName"
)

//go:embed config/template-cp.yaml
var defaultCAPIConfigCP string

//go:embed config/template-md.yaml
var defaultClusterConfigMD string

var requiredEnvs = []string{decoder.CloudStackCloudConfigB64SecretKey}

var (
	eksaCloudStackDatacenterResourceType = fmt.Sprintf("cloudstackdatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaCloudStackMachineResourceType    = fmt.Sprintf("cloudstackmachineconfigs.%s", v1alpha1.GroupVersion.Group)
)

type cloudstackProvider struct {
	datacenterConfig      *v1alpha1.CloudStackDatacenterConfig
	clusterConfig         *v1alpha1.Cluster
	providerKubectlClient ProviderKubectlClient
	writer                filewriter.FileWriter
	selfSigned            bool
	templateBuilder       *TemplateBuilder
	validator             ProviderValidator
	execConfig            *decoder.CloudStackExecConfig
	log                   logr.Logger
}

func (p *cloudstackProvider) PreBootstrapSetup(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) PreCAPIInstallOnBootstrap(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) PostBootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) PostWorkloadInit(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) UpdateSecrets(ctx context.Context, cluster *types.Cluster, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) generateSecrets(ctx context.Context, cluster *types.Cluster) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When a secret already exists with the profile name we skip creating it

func generateSecret(profile decoder.CloudStackProfileConfig) *corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func machineRefSliceToMap(machineRefs []v1alpha1.Ref) map[string]v1alpha1.Ref {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) validateMachineConfigImmutability(ctx context.Context, cluster *types.Cluster, newConfig *v1alpha1.CloudStackMachineConfig, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) ValidateNewSpec(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// ChangeDiff returns the component change diff for the provider.
func (p *cloudstackProvider) ChangeDiff(currentComponents, newComponents *cluster.ManagementComponents) *types.ComponentChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) RunPostControlPlaneUpgrade(ctx context.Context, oldClusterSpec *cluster.Spec, clusterSpec *cluster.Spec, workloadCluster *types.Cluster, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// Nothing to do
	return nil
}

type ProviderKubectlClient interface {
	ApplyKubeSpecFromBytes(ctx context.Context, cluster *types.Cluster, data []byte) error
	CreateNamespaceIfNotPresent(ctx context.Context, kubeconfig string, namespace string) error
	LoadSecret(ctx context.Context, secretObject string, secretObjType string, secretObjectName string, kubeConfFile string) error
	GetEksaCluster(ctx context.Context, cluster *types.Cluster, clusterName string) (*v1alpha1.Cluster, error)
	GetEksaCloudStackDatacenterConfig(ctx context.Context, cloudstackDatacenterConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.CloudStackDatacenterConfig, error)
	GetEksaCloudStackMachineConfig(ctx context.Context, cloudstackMachineConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.CloudStackMachineConfig, error)
	GetKubeadmControlPlane(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...executables.KubectlOpt) (*controlplanev1beta2.KubeadmControlPlane, error)
	GetMachineDeployment(ctx context.Context, workerNodeGroupName string, opts ...executables.KubectlOpt) (*clusterv1beta2.MachineDeployment, error)
	GetEtcdadmCluster(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...executables.KubectlOpt) (*etcdv1beta1.EtcdadmCluster, error)
	GetSecretFromNamespace(ctx context.Context, kubeconfigFile, name, namespace string) (*corev1.Secret, error)
	UpdateAnnotation(ctx context.Context, resourceType, objectName string, annotations map[string]string, opts ...executables.KubectlOpt) error
	SearchCloudStackMachineConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.CloudStackMachineConfig, error)
	SearchCloudStackDatacenterConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.CloudStackDatacenterConfig, error)
	DeleteEksaCloudStackDatacenterConfig(ctx context.Context, cloudstackDatacenterConfigName string, kubeconfigFile string, namespace string) error
	DeleteEksaCloudStackMachineConfig(ctx context.Context, cloudstackMachineConfigName string, kubeconfigFile string, namespace string) error
	SetEksaControllerEnvVar(ctx context.Context, envVar, envVarVal, kubeconfig string) error
}

// NewProvider initializes the CloudStack provider object.
func NewProvider(datacenterConfig *v1alpha1.CloudStackDatacenterConfig, clusterConfig *v1alpha1.Cluster, providerKubectlClient ProviderKubectlClient, validator ProviderValidator, writer filewriter.FileWriter, now types.NowFunc, log logr.Logger) *cloudstackProvider {
	_ = "STUB: not implemented" //nolint:revive
	return nil
}

func (p *cloudstackProvider) UpdateKubeConfig(_ *[]byte, _ string) error {
	_ = "STUB: not implemented"
	// customize generated kube config
	return nil
}

func (p *cloudstackProvider) BootstrapClusterOpts(clusterSpec *cluster.Spec) ([]bootstrapper.BootstrapClusterOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *cloudstackProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *cloudstackProvider) DatacenterResourceType() string { _ = "STUB: not implemented"; return "" }

func (p *cloudstackProvider) MachineResourceType() string { _ = "STUB: not implemented"; return "" }

func (p *cloudstackProvider) generateSSHKeysIfNotSet(machineConfigs map[string]*v1alpha1.CloudStackMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// use same key already generated

// generate new key

func (p *cloudstackProvider) setMachineConfigDefaults(clusterSpec *cluster.Spec) {
	_ = "STUB: not implemented"
	return
}

func (p *cloudstackProvider) validateManagementApiEndpoint(rawurl string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) validateEnv(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) validateClusterSpec(ctx context.Context, clusterSpec *cluster.Spec) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) SetupAndValidateCreateCluster(ctx context.Context, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) SetupAndValidateUpgradeCluster(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec, currentSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) SetupAndValidateDeleteCluster(ctx context.Context, _ *types.Cluster, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupAndValidateUpgradeManagementComponents performs necessary setup for upgrade management components operation.
func (p *cloudstackProvider) SetupAndValidateUpgradeManagementComponents(ctx context.Context, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func isEqualMap[K, V comparable](a, b map[K]V) bool { _ = "STUB: not implemented"; return false }

// Ensure all keys are present in b, and a's values equal b's values.

func hasSameAvailabilityZones(old, nw []v1alpha1.CloudStackAvailabilityZone) bool {
	_ = "STUB: not implemented"
	return false
}

// Equality of availability zones doesn't take into consideration the availability zones
// ManagementApiEndpoint. Its unclear why this is the case. The ManagementApiEndpoint seems
// to only be used for proxy configuration.

func (p *cloudstackProvider) CleanupProviderInfrastructure(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) BootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// Nothing to do
	return nil
}

// Version returns the version of the provider.
func (p *cloudstackProvider) Version(componnets *cluster.ManagementComponents) string {
	_ = "STUB: not implemented"
	return ""
}

// EnvMap returns a map of environment variables required for the cloudstack provider.
func (p *cloudstackProvider) EnvMap(_ *cluster.ManagementComponents, _ *cluster.Spec) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *cloudstackProvider) GetDeployments() map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

// GetInfrastructureBundle returns the infrastructure bundle for the provider.
func (p *cloudstackProvider) GetInfrastructureBundle(components *cluster.ManagementComponents) *types.InfrastructureBundle {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) DatacenterConfig(clusterSpec *cluster.Spec) providers.DatacenterConfig {
	_ = "STUB: not implemented"
	return *new(providers.DatacenterConfig)
}

func (p *cloudstackProvider) MachineConfigs(spec *cluster.Spec) []providers.MachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func annotateMachineConfig(spec *cluster.Spec, machineConfigName, annotationKey, annotationValue string) {
	_ = "STUB: not implemented"
	return
}

func setMachineConfigManagedBy(spec *cluster.Spec, machineConfigName string) {
	_ = "STUB: not implemented"
	return
}

func (p *cloudstackProvider) PostClusterDeleteValidate(_ context.Context, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	// No validations
	return nil
}

func (p *cloudstackProvider) PostMoveManagementToBootstrap(_ context.Context, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	// NOOP
	return nil
}

func (p *cloudstackProvider) validateMachineConfigsNameUniqueness(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) validateMachineConfigNameUniqueness(ctx context.Context, machineConfigName string, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudstackProvider) InstallCustomProviderComponents(ctx context.Context, kubeconfigFile string) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCoreComponentsUpgrade staisfies the Provider interface.
func (p *cloudstackProvider) PreCoreComponentsUpgrade(
	ctx context.Context,
	cluster *types.Cluster,
	managementComponents *cluster.ManagementComponents,
	clusterSpec *cluster.Spec,
) error {
	_ = "STUB: not implemented"
	return nil
}
