package createvalidations

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/validations"
)

func ValidateClusterNameIsUnique(ctx context.Context, k validations.KubectlClient, cluster *types.Cluster, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateManagementCluster(ctx context.Context, k validations.KubectlClient, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
