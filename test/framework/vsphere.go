package framework

import (
	"context"
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/executables"
	anywheretypes "github.com/aws/eks-anywhere/pkg/types"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

const (
	vsphereDatacenterVar        = "T_VSPHERE_DATACENTER"
	vsphereDatastoreVar         = "T_VSPHERE_DATASTORE"
	vsphereFolderVar            = "T_VSPHERE_FOLDER"
	vsphereNetworkVar           = "T_VSPHERE_NETWORK"
	vsphereSecondNetworkVar     = "T_VSPHERE_SECOND_NETWORK"
	vspherePrivateNetworkVar    = "T_VSPHERE_PRIVATE_NETWORK"
	vsphereResourcePoolVar      = "T_VSPHERE_RESOURCE_POOL"
	vsphereServerVar            = "T_VSPHERE_SERVER"
	vsphereSshAuthorizedKeyVar  = "T_VSPHERE_SSH_AUTHORIZED_KEY"
	vsphereStoragePolicyNameVar = "T_VSPHERE_STORAGE_POLICY_NAME"
	vsphereTlsInsecureVar       = "T_VSPHERE_TLS_INSECURE"
	vsphereTlsThumbprintVar     = "T_VSPHERE_TLS_THUMBPRINT"
	vsphereUsernameVar          = "EKSA_VSPHERE_USERNAME"
	vspherePasswordVar          = "EKSA_VSPHERE_PASSWORD"
	cidrVar                     = "T_VSPHERE_CIDR"
	privateNetworkCidrVar       = "T_VSPHERE_PRIVATE_NETWORK_CIDR"
	govcUrlVar                  = "VSPHERE_SERVER"
	govcInsecureVar             = "GOVC_INSECURE"
	govcDatacenterVar           = "GOVC_DATACENTER"
	vsphereTemplateEnvVarPrefix = "T_VSPHERE_TEMPLATE_"
	vsphereTemplatesFolder      = "T_VSPHERE_TEMPLATE_FOLDER"
	vsphereTestTagEnvVar        = "T_VSPHERE_TAG"
)

var requiredEnvVars = []string{
	vsphereDatacenterVar,
	vsphereDatastoreVar,
	vsphereFolderVar,
	vsphereNetworkVar,
	vsphereSecondNetworkVar,
	vspherePrivateNetworkVar,
	vsphereResourcePoolVar,
	vsphereServerVar,
	vsphereSshAuthorizedKeyVar,
	vsphereTlsInsecureVar,
	vsphereTlsThumbprintVar,
	vsphereUsernameVar,
	vspherePasswordVar,
	cidrVar,
	privateNetworkCidrVar,
	govcUrlVar,
	govcInsecureVar,
	govcDatacenterVar,
	vsphereTestTagEnvVar,
}

type VSphere struct {
	t                 *testing.T
	testsConfig       vsphereConfig
	fillers           []api.VSphereFiller
	clusterFillers    []api.ClusterFiller
	cidr              string
	GovcClient        *executables.Govc
	devRelease        *releasev1.EksARelease
	templatesRegistry *templateRegistry
}

type vsphereConfig struct {
	Datacenter        string
	Datastore         string
	Folder            string
	Network           string
	ResourcePool      string
	Server            string
	SSHAuthorizedKey  string
	StoragePolicyName string
	TLSInsecure       bool
	TLSThumbprint     string
	TemplatesFolder   string
}

// VSphereOpt is construction option for the E2E vSphere provider.
type VSphereOpt func(*VSphere)

func NewVSphere(t *testing.T, opts ...VSphereOpt) *VSphere { _ = "STUB: not implemented"; return nil }

// withVSphereKubeVersionAndOS returns a VSphereOpt that adds API fillers to use a vSphere template for
// the specified OS family and version (default if not provided), corresponding to a particular
// Kubernetes version, in addition to configuring all machine configs to use this OS family.
func withVSphereKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease) VSphereOpt {
	_ = "STUB: not implemented"
	return *new(VSphereOpt)
}

