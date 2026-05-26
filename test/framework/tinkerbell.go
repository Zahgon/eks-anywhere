package framework

import (
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

const (
	tinkerbellProviderName                              = "tinkerbell"
	tinkerbellBootstrapIPEnvVar                         = "T_TINKERBELL_BOOTSTRAP_IP"
	tinkerbellControlPlaneNetworkCidrEnvVar             = "T_TINKERBELL_CP_NETWORK_CIDR"
	tinkerbellImageUbuntu2204Kubernetes129EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_29"
	tinkerbellImageUbuntu2204Kubernetes129RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_29_RTOS"
	tinkerbellImageUbuntu2204Kubernetes130RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_30_RTOS"
	tinkerbellImageUbuntu2204Kubernetes131RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_31_RTOS"
	tinkerbellImageUbuntu2204Kubernetes132RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_32_RTOS"
	tinkerbellImageUbuntu2204Kubernetes133RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_33_RTOS"
	tinkerbellImageUbuntu2404Kubernetes133RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_33_RTOS"
	tinkerbellImageUbuntu2404Kubernetes129EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_29"
	tinkerbellImageUbuntu2404Kubernetes130EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_30"
	tinkerbellImageUbuntu2404Kubernetes131EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_31"
	tinkerbellImageUbuntu2404Kubernetes132EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_32"
	tinkerbellImageUbuntu2404Kubernetes133EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_33"
	tinkerbellImageUbuntu2204Kubernetes129GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_29_GENERIC"
	tinkerbellImageUbuntu2204Kubernetes130GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_30_GENERIC"
	tinkerbellImageUbuntu2204Kubernetes131GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_31_GENERIC"
	tinkerbellImageUbuntu2204Kubernetes132GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_32_GENERIC"
	tinkerbellImageUbuntu2204Kubernetes133GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_33_GENERIC"
	tinkerbellImageUbuntu2404Kubernetes133GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_33_GENERIC"
	tinkerbellImageUbuntu2204Kubernetes130EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_30"
	tinkerbellImageUbuntu2204Kubernetes131EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_31"
	tinkerbellImageUbuntu2204Kubernetes132EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_32"
	tinkerbellImageUbuntu2204Kubernetes133EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_33"
	tinkerbellImageRedHat129EnvVar                      = "T_TINKERBELL_IMAGE_REDHAT_1_29"
	tinkerbellImageRedHat130EnvVar                      = "T_TINKERBELL_IMAGE_REDHAT_1_30"
	tinkerbellImageRedHat131EnvVar                      = "T_TINKERBELL_IMAGE_REDHAT_1_31"
	tinkerbellImageRedHat9129EnvVar                     = "T_TINKERBELL_IMAGE_REDHAT_9_1_29"
	tinkerbellImageRedHat9130EnvVar                     = "T_TINKERBELL_IMAGE_REDHAT_9_1_30"
	tinkerbellImageRedHat9131EnvVar                     = "T_TINKERBELL_IMAGE_REDHAT_9_1_31"
	tinkerbellImageRedHat9132EnvVar                     = "T_TINKERBELL_IMAGE_REDHAT_9_1_32"
	tinkerbellImageRedHat9133EnvVar                     = "T_TINKERBELL_IMAGE_REDHAT_9_1_33"
	tinkerbellImageUbuntu2204Kubernetes134EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_34"
	tinkerbellImageUbuntu2204Kubernetes134RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_34_RTOS"
	tinkerbellImageUbuntu2404Kubernetes134RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_34_RTOS"
	tinkerbellImageUbuntu2204Kubernetes134GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_34_GENERIC"
	tinkerbellImageUbuntu2404Kubernetes134GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_34_GENERIC"
	tinkerbellImageUbuntu2404Kubernetes134EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_34"
	tinkerbellImageRedHat9134EnvVar                     = "T_TINKERBELL_IMAGE_REDHAT_9_1_34"
	tinkerbellImageUbuntu2204Kubernetes135EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_35"
	tinkerbellImageUbuntu2204Kubernetes135RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_35_RTOS"
	tinkerbellImageUbuntu2404Kubernetes135RTOSEnvVar    = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_35_RTOS"
	tinkerbellImageUbuntu2204Kubernetes135GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2204_1_35_GENERIC"
	tinkerbellImageUbuntu2404Kubernetes135GenericEnvVar = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_35_GENERIC"
	tinkerbellImageUbuntu2404Kubernetes135EnvVar        = "T_TINKERBELL_IMAGE_UBUNTU_2404_1_35"
	tinkerbellImageRedHat9135EnvVar                     = "T_TINKERBELL_IMAGE_REDHAT_9_1_35"
	tinkerbellInventoryCsvFilePathEnvVar                = "T_TINKERBELL_INVENTORY_CSV"
	tinkerbellSSHAuthorizedKey                          = "T_TINKERBELL_SSH_AUTHORIZED_KEY"
	tinkerbellCIEnvironmentEnvVar                       = "T_TINKERBELL_CI_ENVIRONMENT"
	controlPlaneIdentifier                              = "cp"
	workerIdentifier                                    = "worker"
	tinkerbellHookIsoURLEnvVar                          = "T_TINKERBELL_HOOK_ISO_URL"
)

var requiredTinkerbellEnvVars = []string{
	tinkerbellControlPlaneNetworkCidrEnvVar,
	tinkerbellImageUbuntu2204Kubernetes129EnvVar,
	tinkerbellImageUbuntu2204Kubernetes129RTOSEnvVar,
	tinkerbellImageUbuntu2204Kubernetes130RTOSEnvVar,
	tinkerbellImageUbuntu2204Kubernetes131RTOSEnvVar,
	tinkerbellImageUbuntu2204Kubernetes132RTOSEnvVar,
	tinkerbellImageUbuntu2204Kubernetes133RTOSEnvVar,
	tinkerbellImageUbuntu2404Kubernetes133RTOSEnvVar,
	tinkerbellImageUbuntu2204Kubernetes129GenericEnvVar,
	tinkerbellImageUbuntu2204Kubernetes130GenericEnvVar,
	tinkerbellImageUbuntu2204Kubernetes131GenericEnvVar,
	tinkerbellImageUbuntu2204Kubernetes132GenericEnvVar,
	tinkerbellImageUbuntu2204Kubernetes133GenericEnvVar,
	tinkerbellImageUbuntu2404Kubernetes133GenericEnvVar,
	tinkerbellImageUbuntu2204Kubernetes130EnvVar,
	tinkerbellImageUbuntu2204Kubernetes131EnvVar,
	tinkerbellImageUbuntu2204Kubernetes132EnvVar,
	tinkerbellImageUbuntu2204Kubernetes133EnvVar,
	tinkerbellImageUbuntu2404Kubernetes129EnvVar,
	tinkerbellImageUbuntu2404Kubernetes130EnvVar,
	tinkerbellImageUbuntu2404Kubernetes131EnvVar,
	tinkerbellImageUbuntu2404Kubernetes132EnvVar,
	tinkerbellImageUbuntu2404Kubernetes133EnvVar,
	tinkerbellImageRedHat129EnvVar,
	tinkerbellImageRedHat130EnvVar,
	tinkerbellImageRedHat131EnvVar,
	tinkerbellImageRedHat9129EnvVar,
	tinkerbellImageRedHat9130EnvVar,
	tinkerbellImageRedHat9131EnvVar,
	tinkerbellImageRedHat9132EnvVar,
	tinkerbellImageRedHat9133EnvVar,
	tinkerbellImageUbuntu2204Kubernetes134EnvVar,
	tinkerbellImageUbuntu2204Kubernetes134RTOSEnvVar,
	tinkerbellImageUbuntu2404Kubernetes134RTOSEnvVar,
	tinkerbellImageUbuntu2204Kubernetes134GenericEnvVar,
	tinkerbellImageUbuntu2404Kubernetes134GenericEnvVar,
	tinkerbellImageUbuntu2404Kubernetes134EnvVar,
	tinkerbellImageRedHat9134EnvVar,
	tinkerbellImageUbuntu2204Kubernetes135EnvVar,
	tinkerbellImageUbuntu2204Kubernetes135RTOSEnvVar,
	tinkerbellImageUbuntu2404Kubernetes135RTOSEnvVar,
	tinkerbellImageUbuntu2204Kubernetes135GenericEnvVar,
	tinkerbellImageUbuntu2404Kubernetes135GenericEnvVar,
	tinkerbellImageUbuntu2404Kubernetes135EnvVar,
	tinkerbellImageRedHat9135EnvVar,
	tinkerbellInventoryCsvFilePathEnvVar,
	tinkerbellSSHAuthorizedKey,
	tinkerbellHookIsoURLEnvVar,
}

func RequiredTinkerbellEnvVars() []string { _ = "STUB: not implemented"; return nil }

type TinkerbellOpt func(*Tinkerbell)

type Tinkerbell struct {
	t                    *testing.T
	fillers              []api.TinkerbellFiller
	clusterFillers       []api.ClusterFiller
	serverIP             string
	cidr                 string
	inventoryCsvFilePath string
}

// UpdateTinkerbellMachineSSHAuthorizedKey updates a tinkerbell machine configs SSHAuthorizedKey.
func UpdateTinkerbellMachineSSHAuthorizedKey() api.TinkerbellMachineFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellMachineFiller)
}

