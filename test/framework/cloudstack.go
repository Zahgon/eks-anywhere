package framework

import (
	"context"
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/executables"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

const (
	cloudstackDomainVar                         = "T_CLOUDSTACK_DOMAIN"
	cloudstackMultiLevelDomainVar               = "T_CLOUDSTACK_MULTILEVEL_DOMAIN"
	cloudstackZoneVar                           = "T_CLOUDSTACK_ZONE"
	cloudstackZone2Var                          = "T_CLOUDSTACK_ZONE_2"
	cloudstackZone3Var                          = "T_CLOUDSTACK_ZONE_3"
	cloudstackAccountVar                        = "T_CLOUDSTACK_ACCOUNT"
	cloudstackAccountForMultiLevelDomainVar     = "T_CLOUDSTACK_ACCOUNT_FOR_MULTILEVEL_DOMAIN"
	cloudstackNetworkVar                        = "T_CLOUDSTACK_NETWORK"
	cloudstackNetwork2Var                       = "T_CLOUDSTACK_NETWORK_2"
	cloudstackNetwork3Var                       = "T_CLOUDSTACK_NETWORK_3"
	cloudstackCredentialsVar                    = "T_CLOUDSTACK_CREDENTIALS"
	cloudstackCredentials2Var                   = "T_CLOUDSTACK_CREDENTIALS_2"
	cloudstackCredentials3Var                   = "T_CLOUDSTACK_CREDENTIALS_3"
	cloudstackCredentialsForMultiLevelDomainVar = "T_CLOUDSTACK_CREDENTIALS_FOR_MULTILEVEL_DOMAIN"
	cloudstackManagementServerVar               = "T_CLOUDSTACK_MANAGEMENT_SERVER"
	cloudstackManagementServer2Var              = "T_CLOUDSTACK_MANAGEMENT_SERVER_2"
	cloudstackManagementServer3Var              = "T_CLOUDSTACK_MANAGEMENT_SERVER_3"
	cloudstackSSHAuthorizedKeyVar               = "T_CLOUDSTACK_SSH_AUTHORIZED_KEY"
	cloudstackComputeOfferingLargeVar           = "T_CLOUDSTACK_COMPUTE_OFFERING_LARGE"
	cloudstackComputeOfferingLargerVar          = "T_CLOUDSTACK_COMPUTE_OFFERING_LARGER"
	cloudStackClusterIPPoolEnvVar               = "T_CLOUDSTACK_CLUSTER_IP_POOL"
	cloudStackCidrVar                           = "T_CLOUDSTACK_CIDR"
	podCidrVar                                  = "T_CLOUDSTACK_POD_CIDR"
	serviceCidrVar                              = "T_CLOUDSTACK_SERVICE_CIDR"
	cloudstackFeatureGateEnvVar                 = "CLOUDSTACK_PROVIDER"
	cloudstackB64EncodedSecretEnvVar            = "EKSA_CLOUDSTACK_B64ENCODED_SECRET"
)

var requiredCloudStackEnvVars = []string{
	cloudstackAccountVar,
	cloudstackAccountForMultiLevelDomainVar,
	cloudstackDomainVar,
	cloudstackMultiLevelDomainVar,
	cloudstackZoneVar,
	cloudstackZone2Var,
	cloudstackZone3Var,
	cloudstackCredentialsVar,
	cloudstackCredentials2Var,
	cloudstackCredentials3Var,
	cloudstackCredentialsForMultiLevelDomainVar,
	cloudstackAccountVar,
	cloudstackNetworkVar,
	cloudstackNetwork2Var,
	cloudstackNetwork3Var,
	cloudstackManagementServerVar,
	cloudstackManagementServer2Var,
	cloudstackManagementServer3Var,
	cloudstackSSHAuthorizedKeyVar,
	cloudstackComputeOfferingLargeVar,
	cloudstackComputeOfferingLargerVar,
	cloudStackCidrVar,
	podCidrVar,
	serviceCidrVar,
	cloudstackFeatureGateEnvVar,
	cloudstackB64EncodedSecretEnvVar,
}

type CloudStack struct {
	t                 *testing.T
	fillers           []api.CloudStackFiller
	clusterFillers    []api.ClusterFiller
	cidr              string
	podCidr           string
	serviceCidr       string
	cmkClient         *executables.Cmk
	devRelease        *releasev1.EksARelease
	templatesRegistry *templateRegistry
}

type CloudStackOpt func(*CloudStack)

func UpdateLargerCloudStackComputeOffering() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// UpdateAddCloudStackAz4 add availability zone 4 to the cluster spec.
func UpdateAddCloudStackAz4() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// UpdateAddCloudStackAz3 add availability zone 3 to the cluster spec.
func UpdateAddCloudStackAz3() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

func UpdateAddCloudStackAz2() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

func UpdateAddCloudStackAz1() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

func RemoveAllCloudStackAzs() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// CloudStackCredentialsAz1 returns the value of the environment variable for cloudstackCredentialsVar.
func CloudStackCredentialsAz1() string { _ = "STUB: not implemented"; return "" }

func NewCloudStack(t *testing.T, opts ...CloudStackOpt) *CloudStack {
	_ = "STUB: not implemented"
	return nil
}

func WithCloudStackWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup, fillers ...api.CloudStackMachineConfigFiller) CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

