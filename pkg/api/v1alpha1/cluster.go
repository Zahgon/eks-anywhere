package v1alpha1

import (
	"regexp"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/aws/eks-anywhere/pkg/constants"
)

// constants defined for cluster.go.
const (
	ClusterKind              = "Cluster"
	RegistryMirrorCAKey      = "EKSA_REGISTRY_MIRROR_CA"
	podSubnetNodeMaskMaxDiff = 16
)

var re = regexp.MustCompile(constants.DefaultCuratedPackagesRegistryRegex)

// +kubebuilder:object:generate=false
type ClusterGenerateOpt func(config *ClusterGenerate)

// Used for generating yaml for generate clusterconfig command.
func NewClusterGenerate(clusterName string, opts ...ClusterGenerateOpt) *ClusterGenerate {
	_ = "STUB: not implemented"
	return nil
}

func ControlPlaneConfigCount(count int) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func ExternalETCDConfigCount(count int) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func WorkerNodeConfigCount(count int) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func WorkerNodeConfigName(name string) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func WithClusterEndpoint() ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

// WithCPUpgradeRolloutStrategy allows add UpgradeRolloutStrategy option to cluster config under ControlPlaneConfiguration.
func WithCPUpgradeRolloutStrategy(maxSurge int, maxUnavailable int) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func WithDatacenterRef(ref ProviderRefAccessor) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func WithSharedMachineGroupRef(ref ProviderRefAccessor) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func WithCPMachineGroupRef(ref ProviderRefAccessor) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func WithWorkerMachineGroupRef(ref ProviderRefAccessor) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

// WithWorkerMachineUpgradeRolloutStrategy allows add UpgradeRolloutStrategy option to cluster config under WorkerNodeGroupConfiguration.
func WithWorkerMachineUpgradeRolloutStrategy(maxSurge int, maxUnavailable int) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

func WithEtcdMachineGroupRef(ref ProviderRefAccessor) ClusterGenerateOpt {
	_ = "STUB: not implemented"
	return *new(ClusterGenerateOpt)
}

var clusterConfigValidations = []func(*Cluster) error{
	validateClusterConfigName,
	validateControlPlaneEndpoint,
	validateExternalEtcdSupport,
	validateMachineGroupRefs,
	validateControlPlaneReplicas,
	validateWorkerNodeGroups,
	validateNetworking,
	validateGitOps,
	validateEtcdReplicas,
	validateIdentityProviderRefs,
	validateProxyConfig,
	validateMirrorConfig,
	validatePodIAMConfig,
	validateCPUpgradeRolloutStrategy,
	validateControlPlaneLabels,
	validatePackageControllerConfiguration,
	validateEksaVersion,
	validateControlPlaneCertSANs,
	validateControlPlaneAPIServerExtraArgs,
	validateControlPlaneAPIServerOIDCExtraArgs,
	validateControlPlaneKubeletConfiguration,
	validateWorkerNodeKubeletConfiguration,
	validateAuditPolicyContent,
}

