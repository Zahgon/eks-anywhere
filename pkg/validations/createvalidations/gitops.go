package createvalidations

import (
	"context"
	"fmt"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/validations"
)

var (
	clusterResourceType      = fmt.Sprintf("clusters.%s", v1alpha1.GroupVersion.Group)
	fluxConfigResourceType   = fmt.Sprintf("fluxconfigs.%s", v1alpha1.GroupVersion.Group)
	gitOpsConfigResourceType = fmt.Sprintf("gitopsconfigs.%s", v1alpha1.GroupVersion.Group)
)

func ValidateGitOps(ctx context.Context, k validations.KubectlClient, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// validateGitOpsConfig method will be removed in a future release since gitOpsConfig is deprecated in favor of fluxConfig.
func validateGitOpsConfig(ctx context.Context, k validations.KubectlClient, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFluxConfig(ctx context.Context, k validations.KubectlClient, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// when processing deprecated gitopsConfig, we parse and convert it to fluxConfig.
// for this case both fluxConfig and gitopsConfig can exist in spec. Skip fluxConfig validation.
