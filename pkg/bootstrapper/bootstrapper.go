package bootstrapper

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/types"
)

type Bootstrapper struct {
	clusterClient ClusterClient
}

type ClusterClient interface {
	Apply(ctx context.Context, cluster *types.Cluster, data []byte) error
	CreateNamespace(ctx context.Context, kubeconfig, namespace string) error
	GetCAPIClusterCRD(ctx context.Context, cluster *types.Cluster) error
	GetCAPIClusters(ctx context.Context, cluster *types.Cluster) ([]types.CAPICluster, error)
	KindClusterExists(ctx context.Context, clusterName string) (bool, error)
	GetKindClusterKubeconfig(ctx context.Context, clusterName string) (string, error)
	CreateBootstrapCluster(ctx context.Context, clusterSpec *cluster.Spec, opts ...BootstrapClusterClientOption) (string, error)
	DeleteKindCluster(ctx context.Context, cluster *types.Cluster) error
	WithExtraDockerMounts() BootstrapClusterClientOption
	WithExtraPortMappings([]int) BootstrapClusterClientOption
	WithEnv(env map[string]string) BootstrapClusterClientOption
}

type (
	BootstrapClusterClientOption func() error
	BootstrapClusterOption       func(b *Bootstrapper) BootstrapClusterClientOption
)

// New constructs a new bootstrapper.
func New(clusterClient ClusterClient) *Bootstrapper { _ = "STUB: not implemented"; return nil }

func (b *Bootstrapper) CreateBootstrapCluster(ctx context.Context, clusterSpec *cluster.Spec, opts ...BootstrapClusterOption) (*types.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bootstrapper) DeleteBootstrapCluster(ctx context.Context, cluster *types.Cluster, operationType constants.Operation, isForceCleanup bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bootstrapper) managementInCluster(ctx context.Context, cluster *types.Cluster) (*types.CAPICluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bootstrapper) getClientOptions(opts []BootstrapClusterOption) []BootstrapClusterClientOption {
	_ = "STUB: not implemented"
	return nil
}

func WithExtraDockerMounts() BootstrapClusterOption {
	_ = "STUB: not implemented"
	return *new(BootstrapClusterOption)
}

func WithExtraPortMappings(ports []int) BootstrapClusterOption {
	_ = "STUB: not implemented"
	return *new(BootstrapClusterOption)
}

func WithEnv(env map[string]string) BootstrapClusterOption {
	_ = "STUB: not implemented"
	return *new(BootstrapClusterOption)
}
