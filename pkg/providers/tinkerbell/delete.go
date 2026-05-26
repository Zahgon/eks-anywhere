package tinkerbell

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/types"
)

func (p *Provider) SetupAndValidateDeleteCluster(ctx context.Context, cluster *types.Cluster, _ *cluster.Spec) error {
	_ = "STUB: not implemented"
	// noop
	return nil
}

func (p *Provider) PostClusterDeleteValidate(ctx context.Context, managementCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