// WithRedHat129VSphere vsphere test with Redhat 8 for Kubernetes 1.29.
func WithRedHat129VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat130VSphere vsphere test with Redhat 8 for Kubernetes 1.30.
func WithRedHat130VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat131VSphere vsphere test with Redhat 8 for Kubernetes 1.31.
func WithRedHat131VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat9129VSphere vsphere test with Redhat 9 for Kubernetes 1.29.
func WithRedHat9129VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat9130VSphere vsphere test with Redhat 9 for Kubernetes 1.30.
func WithRedHat9130VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat9131VSphere vsphere test with Redhat 9 for Kubernetes 1.31.
func WithRedHat9131VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat9132VSphere vsphere test with Redhat 9 for Kubernetes 1.32.
func WithRedHat9132VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat9133VSphere vsphere test with Redhat 9 for Kubernetes 1.33.
func WithRedHat9133VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat9134VSphere vsphere test with Redhat 9 for Kubernetes 1.34.
func WithRedHat9134VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithRedHat9135VSphere vsphere test with Redhat 9 for Kubernetes 1.35.
func WithRedHat9135VSphere() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithUbuntu2204129 returns a VSphereOpt that adds API fillers to use a Ubuntu 22.04 vSphere template for k8s 1.29
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204129() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithUbuntu2204130 returns a VSphereOpt that adds API fillers to use a Ubuntu 22.04 vSphere template for k8s 1.30
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204130() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithUbuntu2204131 returns a VSphereOpt that adds API fillers to use a Ubuntu 22.04 vSphere template for k8s 1.31
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204131() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithUbuntu2204132 returns a VSphereOpt that adds API fillers to use a Ubuntu 22.04 vSphere template for k8s 1.32
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204132() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithUbuntu2204133 returns a VSphereOpt that adds API fillers to use a Ubuntu 22.04 vSphere template for k8s 1.33
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204133() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithUbuntu2204134 returns a VSphereOpt that adds API fillers to use a Ubuntu 22.04 vSphere template for k8s 1.34
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204134() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithUbuntu2204135 returns a VSphereOpt that adds API fillers to use a Ubuntu 22.04 vSphere template for k8s 1.35
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2204135() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithUbuntu2404135 returns a VSphereOpt that adds API fillers to use a Ubuntu 24.04 vSphere template for k8s 1.35
// and the "ubuntu" osFamily in all machine configs.
func WithUbuntu2404135() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithBottleRocket129 returns br 1.29 var.
func WithBottleRocket129() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithBottleRocket130 returns br 1.30 var.
func WithBottleRocket130() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithBottleRocket131 returns br 1.31 var.
func WithBottleRocket131() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithBottleRocket132 returns br 1.32 var.
func WithBottleRocket132() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithBottleRocket133 returns br 1.33 var.
func WithBottleRocket133() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithBottleRocket134 returns br 1.34 var.
func WithBottleRocket134() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithBottleRocket135 returns br 1.35 var.
func WithBottleRocket135() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithSecondNetwork returns a VSphereMachineConfigFiller that configures second network.
func WithSecondNetwork() api.VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereMachineConfigFiller)
}

func WithPrivateNetwork() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithLinkedCloneMode sets clone mode to LinkedClone for all the machine.
func WithLinkedCloneMode() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithFullCloneMode sets clone mode to FullClone for all the machine.
func WithFullCloneMode() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithDiskGiBForAllMachines sets diskGiB for all the machines.
func WithDiskGiBForAllMachines(value int) VSphereOpt {
	_ = "STUB: not implemented"
	return *new(VSphereOpt)
}

// WithNTPServersForAllMachines sets NTP servers for all the machines.
func WithNTPServersForAllMachines() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

// WithBottlerocketKubernetesSettingsForAllMachines sets Bottlerocket Kubernetes settings for all the machines.
func WithBottlerocketKubernetesSettingsForAllMachines() VSphereOpt {
	_ = "STUB: not implemented"
	return *new(VSphereOpt)
}