// withKubeVersionAndOS returns a CloudStack Opt that adds API fillers to use a CloudStack template for
// the specified OS family and version (default if not provided), corresponding to a particular
// Kubernetes version, in addition to configuring all machine configs to use this OS family.
func withCloudStackKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease) CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

// WithCloudStackRedhat129 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.29.
func WithCloudStackRedhat129() CloudStackOpt { _ = "STUB: not implemented"; return *new(CloudStackOpt) }

// WithCloudStackRedhat130 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.30.
func WithCloudStackRedhat130() CloudStackOpt { _ = "STUB: not implemented"; return *new(CloudStackOpt) }

// WithCloudStackRedhat131 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.31.
func WithCloudStackRedhat131() CloudStackOpt { _ = "STUB: not implemented"; return *new(CloudStackOpt) }

// WithCloudStackRedhat9Kubernetes129 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.29.
func WithCloudStackRedhat9Kubernetes129() CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

// WithCloudStackRedhat9Kubernetes130 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.30.
func WithCloudStackRedhat9Kubernetes130() CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

// WithCloudStackRedhat9Kubernetes131 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.31.
func WithCloudStackRedhat9Kubernetes131() CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

// WithCloudStackRedhat9Kubernetes132 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.32.
func WithCloudStackRedhat9Kubernetes132() CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

// WithCloudStackRedhat9Kubernetes133 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.33.
func WithCloudStackRedhat9Kubernetes133() CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

// WithCloudStackRedhat9Kubernetes134 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.34 RedHat9.
func WithCloudStackRedhat9Kubernetes134() CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

// WithCloudStackRedhat9Kubernetes135 returns a function which can be invoked to configure the Cloudstack object to be compatible with K8s 1.35 RedHat9.
func WithCloudStackRedhat9Kubernetes135() CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

func WithCloudStackFillers(fillers ...api.CloudStackFiller) CloudStackOpt {
	_ = "STUB: not implemented"
	return *new(CloudStackOpt)
}

func (c *CloudStack) Name() string { _ = "STUB: not implemented"; return "" }

func (c *CloudStack) Setup() {
	_ = "STUB: not implemented"

	// UpdateKubeConfig customizes generated kubeconfig for the provider.
	return
}

func (c *CloudStack) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"

	// ClusterConfigUpdates satisfies the test framework Provider.
	return nil
}

func (c *CloudStack) ClusterConfigUpdates() []api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return nil
}

