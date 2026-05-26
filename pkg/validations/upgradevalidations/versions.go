package upgradevalidations

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/validations"
)

// ValidateServerVersionSkew validates Kubernetes version skew between upgrades for the CLI.
func ValidateServerVersionSkew(ctx context.Context, newCluster *anywherev1.Cluster, cluster *types.Cluster, managementCluster *types.Cluster, kubectl validations.KubectlClient) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateWorkerServerVersionSkew validates worker node group Kubernetes version skew between upgrades.
func ValidateWorkerServerVersionSkew(ctx context.Context, newCluster *anywherev1.Cluster, cluster *types.Cluster, managementCluster *types.Cluster, kubectl validations.KubectlClient) error {
	_ = "STUB: not implemented"
	return nil
}
