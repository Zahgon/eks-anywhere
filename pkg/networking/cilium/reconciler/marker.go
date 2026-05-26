package reconciler

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// EKSACiliumInstalledAnnotation indicates a cluster has previously been observed to have
// EKS-A Cilium installed irrespective of whether its still installed.
const EKSACiliumInstalledAnnotation = "anywhere.eks.amazonaws.com/eksa-cilium"

// ciliumWasInstalled checks cluster for the EKSACiliumInstalledAnnotation.
func ciliumWasInstalled(ctx context.Context, cluster *v1alpha1.Cluster) bool {
	_ = "STUB: not implemented"
	return false
}

// markCiliumInstalled populates the EKSACiliumInstalledAnnotation on cluster. It may trigger
// anothe reconciliation event.
func markCiliumInstalled(ctx context.Context, cluster *v1alpha1.Cluster) {
	_ = "STUB: not implemented"
	return
}