// GetClusterConfig parses a Cluster object from a multiobject yaml file in disk
// and sets defaults if necessary.
func GetClusterConfig(fileName string) (*Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusterConfig parses a Cluster object from a multiobject yaml file in disk
// sets defaults if necessary and validates the Cluster.
func GetAndValidateClusterConfig(fileName string) (*Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusterDefaultKubernetesVersion returns the default kubernetes version for a Cluster.
func GetClusterDefaultKubernetesVersion() KubernetesVersion {
	_ = "STUB: not implemented"

	// ValidateClusterConfigContent validates a Cluster object without modifying it
	// Some of the validations are a bit heavy and need a network connection.
	return *new(KubernetesVersion)
}

func ValidateClusterConfigContent(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// ParseClusterConfig unmarshalls an API object implementing the KindAccessor interface
// from a multiobject yaml file in disk. It doesn't set defaults nor validates the object.
func ParseClusterConfig(fileName string, clusterConfig KindAccessor) error {
	_ = "STUB: not implemented"
	return nil
}

type kindObject struct {
	Kind string `json:"kind,omitempty"`
}

// ParseClusterConfigFromContent unmarshalls an API object implementing the KindAccessor interface
// from a multiobject yaml content. It doesn't set defaults nor validates the object.
func ParseClusterConfigFromContent(content []byte, clusterConfig KindAccessor) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cluster) PauseReconcile() { _ = "STUB: not implemented"; return }

// AllowDeleteWhilePaused adds the allow-delete-when-paused annotation to the cluster.
func (c *Cluster) AllowDeleteWhilePaused() { _ = "STUB: not implemented"; return }

// PreventDeleteWhilePaused removes the allow-delete-when-paused annotation to the cluster.
func (c *Cluster) PreventDeleteWhilePaused() { _ = "STUB: not implemented"; return }

func (c *Cluster) ClearPauseAnnotation() { _ = "STUB: not implemented"; return }

// AddManagedByCLIAnnotation adds the managed-by-cli annotation to the cluster.
func (c *Cluster) AddManagedByCLIAnnotation() { _ = "STUB: not implemented"; return }

// ClearManagedByCLIAnnotation removes the managed-by-cli annotation from the cluster.
func (c *Cluster) ClearManagedByCLIAnnotation() { _ = "STUB: not implemented"; return }

// AddTinkerbellIPAnnotation adds the managed-by-cli annotation to the cluster.
func (c *Cluster) AddTinkerbellIPAnnotation(tinkerbellIP string) { _ = "STUB: not implemented"; return }

// ClearTinkerbellIPAnnotation removes the managed-by-cli annotation from the cluster.
func (c *Cluster) ClearTinkerbellIPAnnotation() { _ = "STUB: not implemented"; return }

// HasTinkerbellIPAnnotation returns the tinkerbell IP value if the annotation exists.
func (c *Cluster) HasTinkerbellIPAnnotation() string { _ = "STUB: not implemented"; return "" }

// RegistryAuth returns whether registry requires authentication or not.
func (c *Cluster) RegistryAuth() bool { _ = "STUB: not implemented"; return false }

func (c *Cluster) ProxyConfiguration() map[string]string { _ = "STUB: not implemented"; return nil }

func (c *Cluster) IsReconcilePaused() bool { _ = "STUB: not implemented"; return false }

func ValidateClusterName(clusterName string) error {
	_ = "STUB: not implemented"
	// this regex will not work for AWS provider as CFN has restrictions with UPPERCASE chars;
	// if you are using AWS provider please use only lowercase chars
	return nil
}

func ValidateClusterNameLength(clusterName string) error {
	_ = "STUB: not implemented"
	// docker container hostname can have a maximum length of 64 characters. we append "-eks-a-cluster"
	// to get the KinD cluster's name and on top of this, KinD also adds a "-control-plane suffix" to
	// the cluster name to arrive at the name for the control plane node (container), which makes the
	// control plane node name 64 characters in length.
	return nil
}

// NormalizeKubernetesVersion searches the YAML content for the specific string
// `kubernetesVersion: <>` and adds double quotes around the version so that it
// is always interpreted as a string instead of a float64 value.
// Ref: https://github.com/aws/eks-anywhere/issues/9184
func NormalizeKubernetesVersion(yamlContent string) string { _ = "STUB: not implemented"; return "" }

func validateClusterConfigName(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }

func validateExternalEtcdSupport(cluster *Cluster) error { _ = "STUB: not implemented"; return nil }

func validateMachineGroupRefs(cluster *Cluster) error { _ = "STUB: not implemented"; return nil }

func validateControlPlaneReplicas(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// For unstacked/external etcd, controlplane replicas can be any number including even numbers.

func validateControlPlaneLabels(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateControlPlaneEndpoint(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

var domainNameRegex = regexp.MustCompile(`(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`)

func validateControlPlaneCertSANs(cfg *Cluster) error { _ = "STUB: not implemented"; return nil }

func validateControlPlaneAPIServerExtraArgs(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateControlPlaneAPIServerOIDCExtraArgs(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateControlPlaneKubeletConfiguration(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateWorkerNodeKubeletConfiguration(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAuditPolicyContent(c *Cluster) error { _ = "STUB: not implemented"; return nil }

func validateKubeletConfiguration(kubeletConfig *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func validateWorkerNodeGroups(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }

// This block should never fire. If it does, it means we have a bug in how we set our defaults.
// When Count == nil it should be set to 1 by SetDefaults method prior to reaching validation.

func validateAutoscalingConfig(w *WorkerNodeGroupConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNodeLabels(labels map[string]string, fldPath *field.Path) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEtcdReplicas(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }

// only log warning about recommended etcd cluster size for providers other than docker

func validateNetworking(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }

// the pod subnet mask needs to allow one or multiple node-masks
// i.e. if it has a /24 the node mask must be between 24 and 32 for ipv4
// the below validations are run by kubeadm and we are bubbling those up here for better customer experience

// PodSubnetNodeMaskMaxDiff is limited to 16 due to an issue with uncompressed IP bitmap in core
// The node subnet mask size must be no more than the pod subnet mask size + 16

func validateCNIPlugin(network ClusterNetwork) error { _ = "STUB: not implemented"; return nil }

func validateCNIConfig(cniConfig *CNIConfig) error { _ = "STUB: not implemented"; return nil }

func validateCiliumConfig(cilium *CiliumConfig) error { _ = "STUB: not implemented"; return nil }

func validateProxyConfig(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }

func validateProxyData(proxy string) error { _ = "STUB: not implemented"; return nil }

func validateMirrorConfig(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }

// More than one mirror for curated package would introduce ambiguity in the package controller

func validateIdentityProviderRefs(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGitOps(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }

func validatePodIAMConfig(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }

func validateCPUpgradeRolloutStrategy(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMDUpgradeRolloutStrategy(w *WorkerNodeGroupConfiguration, datacenterRefKind string) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePackageControllerConfiguration(clusterConfig *Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEksaVersion(clusterConfig *Cluster) error { _ = "STUB: not implemented"; return nil }