// WithSSHAuthorizedKeyForAllMachines sets SSH authorized keys for all the machines.
func WithSSHAuthorizedKeyForAllMachines(sshKey string) VSphereOpt {
	_ = "STUB: not implemented"
	return *new(VSphereOpt)
}

// WithVSphereTags with vsphere tags option.
func WithVSphereTags() VSphereOpt { _ = "STUB: not implemented"; return *new(VSphereOpt) }

func WithVSphereWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup, fillers ...api.VSphereMachineConfigFiller) VSphereOpt {
	_ = "STUB: not implemented"
	return *new(VSphereOpt)
}

// WithMachineTemplate returns an api.ClusterConfigFiller that changes template in machine template.
func (v *VSphere) WithMachineTemplate(machineName, template string) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithNewWorkerNodeGroup returns an api.ClusterFiller that adds a new workerNodeGroupConfiguration and
// a corresponding VSphereMachineConfig to the cluster config.
func (v *VSphere) WithNewWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithWorkerNodeGroupConfiguration returns an api.ClusterFiller that adds a new workerNodeGroupConfiguration item to the cluster config.
func (v *VSphere) WithWorkerNodeGroupConfiguration(name string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// updateMachineSSHAuthorizedKey updates a vsphere machine configs SSHAuthorizedKey.
func updateMachineSSHAuthorizedKey() api.VSphereMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereMachineConfigFiller)
}

// WithVSphereFillers adds VSphereFiller to the provider default fillers.
func WithVSphereFillers(fillers ...api.VSphereFiller) VSphereOpt {
	_ = "STUB: not implemented"
	return *new(VSphereOpt)
}

// Name returns the provider name. It satisfies the test framework Provider.
func (v *VSphere) Name() string {
	_ = "STUB: not implemented"

	// Setup does nothing. It satisfies the test framework Provider.
	return ""
}

func (v *VSphere) Setup() {
	_ = "STUB: not implemented"

	// UpdateKubeConfig customizes generated kubeconfig for the provider.
	return
}

func (v *VSphere) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"

	// ClusterConfigUpdates satisfies the test framework Provider.
	return nil
}

func (v *VSphere) ClusterConfigUpdates() []api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return nil
}

