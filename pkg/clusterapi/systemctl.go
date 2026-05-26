package clusterapi

import (
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

var buildContainerdConfigCommands = []string{
	"cat /etc/containerd/config_append.toml >> /etc/containerd/config.toml",
}

var restartContainerdCommands = []string{
	"sudo systemctl daemon-reload",
	"sudo systemctl restart containerd",
}

// CreateContainerdConfigFileInKubeadmControlPlane adds the prekubeadm command to create containerd config file in kubeadmControlPlane if registry mirror config exists.
func CreateContainerdConfigFileInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, cluster *v1alpha1.Cluster) {
	_ = "STUB: not implemented"
	return
}

// CreateContainerdConfigFileInKubeadmConfigTemplate adds the prekubeadm command to create containerd config file in kubeadmConfigTemplate if registry mirror config exists.
func CreateContainerdConfigFileInKubeadmConfigTemplate(kct *bootstrapv1beta2.KubeadmConfigTemplate, cluster *v1alpha1.Cluster) {
	_ = "STUB: not implemented"
	return
}

// RestartContainerdInKubeadmControlPlane adds the prekubeadm command to restart containerd daemon in kubeadmControlPlane if registry mirror or proxy config exists.
func RestartContainerdInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, cluster *v1alpha1.Cluster) {
	_ = "STUB: not implemented"
	return
}

// RestartContainerdInKubeadmConfigTemplate adds the prekubeadm command to restart containerd daemon in kubeadmConfigTemplate if registry mirror or proxy config exists.
func RestartContainerdInKubeadmConfigTemplate(kct *bootstrapv1beta2.KubeadmConfigTemplate, cluster *v1alpha1.Cluster) {
	_ = "STUB: not implemented"
	return
}

func restartContainerdNeeded(cluster *v1alpha1.Cluster) bool {
	_ = "STUB: not implemented"
	return false
}