func NewTinkerbell(t *testing.T, opts ...TinkerbellOpt) *Tinkerbell {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tinkerbell) Name() string { _ = "STUB: not implemented"; return "" }

func (t *Tinkerbell) Setup() {
	_ = "STUB: not implemented"

	// UpdateKubeConfig customizes generated kubeconfig for the provider.
	return
}

func (t *Tinkerbell) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"

	// ClusterConfigUpdates satisfies the test framework Provider.
	return nil
}

func (t *Tinkerbell) ClusterConfigUpdates() []api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tinkerbell) WithProviderUpgrade(fillers ...api.TinkerbellFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// CleanupResources runs a clean up the Tinkerbell machines which simply powers them down.
func (t *Tinkerbell) CleanupResources(_ string) error { _ = "STUB: not implemented"; return nil }

// WithKubeVersionAndOS returns a cluster config filler that sets the cluster kube version and the right image for all
// tinkerbell machine configs.
func (t *Tinkerbell) WithKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, _ *releasev1.EksARelease, kernelVariant ...string) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithCPKubeVersionAndOS returns a cluster config filler that sets the cluster kube version and the right image for CP
// tinkerbell machine configs.
func (t *Tinkerbell) WithCPKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithWorkerKubeVersionAndOS returns a cluster config filler that sets the cluster kube version and the right image for all
// Worker tinkerbell machine configs.
func (t *Tinkerbell) WithWorkerKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// WithNewWorkerNodeGroup returns an api.ClusterFiller that adds a new workerNodeGroupConfiguration and
// a corresponding TinkerbellMachineConfig to the cluster config.
func (t *Tinkerbell) WithNewWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	// TODO: Implement for Tinkerbell provider
	return *new(api.ClusterConfigFiller)
}