// WithKubeVersionAndOS returns a cluster config filler that sets the cluster kube version and the right template for all
// vsphere machine configs.
func (v *VSphere) WithKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease, _ ...string) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithKubeVersionAndOSMachineConfig returns a cluster config filler that sets the cluster kube version and the right template for a specific
// vsphere machine config.
func (v *VSphere) WithKubeVersionAndOSMachineConfig(name string, kubeVersion anywherev1.KubernetesVersion, os OS) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2204135 returns a cluster config filler that sets the kubernetes version of the cluster to 1.35
// as well as the right Ubuntu 22.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2204135() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2404135 returns a cluster config filler that sets the kubernetes version of the cluster to 1.35
// as well as the right Ubuntu 24.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2404135() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2204129 returns a cluster config filler that sets the kubernetes version of the cluster to 1.29
// as well as the right Ubuntu 22.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2204129() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2204130 returns a cluster config filler that sets the kubernetes version of the cluster to 1.30
// as well as the right Ubuntu 22.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2204130() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2204131 returns a cluster config filler that sets the kubernetes version of the cluster to 1.31
// as well as the right Ubuntu 22.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2204131() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2204132 returns a cluster config filler that sets the kubernetes version of the cluster to 1.32
// as well as the right Ubuntu 22.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2204132() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2204133 returns a cluster config filler that sets the kubernetes version of the cluster to 1.33
// as well as the right Ubuntu 22.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2204133() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2204134 returns a cluster config filler that sets the kubernetes version of the cluster to 1.34
// as well as the right Ubuntu 22.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2204134() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2404133 returns a cluster config filler that sets the kubernetes version of the cluster to 1.33
// as well as the right Ubuntu 24.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2404133() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu2404134 returns a cluster config filler that sets the kubernetes version of the cluster to 1.34
// as well as the right Ubuntu 24.04 template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu2404134() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithBottleRocket134 returns a cluster config filler that sets the kubernetes version of the cluster to 1.34
// as well as the right bottlerocket template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithBottleRocket134() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithBottleRocket131 returns a cluster config filler that sets the kubernetes version of the cluster to 1.31
// as well as the right bottlerocket template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithBottleRocket131() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithBottleRocket132 returns a cluster config filler that sets the kubernetes version of the cluster to 1.32
// as well as the right bottlerocket template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithBottleRocket132() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithBottleRocket133 returns a cluster config filler that sets the kubernetes version of the cluster to 1.33
// as well as the right botlterocket template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithBottleRocket133() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithBottleRocket135 returns a cluster config filler that sets the kubernetes version of the cluster to 1.35
// as well as the right bottlerocket template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithBottleRocket135() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu129 returns a cluster config filler that sets the kubernetes version of the cluster to 1.29
// as well as the right Ubuntu template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu129() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu130 returns a cluster config filler that sets the kubernetes version of the cluster to 1.30
// as well as the right Ubuntu template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu130() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu131 returns a cluster config filler that sets the kubernetes version of the cluster to 1.31
// as well as the right Ubuntu template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu131() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu132 returns a cluster config filler that sets the kubernetes version of the cluster to 1.32
// as well as the right Ubuntu template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu132() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu133 returns a cluster config filler that sets the kubernetes version of the cluster to 1.33
// as well as the right Ubuntu template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu133() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu134 returns a cluster config filler that sets the kubernetes version of the cluster to 1.34
// as well as the right Ubuntu template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu134() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu135 returns a cluster config filler that sets the kubernetes version of the cluster to 1.35
// as well as the right Ubuntu template and osFamily for all VSphereMachineConfigs.
func (v *VSphere) WithUbuntu135() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// CleanupResources deletes all the VMs owned by the test EKS-A cluster. It satisfies the test framework Provider.
func (v *VSphere) CleanupResources(clusterName string) error { _ = "STUB: not implemented"; return nil }

