package awsiamauth

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/types"
)

// Client is a Kubernetes client.
type Client interface {
	ApplyKubeSpecFromBytes(ctx context.Context, cluster *types.Cluster, data []byte) error
	GetApiServerUrl(ctx context.Context, cluster *types.Cluster) (string, error)
	GetObject(ctx context.Context, resourceType string, name string, namespace string, kubeconfig string, obj runtime.Object) error
}

// RetrierClient wraps basic kubernetes API operations around a retrier.
type RetrierClient struct {
	client  Client
	retrier retrier.Retrier
}

// RetrierClientOpt allows to customize a RetrierClient
// on construction.
type RetrierClientOpt func(*RetrierClient)

// RetrierClientRetrier allows to use a custom retrier.
func RetrierClientRetrier(retrier retrier.Retrier) RetrierClientOpt {
	_ = "STUB: not implemented"
	return *new(RetrierClientOpt)
}

// NewRetrierClient constructs a new RetrierClient.
func NewRetrierClient(client Client, opts ...RetrierClientOpt) RetrierClient {
	_ = "STUB: not implemented"
	return *new(RetrierClient)
}

// Apply creates/updates the data objects for a cluster.
func (c RetrierClient) Apply(ctx context.Context, cluster *types.Cluster, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAPIServerURL gets the api server url from K8s config.
func (c RetrierClient) GetAPIServerURL(ctx context.Context, cluster *types.Cluster) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetClusterCACert gets the ca cert for a cluster from a secret.
func (c RetrierClient) GetClusterCACert(ctx context.Context, cluster *types.Cluster, clusterName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAWSIAMKubeconfigSecretValue gets the AWS IAM kubeconfig value for a cluster from a secret.
func (c RetrierClient) GetAWSIAMKubeconfigSecretValue(ctx context.Context, cluster *types.Cluster, clusterName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
