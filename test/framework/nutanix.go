package framework

import (
	"context"
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/internal/pkg/nutanix"
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/constants"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

const (
	nutanixEndpoint                     = "T_NUTANIX_ENDPOINT"
	nutanixPort                         = "T_NUTANIX_PORT"
	nutanixAdditionalTrustBundle        = "T_NUTANIX_ADDITIONAL_TRUST_BUNDLE"
	nutanixInsecure                     = "T_NUTANIX_INSECURE"
	nutanixMachineBootType              = "T_NUTANIX_MACHINE_BOOT_TYPE"
	nutanixMachineMemorySize            = "T_NUTANIX_MACHINE_MEMORY_SIZE"
	nutanixSystemDiskSize               = "T_NUTANIX_SYSTEMDISK_SIZE"
	nutanixMachineVCPUsPerSocket        = "T_NUTANIX_MACHINE_VCPU_PER_SOCKET"
	nutanixMachineVCPUSocket            = "T_NUTANIX_MACHINE_VCPU_SOCKET"
	nutanixPrismElementClusterName      = "T_NUTANIX_PRISM_ELEMENT_CLUSTER_NAME"
	nutanixSSHAuthorizedKey             = "T_NUTANIX_SSH_AUTHORIZED_KEY"
	nutanixSubnetName                   = "T_NUTANIX_SUBNET_NAME"
	nutanixControlPlaneEndpointIP       = "T_NUTANIX_CONTROL_PLANE_ENDPOINT_IP"
	nutanixControlPlaneCidrVar          = "T_NUTANIX_CONTROL_PLANE_CIDR"
	nutanixPodCidrVar                   = "T_NUTANIX_POD_CIDR"
	nutanixServiceCidrVar               = "T_NUTANIX_SERVICE_CIDR"
	nutanixTemplateNameUbuntu2204129Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2204_1_29"
	nutanixTemplateNameUbuntu2204130Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2204_1_30"
	nutanixTemplateNameUbuntu2204131Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2204_1_31"
	nutanixTemplateNameUbuntu2204132Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2204_1_32"
	nutanixTemplateNameUbuntu2204133Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2204_1_33"
	nutanixTemplateNameUbuntu2204134Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2204_1_34"
	nutanixTemplateNameUbuntu2204135Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2204_1_35"
	nutanixTemplateNameUbuntu2404129Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2404_1_29"
	nutanixTemplateNameUbuntu2404130Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2404_1_30"
	nutanixTemplateNameUbuntu2404131Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2404_1_31"
	nutanixTemplateNameUbuntu2404132Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2404_1_32"
	nutanixTemplateNameUbuntu2404133Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2404_1_33"
	nutanixTemplateNameUbuntu2404134Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2404_1_34"
	nutanixTemplateNameUbuntu2404135Var = "T_NUTANIX_TEMPLATE_NAME_UBUNTU_2404_1_35"
	nutanixTemplateNameRedHat129Var     = "T_NUTANIX_TEMPLATE_NAME_REDHAT_1_29"
	nutanixTemplateNameRedHat130Var     = "T_NUTANIX_TEMPLATE_NAME_REDHAT_1_30"
	nutanixTemplateNameRedHat131Var     = "T_NUTANIX_TEMPLATE_NAME_REDHAT_1_31"
	nutanixTemplateNameRedHat9129Var    = "T_NUTANIX_TEMPLATE_NAME_REDHAT_9_1_29"
	nutanixTemplateNameRedHat9130Var    = "T_NUTANIX_TEMPLATE_NAME_REDHAT_9_1_30"
	nutanixTemplateNameRedHat9131Var    = "T_NUTANIX_TEMPLATE_NAME_REDHAT_9_1_31"
	nutanixTemplateNameRedHat9132Var    = "T_NUTANIX_TEMPLATE_NAME_REDHAT_9_1_32"
	nutanixTemplateNameRedHat9133Var    = "T_NUTANIX_TEMPLATE_NAME_REDHAT_9_1_33"
	nutanixTemplateNameRedHat9134Var    = "T_NUTANIX_TEMPLATE_NAME_REDHAT_9_1_34"
	nutanixTemplateNameRedHat9135Var    = "T_NUTANIX_TEMPLATE_NAME_REDHAT_9_1_35"
)

var requiredNutanixEnvVars = []string{
	constants.EksaNutanixUsernameKey,
	constants.EksaNutanixPasswordKey,
	nutanixEndpoint,
	nutanixPort,
	nutanixAdditionalTrustBundle,
	nutanixMachineBootType,
	nutanixMachineMemorySize,
	nutanixSystemDiskSize,
	nutanixMachineVCPUsPerSocket,
	nutanixMachineVCPUSocket,
	nutanixPrismElementClusterName,
	nutanixSSHAuthorizedKey,
	nutanixSubnetName,
	nutanixPodCidrVar,
	nutanixServiceCidrVar,
	nutanixTemplateNameUbuntu2204129Var,
	nutanixTemplateNameUbuntu2204130Var,
	nutanixTemplateNameUbuntu2204131Var,
	nutanixTemplateNameUbuntu2204132Var,
	nutanixTemplateNameUbuntu2204133Var,
	nutanixTemplateNameUbuntu2204134Var,
	nutanixTemplateNameUbuntu2204135Var,
	nutanixTemplateNameUbuntu2404129Var,
	nutanixTemplateNameUbuntu2404130Var,
	nutanixTemplateNameUbuntu2404131Var,
	nutanixTemplateNameUbuntu2404132Var,
	nutanixTemplateNameUbuntu2404133Var,
	nutanixTemplateNameUbuntu2404134Var,
	nutanixTemplateNameUbuntu2404135Var,
	nutanixTemplateNameRedHat129Var,
	nutanixTemplateNameRedHat130Var,
	nutanixTemplateNameRedHat131Var,
	nutanixTemplateNameRedHat9129Var,
	nutanixTemplateNameRedHat9130Var,
	nutanixTemplateNameRedHat9131Var,
	nutanixTemplateNameRedHat9132Var,
	nutanixTemplateNameRedHat9133Var,
	nutanixTemplateNameRedHat9134Var,
	nutanixTemplateNameRedHat9135Var,
	nutanixInsecure,
}

type Nutanix struct {
	t                      *testing.T
	fillers                []api.NutanixFiller
	clusterFillers         []api.ClusterFiller
	client                 nutanix.PrismClient
	controlPlaneEndpointIP string
	cpCidr                 string
	podCidr                string
	serviceCidr            string
	devRelease             *releasev1.EksARelease
	templatesRegistry      *templateRegistry
}

type NutanixOpt func(*Nutanix)

func NewNutanix(t *testing.T, opts ...NutanixOpt) *Nutanix { _ = "STUB: not implemented"; return nil }

// Assumption: generated clusterconfig by nutanix provider sets name as id type by default.
// for uuid specific id type, we will set it thru each specific test so that current CI
// works as is with name id type for following resources

// RequiredNutanixEnvVars returns a list of environment variables needed for Nutanix tests.
func RequiredNutanixEnvVars() []string { _ = "STUB: not implemented"; return nil }

// Name returns the provider name. It satisfies the test framework Provider.
func (n *Nutanix) Name() string {
	_ = "STUB: not implemented"

	// Setup does nothing. It satisfies the test framework Provider.
	return ""
}

func (n *Nutanix) Setup() {
	_ = "STUB: not implemented"

	// UpdateKubeConfig customizes generated kubeconfig for the provider.
	return
}

func (n *Nutanix) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"

	// CleanupResources satisfies the test framework Provider.
	return nil
}