func (v *VSphere) WithProviderUpgrade(fillers ...api.VSphereFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func (v *VSphere) WithProviderUpgradeGit(fillers ...api.VSphereFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithNewVSphereWorkerNodeGroup adds a new worker node group to the cluster config.
func (v *VSphere) WithNewVSphereWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// templateForKubeVersionAndOS returns a vSphere filler for the given OS and Kubernetes version.
func (v *VSphere) templateForKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease) api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// templateForKubeVersionAndOSMachineConfig returns a vSphere filler for the given OS and Kubernetes version for a specific machine config.
func (v *VSphere) templateForKubeVersionAndOSMachineConfig(name string, kubeVersion anywherev1.KubernetesVersion, os OS) api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2204Kubernetes129Template returns vsphere filler for 1.29 Ubuntu 22.04.
func (v *VSphere) Ubuntu2204Kubernetes129Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2204Kubernetes130Template returns vsphere filler for 1.30 Ubuntu 22.04.
func (v *VSphere) Ubuntu2204Kubernetes130Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2204Kubernetes131Template returns vsphere filler for 1.31 Ubuntu 22.04.
func (v *VSphere) Ubuntu2204Kubernetes131Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2204Kubernetes132Template returns vsphere filler for 1.32 Ubuntu 22.04.
func (v *VSphere) Ubuntu2204Kubernetes132Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2204Kubernetes133Template returns vsphere filler for 1.33 Ubuntu 22.04.
func (v *VSphere) Ubuntu2204Kubernetes133Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2204Kubernetes134Template returns vsphere filler for 1.34 Ubuntu 22.04.
func (v *VSphere) Ubuntu2204Kubernetes134Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2204Kubernetes135Template returns vsphere filler for 1.35 Ubuntu 22.04.
func (v *VSphere) Ubuntu2204Kubernetes135Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2404Kubernetes129Template returns vsphere filler for 1.29 Ubuntu 24.04.
func (v *VSphere) Ubuntu2404Kubernetes129Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2404Kubernetes130Template returns vsphere filler for 1.30 Ubuntu 24.04.
func (v *VSphere) Ubuntu2404Kubernetes130Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2404Kubernetes131Template returns vsphere filler for 1.31 Ubuntu 24.04.
func (v *VSphere) Ubuntu2404Kubernetes131Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2404Kubernetes132Template returns vsphere filler for 1.32 Ubuntu 24.04.
func (v *VSphere) Ubuntu2404Kubernetes132Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2404Kubernetes133Template returns vsphere filler for 1.33 Ubuntu 24.04.
func (v *VSphere) Ubuntu2404Kubernetes133Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2404Kubernetes134Template returns vsphere filler for 1.34 Ubuntu 24.04.
func (v *VSphere) Ubuntu2404Kubernetes134Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu2404Kubernetes135Template returns vsphere filler for 1.35 Ubuntu 24.04.
func (v *VSphere) Ubuntu2404Kubernetes135Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu129Template returns vsphere filler for 1.29 Ubuntu.
func (v *VSphere) Ubuntu129Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu130Template returns vsphere filler for 1.30 Ubuntu.
func (v *VSphere) Ubuntu130Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu131Template returns vsphere filler for 1.31 Ubuntu.
func (v *VSphere) Ubuntu131Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu132Template returns vsphere filler for 1.32 Ubuntu.
func (v *VSphere) Ubuntu132Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu133Template returns vsphere filler for 1.33 Ubuntu.
func (v *VSphere) Ubuntu133Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu134Template returns vsphere filler for 1.34 Ubuntu.
func (v *VSphere) Ubuntu134Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Ubuntu135Template returns vsphere filler for 1.35 Ubuntu.
func (v *VSphere) Ubuntu135Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Bottlerocket129Template returns vsphere filler for 1.29 BR.
func (v *VSphere) Bottlerocket129Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Bottlerocket130Template returns vsphere filler for 1.30 BR.
func (v *VSphere) Bottlerocket130Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Bottlerocket131Template returns vsphere filler for 1.31 BR.
func (v *VSphere) Bottlerocket131Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Bottlerocket132Template returns vsphere filler for 1.32 BR.
func (v *VSphere) Bottlerocket132Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Bottlerocket133Template returns vsphere filler for 1.33 BR.
func (v *VSphere) Bottlerocket133Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Bottlerocket134Template returns vsphere filler for 1.34 BR.
func (v *VSphere) Bottlerocket134Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Bottlerocket135Template returns vsphere filler for 1.35 BR.
func (v *VSphere) Bottlerocket135Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat129Template returns vsphere filler for 1.29 Redhat.
func (v *VSphere) Redhat129Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat130Template returns vsphere filler for 1.30 Redhat.
func (v *VSphere) Redhat130Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat131Template returns vsphere filler for 1.31 Redhat.
func (v *VSphere) Redhat131Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat9129Template returns vsphere filler for 1.29 Redhat 9.
func (v *VSphere) Redhat9129Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat9130Template returns vsphere filler for 1.30 Redhat 9.
func (v *VSphere) Redhat9130Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat9131Template returns vsphere filler for 1.31 Redhat 9.
func (v *VSphere) Redhat9131Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat9132Template returns vsphere filler for 1.32 Redhat 9.
func (v *VSphere) Redhat9132Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat9133Template returns vsphere filler for 1.33 Redhat 9.
func (v *VSphere) Redhat9133Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat9134Template returns vsphere filler for 1.34 Redhat 9.
func (v *VSphere) Redhat9134Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Redhat9135Template returns vsphere filler for 1.35 Redhat 9.
func (v *VSphere) Redhat9135Template() api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

func (v *VSphere) getDevRelease() *releasev1.EksARelease { _ = "STUB: not implemented"; return nil }

func (v *VSphere) templateForDevRelease(kubeVersion anywherev1.KubernetesVersion, os OS, useBundlesOverride bool) string {
	_ = "STUB: not implemented"
	return ""
}

func RequiredVsphereEnvVars() []string { _ = "STUB: not implemented"; return nil }

// VSphereExtraEnvVarPrefixes returns prefixes for env vars that although not always required,
// might be necessary for certain tests.
func VSphereExtraEnvVarPrefixes() []string { _ = "STUB: not implemented"; return nil }

func vSphereMachineConfig(name string, fillers ...api.VSphereMachineConfigFiller) api.VSphereFiller {
	_ = "STUB: not implemented"
	return *new(api.VSphereFiller)
}

// Need to add these because at this point the default fillers that assign these
// values to all machines have already ran

func buildVSphereWorkerNodeGroupClusterFiller(machineConfigName string, workerNodeGroup *WorkerNodeGroup) api.ClusterFiller {
	_ = "STUB: not implemented"
	// Set worker node group ref to vsphere machine config
	return *new(api.ClusterFiller)
}

// WithKubeVersionAndOSForRelease returns a vSphereOpt that sets the cluster kube version and the right template for all
// vsphere machine configs based on the EKS-A release.
func WithKubeVersionAndOSForRelease(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease, useBundlesOverride bool) VSphereOpt {
	_ = "STUB: not implemented"
	return *new(VSphereOpt)
}

// WithKubeVersionAndOSForRelease returns a cluster config filler that sets the cluster kube version and the right template for all
// vsphere machine configs based on the EKS-A release.
func (v *VSphere) WithKubeVersionAndOSForRelease(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease, useBundlesOverride bool) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

func optionToSetTemplateForRelease(kubeVersion anywherev1.KubernetesVersion, os OS, release *releasev1.EksARelease, useBundlesOverride bool) VSphereOpt {
	_ = "STUB: not implemented"
	return *new(VSphereOpt)
}

// envVarForTemplate looks for explicit configuration through an env var: "T_VSPHERE_TEMPLATE_{osFamily}_{eks-d version}"
// eg: T_VSPHERE_TEMPLATE_REDHAT_KUBERNETES_1_27_EKS_22.
func (v *VSphere) envVarForTemplate(os OS, eksDName string) string {
	_ = "STUB: not implemented"
	return ""
}

// defaultNameForTemplate looks for a template with the name path: "{folder}/{eks-d version}-{osFamily}"
// eg: /SDDC-Datacenter/vm/Templates/kubernetes-1-27-eks-22-redhat.
func (v *VSphere) defaultNameForTemplate(os OS, eksDName string) string {
	_ = "STUB: not implemented"
	return ""
}

// defaultEnvVarForTemplate returns the value of the default template env vars: "T_VSPHERE_TEMPLATE_{osFamily}_{kubeVersion}"
// eg. T_VSPHERE_TEMPLATE_REDHAT_1_27.
func (v *VSphere) defaultEnvVarForTemplate(os OS, kubeVersion anywherev1.KubernetesVersion) string {
	_ = "STUB: not implemented"
	return ""
}

// searchTemplate returns template name if the given template exists in the datacenter.
func (v *VSphere) searchTemplate(ctx context.Context, template string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func readVersionsBundles(t testing.TB, release *releasev1.EksARelease, kubeVersion anywherev1.KubernetesVersion, useBundlesOverride bool) *releasev1.VersionsBundle {
	_ = "STUB: not implemented"
	return nil
}

func readVSphereConfig() (vsphereConfig, error) {
	_ = "STUB: not implemented"
	return *new(vsphereConfig), nil
}

// ClusterStateValidations returns a list of provider specific validations.
func (v *VSphere) ClusterStateValidations() []clusterf.StateValidation {
	_ = "STUB: not implemented"
	return nil
}

// ValidateNodesDiskGiB validates DiskGiB for all the machines.
func (v *VSphere) ValidateNodesDiskGiB(machines map[string]anywheretypes.Machine, expectedDiskSize int) error {
	_ = "STUB: not implemented"
	return nil
}