// WithTinkerbellTemplateConfig returns a cluster config filler that sets a custom TinkerbellTemplateConfig.
func (t *Tinkerbell) WithTinkerbellTemplateConfig(templateConfig *anywherev1.TinkerbellTemplateConfig) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

func envVarForImage(os OS, kubeVersion anywherev1.KubernetesVersion, kernelVariant ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// withKubeVersionAndOS returns a cluster config filler that sets the cluster kube version and the right image for all
// tinkerbell machine configs.
func withKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, os OS, machineConfigType string, release *releasev1.EksARelease) TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithUbuntu129Tinkerbell tink test with ubuntu 1.29.
func WithUbuntu129Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithUbuntu130Tinkerbell tink test with ubuntu 1.30.
func WithUbuntu130Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithUbuntu131Tinkerbell tink test with ubuntu 1.31.
func WithUbuntu131Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithUbuntu132Tinkerbell tink test with ubuntu 1.32.
func WithUbuntu132Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithUbuntu133Tinkerbell tink test with ubuntu 1.33.
func WithUbuntu133Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithUbuntu134Tinkerbell tink test with ubuntu 1.34.
func WithUbuntu134Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithRedHat129Tinkerbell tink test with redhat 1.29.
func WithRedHat129Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithRedHat130Tinkerbell tink test with redhat 1.30.
func WithRedHat130Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithRedHat131Tinkerbell tink test with redhat 1.31.
func WithRedHat131Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithRedHat9129Tinkerbell tink test with redhat9 efi 1.29.
func WithRedHat9129Tinkerbell() TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithRedHat9130Tinkerbell tink test with redhat9 efi 1.30.
func WithRedHat9130Tinkerbell() TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithRedHat9131Tinkerbell tink test with redhat9 efi 1.31.
func WithRedHat9131Tinkerbell() TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithRedHat9132Tinkerbell tink test with redhat9 efi 1.32.
func WithRedHat9132Tinkerbell() TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithRedHat9133Tinkerbell tink test with redhat9 efi 1.33.
func WithRedHat9133Tinkerbell() TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithRedHat9134Tinkerbell tink test with redhat9 efi 1.34.
func WithRedHat9134Tinkerbell() TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithRedHat9135Tinkerbell tink test with redhat9 efi 1.35.
func WithRedHat9135Tinkerbell() TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