func (n *Nutanix) CleanupResources(clustername string) error { _ = "STUB: not implemented"; return nil }

// ClusterConfigUpdates satisfies the test framework Provider.
func (n *Nutanix) ClusterConfigUpdates() []api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return nil
}

// WithProviderUpgrade returns a ClusterE2EOpt that updates the cluster config for provider-specific upgrade.
func (n *Nutanix) WithProviderUpgrade(fillers ...api.NutanixFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithKubeVersionAndOS returns a cluster config filler that sets the cluster kube version and the right template for all
// nutanix machine configs.
func (n *Nutanix) WithKubeVersionAndOS(_ anywherev1.KubernetesVersion, _ OS, _ *releasev1.EksARelease, _ ...string) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	// TODO: Update tests to use this
	return *new(api.ClusterConfigFiller)
}

// WithNewWorkerNodeGroup returns an api.ClusterFiller that adds a new workerNodeGroupConfiguration and
// a corresponding NutanixMachineConfig to the cluster config.
func (n *Nutanix) WithNewWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	// TODO: Implement for Nutanix provider
	return *new(api.ClusterConfigFiller)
}

// withNutanixKubeVersionAndOS returns a NutanixOpt that adds API fillers to use a Nutanix template for
// the specified OS family and version (default if not provided), corresponding to a particular
// Kubernetes version, in addition to configuring all machine configs to use this OS family.
func withNutanixKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease) NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes129Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template for k8s 1.29
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes129Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes130Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template for k8s 1.30
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes130Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes131Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template for k8s 1.31
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes131Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes132Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template for k8s 1.32
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes132Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes133Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template for k8s 1.33
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes133Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes134Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template for k8s 1.34
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes134Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes129Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template for k8s 1.29
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes129Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes130Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template for k8s 1.30
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes130Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes131Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template for k8s 1.31
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes131Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes132Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template for k8s 1.32
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes132Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes133Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template for k8s 1.33
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes133Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes134Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template for k8s 1.34
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes134Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat129Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 8 Nutanix template for k8s 1.29
// and the "redhat" osFamily in all machine configs.
func WithRedHat129Nutanix() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithRedHat130Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 8 Nutanix template for k8s 1.30
// and the "redhat" osFamily in all machine configs.
func WithRedHat130Nutanix() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithRedHat131Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 8 Nutanix template for k8s 1.31
// and the "redhat" osFamily in all machine configs.
func WithRedHat131Nutanix() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithRedHat9Kubernetes129Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template for k8s 1.29
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes129Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes130Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template for k8s 1.30
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes130Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes131Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template for k8s 1.31
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes131Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes132Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template for k8s 1.32
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes132Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes133Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template for k8s 1.33
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes133Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes134Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template for k8s 1.34
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes134Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204135Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template for k8s 1.35
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204135Nutanix() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithUbuntu2404135Nutanix returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template for k8s 1.35
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404135Nutanix() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithRedHat9135Nutanix returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template for k8s 1.35
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes135Nutanix() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// withNutanixKubeVersionAndOSForUUID returns a NutanixOpt that adds API fillers to use a Nutanix template UUID
// corresponding to the provided OS family and Kubernetes version, in addition to configuring all machine configs
// to use this OS family.
func withNutanixKubeVersionAndOSForUUID(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease) NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat129NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat 8 Nutanix template UUID for k8s 1.29
// and the "redhat" osFamily in all machine configs.
func WithRedHat129NutanixUUID() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithRedHat130NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat Nutanix template UUID for k8s 1.30
// and the "redhat" osFamily in all machine configs.
func WithRedHat130NutanixUUID() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithRedHat131NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat Nutanix template UUID for k8s 1.31
// and the "redhat" osFamily in all machine configs.
func WithRedHat131NutanixUUID() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithRedHat9Kubernetes129NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template UUID for k8s 1.28
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes129NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes130NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template UUID for k8s 1.30
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes130NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes131NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template UUID for k8s 1.31
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes131NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes132NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template UUID for k8s 1.32
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes132NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes133NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template UUID for k8s 1.33
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes133NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithRedHat9Kubernetes134NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template UUID for k8s 1.34
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes134NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204135NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template UUID for k8s 1.35
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204135NutanixUUID() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithUbuntu2404135NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template UUID for k8s 1.35
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404135NutanixUUID() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithRedHat9135NutanixUUID returns a NutanixOpt that adds API fillers to use a RedHat 9 Nutanix template UUID for k8s 1.35
// and the "redhat" osFamily in all machine configs.
func WithRedHat9Kubernetes135NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes129NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template UUID for k8s 1.29
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes129NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes130NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template UUID for k8s 1.30
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes130NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes131NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template UUID for k8s 1.31
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes131NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes132NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template UUID for k8s 1.32
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes132NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes133NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template UUID for k8s 1.33
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes133NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2204Kubernetes134NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 22.04 Nutanix template UUID for k8s 1.34
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204Kubernetes134NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes129NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template UUID for k8s 1.29
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes129NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes130NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template UUID for k8s 1.30
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes130NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes131NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template UUID for k8s 1.31
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes131NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes132NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template UUID for k8s 1.32
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes132NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes133NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template UUID for k8s 1.33
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes133NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

