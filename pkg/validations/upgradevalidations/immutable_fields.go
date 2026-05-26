package upgradevalidations

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/validations"
)

func ValidateImmutableFields(ctx context.Context, k validations.KubectlClient, cluster *types.Cluster, spec *cluster.Spec, provider providers.Provider) error {
	_ = "STUB: not implemented"
	return nil
}

/* compare all clusterNetwork fields individually, since we do allow updating updating fields for configuring plugins such as CiliumConfig through the cli*/

func ValidateGitOpsImmutableFields(ctx context.Context, k validations.KubectlClient, cluster *types.Cluster, clusterSpec *cluster.Spec, oldCluster *v1alpha1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