func WithBottleRocketTinkerbell() TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

func WithTinkerbellExternalEtcdTopology(count int) TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

func WithCustomTinkerbellMachineConfig(selector string) TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// ClusterStateValidations returns a list of provider specific validations.
func (t *Tinkerbell) ClusterStateValidations() []clusterf.StateValidation {
	_ = "STUB: not implemented"
	return nil
}

// WithOSImageURL Modify OS Image url.
func WithOSImageURL(url string) TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithHookImagesURLPath Modify Hook OS Image url.
func WithHookImagesURLPath(url string) TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// WithHookIsoBoot sets IsoBoot to true.
func WithHookIsoBoot() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// WithHookIsoURLPath helps in setting the HookOS ISO URL value.
func WithHookIsoURLPath(url string) TinkerbellOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellOpt)
}

// imageForKubeVersionAndOS sets osImageURL on the appropriate field in the Machine Config based on the machineConfigType string provided else sets it at Data Center config.
func imageForKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, operatingSystem OS, machineConfigType string, kernelVariant ...string) api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu129Image represents an Ubuntu raw image corresponding to Kubernetes 1.29.
func Ubuntu129Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu130Image represents an Ubuntu raw image corresponding to Kubernetes 1.30.
func Ubuntu130Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu131Image represents an Ubuntu raw image corresponding to Kubernetes 1.31.
func Ubuntu131Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu132Image represents an Ubuntu raw image corresponding to Kubernetes 1.32.
func Ubuntu132Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu133Image represents an Ubuntu raw image corresponding to Kubernetes 1.33.
func Ubuntu133Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu134Image represents an Ubuntu raw image corresponding to Kubernetes 1.34.
func Ubuntu134Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu129ImageForCP represents an Ubuntu raw image corresponding to Kubernetes 1.29 and is set for CP machine config.
func Ubuntu129ImageForCP() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu130ImageForCP represents an Ubuntu raw image corresponding to Kubernetes 1.30 and is set for CP machine config.
func Ubuntu130ImageForCP() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu131ImageForCP represents an Ubuntu raw image corresponding to Kubernetes 1.31 and is set for CP machine config.
func Ubuntu131ImageForCP() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu132ImageForCP represents an Ubuntu raw image corresponding to Kubernetes 1.32 and is set for CP machine config.
func Ubuntu132ImageForCP() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu133ImageForCP represents an Ubuntu raw image corresponding to Kubernetes 1.33 and is set for CP machine config.
func Ubuntu133ImageForCP() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu134ImageForCP represents an Ubuntu raw image corresponding to Kubernetes 1.34 and is set for CP machine config.
func Ubuntu134ImageForCP() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu129ImageForWorker represents an Ubuntu raw image corresponding to Kubernetes 1.29 and is set for worker machine config.
func Ubuntu129ImageForWorker() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu130ImageForWorker represents an Ubuntu raw image corresponding to Kubernetes 1.30 and is set for worker machine config.
func Ubuntu130ImageForWorker() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu131ImageForWorker represents an Ubuntu raw image corresponding to Kubernetes 1.31 and is set for worker machine config.
func Ubuntu131ImageForWorker() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu132ImageForWorker represents an Ubuntu raw image corresponding to Kubernetes 1.32 and is set for worker machine config.
func Ubuntu132ImageForWorker() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu133ImageForWorker represents an Ubuntu raw image corresponding to Kubernetes 1.33 and is set for worker machine config.
func Ubuntu133ImageForWorker() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu134ImageForWorker represents an Ubuntu raw image corresponding to Kubernetes 1.34 and is set for worker machine config.
func Ubuntu134ImageForWorker() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes129Image represents an Ubuntu 22.04 raw image corresponding to Kubernetes 1.29.
func Ubuntu2204Kubernetes129Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes130Image represents an Ubuntu 22.04 raw image corresponding to Kubernetes 1.30.
func Ubuntu2204Kubernetes130Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes131Image represents an Ubuntu 22.04 raw image corresponding to Kubernetes 1.31.
func Ubuntu2204Kubernetes131Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes132Image represents an Ubuntu 22.04 raw image corresponding to Kubernetes 1.32.
func Ubuntu2204Kubernetes132Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes133Image represents an Ubuntu 22.04 raw image corresponding to Kubernetes 1.33.
func Ubuntu2204Kubernetes133Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes134Image represents an Ubuntu 22.04 raw image corresponding to Kubernetes 1.34.
func Ubuntu2204Kubernetes134Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes135Image represents an Ubuntu 22.04 raw image corresponding to Kubernetes 1.35.
func Ubuntu2204Kubernetes135Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes129Image represents an Ubuntu 24.04 raw image corresponding to Kubernetes 1.29.
func Ubuntu2404Kubernetes129Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes130Image represents an Ubuntu 24.04 raw image corresponding to Kubernetes 1.30.
func Ubuntu2404Kubernetes130Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes131Image represents an Ubuntu 24.04 raw image corresponding to Kubernetes 1.31.
func Ubuntu2404Kubernetes131Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes132Image represents an Ubuntu 24.04 raw image corresponding to Kubernetes 1.32.
func Ubuntu2404Kubernetes132Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes133Image represents an Ubuntu 24.04 raw image corresponding to Kubernetes 1.33.
func Ubuntu2404Kubernetes133Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes134Image represents an Ubuntu 24.04 raw image corresponding to Kubernetes 1.34.
func Ubuntu2404Kubernetes134Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes129RTOSImage represents an Ubuntu 22.04 raw image with RTOS kernel corresponding to Kubernetes 1.29.
func Ubuntu2204Kubernetes129RTOSImage() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes133RTOSImage represents an Ubuntu 24.04 raw image with RTOS kernel corresponding to Kubernetes 1.33.
func Ubuntu2404Kubernetes133RTOSImage() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes134RTOSImage represents an Ubuntu 24.04 raw image with RTOS kernel corresponding to Kubernetes 1.34.
func Ubuntu2404Kubernetes134RTOSImage() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2204Kubernetes129GenericImage represents an Ubuntu 22.04 raw image with Generic kernel corresponding to Kubernetes 1.29.
func Ubuntu2204Kubernetes129GenericImage() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes133GenericImage represents an Ubuntu 24.04 raw image with Generic kernel corresponding to Kubernetes 1.33.
func Ubuntu2404Kubernetes133GenericImage() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes134GenericImage represents an Ubuntu 24.04 raw image with Generic kernel corresponding to Kubernetes 1.34.
func Ubuntu2404Kubernetes134GenericImage() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// RedHat9Kubernetes134Image represents a RedHat 9 raw image corresponding to Kubernetes 1.34.
func RedHat9Kubernetes134Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// WithUbuntu135Tinkerbell returns a TinkerbellOpt that adds API fillers to use a Ubuntu Tinkerbell template for k8s 1.35.
func WithUbuntu135Tinkerbell() TinkerbellOpt { _ = "STUB: not implemented"; return *new(TinkerbellOpt) }