// CleanupResources satisfies the test framework Provider.
func (c *CloudStack) CleanupResources(clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CloudStack) WithProviderUpgrade(fillers ...api.CloudStackFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func (c *CloudStack) WithProviderUpgradeGit(fillers ...api.CloudStackFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func (c *CloudStack) getControlPlaneIP() (string, error) { _ = "STUB: not implemented"; return "", nil }

func RequiredCloudstackEnvVars() []string { _ = "STUB: not implemented"; return nil }

func (c *CloudStack) WithNewCloudStackWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup, fillers ...api.CloudStackMachineConfigFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithNewWorkerNodeGroup returns an api.ClusterFiller that adds a new workerNodeGroupConfiguration and
// a corresponding CloudStackMachineConfig to the cluster config.
func (c *CloudStack) WithNewWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithWorkerNodeGroupConfiguration returns an api.ClusterFiller that adds a new workerNodeGroupConfiguration item to the cluster config.
func (c *CloudStack) WithWorkerNodeGroupConfiguration(name string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

func cloudStackMachineConfig(name string, fillers ...api.CloudStackMachineConfigFiller) api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Need to add these because at this point the default fillers that assign these
// values to all machines have already ran

// templateForKubeVersionAndOS returns a CloudStack filler for the given OS and Kubernetes version.
func (c *CloudStack) templateForKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease) api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat129Template returns cloudstack filler for 1.29 RedHat.
func (c *CloudStack) Redhat129Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat130Template returns cloudstack filler for 1.30 RedHat.
func (c *CloudStack) Redhat130Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat131Template returns cloudstack filler for 1.31 RedHat.
func (c *CloudStack) Redhat131Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat9Kubernetes129Template returns cloudstack filler for 1.29 RedHat.
func (c *CloudStack) Redhat9Kubernetes129Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat9Kubernetes130Template returns cloudstack filler for 1.30 RedHat.
func (c *CloudStack) Redhat9Kubernetes130Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat9Kubernetes131Template returns cloudstack filler for 1.31 RedHat.
func (c *CloudStack) Redhat9Kubernetes131Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat9Kubernetes132Template returns cloudstack filler for 1.32 RedHat.
func (c *CloudStack) Redhat9Kubernetes132Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat9Kubernetes133Template returns cloudstack filler for 1.33 RedHat.
func (c *CloudStack) Redhat9Kubernetes133Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat9Kubernetes134Template returns cloudstack filler for 1.34 RedHat9.
func (c *CloudStack) Redhat9Kubernetes134Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

// Redhat9Kubernetes135Template returns cloudstack filler for 1.35 RedHat9.
func (c *CloudStack) Redhat9Kubernetes135Template() api.CloudStackFiller {
	_ = "STUB: not implemented"
	return *new(api.CloudStackFiller)
}

func buildCloudStackWorkerNodeGroupClusterFiller(machineConfigName string, workerNodeGroup *WorkerNodeGroup) api.ClusterFiller {
	_ = "STUB: not implemented"
	// Set worker node group ref to cloudstack machine config
	return *new(api.ClusterFiller)
}

// ClusterStateValidations returns a list of provider specific validations.
func (c *CloudStack) ClusterStateValidations() []clusterf.StateValidation {
	_ = "STUB: not implemented"
	return nil
}

// WithKubeVersionAndOS returns a cluster config filler that sets the cluster kube version and the right template for all
// cloudstack machine configs.
func (c *CloudStack) WithKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease, _ ...string) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat129 returns a cluster config filler that sets the kubernetes version of the cluster to 1.29
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat129() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat130 returns a cluster config filler that sets the kubernetes version of the cluster to 1.30
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat130() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat131 returns a cluster config filler that sets the kubernetes version of the cluster to 1.31
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat131() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat9Kubernetes129 returns a cluster config filler that sets the kubernetes version of the cluster to 1.29
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat9Kubernetes129() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat9Kubernetes130 returns a cluster config filler that sets the kubernetes version of the cluster to 1.30
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat9Kubernetes130() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat9Kubernetes131 returns a cluster config filler that sets the kubernetes version of the cluster to 1.31
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat9Kubernetes131() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat9Kubernetes132 returns a cluster config filler that sets the kubernetes version of the cluster to 1.32
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat9Kubernetes132() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat9Kubernetes133 returns a cluster config filler that sets the kubernetes version of the cluster to 1.33
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat9Kubernetes133() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat9Kubernetes134 returns a cluster config filler that sets the kubernetes version of the cluster to 1.34
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat9Kubernetes134() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhat9Kubernetes135 returns a cluster config filler that sets the kubernetes version of the cluster to 1.35
// as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhat9Kubernetes135() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithRedhatVersion returns a cluster config filler that sets the kubernetes version of the cluster to the k8s
// version provider, as well as the right redhat template for all CloudStackMachineConfigs.
func (c *CloudStack) WithRedhatVersion(version anywherev1.KubernetesVersion) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

func (c *CloudStack) getDevRelease() *releasev1.EksARelease { _ = "STUB: not implemented"; return nil }

func (c *CloudStack) templateForDevRelease(kubeVersion anywherev1.KubernetesVersion, os OS, useBundlesOverride bool) string {
	_ = "STUB: not implemented"
	return ""
}

// envVarForTemplate Looks for explicit configuration through an env var: "T_CLOUDSTACK_TEMPLATE_{osFamily}_{eks-d version}"
// eg: T_CLOUDSTACK_TEMPLATE_REDHAT_KUBERNETES_1_27_EKS_22.
func (c *CloudStack) envVarForTemplate(os OS, eksDName string) string {
	_ = "STUB: not implemented"
	return ""
}

// defaultNameForTemplate looks for a template: "{eks-d version}-{osFamily}"
// eg: kubernetes-1-27-eks-22-redhat.
func (c *CloudStack) defaultNameForTemplate(os OS, eksDName string) string {
	_ = "STUB: not implemented"
	return ""
}

// defaultEnvVarForTemplate returns the value of the default template env vars: "T_CLOUDSTACK_TEMPLATE_{osFamily}_{kubeVersion}"
// eg. T_CLOUDSTACK_TEMPLATE_REDHAT_1_27.
func (c *CloudStack) defaultEnvVarForTemplate(os OS, kubeVersion anywherev1.KubernetesVersion) string {
	_ = "STUB: not implemented"
	return ""
}

// searchTemplate returns template name if the given template exists in the datacenter.
func (c *CloudStack) searchTemplate(ctx context.Context, template string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
