package framework

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

const (
	snowAMIIDUbuntu128   = "T_SNOW_AMIID_UBUNTU_1_28"
	snowDevices          = "T_SNOW_DEVICES"
	snowControlPlaneCidr = "T_SNOW_CONTROL_PLANE_CIDR"
	snowPodCidr          = "T_SNOW_POD_CIDR"
	snowCredentialsFile  = "EKSA_AWS_CREDENTIALS_FILE"
	snowCertificatesFile = "EKSA_AWS_CA_BUNDLES_FILE"
	snowIPPoolIPStart    = "T_SNOW_IPPOOL_IPSTART"
	snowIPPoolIPEnd      = "T_SNOW_IPPOOL_IPEND"
	snowIPPoolGateway    = "T_SNOW_IPPOOL_GATEWAY"
	snowIPPoolSubnet     = "T_SNOW_IPPOOL_SUBNET"

	snowEc2TagPrefix = "sigs.k8s.io/cluster-api-provider-aws-snow/cluster/"
)

var requiredSnowEnvVars = []string{
	snowDevices,
	snowControlPlaneCidr,
	snowCredentialsFile,
	snowCertificatesFile,
}

type Snow struct {
	t              *testing.T
	fillers        []api.SnowFiller
	clusterFillers []api.ClusterFiller
	cpCidr         string
	podCidr        string
}

type SnowOpt func(*Snow)

func NewSnow(t *testing.T, opts ...SnowOpt) *Snow { _ = "STUB: not implemented"; return nil }

func (s *Snow) Name() string { _ = "STUB: not implemented"; return "" }

func (s *Snow) Setup() {
	_ = "STUB: not implemented"

	// UpdateKubeConfig customizes generated kubeconfig for the provider.
	return
}

func (s *Snow) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"

	// ClusterConfigUpdates satisfies the test framework Provider.
	return nil
}

func (s *Snow) ClusterConfigUpdates() []api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return nil
}

// CleanupResources  satisfies the test framework Provider.
func (s *Snow) CleanupResources(clusterName string) error { _ = "STUB: not implemented"; return nil }

// snow device doesn't support filter hitherto

func cleanupKeypairs(ec2Client *ec2.EC2, clusterName string) ([]*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isNotTerminatedAndHasTag(instance *ec2.Instance, tag string) bool {
	_ = "STUB: not implemented"
	return false
}

func newSession(ip string) (*session.Session, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Snow) WithProviderUpgrade(fillers ...api.SnowFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithBottlerocket128 returns a cluster config filler that sets the kubernetes version of the cluster to 1.28
// as well as the right devices and osFamily for all SnowMachineConfigs. It also sets any
// necessary machine config default required for BR, like the container volume size. If the env var is set, this will
// also set the AMI ID. Otherwise, it will leave it empty and let CAPAS select one.
func (s *Snow) WithBottlerocket128() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithBottlerocketStaticIP128 returns a cluster config filler that sets the kubernetes version of the cluster to 1.28
// as well as the right devices, osFamily and static ip config for all SnowMachineConfigs. Comparing to WithBottlerocket128,
// this method also adds a snow ip pool to support static ip configuration.
func (s *Snow) WithBottlerocketStaticIP128() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithUbuntu128 returns a cluster config filler that sets the kubernetes version of the cluster to 1.28
// as well as the right devices and osFamily for all SnowMachineConfigs. If the env var is set, this will
// also set the AMI ID. Otherwise, it will leave it empty and let CAPAS select one.
func (s *Snow) WithUbuntu128() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

func (s *Snow) withBottlerocketForKubeVersion(kubeVersion anywherev1.KubernetesVersion) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

func (s *Snow) withBottlerocketStaticIPForKubeVersion(kubeVersion anywherev1.KubernetesVersion) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithKubeVersionAndOS returns a cluster config filler that sets the cluster kube version and the correct AMI ID
// and devices for the Snow machine configs.
func (s *Snow) WithKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, _ *releasev1.EksARelease, _ ...string) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

func (s *Snow) withAMIIDFromEnvVar(envvar string) api.SnowFiller {
	_ = "STUB: not implemented"
	return *new(api.SnowFiller)
}

func (s *Snow) withIPPoolFromEnvVar(name string) api.SnowFiller {
	_ = "STUB: not implemented"
	return *new(api.SnowFiller)
}

// WithSnowUbuntu128 returns SnowOpt that sets the right devices and osFamily for all SnowMachineConfigs.
// If the env var is set, this will also set the AMI ID.
// Otherwise, it will leave it empty and let CAPAS select one.
func WithSnowUbuntu128() SnowOpt { _ = "STUB: not implemented"; return *new(SnowOpt) }

// WithSnowWorkerNodeGroup stores the necessary fillers to update/create the provided worker node group with its corresponding SnowMachineConfig
// and apply the given changes to that machine config.
func WithSnowWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup, fillers ...api.SnowMachineConfigFiller) SnowOpt {
	_ = "STUB: not implemented"
	return *new(SnowOpt)
}

// WithWorkerNodeGroup returns a filler that updates/creates the provided worker node group with its corresponding SnowMachineConfig
// and applies the given changes to that machine config.
func (s *Snow) WithWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup, fillers ...api.SnowMachineConfigFiller) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithNewWorkerNodeGroup returns a filler that updates/creates the provided worker node group with its corresponding SnowMachineConfig.
func (s *Snow) WithNewWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithNewSnowWorkerNodeGroup updates the test cluster Config with the fillers for an specific snow worker node group.
// It applies the fillers in WorkerNodeGroup to the named worker node group and the ones for the corresponding SnowMachineConfig.
func (s *Snow) WithNewSnowWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup, fillers ...api.SnowMachineConfigFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func snowMachineConfig(name string, fillers ...api.SnowMachineConfigFiller) api.SnowFiller {
	_ = "STUB: not implemented"
	return *new(api.SnowFiller)
}

func buildSnowWorkerNodeGroupClusterFiller(machineConfigName string, workerNodeGroup *WorkerNodeGroup) api.ClusterFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterFiller)
}

// ClusterStateValidations returns a list of provider specific validations.
func (s *Snow) ClusterStateValidations() []clusterf.StateValidation {
	_ = "STUB: not implemented"
	return nil
}
