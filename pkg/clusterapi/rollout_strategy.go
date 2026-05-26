package clusterapi

import (
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// SetUpgradeRolloutStrategyInKubeadmControlPlane updates the kubeadm control plane with the upgrade rollout strategy defined in an eksa cluster.
func SetUpgradeRolloutStrategyInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, rolloutStrategy *anywherev1.ControlPlaneUpgradeRolloutStrategy) {
	_ = "STUB: not implemented"
	return
}

// SetUpgradeRolloutStrategyInMachineDeployment updates the machine deployment with the upgrade rollout strategy defined in an eksa cluster.
func SetUpgradeRolloutStrategyInMachineDeployment(md *clusterv1beta2.MachineDeployment, rolloutStrategy *anywherev1.WorkerNodesUpgradeRolloutStrategy) {
	_ = "STUB: not implemented"
	return
}
