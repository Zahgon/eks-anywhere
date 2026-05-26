package snow

import (
	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"

	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const bottlerocketBootstrapImage = "bottlerocket-bootstrap-snow"

func bottlerocketBootstrapSnow(image releasev1.Image) bootstrapv1beta2.BottlerocketBootstrapContainer {
	_ = "STUB: not implemented"
	return *new(bootstrapv1beta2.BottlerocketBootstrapContainer)
}

func addBottlerocketBootstrapSnowInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, image releasev1.Image) {
	_ = "STUB: not implemented"
	return
}

func addBottlerocketBootstrapSnowInKubeadmConfigTemplate(kct *bootstrapv1beta2.KubeadmConfigTemplate, image releasev1.Image) {
	_ = "STUB: not implemented"
	return
}

func addBottlerocketBootstrapSnowInEtcdCluster(etcd *etcdv1.EtcdadmCluster, image releasev1.Image) {
	_ = "STUB: not implemented"
	return
}
