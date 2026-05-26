package clusterapi

import (
	_ "embed"

	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

//go:embed config/http-proxy.conf
var proxyConfig string

func proxy(cluster *v1alpha1.Cluster) bootstrapv1beta2.ProxyConfiguration {
	_ = "STUB: not implemented"
	return *new(bootstrapv1beta2.ProxyConfiguration)
}

// SetProxyConfigInKubeadmControlPlaneForBottlerocket sets up proxy configuration in kubeadmControlPlane for bottlerocket.
func SetProxyConfigInKubeadmControlPlaneForBottlerocket(kcp *controlplanev1beta2.KubeadmControlPlane, cluster *v1alpha1.Cluster) {
	_ = "STUB: not implemented"
	return
}

// SetProxyConfigInKubeadmControlPlaneForUbuntu sets up proxy configuration in kubeadmControlPlane for ubuntu.
func SetProxyConfigInKubeadmControlPlaneForUbuntu(kcp *controlplanev1beta2.KubeadmControlPlane, cluster *v1alpha1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// SetProxyConfigInKubeadmConfigTemplateForBottlerocket sets up proxy configuration in kubeadmConfigTemplate for bottlerocket.
func SetProxyConfigInKubeadmConfigTemplateForBottlerocket(kct *bootstrapv1beta2.KubeadmConfigTemplate, cluster *v1alpha1.Cluster) {
	_ = "STUB: not implemented"
	return
}

// SetProxyConfigInKubeadmConfigTemplateForUbuntu sets up proxy configuration in kubeadmConfigTemplate for ubuntu.
func SetProxyConfigInKubeadmConfigTemplateForUbuntu(kct *bootstrapv1beta2.KubeadmConfigTemplate, cluster *v1alpha1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// setProxyConfigInEtcdCluster sets up proxy configuration in etcdadmCluster.
func setProxyConfigInEtcdCluster(etcd *etcdv1.EtcdadmCluster, cluster *v1alpha1.Cluster) {
	_ = "STUB: not implemented"
	return
}

func NoProxyDefaults() []string { _ = "STUB: not implemented"; return nil }

func noProxyList(cluster *v1alpha1.Cluster) []string { _ = "STUB: not implemented"; return nil }

// Add no-proxy defaults

func proxyConfigContent(cluster *v1alpha1.Cluster) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func proxyConfigFile(cluster *v1alpha1.Cluster) (bootstrapv1beta2.File, error) {
	_ = "STUB: not implemented"
	return *new(bootstrapv1beta2.File), nil
}

func addProxyConfigInKubeadmConfigSpecFiles(kcs *bootstrapv1beta2.KubeadmConfigSpec, cluster *v1alpha1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
