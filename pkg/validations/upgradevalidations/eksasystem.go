package upgradevalidations

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	eksaControllerDeploymentName = "eksa-controller-manager"
)

func ValidateEksaSystemComponents(ctx context.Context, k *executables.Kubectl, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
