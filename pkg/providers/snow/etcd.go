package snow

import (
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// FgEtcdLearner is a Kubeadm feature gate for etcd learner mode.
const FgEtcdLearner = "EtcdLearnerMode"

func addStackedEtcdExtraArgsInKubeadmControlPlane(kcp *controlplanev1beta2.KubeadmControlPlane, externalEtcdConfig *v1alpha1.ExternalEtcdConfiguration) {
	_ = "STUB: not implemented"
	return
}

func disableEtcdLearnerMode(kcp *controlplanev1beta2.KubeadmControlPlane) {
	_ = "STUB: not implemented"
	return
}