// WithUbuntu2404Kubernetes134NutanixUUID returns a NutanixOpt that adds API fillers to use a Ubuntu 24.04 Nutanix template UUID for k8s 1.34
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404Kubernetes134NutanixUUID() NutanixOpt {
	_ = "STUB: not implemented"
	return *new(NutanixOpt)
}

func (n *Nutanix) withNutanixUUID(name string, osFamily anywherev1.OSFamily) []api.NutanixFiller {
	_ = "STUB: not implemented"
	return nil
}

// WithPrismElementClusterUUID returns a NutanixOpt that adds API fillers to use a PE Cluster UUID.
func WithPrismElementClusterUUID() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// WithNutanixSubnetUUID returns a NutanixOpt that adds API fillers to use a Subnet UUID.
func WithNutanixSubnetUUID() NutanixOpt { _ = "STUB: not implemented"; return *new(NutanixOpt) }

// templateForKubeVersionAndOS returns a Nutanix filler for the given OS and Kubernetes version.
func (n *Nutanix) templateForKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease) api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2204Kubernetes129Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2204Kubernetes129Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2204Kubernetes130Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2204Kubernetes130Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2204Kubernetes131Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2204Kubernetes131Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2204Kubernetes132Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2204Kubernetes132Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2204Kubernetes133Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2204Kubernetes133Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2204Kubernetes134Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2204Kubernetes134Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2404Kubernetes129Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2404Kubernetes129Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2404Kubernetes130Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2404Kubernetes130Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2404Kubernetes131Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2404Kubernetes131Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2404Kubernetes132Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2404Kubernetes132Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2404Kubernetes133Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2404Kubernetes133Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2404Kubernetes134Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2404Kubernetes134Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat129Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat129Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat130Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat130Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat131Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat131Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat9Kubernetes129Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat9Kubernetes129Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat9Kubernetes130Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat9Kubernetes130Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat9Kubernetes131Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat9Kubernetes131Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat9Kubernetes132Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat9Kubernetes132Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat9Kubernetes133Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat9Kubernetes133Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat9Kubernetes134Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat9Kubernetes134Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2204135Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2204135Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// Ubuntu2404135Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) Ubuntu2404135Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// RedHat9135Template returns NutanixFiller by reading the env var and setting machine config's
// image name parameter in the spec.
func (n *Nutanix) RedHat9135Template() api.NutanixFiller {
	_ = "STUB: not implemented"
	return *new(api.NutanixFiller)
}

