package upgradevalidations

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/validations"
)

// ValidatePodDisruptionBudgets returns an error if any pdbs are detected on a cluster.
func ValidatePodDisruptionBudgets(ctx context.Context, k validations.KubectlClient, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
