package clusterapi

import (
	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/release/api/v1alpha1"
)

func bottlerocketBootstrap(image v1alpha1.Image) bootstrapv1beta2.BottlerocketBootstrap {
	_ = "STUB: not implemented"
	return *new(bootstrapv1beta2.BottlerocketBootstrap)
}

func bottlerocketAdmin(image v1alpha1.Image) bootstrapv1beta2.BottlerocketAdmin {
	_ = "STUB: not implemented"
	return *new(bootstrapv1beta2.BottlerocketAdmin)
}

func bottlerocketControl(image v1alpha1.Image) bootstrapv1beta2.BottlerocketControl {
	_ = "STUB: not implemented"
	return *new(bootstrapv1beta2.BottlerocketControl)
}

func pause(image v1alpha1.Image) bootstrapv1beta2.Pause {
	_ = "STUB: not implemented"
	return *new(bootstrapv1beta2.Pause)
}

func hostConfig(config *anywherev1.HostOSConfiguration) *bootstrapv1beta2.BottlerocketSettings {
	_ = "STUB: not implemented"
	return nil
}

// SetBottlerocketInKubeadmControlPlane adds bottlerocket bootstrap image metadata in kubeadmControlPlane.
func SetBottlerocketInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, versionsBundle *cluster.VersionsBundle) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketAdminContainerImageInKubeadmControlPlane overrides the default bottlerocket admin container image metadata in kubeadmControlPlane.
func SetBottlerocketAdminContainerImageInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, versionsBundle *cluster.VersionsBundle) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketControlContainerImageInKubeadmControlPlane overrides the default bottlerocket control container image metadata in kubeadmControlPlane.
func SetBottlerocketControlContainerImageInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, versionsBundle *cluster.VersionsBundle) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketInKubeadmConfigTemplate adds bottlerocket bootstrap image metadata in kubeadmConfigTemplate.
func SetBottlerocketInKubeadmConfigTemplate(kct *bootstrapv1beta2.KubeadmConfigTemplate, versionsBundle *cluster.VersionsBundle) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketAdminContainerImageInKubeadmConfigTemplate overrides the default bottlerocket admin container image metadata in kubeadmConfigTemplate.
func SetBottlerocketAdminContainerImageInKubeadmConfigTemplate(kct *bootstrapv1beta2.KubeadmConfigTemplate, versionsBundle *cluster.VersionsBundle) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketControlContainerImageInKubeadmConfigTemplate overrides the default bottlerocket control container image metadata in kubeadmConfigTemplate.
func SetBottlerocketControlContainerImageInKubeadmConfigTemplate(kct *bootstrapv1beta2.KubeadmConfigTemplate, versionsBundle *cluster.VersionsBundle) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketHostConfigInKubeadmControlPlane sets bottlerocket specific kernel settings in kubeadmControlPlane.
func SetBottlerocketHostConfigInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, hostOSConfig *anywherev1.HostOSConfiguration) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketHostConfigInKubeadmConfigTemplate sets bottlerocket specific kernel settings in kubeadmConfigTemplate.
func SetBottlerocketHostConfigInKubeadmConfigTemplate(kct *bootstrapv1beta2.KubeadmConfigTemplate, hostOSConfig *anywherev1.HostOSConfiguration) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketInEtcdCluster adds bottlerocket config in etcdadmCluster.
func SetBottlerocketInEtcdCluster(etcd *etcdv1.EtcdadmCluster, versionsBundle *cluster.VersionsBundle) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketAdminContainerImageInEtcdCluster overrides the default bottlerocket admin container image metadata in etcdadmCluster.
func SetBottlerocketAdminContainerImageInEtcdCluster(etcd *etcdv1.EtcdadmCluster, adminImage v1alpha1.Image) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketControlContainerImageInEtcdCluster overrides the default bottlerocket control container image metadata in etcdadmCluster.
func SetBottlerocketControlContainerImageInEtcdCluster(etcd *etcdv1.EtcdadmCluster, controlImage v1alpha1.Image) {
	_ = "STUB: not implemented"
	return
}

// SetBottlerocketHostConfigInEtcdCluster sets bottlerocket specific kernel settings in etcdadmCluster.
func SetBottlerocketHostConfigInEtcdCluster(etcd *etcdv1.EtcdadmCluster, hostOSConfig *anywherev1.HostOSConfiguration) {
	_ = "STUB: not implemented"
	return
}
