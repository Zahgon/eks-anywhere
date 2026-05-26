package vsphere

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"time"

	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	CredentialsObjectName    = "vsphere-credentials"
	eksaLicense              = "EKSA_LICENSE"
	vSphereUsernameKey       = "VSPHERE_USERNAME"
	vSpherePasswordKey       = "VSPHERE_PASSWORD"
	vSphereServerKey         = "VSPHERE_SERVER"
	govcDatacenterKey        = "GOVC_DATACENTER"
	govcInsecure             = "GOVC_INSECURE"
	expClusterResourceSetKey = "EXP_CLUSTER_RESOURCE_SET"
	defaultTemplateLibrary   = "eks-a-templates"
	defaultTemplatesFolder   = "vm/Templates"
	maxRetries               = 30
	backOffPeriod            = 5 * time.Second
	disk1                    = "Hard disk 1"
	disk2                    = "Hard disk 2"
	MemoryAvailable          = "Memory_Available"
)

const (
	// Documentation URLs.
	vSpherePermissionDoc = "https://anywhere.eks.amazonaws.com/docs/getting-started/vsphere/vsphere-preparation/"
)

//go:embed config/template-cp.yaml
var defaultCAPIConfigCP string

//go:embed config/template-md.yaml
var defaultClusterConfigMD string

//go:embed config/secret.yaml
var defaultSecretObject string

//go:embed config/template-failuredomain.yaml
var defaultFailureDomainConfig string

