package clusterapi

import (
	corev1 "k8s.io/api/core/v1"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
)

// SetKubeVipInKubeadmControlPlane appends kube-vip manifest to kubeadmControlPlane's kubeadmConfigSpec files.
func SetKubeVipInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, address, image string) error {
	_ = "STUB: not implemented"
	return nil
}

func kubeVip(address, image string) *corev1.Pod { _ = "STUB: not implemented"; return nil }
