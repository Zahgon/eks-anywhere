package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/intstr"
	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta1"

	"github.com/aws/eks-anywhere/pkg/semver"
)

const (
	// PausedAnnotation is an annotation that can be applied to EKS-A cluster
	// object to prevent a controller from processing a resource.
	pausedAnnotation = "anywhere.eks.amazonaws.com/paused"

	// ManagedByCLIAnnotation can be applied to an EKS-A Cluster to signal when the CLI is currently
	// performing an operation so the controller should not take any action. When marked for deletion,
	// the controller will remove the finalizer and let the cluster be deleted.
	ManagedByCLIAnnotation = "anywhere.eks.amazonaws.com/managed-by-cli"

	// tinkerbellIPAnnotation can be applied to an EKS-A Cluster to convey the tinkerbell bootstrap ip to the
	// EKSA controller. When marked for deletion, the controller will remove the IP annotation.
	tinkerbellIPAnnotation = "anywhere.eks.amazonaws.com/tinkerbell-bootstrap-ip"

	// ControlPlaneAnnotation is an annotation that can be applied to EKS-A machineconfig
	// object to prevent a controller from making changes to that resource.
	controlPlaneAnnotation = "anywhere.eks.amazonaws.com/control-plane"

	clusterResourceType = "clusters.anywhere.eks.amazonaws.com"

	// etcdAnnotation can be applied to EKS-A machineconfig CR for etcd, to prevent controller from making changes to it.
	etcdAnnotation = "anywhere.eks.amazonaws.com/etcd"

	// skipIPCheckAnnotation skips the availability control plane IP validation during cluster creation. Use only if your network configuration is conflicting with the default port scan.
	skipIPCheckAnnotation = "anywhere.eks.amazonaws.com/skip-ip-check"

	// skipEksaVersionSkewCheck is an annotation that bypasses the EKS-A version compatibility check during upgrades.
	// WARNING: Use this with caution as skipping minor versions could break component compatibility and cause cluster instability or failures.
	skipEksaVersionSkewCheck = "anywhere.eks.amazonaws.com/skip-eksa-version-skew-check"

	// managementAnnotation points to the name of a management cluster
	// cluster object.
	managementAnnotation = "anywhere.eks.amazonaws.com/managed-by"

	// managementComponentsVersionAnnotation is an annotation applied to an EKS-A management cluster pointing to the current version of the management components.
	// The value for this annotation is expected to correspond to an EKSARelease object version, following semantic version convention: e.g. v0.18.3
	// This is an internal EKS-A managed annotation, not meant to be updated manually.
	managementComponentsVersionAnnotation = "anywhere.eks.amazonaws.com/management-components-version"

	// defaultEksaNamespace is the default namespace for EKS-A resources when not specified.
	defaultEksaNamespace = "default"

	// ControlEndpointDefaultPort defaults cluster control plane endpoint port if not specified.
	ControlEndpointDefaultPort = "6443"

	// AllowDeleteWhenPausedAnnotation is an annotation applied to an EKS-A cluster that allows the deletion of the cluster
	// when paused.
	AllowDeleteWhenPausedAnnotation = "anywhere.eks.amazonaws.com/allow-delete-when-paused"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterSpec defines the desired state of Cluster.
type ClusterSpec struct {
	KubernetesVersion             KubernetesVersion              `json:"kubernetesVersion,omitempty"`
	ControlPlaneConfiguration     ControlPlaneConfiguration      `json:"controlPlaneConfiguration,omitempty"`
	WorkerNodeGroupConfigurations []WorkerNodeGroupConfiguration `json:"workerNodeGroupConfigurations,omitempty"`
	DatacenterRef                 Ref                            `json:"datacenterRef,omitempty"`
	IdentityProviderRefs          []Ref                          `json:"identityProviderRefs,omitempty"`
	GitOpsRef                     *Ref                           `json:"gitOpsRef,omitempty"`
	ClusterNetwork                ClusterNetwork                 `json:"clusterNetwork,omitempty"`
	// +kubebuilder:validation:Optional
	ExternalEtcdConfiguration   *ExternalEtcdConfiguration   `json:"externalEtcdConfiguration,omitempty"`
	ProxyConfiguration          *ProxyConfiguration          `json:"proxyConfiguration,omitempty"`
	RegistryMirrorConfiguration *RegistryMirrorConfiguration `json:"registryMirrorConfiguration,omitempty"`
	ManagementCluster           ManagementCluster            `json:"managementCluster,omitempty"`
	PodIAMConfig                *PodIAMConfig                `json:"podIamConfig,omitempty"`
	Packages                    *PackageConfiguration        `json:"packages,omitempty"`
	// BundlesRef contains a reference to the Bundles containing the desired dependencies for the cluster.
	// DEPRECATED: Use EksaVersion instead.
	BundlesRef         *BundlesRef         `json:"bundlesRef,omitempty"`
	EksaVersion        *EksaVersion        `json:"eksaVersion,omitempty"`
	MachineHealthCheck *MachineHealthCheck `json:"machineHealthCheck,omitempty"`
	EtcdEncryption     *[]EtcdEncryption   `json:"etcdEncryption,omitempty"`
	LicenseToken       string              `json:"licenseToken,omitempty"`
}

// EksaVersion is the semver identifying the release of eks-a used to populate the cluster components.
type EksaVersion string

const (
	// DevBuildVersion is the version string for the dev build of EKS-A.
	DevBuildVersion = "v0.19.0-dev+latest"

	// MinEksAVersionWithEtcdURL is the version from which the etcd url will be set
	// for etcdadm to pull the etcd tarball if that binary isnt cached.
	MinEksAVersionWithEtcdURL = "v0.19.0"
)

// Equal checks if two EksaVersions are equal.
func (n *EksaVersion) Equal(o *EksaVersion) bool { _ = "STUB: not implemented"; return false }

// HasAWSIamConfig checks if AWSIamConfig is configured for the cluster.
func (c *Cluster) HasAWSIamConfig() bool { _ = "STUB: not implemented"; return false }

// IsPackagesEnabled checks if the user has opted out of curated packages
// installation.
func (c *Cluster) IsPackagesEnabled() bool { _ = "STUB: not implemented"; return false }

func (n *Cluster) Equal(o *Cluster) bool { _ = "STUB: not implemented"; return false }

func (n *Cluster) Validate() error { _ = "STUB: not implemented"; return nil }

func (n *Cluster) SetDefaults() {
	_ = "STUB: not implemented"
	// TODO: move any defaults that can return error out of this package
	// All the defaults here should be context unaware
	return
}

type ProxyConfiguration struct {
	HttpProxy  string   `json:"httpProxy,omitempty"`
	HttpsProxy string   `json:"httpsProxy,omitempty"`
	NoProxy    []string `json:"noProxy,omitempty"`
}

func (n *ProxyConfiguration) Equal(o *ProxyConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// RegistryMirrorConfiguration defines the settings for image registry mirror.
type RegistryMirrorConfiguration struct {
	// Endpoint defines the registry mirror endpoint to use for pulling images
	Endpoint string `json:"endpoint,omitempty"`

	// Port defines the port exposed for registry mirror endpoint
	Port string `json:"port,omitempty"`

	// OCINamespaces defines the mapping from an upstream registry to a local namespace where upstream
	// artifacts are placed into
	OCINamespaces []OCINamespace `json:"ociNamespaces,omitempty"`

	// CACertContent defines the contents registry mirror CA certificate
	CACertContent string `json:"caCertContent,omitempty"`

	// Authenticate defines if registry requires authentication
	Authenticate bool `json:"authenticate,omitempty"`

	// InsecureSkipVerify skips the registry certificate verification.
	// Only use this solution for isolated testing or in a tightly controlled, air-gapped environment.
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

// OCINamespace represents an entity in a local reigstry to group related images.
type OCINamespace struct {
	// Registry refers to the name of the upstream registry
	Registry string `json:"registry"`
	// Namespace refers to the name of a namespace in the local registry
	Namespace string `json:"namespace"`
}

func (n *RegistryMirrorConfiguration) Equal(o *RegistryMirrorConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// OCINamespacesSliceEqual is used to check equality of the OCINamespaces fields of two RegistryMirrorConfiguration.
func OCINamespacesSliceEqual(a, b []OCINamespace) bool { _ = "STUB: not implemented"; return false }

func generateOCINamespaceKey(n OCINamespace) (key string) { _ = "STUB: not implemented"; return "" }

type ControlPlaneConfiguration struct {
	// Count defines the number of desired control plane nodes. Defaults to 1.
	Count int `json:"count,omitempty"`
	// Endpoint defines the host ip and port to use for the control plane.
	Endpoint *Endpoint `json:"endpoint,omitempty"`
	// MachineGroupRef defines the machine group configuration for the control plane.
	MachineGroupRef *Ref `json:"machineGroupRef,omitempty"`
	// Taints define the set of taints to be applied on control plane nodes
	Taints []corev1.Taint `json:"taints,omitempty"`
	// Labels define the labels to assign to the node
	Labels map[string]string `json:"labels,omitempty"`
	// UpgradeRolloutStrategy determines the rollout strategy to use for rolling upgrades
	// and related parameters/knobs
	UpgradeRolloutStrategy *ControlPlaneUpgradeRolloutStrategy `json:"upgradeRolloutStrategy,omitempty"`
	// SkipLoadBalancerDeployment skip deploying control plane load balancer.
	// Make sure your infrastructure can handle control plane load balancing when you set this field to true.
	SkipLoadBalancerDeployment bool `json:"skipLoadBalancerDeployment,omitempty"`
	// CertSANs is a slice of domain names or IPs to be added as Subject Name Alternatives of the
	// Kube API Servers Certificate.
	CertSANs []string `json:"certSans,omitempty"`
	// MachineHealthCheck is a control-plane level override for the timeouts and maxUnhealthy specified in the top-level MHC configuration. If not configured, the defaults in the top-level MHC configuration are used.
	MachineHealthCheck *MachineHealthCheck `json:"machineHealthCheck,omitempty"`
	// APIServerExtraArgs defines the flags to configure for the API server.
	APIServerExtraArgs map[string]string `json:"apiServerExtraArgs,omitempty"`
	// KubeletConfiguration is a struct that exposes the Kubelet settings for the user to set on control plane nodes.
	// +kubebuilder:pruning:PreserveUnknownFields
	KubeletConfiguration *unstructured.Unstructured `json:"kubeletConfiguration,omitempty"`
	// AuditPolicyContent defines the audit policy configuration as a string.
	// If not specified, the default audit policy will be used.
	// +optional
	AuditPolicyContent string `json:"auditPolicyContent,omitempty"`
	// SkipAdmissionForSystemResources skips admission plugin checks for system-level Kubernetes resources
	// When enabled, operations on system resources (such as kube-system ns resources Pods,
	// RBAC, API service registrations, flow control, etc. and system user operations) will bypass admission plugins
	// will bypass admission plugins to prevent potential deadlocks or failures for cluster operations.
	// +optional
	SkipAdmissionForSystemResources *bool `json:"skipAdmissionForSystemResources,omitempty"`
}

// MachineHealthCheck allows to configure timeouts for machine health checks. Machine Health Checks are responsible for remediating unhealthy Machines.
// Configuring these values will decide how long to wait to remediate unhealthy machine or determine health of nodes' machines.
type MachineHealthCheck struct {
	// NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to "10m0s" (10 minutes) for all providers. For Tinkerbell provider the default is "20m0s".
	NodeStartupTimeout *metav1.Duration `json:"nodeStartupTimeout,omitempty"`
	// UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to "5m0s" (5 minutes).
	UnhealthyMachineTimeout *metav1.Duration `json:"unhealthyMachineTimeout,omitempty"`
	// MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to "100%" for controlplane machines and "40%" for worker machines.
	MaxUnhealthy *intstr.IntOrString `json:"maxUnhealthy,omitempty"`
}

func TaintsSliceEqual(s1, s2 []corev1.Taint) bool { _ = "STUB: not implemented"; return false }

// MapEqual compares two maps to check whether or not they are equal.
func MapEqual(s1, s2 map[string]string) bool { _ = "STUB: not implemented"; return false }

func (n *ControlPlaneConfiguration) Equal(o *ControlPlaneConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

type Endpoint struct {
	// Host defines the ip that you want to use to connect to the control plane
	Host string `json:"host"`
}

// Equal compares if expected endpoint and existing endpoint are equal for non CloudStack clusters.
func (n *Endpoint) Equal(o *Endpoint, kind string) bool { _ = "STUB: not implemented"; return false }

// CloudStackEqual makes CloudStack cluster upgrade to new release backward compatible by striping CloudStack cluster existing endpoint default port
// and comparing if expected endpoint and existing endpoint are equal.
// Cloudstack CLI used to add default port to cluster object.
// Now cluster object stays the same with customer input and port is defaulted only in CAPI template.
func (n *Endpoint) CloudStackEqual(o *Endpoint) bool { _ = "STUB: not implemented"; return false }

// GetControlPlaneHostPort retrieves the ControlPlaneConfiguration host and port split defined in the cluster.Spec.
func GetControlPlaneHostPort(pHost string, defaultPort string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type WorkerNodeGroupConfiguration struct {
	// Name refers to the name of the worker node group
	Name string `json:"name,omitempty"`
	// Count defines the number of desired worker nodes. Defaults to 1.
	Count *int `json:"count,omitempty"`
	// AutoScalingConfiguration defines the auto scaling configuration
	AutoScalingConfiguration *AutoScalingConfiguration `json:"autoscalingConfiguration,omitempty"`
	// MachineGroupRef defines the machine group configuration for the worker nodes.
	MachineGroupRef *Ref `json:"machineGroupRef,omitempty"`
	// Taints define the set of taints to be applied on worker nodes
	Taints []corev1.Taint `json:"taints,omitempty"`
	// Labels define the labels to assign to the node
	Labels map[string]string `json:"labels,omitempty"`
	// UpgradeRolloutStrategy determines the rollout strategy to use for rolling upgrades
	// and related parameters/knobs
	UpgradeRolloutStrategy *WorkerNodesUpgradeRolloutStrategy `json:"upgradeRolloutStrategy,omitempty"`
	// KubernetesVersion defines the version for worker nodes. If not set, the top level spec kubernetesVersion will be used.
	KubernetesVersion *KubernetesVersion `json:"kubernetesVersion,omitempty"`
	// MachineHealthCheck is a worker node level override for the timeouts and maxUnhealthy specified in the top-level MHC configuration. If not configured, the defaults in the top-level MHC configuration are used.
	MachineHealthCheck *MachineHealthCheck `json:"machineHealthCheck,omitempty"`
	// KubeletConfiguration is a struct that exposes the Kubelet settings for the user to set on worker nodes.
	// +kubebuilder:pruning:PreserveUnknownFields
	KubeletConfiguration *unstructured.Unstructured `json:"kubeletConfiguration,omitempty"`
	// FailureDomains is the optional list of failure domains to distribute worker nodes across the infrastructure.
	FailureDomains []string `json:"failureDomains,omitempty"`
}

// Equal compares two WorkerNodeGroupConfigurations.
func (w WorkerNodeGroupConfiguration) Equal(other WorkerNodeGroupConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// Equal compares two KubernetesVersions.
func (k *KubernetesVersion) Equal(other *KubernetesVersion) bool {
	_ = "STUB: not implemented"
	return false
}

func intPtrEqual(a, b *int) bool { _ = "STUB: not implemented"; return false }

func WorkerNodeGroupConfigurationsSliceEqual(a, b []WorkerNodeGroupConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// WorkerNodeGroupConfigurationKubeVersionUnchanged checks if a worker node group's k8s version has not changed. The ClusterVersions are the top level kubernetes version of a cluster.
func WorkerNodeGroupConfigurationKubeVersionUnchanged(o, n *WorkerNodeGroupConfiguration, oldCluster, newCluster *Cluster) bool {
	_ = "STUB: not implemented"
	return false
}

type ClusterNetwork struct {
	// Comma-separated list of CIDR blocks to use for pod and service subnets.
	// Defaults to 192.168.0.0/16 for pod subnet.
	Pods     Pods     `json:"pods,omitempty"`
	Services Services `json:"services,omitempty"`
	// Deprecated. Use CNIConfig
	CNI CNI `json:"cni,omitempty"`
	// CNIConfig specifies the CNI plugin to be installed in the cluster
	CNIConfig *CNIConfig `json:"cniConfig,omitempty"`
	DNS       DNS        `json:"dns,omitempty"`
	Nodes     *Nodes     `json:"nodes,omitempty"`
}

func (n *ClusterNetwork) Equal(o *ClusterNetwork) bool { _ = "STUB: not implemented"; return false }

func getCNIConfig(cn *ClusterNetwork) *CNIConfig {
	_ = "STUB: not implemented"
	/* Only needed since we're introducing CNIConfig to replace the deprecated CNI field. This way we can compare the individual fields
	   for the CNI plugin configuration*/return nil
}

// This is for upgrading from release-0.7, to ensure that all oCNIConfig fields, such as policyEnforcementMode have the default values

func (n *Pods) Equal(o *Pods) bool { _ = "STUB: not implemented"; return false }

func (n *Services) Equal(o *Services) bool { _ = "STUB: not implemented"; return false }

func (n *DNS) Equal(o *DNS) bool { _ = "STUB: not implemented"; return false }

func (n *CNIConfig) Equal(o *CNIConfig) bool { _ = "STUB: not implemented"; return false }

func (n *CiliumConfig) Equal(o *CiliumConfig) bool { _ = "STUB: not implemented"; return false }

// Compare CNIExclusive field

// We consider nil to be false in equality checks. Here we're checking if o is false then
// n must be false and vice-versa. If neither of these are true, then both o and n must be
// true so we don't need an explicit check.

// Compare HelmValues field

func (n *KindnetdConfig) Equal(o *KindnetdConfig) bool { _ = "STUB: not implemented"; return false }

func UsersSliceEqual(a, b []UserConfiguration) bool { _ = "STUB: not implemented"; return false }

func CNIPluginSame(n ClusterNetwork, o ClusterNetwork) bool {
	_ = "STUB: not implemented"

	/*This shouldn't be required since we set CNIConfig and unset CNI as part of cluster_defaults. However, while upgrading an existing cluster, the eks-a controller
	does not set any defaults (no mutating webhook), so it gets stuck in an error loop. Adding these checks to avoid that. We can remove it when removing the CNI field
	in a later release*/return false
}

func SliceEqual(a, b []string) bool { _ = "STUB: not implemented"; return false }

func RefSliceEqual(a, b []Ref) bool { _ = "STUB: not implemented"; return false }

type Pods struct {
	CidrBlocks []string `json:"cidrBlocks,omitempty"`
}

type Services struct {
	CidrBlocks []string `json:"cidrBlocks,omitempty"`
}

type DNS struct {
	// ResolvConf refers to the DNS resolver configuration
	ResolvConf *ResolvConf `json:"resolvConf,omitempty"`
}

type ResolvConf struct {
	// Path defines the path to the file that contains the DNS resolver configuration
	Path string `json:"path,omitempty"`
}

type Nodes struct {
	// CIDRMaskSize defines the mask size for node cidr in the cluster, default for ipv4 is 24. This is an optional field
	CIDRMaskSize *int `json:"cidrMaskSize,omitempty"`
}

// Equal compares two Nodes definitions and return true if the are equivalent.
func (n *Nodes) Equal(o *Nodes) bool { _ = "STUB: not implemented"; return false }

func (n *ResolvConf) Equal(o *ResolvConf) bool { _ = "STUB: not implemented"; return false }

type KubernetesVersion string

const (
	Kube118 KubernetesVersion = "1.18"
	Kube119 KubernetesVersion = "1.19"
	Kube120 KubernetesVersion = "1.20"
	Kube121 KubernetesVersion = "1.21"
	Kube122 KubernetesVersion = "1.22"
	Kube123 KubernetesVersion = "1.23"
	Kube124 KubernetesVersion = "1.24"
	Kube125 KubernetesVersion = "1.25"
	Kube126 KubernetesVersion = "1.26"
	Kube127 KubernetesVersion = "1.27"
	Kube128 KubernetesVersion = "1.28"
	Kube129 KubernetesVersion = "1.29"
	Kube130 KubernetesVersion = "1.30"
	Kube131 KubernetesVersion = "1.31"
	Kube132 KubernetesVersion = "1.32"
	Kube133 KubernetesVersion = "1.33"
	Kube134 KubernetesVersion = "1.34"
	Kube135 KubernetesVersion = "1.35"
)

// KubeVersionToSemver converts kube version to semver for comparisons.
func KubeVersionToSemver(kubeVersion KubernetesVersion) (*semver.Version, error) {
	_ = "STUB: not implemented"
	// appending the ".0" as the patch version to have a valid semver string and use those semvers for comparison
	return nil, nil
}

type CNI string

type CiliumPolicyEnforcementMode string

type CiliumRoutingMode string

type CNIConfig struct {
	Cilium   *CiliumConfig   `json:"cilium,omitempty"`
	Kindnetd *KindnetdConfig `json:"kindnetd,omitempty"`
}

// IsManaged indicates if EKS-A is responsible for the CNI installation.
func (n *CNIConfig) IsManaged() bool { _ = "STUB: not implemented"; return false }

// CiliumConfig contains configuration specific to the Cilium CNI.
type CiliumConfig struct {
	// DEPRECATED: Use HelmValues instead. This field will be ignored when HelmValues is set.
	// PolicyEnforcementMode determines communication allowed between pods. Accepted values are default, always, never.
	PolicyEnforcementMode CiliumPolicyEnforcementMode `json:"policyEnforcementMode,omitempty"`

	// DEPRECATED: Use HelmValues instead. This field will be ignored when HelmValues is set.
	// EgressMasquaradeInterfaces determines which network interfaces are used for masquerading. Accepted values are a valid interface name or interface prefix.
	// +optional
	EgressMasqueradeInterfaces string `json:"egressMasqueradeInterfaces,omitempty"`

	// DEPRECATED: Use HelmValues instead. This field will be ignored when HelmValues is set.
	// SkipUpgrade indicicates that Cilium maintenance should be skipped during upgrades. This can
	// be used when operators wish to self manage the Cilium installation.
	// +optional
	SkipUpgrade *bool `json:"skipUpgrade,omitempty"`

	// DEPRECATED: Use HelmValues instead. This field will be ignored when HelmValues is set.
	// RoutingMode indicates the routing tunnel mode to use for Cilium. Accepted values are overlay (geneve tunnel with overlay)
	// or direct (tunneling disabled with direct routing)
	// Defaults to overlay.
	// +optional
	RoutingMode CiliumRoutingMode `json:"routingMode,omitempty"`

	// DEPRECATED: Use HelmValues instead. This field will be ignored when HelmValues is set.
	// IPv4NativeRoutingCIDR specifies the CIDR to use when RoutingMode is set to direct.
	// When specified, Cilium assumes networking for this CIDR is preconfigured and
	// hands traffic destined for that range to the Linux network stack without
	// applying any SNAT.
	// If this is not set autoDirectNodeRoutes will be set to true
	// +optional
	IPv4NativeRoutingCIDR string `json:"ipv4NativeRoutingCIDR,omitempty"`

	// DEPRECATED: Use HelmValues instead. This field will be ignored when HelmValues is set.
	// IPv6NativeRoutingCIDR specifies the IPv6 CIDR to use when RoutingMode is set to direct.
	// When specified, Cilium assumes networking for this CIDR is preconfigured and
	// hands traffic destined for that range to the Linux network stack without
	// applying any SNAT.
	// If this is not set autoDirectNodeRoutes will be set to true
	// +optional
	IPv6NativeRoutingCIDR string `json:"ipv6NativeRoutingCIDR,omitempty"`

	// DEPRECATED: Use HelmValues instead. This field will be ignored when HelmValues is set.
	// CNIExclusive controls whether Cilium should remove other CNI configuration files.
	// When true (default), Cilium removes other CNI configs; when false, it leaves them alone.
	// +optional
	CNIExclusive *bool `json:"cniExclusive,omitempty"`

	// HelmValues specifies the complete Helm values configuration for Cilium in YAML format.
	// When set, this parameter takes precedence over all other Cilium-specific fields in this configuration.
	// All other Cilium properties (CNIExclusive, EgressMasqueradeInterfaces, IPv4NativeRoutingCIDR, etc.)
	// will be ignored when HelmValues is specified.
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	// +optional
	HelmValues *apiextensionsv1.JSON `json:"helmValues,omitempty"`
}

// IsManaged returns true if SkipUpgrade is nil or false indicating EKS-A is responsible for
// the Cilium installation.
func (n *CiliumConfig) IsManaged() bool { _ = "STUB: not implemented"; return false }

// KindnetdConfig contains configuration specific to the Kindnetd CNI.
type KindnetdConfig struct{}

const (
	// Cilium is the EKS-A Cilium.
	Cilium CNI = "cilium"

	// CiliumEnterprise is Isovalents Cilium.
	CiliumEnterprise CNI = "cilium-enterprise"

	// Kindnetd is the CNI shipped with KinD.
	Kindnetd CNI = "kindnetd"
)

var validCNIs = map[CNI]struct{}{
	Cilium:   {},
	Kindnetd: {},
}

// Policy enforcement modes for Cilium.
const (
	CiliumPolicyModeDefault CiliumPolicyEnforcementMode = "default"
	CiliumPolicyModeAlways  CiliumPolicyEnforcementMode = "always"
	CiliumPolicyModeNever   CiliumPolicyEnforcementMode = "never"
)

var validCiliumPolicyEnforcementModes = map[CiliumPolicyEnforcementMode]bool{
	CiliumPolicyModeAlways:  true,
	CiliumPolicyModeDefault: true,
	CiliumPolicyModeNever:   true,
}

// Routing modes for Cilium.
const (
	CiliumRoutingModeOverlay CiliumRoutingMode = "overlay"
	CiliumRoutingModeDirect  CiliumRoutingMode = "direct"
)

// FailureReasonType is a type for defining failure reasons.
type FailureReasonType string

// Reasons for the terminal failures while reconciling the Cluster object.
const (
	// MissingDependentObjectsReason reports that the Cluster is missing dependent objects.
	MissingDependentObjectsReason FailureReasonType = "MissingDependentObjects"

	// ManagementClusterRefInvalidReason reports that the Cluster management cluster reference is invalid. This
	// can whether if it does not exist or the cluster referenced is not a management cluster.
	ManagementClusterRefInvalidReason FailureReasonType = "ManagementClusterRefInvalid"

	// ClusterInvalidReason reports that the Cluster spec validation has failed.
	ClusterInvalidReason FailureReasonType = "ClusterInvalid"

	// DatacenterConfigInvalidReason reports that the Cluster datacenterconfig validation has failed.
	DatacenterConfigInvalidReason FailureReasonType = "DatacenterConfigInvalid"

	// MachineConfigInvalidReason reports that the Cluster machineconfig validation has failed.
	MachineConfigInvalidReason FailureReasonType = "MachineConfigInvalid"

	// FailureDomainInvalidReason reports that the Cluster failure domain validation has failed.
	FailureDomainInvalidReason FailureReasonType = "FailureDomainInvalid"

	// UnavailableControlPlaneIPReason reports that the Cluster controlPlaneIP is already in use.
	UnavailableControlPlaneIPReason FailureReasonType = "UnavailableControlPlaneIP"

	// EksaVersionInvalidReason reports that the Cluster eksaVersion validation has failed.
	EksaVersionInvalidReason FailureReasonType = "EksaVersionInvalid"

	// BundleNotFoundReason reports that the bundle related to the cluster not found.
	BundleNotFoundReason FailureReasonType = "BundleNotFoundForCluster"

	// ExtendedK8sVersionSupportNotSupportedReason reports that validation for supporting extended kubernetes version failed.
	ExtendedK8sVersionSupportNotSupportedReason FailureReasonType = "ExtendedKubernetesVersionSupportNotSupported"
)

// Reasons for the terminal failures while reconciling the Cluster object specific for Tinkerbell.
const (
	// HardwareInvalidReason reports that the hardware validation has failed.
	HardwareInvalidReason FailureReasonType = "HardwareInvalid"

	// MachineInvalidReason reports that the baremetal machine validation has failed.
	MachineInvalidReason FailureReasonType = "MachineInvalid"
)

// ClusterCertificateInfo contains information about certificate expiration for cluster components.
type ClusterCertificateInfo struct {
	// Machine defines the machine name.
	Machine string `json:"machine"`
	// ExpiresInDays defines the number of days until the certificate expires.
	ExpiresInDays int `json:"expiresInDays"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	// Descriptive message about a fatal problem while reconciling a cluster
	// +optional
	FailureMessage *string `json:"failureMessage,omitempty"`

	// Machine readable value about a terminal problem while reconciling the cluster
	// set at the same time as failureMessage
	// +optional
	FailureReason *FailureReasonType `json:"failureReason,omitempty"`

	// EksdReleaseRef defines the properties of the EKS-D object on the cluster
	EksdReleaseRef *EksdReleaseRef `json:"eksdReleaseRef,omitempty"`
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`

	// ClusterCertificateInfo contains information about all the control plane and external etcd certificates
	// +optional
	ClusterCertificateInfo []ClusterCertificateInfo `json:"clusterCertificateInfo,omitempty"`

	// ReconciledGeneration represents the .metadata.generation the last time the
	// cluster was successfully reconciled. It is the latest generation observed
	// by the controller.
	// NOTE: This field was added for internal use and we do not provide guarantees
	// to its behavior if changed externally. Its meaning and implementation are
	// subject to change in the future.
	ReconciledGeneration int64 `json:"reconciledGeneration,omitempty"`

	// ChildrenReconciledGeneration represents the sum of the .metadata.generation
	// for all the linked objects for the cluster, observed the last time the
	// cluster was successfully reconciled.
	// NOTE: This field was added for internal use and we do not provide guarantees
	// to its behavior if changed externally. Its meaning and implementation are
	// subject to change in the future.
	ChildrenReconciledGeneration int64 `json:"childrenReconciledGeneration,omitempty"`

	// ObservedGeneration is the latest generation observed by the controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

type EksdReleaseRef struct {
	// ApiVersion refers to the EKS-D API version
	ApiVersion string `json:"apiVersion"`
	// Kind refers to the Release kind for the EKS-D object
	Kind string `json:"kind"`
	// Name refers to the name of the EKS-D object on the cluster
	Name string `json:"name"`
	// Namespace refers to the namespace for the EKS-D release resources
	Namespace string `json:"namespace"`
}

type BundlesRef struct {
	// APIVersion refers to the Bundles APIVersion
	APIVersion string `json:"apiVersion"`
	// Name refers to the name of the Bundles object in the cluster
	Name string `json:"name"`
	// Namespace refers to the Bundles's namespace
	Namespace string `json:"namespace"`
}

func (b *BundlesRef) Equal(o *BundlesRef) bool { _ = "STUB: not implemented"; return false }

type Ref struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
}

func (n *Ref) Equal(o *Ref) bool { _ = "STUB: not implemented"; return false }

// IsEmpty checks if the given ref object is empty.
func (n Ref) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// +kubebuilder:object:generate=false
// Interface for getting DatacenterRef fields for Cluster type.
type ProviderRefAccessor interface {
	Kind() string
	Name() string
}

// +kubebuilder:object:generate=false
// Interface for getting Kind field for Cluster type.
type KindAccessor interface {
	Kind() string
	ExpectedKind() string
}

// PackageConfiguration for installing EKS Anywhere curated packages.
type PackageConfiguration struct {
	// Disable package controller on cluster
	Disable bool `json:"disable,omitempty"`

	// Controller package controller configuration
	Controller *PackageControllerConfiguration `json:"controller,omitempty"`

	// Cronjob for ecr token refresher
	CronJob *PackageControllerCronJob `json:"cronjob,omitempty"`
}

// Equal for PackageConfiguration.
func (n *PackageConfiguration) Equal(o *PackageConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// PackageControllerConfiguration configure aspects of package controller.
type PackageControllerConfiguration struct {
	// Repository package controller repository
	Repository string `json:"repository,omitempty"`

	// Tag package controller tag
	Tag string `json:"tag,omitempty"`

	// Digest package controller digest
	Digest string `json:"digest,omitempty"`

	// DisableWebhooks on package controller
	DisableWebhooks bool `json:"disableWebhooks,omitempty"`

	// Env of package controller in the format `key=value`
	Env []string `json:"env,omitempty"`

	// Resources of package controller
	Resources PackageControllerResources `json:"resources,omitempty"`
}

// Equal for PackageControllerConfiguration.
func (n *PackageControllerConfiguration) Equal(o *PackageControllerConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// PackageControllerResources resource aspects of package controller.
type PackageControllerResources struct {
	// Requests for image resources
	Requests ImageResource `json:"requests,omitempty"`
	Limits   ImageResource `json:"limits,omitempty"`
}

// Equal for PackageControllerResources.
func (n *PackageControllerResources) Equal(o *PackageControllerResources) bool {
	_ = "STUB: not implemented"
	return false
}

// ImageResource resources for container image.
type ImageResource struct {
	// CPU image cpu
	CPU string `json:"cpu,omitempty"`

	// Memory image memory
	Memory string `json:"memory,omitempty"`
}

// Equal for ImageResource.
func (n *ImageResource) Equal(o *ImageResource) bool { _ = "STUB: not implemented"; return false }

// PackageControllerCronJob configure aspects of package controller.
type PackageControllerCronJob struct {
	// Repository ecr token refresher repository
	Repository string `json:"repository,omitempty"`

	// Tag ecr token refresher tag
	Tag string `json:"tag,omitempty"`

	// Digest ecr token refresher digest
	Digest string `json:"digest,omitempty"`

	// Disable on cron job
	Disable bool `json:"disable,omitempty"`
}

// Equal for PackageControllerCronJob.
func (n *PackageControllerCronJob) Equal(o *PackageControllerCronJob) bool {
	_ = "STUB: not implemented"
	return false
}

// ExternalEtcdConfiguration defines the configuration options for using unstacked etcd topology.
type ExternalEtcdConfiguration struct {
	Count int `json:"count,omitempty"`
	// MachineGroupRef defines the machine group configuration for the etcd machines.
	MachineGroupRef *Ref `json:"machineGroupRef,omitempty"`
}

func (n *ExternalEtcdConfiguration) Equal(o *ExternalEtcdConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

type ManagementCluster struct {
	Name string `json:"name,omitempty"`
}

func (n *ManagementCluster) Equal(o ManagementCluster) bool {
	_ = "STUB: not implemented"
	return false
}

type PodIAMConfig struct {
	ServiceAccountIssuer string `json:"serviceAccountIssuer"`
}

func (n *PodIAMConfig) Equal(o *PodIAMConfig) bool { _ = "STUB: not implemented"; return false }

// AutoScalingConfiguration defines the configuration for the node autoscaling feature.
type AutoScalingConfiguration struct {
	// MinCount defines the minimum number of nodes for the associated resource group.
	// +optional
	MinCount int `json:"minCount,omitempty"`

	// MaxCount defines the maximum number of nodes for the associated resource group.
	// +optional
	MaxCount int `json:"maxCount,omitempty"`
}

// Equal compares two AutoScalingConfigurations.
func (a *AutoScalingConfiguration) Equal(other *AutoScalingConfiguration) bool {
	_ = "STUB: not implemented"
	return false
}

// UpgradeRolloutStrategyType defines the types of upgrade rollout strategies.
type UpgradeRolloutStrategyType string

const (
	// RollingUpdateStrategyType replaces the old machine by new one using rolling update.
	RollingUpdateStrategyType UpgradeRolloutStrategyType = "RollingUpdate"

	// InPlaceStrategyType upgrades the machines in-place without rolling out any new nodes.
	InPlaceStrategyType UpgradeRolloutStrategyType = "InPlace"
)

// ControlPlaneUpgradeRolloutStrategy indicates rollout strategy for cluster.
type ControlPlaneUpgradeRolloutStrategy struct {
	Type          UpgradeRolloutStrategyType       `json:"type,omitempty"`
	RollingUpdate *ControlPlaneRollingUpdateParams `json:"rollingUpdate,omitempty"`
}

// ControlPlaneRollingUpdateParams is API for rolling update strategy knobs.
type ControlPlaneRollingUpdateParams struct {
	MaxSurge int `json:"maxSurge"`
}

// WorkerNodesUpgradeRolloutStrategy indicates rollout strategy for cluster.
type WorkerNodesUpgradeRolloutStrategy struct {
	Type          UpgradeRolloutStrategyType      `json:"type,omitempty"`
	RollingUpdate *WorkerNodesRollingUpdateParams `json:"rollingUpdate,omitempty"`
}

// Equal compares two WorkerNodesUpgradeRolloutStrategies.
func (w *WorkerNodesUpgradeRolloutStrategy) Equal(other *WorkerNodesUpgradeRolloutStrategy) bool {
	_ = "STUB: not implemented"
	return false
}

// WorkerNodesRollingUpdateParams is API for rolling update strategy knobs.
type WorkerNodesRollingUpdateParams struct {
	MaxSurge       int `json:"maxSurge"`
	MaxUnavailable int `json:"maxUnavailable"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Cluster is the Schema for the clusters API.
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

func (c *Cluster) GetConditions() clusterv1.Conditions {
	_ = "STUB: not implemented"
	return *new(clusterv1.Conditions)
}

func (c *Cluster) SetConditions(conditions clusterv1.Conditions) { _ = "STUB: not implemented"; return }

// +kubebuilder:object:generate=false
// Same as Cluster except stripped down for generation of yaml file during generate clusterconfig.
type ClusterGenerate struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty"`

	Spec ClusterSpec `json:"spec,omitempty"`
}

func (c *Cluster) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *Cluster) ExpectedKind() string { _ = "STUB: not implemented"; return "" }

func (c *Cluster) PausedAnnotation() string { _ = "STUB: not implemented"; return "" }

func (c *Cluster) ControlPlaneAnnotation() string { _ = "STUB: not implemented"; return "" }

// ManagementComponentsVersion returns `management-components version`annotation value on the Cluster object.
func (c *Cluster) ManagementComponentsVersion() string { _ = "STUB: not implemented"; return "" }

// SetManagementComponentsVersion sets the `management-components version` annotation on the Cluster object.
func (c *Cluster) SetManagementComponentsVersion(version string) { _ = "STUB: not implemented"; return }

// DisableControlPlaneIPCheck sets the `skip-ip-check` annotation on the Cluster object.
func (c *Cluster) DisableControlPlaneIPCheck() { _ = "STUB: not implemented"; return }

// ControlPlaneIPCheckDisabled checks it the `skip-ip-check` annotation is set on the Cluster object.
func (c *Cluster) ControlPlaneIPCheckDisabled() bool { _ = "STUB: not implemented"; return false }

// DisableEksaVersionSkewCheck sets the `skip-eksa-version-skew-check` annotation on the Cluster object.
func (c *Cluster) DisableEksaVersionSkewCheck() { _ = "STUB: not implemented"; return }

// EksaVersionSkewCheckDisabled checks if the `skip-eksa-version-skew-check` annotation is set on the Cluster object.
func (c *Cluster) EksaVersionSkewCheckDisabled() bool { _ = "STUB: not implemented"; return false }

func (c *Cluster) ResourceType() string { _ = "STUB: not implemented"; return "" }

func (c *Cluster) EtcdAnnotation() string { _ = "STUB: not implemented"; return "" }

func (c *Cluster) IsSelfManaged() bool { _ = "STUB: not implemented"; return false }

func (c *Cluster) SetManagedBy(managementClusterName string) { _ = "STUB: not implemented"; return }

func (c *Cluster) SetSelfManaged() { _ = "STUB: not implemented"; return }

func (c *ClusterGenerate) SetSelfManaged() { _ = "STUB: not implemented"; return }

func (c *Cluster) ManagementClusterEqual(s2 *Cluster) bool { _ = "STUB: not implemented"; return false }

// IsSingleNode checks if the cluster has only a single node specified between the controlplane and worker nodes.
func (c *Cluster) IsSingleNode() bool { _ = "STUB: not implemented"; return false }

func (c *Cluster) MachineConfigRefs() []Ref { _ = "STUB: not implemented"; return nil }

// SetFailure sets the failureMessage and failureReason of the Cluster status.
func (c *Cluster) SetFailure(failureReason FailureReasonType, failureMessage string) {
	_ = "STUB: not implemented"
	return
}

// ClearFailure clears the failureMessage and failureReason of the Cluster status by setting them to nil.
func (c *Cluster) ClearFailure() { _ = "STUB: not implemented"; return }

// HasFailure checks whether there is a failureMessage and/or failureReason set on the Cluster status.
func (c *Cluster) HasFailure() bool { _ = "STUB: not implemented"; return false }

// KubernetesVersions returns a set of all unique k8s versions specified in the cluster
// for both CP and workers.
func (c *Cluster) KubernetesVersions() []KubernetesVersion { _ = "STUB: not implemented"; return nil }

type refSet map[Ref]struct{}

func (r refSet) addIfNotNil(ref *Ref) bool { _ = "STUB: not implemented"; return false }

func (r refSet) add(ref Ref) bool { _ = "STUB: not implemented"; return false }

func (r refSet) toSlice() []Ref { _ = "STUB: not implemented"; return nil }

func (c *Cluster) ConvertConfigToConfigGenerateStruct() *ClusterGenerate {
	_ = "STUB: not implemented"
	return nil
}

// IsManaged returns true if the Cluster is not self managed.
func (c *Cluster) IsManaged() bool { _ = "STUB: not implemented"; return false }

// ManagedBy returns the Cluster's management cluster's name.
func (c *Cluster) ManagedBy() string { _ = "STUB: not implemented"; return "" }

// IsManagedByCLI returns true if the cluster has the managed-by-cli annotation.
func (c *Cluster) IsManagedByCLI() bool { _ = "STUB: not implemented"; return false }

// CanDeleteWhenPaused returns true if the cluster has the allow-delete-when-paused annotation.
func (c *Cluster) CanDeleteWhenPaused() bool { _ = "STUB: not implemented"; return false }

// +kubebuilder:object:root=true
// ClusterList contains a list of Cluster.
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
