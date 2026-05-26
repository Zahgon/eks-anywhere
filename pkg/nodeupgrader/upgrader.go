package nodeupgrader

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	upgradeBin = "/foo/eksa-upgrades/tools/upgrader"

	// CopierContainerName holds the name of the components copier container.
	CopierContainerName = "components-copier"

	// ContainerdUpgraderContainerName holds the name of the containerd upgrader container.
	ContainerdUpgraderContainerName = "containerd-upgrader"

	// CNIPluginsUpgraderContainerName holds the name of the CNI plugins upgrader container.
	CNIPluginsUpgraderContainerName = "cni-plugins-upgrader"

	// KubeadmUpgraderContainerName holds the name of the kubeadm upgrader container.
	KubeadmUpgraderContainerName = "kubeadm-upgrader"

	// KubeletUpgradeContainerName holds the name of the kubelet/kubectl upgrader container.
	KubeletUpgradeContainerName = "kubelet-kubectl-upgrader"

	// PostUpgradeContainerName holds the name of the post upgrade cleanup/status report container.
	PostUpgradeContainerName = "post-upgrade-status"
)

// PodName returns the name of the upgrader pod based on the nodeName.
func PodName(nodeName string) string { _ = "STUB: not implemented"; return "" }

// UpgradeFirstControlPlanePod returns an upgrader pod that should be deployed on the first control plane node.
func UpgradeFirstControlPlanePod(nodeName, image, kubernetesVersion, etcdVersion string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// UpgradeSecondaryControlPlanePod returns an upgrader pod that can be deployed on the remaining control plane nodes.
func UpgradeSecondaryControlPlanePod(nodeName, image, kubernetesVersion string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// UpgradeWorkerPod returns an upgrader pod that can be deployed on worker nodes.
func UpgradeWorkerPod(nodeName, image string) *corev1.Pod { _ = "STUB: not implemented"; return nil }

func upgraderPod(nodeName, image string, isCP bool) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func containersForUpgrade(isCP bool, image, nodeName string, kubeadmUpgradeCommand ...string) []corev1.Container {
	_ = "STUB: not implemented"
	return nil
}

func copierContainer(image string, isCP bool) corev1.Container {
	_ = "STUB: not implemented"
	return *new(corev1.Container)
}

func nsenterContainer(image, name string, extraArgs ...string) corev1.Container {
	_ = "STUB: not implemented"
	return *new(corev1.Container)
}

func hostComponentsVolume() corev1.Volume { _ = "STUB: not implemented"; return *new(corev1.Volume) }

func kubeVipVolume() corev1.Volume { _ = "STUB: not implemented"; return *new(corev1.Volume) }
