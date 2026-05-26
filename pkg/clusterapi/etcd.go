package clusterapi

import (
	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

// SetUbuntuConfigInEtcdCluster sets up the etcd config in EtcdadmCluster.
func SetUbuntuConfigInEtcdCluster(etcd *etcdv1.EtcdadmCluster, versionsBundle *cluster.VersionsBundle, eksaVersion *v1alpha1.EksaVersion) {
	_ = "STUB: not implemented"
	return
}

// SetEtcdConfigInCluster sets up the etcd config in CAPI Cluster.
func setUnstackedEtcdConfigInCluster(cluster *clusterv1beta2.Cluster, unstackedEtcdObject APIObject) {
	_ = "STUB: not implemented"
	return
}

// SetUnstackedEtcdConfigInKubeadmControlPlaneForBottlerocket sets up unstacked etcd configuration in kubeadmControlPlane for bottlerocket.
func SetUnstackedEtcdConfigInKubeadmControlPlaneForBottlerocket(kcp *controlplanev1beta2.KubeadmControlPlane, externalEtcdConfig *v1alpha1.ExternalEtcdConfiguration) {
	_ = "STUB: not implemented"
	return
}

// SetUnstackedEtcdConfigInKubeadmControlPlaneForUbuntu sets up unstacked etcd configuration in kubeadmControlPlane for ubuntu.
func SetUnstackedEtcdConfigInKubeadmControlPlaneForUbuntu(kcp *controlplanev1beta2.KubeadmControlPlane, externalEtcdConfig *v1alpha1.ExternalEtcdConfiguration) {
	_ = "STUB: not implemented"
	return
}

// setStackedEtcdConfigInKubeadmControlPlane sets up stacked etcd configuration in kubeadmControlPlane.
func setStackedEtcdConfigInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, etcd cluster.VersionedRepository) {
	_ = "STUB: not implemented"
	return
}
