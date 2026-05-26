package framework

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/internal/pkg/api"
)

const (
	maxPod50 = 50
	maxPod60 = 60
)

// WithKubeletConfig returns the default kubelet config set for e2e test.
func WithKubeletConfig() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithKubeletClusterConfig returns a ClusterConfigFiller that adds the default
// KubeletConfig for E2E tests to the cluster Config.
func WithKubeletClusterConfig() api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// ValidateKubeletConfig validates the kubelet config specified in the cluster spec has been applied to the nodes.
func (e *ClusterE2ETest) ValidateKubeletConfig() { _ = "STUB: not implemented"; return }

func getWorkerNodes(all, cpNodes []corev1.Node) *corev1.Node { _ = "STUB: not implemented"; return nil }
