package bootstrapper

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
)

// KindClient is a Kind client.
type KindClient interface {
	CreateBootstrapCluster(ctx context.Context, clusterSpec *cluster.Spec, opts ...BootstrapClusterClientOption) (kubeconfig string, err error)
	DeleteBootstrapCluster(ctx context.Context, cluster *types.Cluster) error
	WithExtraDockerMounts() BootstrapClusterClientOption
	WithExtraPortMappings([]int) BootstrapClusterClientOption
	WithEnv(env map[string]string) BootstrapClusterClientOption
	GetKubeconfig(ctx context.Context, clusterName string) (string, error)
	ClusterExists(ctx context.Context, clusterName string) (bool, error)
}

// KubernetesClient is a Kubernetes client.
type KubernetesClient interface {
	ApplyKubeSpecFromBytes(ctx context.Context, cluster *types.Cluster, data []byte) error
	GetClusters(ctx context.Context, cluster *types.Cluster) ([]types.CAPICluster, error)
	ValidateClustersCRD(ctx context.Context, cluster *types.Cluster) error
	CreateNamespaceIfNotPresent(ctx context.Context, kubeconfig string, namespace string) error
}

// RetrierClientOpt allows to customize a RetrierClient
// on construction.
type RetrierClientOpt func(*RetrierClient)

// WithRetrierClientRetrier allows to use a custom retrier.
func WithRetrierClientRetrier(retrier retrier.Retrier) RetrierClientOpt {
	_ = "STUB: not implemented"
	return *new(RetrierClientOpt)
}

// RetrierClient wraps kind and kubernetes APIs around a retrier.
type RetrierClient struct {
	KindClient
	k8s     KubernetesClient
	retrier retrier.Retrier
}

// NewRetrierClient constructs a new RetrierClient.
func NewRetrierClient(kind KindClient, k8s KubernetesClient, opts ...RetrierClientOpt) RetrierClient {
	_ = "STUB: not implemented"
	return *new(RetrierClient)
}

// Apply creates/updates the data objects for a cluster.
func (c RetrierClient) Apply(ctx context.Context, cluster *types.Cluster, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateNamespace creates a namespace if the namespace does not exist.
func (c RetrierClient) CreateNamespace(ctx context.Context, kubeconfig, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCAPIClusterCRD gets the capi cluster crd in a K8s cluster.
func (c RetrierClient) GetCAPIClusterCRD(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCAPIClusters gets all the capi clusters in a K8s cluster.
func (c RetrierClient) GetCAPIClusters(ctx context.Context, cluster *types.Cluster) ([]types.CAPICluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KindClusterExists checks whether a kind cluster exists by a cluster name.
func (c RetrierClient) KindClusterExists(ctx context.Context, clusterName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetKindClusterKubeconfig gets the kubeconfig for a kind cluster by cluster name.
func (c RetrierClient) GetKindClusterKubeconfig(ctx context.Context, clusterName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DeleteKindCluster deletes a kind cluster by cluster name.
func (c RetrierClient) DeleteKindCluster(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
