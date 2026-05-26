package nutanix

import (
	"context"
	"net/http"

	"github.com/nutanix-cloud-native/prism-go-client/environment/credentials"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/crypto"
)

const (
	minNutanixCPUSockets   = 1
	minNutanixCPUPerSocket = 1
	minNutanixMemoryMiB    = 2048
	minNutanixDiskGiB      = 20
)

// IPValidator is an interface that defines methods to validate the control plane IP.
type IPValidator interface {
	ValidateControlPlaneIPUniqueness(cluster *anywherev1.Cluster) error
}

// Validator is a client to validate nutanix resources.
type Validator struct {
	httpClient    *http.Client
	certValidator crypto.TlsValidator
	clientCache   *ClientCache
}

// NewValidator returns a new validator client.
func NewValidator(clientCache *ClientCache, certValidator crypto.TlsValidator, httpClient *http.Client) *Validator {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateControlPlaneIP(ip string) error {
	_ = "STUB: not implemented"
	// check if controlPlaneEndpointIp is valid
	return nil
}

// ValidateClusterSpec validates the cluster spec.
func (v *Validator) ValidateClusterSpec(ctx context.Context, spec *cluster.Spec, creds credentials.BasicAuthCredential) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) checkImageNameMatchesKubernetesVersion(ctx context.Context, spec *cluster.Spec, client Client) error {
	_ = "STUB: not implemented"
	return nil
}

// validate template field name contains cluster kubernetes version for the control plane machine.

// validate template field name contains cluster kubernetes version for the external etcd machine.

// validate template field name contains cluster kubernetes version for the control plane machine.

// ValidateDatacenterConfig validates the datacenter config.
func (v *Validator) ValidateDatacenterConfig(ctx context.Context, client Client, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateIPRangeForCcmExcludeNodeIPs(ipRange string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateCcmExcludeNodeIPs(ccmExcludeNodeIPs []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateFailureDomains(ctx context.Context, client Client, spec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateWorkerMachineGroup(workerMachineGroups map[string]anywherev1.WorkerNodeGroupConfiguration, workerMachineGroupName string, fdCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateCredentialRef(config *anywherev1.NutanixDatacenterConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateEndpointAndPort(dcConf anywherev1.NutanixDatacenterConfigSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateCredentials(ctx context.Context, client Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateTrustBundleConfig(dcConf anywherev1.NutanixDatacenterConfigSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateMachineSpecs(machineSpec anywherev1.NutanixMachineConfigSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateMachineConfig validates the Prism Element cluster, subnet, and image for the machine.
func (v *Validator) ValidateMachineConfig(ctx context.Context, client Client, cluster *anywherev1.Cluster, config *anywherev1.NutanixMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateGPUInMachineConfig(cluster *anywherev1.Cluster, config *anywherev1.NutanixMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateClusterConfig(ctx context.Context, client Client, identifier anywherev1.NutanixResourceIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateImageConfig(ctx context.Context, client Client, identifier anywherev1.NutanixResourceIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateTemplateMatchesKubernetesVersion(ctx context.Context, identifier anywherev1.NutanixResourceIdentifier, client Client, kubernetesVersionName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Replace 1.23, 1-23, 1_23 to 123 in the template name string.

// Replace 1-23 to 123 in the kubernetesversion string.

// This will return an error if the template name does not contain specified kubernetes version.
// For ex if the kubernetes version is 1.23,
// the template name should include 1.23 or 1-23, 1_23 or 123 i.e. kubernetes-1-23-eks in the string.

func (v *Validator) validateSubnetConfig(ctx context.Context, client Client, cluster, subnet anywherev1.NutanixResourceIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateProjectConfig(ctx context.Context, client Client, identifier anywherev1.NutanixResourceIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateAdditionalCategories(ctx context.Context, client Client, categories []anywherev1.NutanixCategoryIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateGPUConfig(gpu anywherev1.NutanixGPUIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

func getRequestedGPUsForAllMachines(machineCount int, requestedGpus []anywherev1.NutanixGPUIdentifier) []anywherev1.NutanixGPUIdentifier {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) tryAssignGPUsToMachineConfig(machineCount int, requestedGpus []anywherev1.NutanixGPUIdentifier, clusterGpuList []v3.GPU, cluster anywherev1.NutanixResourceIdentifier) ([]v3.GPU, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) isGPURequested(configs map[string]*anywherev1.NutanixMachineConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *Validator) initAvailableGPUsMap(hosts []*v3.HostResponse) map[string][]v3.GPU {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) getAvailableGPUs(hosts []*v3.HostResponse) (map[string][]v3.GPU, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *Validator) tryAssignGPUsToAllMachineConfigs(ctx context.Context, v3Client Client, cluster *cluster.Spec, availableGpu map[string][]v3.GPU) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) getMachineCountForAllMachineConfigs(clusterSpec *cluster.Spec) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) getGPUModeMapping(hosts []*v3.HostResponse) (map[int64]string, map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (v *Validator) validateGPUModeNotMixed(hosts []*v3.HostResponse, cluster *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func createGetGpuModeFunc(gpuDeviceIDToMode map[int64]string, gpuNameToMode map[string]string) func(gpu anywherev1.NutanixGPUIdentifier) string {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateFreeGPU(ctx context.Context, v3Client Client, cluster *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) validateUpgradeRolloutStrategy(clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func checkMachineConfigIsForWorker(config *anywherev1.NutanixMachineConfig, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// findSubnetUUIDByName retrieves the subnet uuid by the given subnet name.
func findSubnetUUIDByName(ctx context.Context, v3Client Client, clusterUUID, subnetName string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getWorkerMachineGroups retrieves the worker machine group names from the cluster spec.
func getWorkerMachineGroups(spec *cluster.Spec) map[string]anywherev1.WorkerNodeGroupConfiguration {
	_ = "STUB: not implemented"
	return nil
}

// getClusterUUID retrieves the cluster uuid by the given cluster identifier.
func getClusterUUID(ctx context.Context, v3Client Client, cluster anywherev1.NutanixResourceIdentifier) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// findClusterUUIDByName retrieves the cluster uuid by the given cluster name.
func findClusterUUIDByName(ctx context.Context, v3Client Client, clusterName string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prism Central is also internally a cluster, but we filter that out here as we only care about prism element clusters

// findImageUUIDByName retrieves the image uuid by the given image name.
func findImageUUIDByName(ctx context.Context, v3Client Client, imageName string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findProjectUUIDByName retrieves the project uuid by the given image name.
func findProjectUUIDByName(ctx context.Context, v3Client Client, projectName string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isRequestedGPUAssignable(gpu v3.GPU, requestedGpu anywherev1.NutanixGPUIdentifier) bool {
	_ = "STUB: not implemented"
	return false
}

func errorGPUNotFound(gpu anywherev1.NutanixGPUIdentifier, cluster anywherev1.NutanixResourceIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}