// Ubuntu135Image represents a Ubuntu raw image corresponding to Kubernetes 1.35.
func Ubuntu135Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu135ImageForCP represents a Ubuntu raw image corresponding to Kubernetes 1.35 for control plane nodes.
func Ubuntu135ImageForCP() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu135ImageForWorker represents a Ubuntu raw image corresponding to Kubernetes 1.35 for worker nodes.
func Ubuntu135ImageForWorker() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes135Image represents a Ubuntu 24.04 raw image corresponding to Kubernetes 1.35.
func Ubuntu2404Kubernetes135Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes135RTOSImage represents a Ubuntu 24.04 RTOS raw image corresponding to Kubernetes 1.35.
func Ubuntu2404Kubernetes135RTOSImage() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// Ubuntu2404Kubernetes135GenericImage represents a Ubuntu 24.04 Generic raw image corresponding to Kubernetes 1.35.
func Ubuntu2404Kubernetes135GenericImage() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// RedHat9Kubernetes135Image represents a RedHat 9 raw image corresponding to Kubernetes 1.35.
func RedHat9Kubernetes135Image() api.TinkerbellFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellFiller)
}

// HookIsoURLOverride returns the Hook ISO URL from the environment variable.
func HookIsoURLOverride() string { _ = "STUB: not implemented"; return "" }

// WithTemplateRef returns a TinkerbellMachineFiller that adds a TemplateRef to a TinkerbellMachineConfig.
func WithTemplateRef(name, kind string) api.TinkerbellMachineFiller {
	_ = "STUB: not implemented"
	return *new(api.TinkerbellMachineFiller)
}