// ClusterStateValidations returns a list of provider specific ClusterStateValidations.
func (n *Nutanix) ClusterStateValidations() []clusterf.StateValidation {
	_ = "STUB: not implemented"
	return nil
}

func (n *Nutanix) getDevRelease() *releasev1.EksARelease { _ = "STUB: not implemented"; return nil }

func (n *Nutanix) templateForDevRelease(kubeVersion anywherev1.KubernetesVersion, os OS, useBundlesOverride bool) string {
	_ = "STUB: not implemented"
	return ""
}

// envVarForTemplate looks for explicit configuration through an env var: "T_NUTANIX_TEMPLATE_{osFamily}_{eks-d version}"
// eg: T_NUTANIX_TEMPLATE_UBUNTU_KUBERNETES_1_27_EKS_22.
func (n *Nutanix) envVarForTemplate(os OS, eksDName string) string {
	_ = "STUB: not implemented"
	return ""
}

// defaultNameForTemplate looks for a template: "{eks-d version}-{osFamily}"
// eg: kubernetes-1-27-eks-22-ubuntu.
func (n *Nutanix) defaultNameForTemplate(os OS, eksDName string) string {
	_ = "STUB: not implemented"
	return ""
}

// defaultEnvVarForTemplate returns the value of the default template env vars: "T_NUTANIX_TEMPLATE_{osFamily}_{kubeVersion}"
// eg. T_NUTANIX_TEMPLATE_UBUNTU_1_27.
func (n *Nutanix) defaultEnvVarForTemplate(os OS, kubeVersion anywherev1.KubernetesVersion) string {
	_ = "STUB: not implemented"
	return ""
}

// searchTemplate returns template name if the given template exists in Prism Central.
func (n *Nutanix) searchTemplate(ctx context.Context, template string) (string, error) {
	_ = "STUB: not implemented"
	// TODO: implement search functionality for Nutanix templates
	return "", nil
}
