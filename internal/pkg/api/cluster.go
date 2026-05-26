package api

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type ClusterFiller func(c *anywherev1.Cluster)

// ClusterToConfigFiller updates the Cluster in the cluster.Config by applying all the fillers.
func ClusterToConfigFiller(fillers ...ClusterFiller) ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(ClusterConfigFiller)
}

// JoinClusterConfigFillers creates one single ClusterConfigFiller from a collection of fillers.
func JoinClusterConfigFillers(fillers ...ClusterConfigFiller) ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(ClusterConfigFiller)
}

func WithKubernetesVersion(v anywherev1.KubernetesVersion) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithLicenseToken sets LicenseToken with the provided token value to use.
func WithLicenseToken(licenseToken string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithBundlesRef sets BundlesRef with the provided name to use.
func WithBundlesRef(name string, namespace string, apiVersion string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithEksaVersion sets EksaVersion with the provided name to use.
func WithEksaVersion(eksaVersion *anywherev1.EksaVersion) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithCiliumPolicyEnforcementMode(mode anywherev1.CiliumPolicyEnforcementMode) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithCiliumEgressMasqueradeInterfaces sets the egressMasqueradeInterfaces with the provided interface option to use.
func WithCiliumEgressMasqueradeInterfaces(interfaceName string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithCiliumSkipUpgrade enables skip upgrade for EKSA Cilium installations.
func WithCiliumSkipUpgrade() ClusterFiller { _ = "STUB: not implemented"; return *new(ClusterFiller) }

// WithCiliumRoutingMode sets the tunnel mode with the provided mode option to use.
func WithCiliumRoutingMode(mode anywherev1.CiliumRoutingMode) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithCiliumHelmValues sets the Helm values configuration for Cilium.
func WithCiliumHelmValues(helmValues map[string]interface{}) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// Marshal to JSON

func WithClusterNamespace(ns string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithControlPlaneCount(r int) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithControlPlaneEndpointIP(value string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithControlPlaneTaints(taints []corev1.Taint) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithControlPlaneLabel(key string, val string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithControlPlaneAPIServerExtraArgs adds the APIServerExtraArgs to the cluster spec.
func WithControlPlaneAPIServerExtraArgs() ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithControlPlaneKubeletConfig adds the Kubelet config to the control plane in cluster spec.
func WithControlPlaneKubeletConfig(kc *unstructured.Unstructured) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// RemoveAllAPIServerExtraArgs removes all the API server flags from the cluster spec.
func RemoveAllAPIServerExtraArgs() ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithPodCidr sets an explicit pod CIDR, overriding the provider's default.
func WithPodCidr(podCidr string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithServiceCidr(svcCidr string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithWorkerKubernetesVersion sets the kubernetes version field for the given worker group.
func WithWorkerKubernetesVersion(name string, version *anywherev1.KubernetesVersion) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// Append the worker node group if not already found in existing configuration

func workerNodeWithKubernetesVersion(name string, version *anywherev1.KubernetesVersion) anywherev1.WorkerNodeGroupConfiguration {
	_ = "STUB: not implemented"
	return *new(anywherev1.WorkerNodeGroupConfiguration)
}

// WithWorkerNodeKubeletConfig adds the Kubelet config to the worker node groups in cluster spec.
func WithWorkerNodeKubeletConfig(kc *unstructured.Unstructured) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithWorkerNodeCount(r int) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithWorkerNodeAutoScalingConfig adds an autoscaling configuration with a given min and max count.
func WithWorkerNodeAutoScalingConfig(min int, max int) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithOIDCIdentityProviderRef(name string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithGitOpsRef(name, kind string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithExternalEtcdTopology(count int) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithEtcdCountIfExternal(count int) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithExternalEtcdMachineRef(kind string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithStackedEtcdTopology() ClusterFiller { _ = "STUB: not implemented"; return *new(ClusterFiller) }

func WithProxyConfig(httpProxy, httpsProxy string, noProxy []string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithRegistryMirror adds a registry mirror configuration.
func WithRegistryMirror(endpoint, port string, caCert string, authenticate bool, insecureSkipVerify bool, ociNamespaces ...anywherev1.OCINamespace) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithManagementCluster(name string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithAWSIamIdentityProviderRef(name string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func RemoveAllWorkerNodeGroups() ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func RemoveWorkerNodeGroup(name string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

func WithWorkerNodeGroup(name string, fillers ...WorkerNodeGroupFiller) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithPodIamFiller configures pod IAM config to enable IRSA.
func WithPodIamFiller(issuerURL string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithEtcdEncryptionFiller configures EtcdEncyption on the cluster.
func WithEtcdEncryptionFiller(kms *anywherev1.KMS, resources []string) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithInPlaceUpgradeStrategy configures the UpgradeStrategy on Control-plane and Worker node groups to InPlace.
func WithInPlaceUpgradeStrategy() ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithSkipAdmissionForSystemResources enables or disables skipping admission plugins for system resources.
func WithSkipAdmissionForSystemResources(skip bool) ClusterFiller {
	_ = "STUB: not implemented"
	return *new(ClusterFiller)
}

// WithPackagesDisabled disables the package controller installation on the cluster.
func WithPackagesDisabled() ClusterFiller { _ = "STUB: not implemented"; return *new(ClusterFiller) }