var (
	eksaVSphereDatacenterResourceType = fmt.Sprintf("vspheredatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaVSphereMachineResourceType    = fmt.Sprintf("vspheremachineconfigs.%s", v1alpha1.GroupVersion.Group)
)

var requiredEnvs = []string{vSphereUsernameKey, vSpherePasswordKey, expClusterResourceSetKey}

type vsphereProvider struct {
	clusterConfig         *v1alpha1.Cluster
	providerGovcClient    ProviderGovcClient
	providerKubectlClient ProviderKubectlClient
	writer                filewriter.FileWriter
	templateBuilder       *VsphereTemplateBuilder
	skipIPCheck           bool
	Retrier               *retrier.Retrier
	validator             *Validator
	defaulter             *Defaulter
	ipValidator           IPValidator
	skippedValidations    map[string]bool
}

type ProviderGovcClient interface {
	SearchTemplate(ctx context.Context, datacenter, template string) (string, error)
	LibraryElementExists(ctx context.Context, library string) (bool, error)
	GetLibraryElementContentVersion(ctx context.Context, element string) (string, error)
	DeleteLibraryElement(ctx context.Context, element string) error
	TemplateHasSnapshot(ctx context.Context, template string) (bool, error)
	GetWorkloadAvailableSpace(ctx context.Context, datastore string) (float64, error)
	ValidateVCenterSetupMachineConfig(ctx context.Context, datacenterConfig *v1alpha1.VSphereDatacenterConfig, machineConfig *v1alpha1.VSphereMachineConfig, selfSigned *bool) error
	ValidateFailureDomainConfig(ctx context.Context, datacenterConfig *v1alpha1.VSphereDatacenterConfig, failureDomain *v1alpha1.FailureDomain) error
	ValidateVCenterConnection(ctx context.Context, server string) error
	ValidateVCenterAuthentication(ctx context.Context) error
	IsCertSelfSigned(ctx context.Context) bool
	GetCertThumbprint(ctx context.Context) (string, error)
	ConfigureCertThumbprint(ctx context.Context, server, thumbprint string) error
	DatacenterExists(ctx context.Context, datacenter string) (bool, error)
	NetworkExists(ctx context.Context, network string) (bool, error)
	GetFolderPath(ctx context.Context, datacenter string, folder string, envMap map[string]string) (string, error)
	GetDatastorePath(ctx context.Context, datacenter string, datastorePath string, envMap map[string]string) (string, error)
	GetResourcePoolPath(ctx context.Context, datacenter string, resourcePool string, envMap map[string]string) (string, error)
	GetComputeClusterPath(ctx context.Context, datacenter string, computeCluster string, envMap map[string]string) (string, error)
	CreateLibrary(ctx context.Context, datastore, library string) error
	DeployTemplateFromLibrary(ctx context.Context, templateDir, templateName, library, datacenter, datastore, network, resourcePool string, resizeDisk2 bool) error
	ImportTemplate(ctx context.Context, library, ovaURL, name string) error
	GetVMDiskSizeInGB(ctx context.Context, vm, datacenter string) (int, error)
	GetTags(ctx context.Context, path string) (tags []string, err error)
	ListTags(ctx context.Context) ([]executables.Tag, error)
	CreateTag(ctx context.Context, tag, category string) error
	AddTag(ctx context.Context, path, tag string) error
	ListCategories(ctx context.Context) ([]string, error)
	CreateCategoryForVM(ctx context.Context, name string) error
	CreateUser(ctx context.Context, username, password string) error
	UserExists(ctx context.Context, username string) (bool, error)
	CreateGroup(ctx context.Context, name string) error
	GroupExists(ctx context.Context, name string) (bool, error)
	AddUserToGroup(ctx context.Context, name, username string) error
	RoleExists(ctx context.Context, name string) (bool, error)
	CreateRole(ctx context.Context, name string, privileges []string) error
	SetGroupRoleOnObject(ctx context.Context, principal, role, object, domain string) error
	GetHardDiskSize(ctx context.Context, vm, datacenter string) (map[string]float64, error)
	GetResourcePoolInfo(ctx context.Context, datacenter, resourcepool string, args ...string) (map[string]int, error)
}

type ProviderKubectlClient interface {
	ApplyKubeSpecFromBytes(ctx context.Context, cluster *types.Cluster, data []byte) error
	CreateNamespaceIfNotPresent(ctx context.Context, kubeconfig, namespace string) error
	LoadSecret(ctx context.Context, secretObject, secretObjType, secretObjectName, kubeConfFile string) error
	GetEksaCluster(ctx context.Context, cluster *types.Cluster, clusterName string) (*v1alpha1.Cluster, error)
	GetEksaVSphereDatacenterConfig(ctx context.Context, vsphereDatacenterConfigName, kubeconfigFile, namespace string) (*v1alpha1.VSphereDatacenterConfig, error)
	GetEksaVSphereMachineConfig(ctx context.Context, vsphereMachineConfigName, kubeconfigFile, namespace string) (*v1alpha1.VSphereMachineConfig, error)
	GetMachineDeployment(ctx context.Context, machineDeploymentName string, opts ...executables.KubectlOpt) (*clusterv1beta2.MachineDeployment, error)
	GetKubeadmControlPlane(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...executables.KubectlOpt) (*controlplanev1beta2.KubeadmControlPlane, error)
	GetEtcdadmCluster(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...executables.KubectlOpt) (*etcdv1.EtcdadmCluster, error)
	GetSecretFromNamespace(ctx context.Context, kubeconfigFile, name, namespace string) (*corev1.Secret, error)
	UpdateAnnotation(ctx context.Context, resourceType, objectName string, annotations map[string]string, opts ...executables.KubectlOpt) error
	RemoveAnnotationInNamespace(ctx context.Context, resourceType, objectName, key string, cluster *types.Cluster, namespace string) error
	SearchVsphereMachineConfig(ctx context.Context, name, kubeconfigFile, namespace string) ([]*v1alpha1.VSphereMachineConfig, error)
	SearchVsphereDatacenterConfig(ctx context.Context, name, kubeconfigFile, namespace string) ([]*v1alpha1.VSphereDatacenterConfig, error)
	SetDaemonSetImage(ctx context.Context, kubeconfigFile, name, namespace, container, image string) error
	DeleteEksaDatacenterConfig(ctx context.Context, vsphereDatacenterResourceType, vsphereDatacenterConfigName, kubeconfigFile, namespace string) error
	DeleteEksaMachineConfig(ctx context.Context, vsphereMachineResourceType, vsphereMachineConfigName, kubeconfigFile, namespace string) error
	ApplyTolerationsFromTaintsToDaemonSet(ctx context.Context, oldTaints, newTaints []corev1.Taint, dsName, kubeconfigFile string) error
}

// IPValidator is an interface that defines methods to validate the control plane IP.
type IPValidator interface {
	ValidateControlPlaneIPUniqueness(cluster *v1alpha1.Cluster) error
}

// NewProvider initializes and returns a new vsphereProvider.
func NewProvider(
	datacenterConfig *v1alpha1.VSphereDatacenterConfig,
	clusterConfig *v1alpha1.Cluster,
	providerGovcClient ProviderGovcClient,
	providerKubectlClient ProviderKubectlClient,
	writer filewriter.FileWriter,
	ipValidator IPValidator,
	now types.NowFunc,
	skipIPCheck bool,
	skippedValidations map[string]bool,
) *vsphereProvider {
	_ = "STUB: not implemented" //nolint:revive
	// TODO(g-gaston): ignoring linter error for exported function returning unexported member
	// We should make it exported, but that would involve a bunch of changes, so will do it separately
	return nil
}

// NewProviderCustomNet initializes and returns a new vsphereProvider.
func NewProviderCustomNet(
	datacenterConfig *v1alpha1.VSphereDatacenterConfig,
	clusterConfig *v1alpha1.Cluster,
	providerGovcClient ProviderGovcClient,
	providerKubectlClient ProviderKubectlClient,
	writer filewriter.FileWriter,
	ipValidator IPValidator,
	now types.NowFunc,
	skipIPCheck bool,
	v *Validator,
	skippedValidations map[string]bool,
) *vsphereProvider {
	_ = "STUB: not implemented" //nolint:revive
	// TODO(g-gaston): ignoring linter error for exported function returning unexported member
	// We should make it exported, but that would involve a bunch of changes, so will do it separately
	return nil
}

func (p *vsphereProvider) UpdateKubeConfig(_ *[]byte, _ string) error {
	_ = "STUB: not implemented"
	// customize generated kube config
	return nil
}

func (p *vsphereProvider) BootstrapClusterOpts(spec *cluster.Spec) ([]bootstrapper.BootstrapClusterOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *vsphereProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *vsphereProvider) DatacenterResourceType() string { _ = "STUB: not implemented"; return "" }

func (p *vsphereProvider) MachineResourceType() string { _ = "STUB: not implemented"; return "" }

func (p *vsphereProvider) generateSSHKeysIfNotSet(machineConfigs map[string]*v1alpha1.VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// use the same key

func (p *vsphereProvider) PostClusterDeleteValidate(_ context.Context, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	// No validations
	return nil
}

func (p *vsphereProvider) PostMoveManagementToBootstrap(_ context.Context, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	// NOOP
	return nil
}

func (p *vsphereProvider) SetupAndValidateCreateCluster(ctx context.Context, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate machine config networks right after basic vCenter validation

// TODO: move this to validator

func (p *vsphereProvider) SetupAndValidateUpgradeCluster(ctx context.Context, cluster *types.Cluster, clusterSpec, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate machine config networks right after basic vCenter validation

// SetupAndValidateUpgradeManagementComponents performs necessary setup for upgrade management components operation.
func (p *vsphereProvider) SetupAndValidateUpgradeManagementComponents(ctx context.Context, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) validateMachineConfigsNameUniqueness(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

type datastoreUsage struct {
	availableSpace float64
	needGiBSpace   int
}

func (p *vsphereProvider) getPrevMachineConfigDatastoreUsage(ctx context.Context, machineConfig *v1alpha1.VSphereMachineConfig, cluster *types.Cluster, count int) (diskGiB float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *vsphereProvider) getMachineConfigDatastoreRequirements(ctx context.Context, machineConfig *v1alpha1.VSphereMachineConfig, count int) (available float64, need int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// TODO: remove dependency on machineConfig

func (p *vsphereProvider) calculateDatastoreUsage(ctx context.Context, machineConfig *v1alpha1.VSphereMachineConfig, cluster *types.Cluster, usage map[string]*datastoreUsage, prevCount, newCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func updateDatastoreUsageMap(machineConfig *v1alpha1.VSphereMachineConfig, needGiB int, availableSpace, prevUsage float64, usage map[string]*datastoreUsage) {
	_ = "STUB: not implemented"
	return
}

func (p *vsphereProvider) validateDatastoreUsageForUpgrade(ctx context.Context, currentClusterSpec *Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) validateDatastoreUsageForCreate(ctx context.Context, vsphereClusterSpec *Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// getPrevMachineConfigMemoryUsage returns the memoryMiB freed up from the given machineConfig based on the count.
func (p *vsphereProvider) getPrevMachineConfigMemoryUsage(ctx context.Context, mc *v1alpha1.VSphereMachineConfig, cluster *types.Cluster, machineConfigCount int) (memoryMiB int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// getMachineConfigMemoryAvailability accepts a machine config and returns available memory in the config's resource pool along with needed memory for the machine config.
func (p *vsphereProvider) getMachineConfigMemoryAvailability(ctx context.Context, datacenter string, mc *v1alpha1.VSphereMachineConfig, machineConfigCount int) (availableMemoryMiB, needMemoryMiB int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// updateMemoryUsageMap updates the memory availability for the machine config's resource pool.
func updateMemoryUsageMap(mc *v1alpha1.VSphereMachineConfig, needMiB, availableMiB int, mu map[string]int) {
	_ = "STUB: not implemented"
	return
}

// needMiB can be ignored when the resource pool memory limit is unset

func addPrevMachineConfigMemoryUsage(mc *v1alpha1.VSphereMachineConfig, prevUsage int, memoryUsage map[string]int) {
	_ = "STUB: not implemented"
	// when the memory limit for the respective resource pool is unset, skip accounting for previous usage and validating the needed memory
	return
}

func (p *vsphereProvider) validateMemoryUsage(ctx context.Context, clusterSpec *Spec, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// account for previous cluster resources that are freed up during upgrade.

// updatePrevClusterMemoryUsage calculates memory freed up from previous CP and worker nodes during upgrade and adds up the memory usage for the specific resource pool.
func (p *vsphereProvider) updatePrevClusterMemoryUsage(ctx context.Context, clusterSpec *Spec, cluster *types.Cluster, memoryUsage map[string]int) error {
	_ = "STUB: not implemented"
	return nil
}

// The last CP machine is deleted only after the desired number of new worker machines are rolled out, so don't add it's memory

// The last worker machine is deleted only after the desired number of new worker machines are rolled out, so don't add it's memory

func (p *vsphereProvider) UpdateSecrets(ctx context.Context, cluster *types.Cluster, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) SetupAndValidateDeleteCluster(ctx context.Context, _ *types.Cluster, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) createSecret(ctx context.Context, cluster *types.Cluster, contents *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) PreCAPIInstallOnBootstrap(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) PostBootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) PostWorkloadInit(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"

	// EnvMap returns a map of environment variables required for the vsphere provider.
	return nil
}

func (p *vsphereProvider) EnvMap(_ *cluster.ManagementComponents, _ *cluster.Spec) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *vsphereProvider) GetDeployments() map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) DatacenterConfig(spec *cluster.Spec) providers.DatacenterConfig {
	_ = "STUB: not implemented"
	return *new(providers.DatacenterConfig)
}

func (p *vsphereProvider) MachineConfigs(spec *cluster.Spec) []providers.MachineConfig {
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

func (p *vsphereProvider) ValidateNewSpec(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) validateMachineConfigImmutability(ctx context.Context, cluster *types.Cluster, newConfig *v1alpha1.VSphereMachineConfig, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) secretContentsChanged(ctx context.Context, workloadCluster *types.Cluster) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ChangeDiff returns the component change diff for the provider.
func (p *vsphereProvider) ChangeDiff(currentComponents, newComponents *cluster.ManagementComponents) *types.ComponentChangeDiff {
	_ = "STUB: not implemented"
	return nil
}

// GetInfrastructureBundle returns the infrastructure bundle for the provider.
func (p *vsphereProvider) GetInfrastructureBundle(components *cluster.ManagementComponents) *types.InfrastructureBundle {
	_ = "STUB: not implemented"
	return nil
}

// Version returns the version of the provider.
func (p *vsphereProvider) Version(components *cluster.ManagementComponents) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *vsphereProvider) RunPostControlPlaneUpgrade(_ context.Context, _, _ *cluster.Spec, _, _ *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func cpiResourceSetName(clusterSpec *cluster.Spec) string { _ = "STUB: not implemented"; return "" }

func machineRefSliceToMap(machineRefs []v1alpha1.Ref) map[string]v1alpha1.Ref {
	_ = "STUB: not implemented"
	return nil
}

func (p *vsphereProvider) InstallCustomProviderComponents(ctx context.Context, kubeconfigFile string) error {
	_ = "STUB: not implemented"

	// PreCoreComponentsUpgrade satisfies the Provider interface.
	return nil
}

func (p *vsphereProvider) PreCoreComponentsUpgrade(
	ctx context.Context,
	cluster *types.Cluster,
	managementComponents *cluster.ManagementComponents,
	clusterSpec *cluster.Spec,
) error {
	_ = "STUB: not implemented"
	return nil
}
