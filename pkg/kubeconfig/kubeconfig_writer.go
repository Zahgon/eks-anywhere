package kubeconfig

import (
	"context"
	"io"
	"time"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// ClientFactory builds Kubernetes clients.
type ClientFactory interface {
	// BuildClientFromKubeconfig builds a Kubernetes client from a kubeconfig file.
	BuildClientFromKubeconfig(kubeconfigPath string) (kubernetes.Client, error)
}

// ClusterAPIKubeconfigSecretWriter reads the kubeconfig secret on a cluster and copies the contents to a writer.
type ClusterAPIKubeconfigSecretWriter struct {
	client  ClientFactory
	timeout time.Duration
	backoff time.Duration
}

// WriterOpt allows to configure [KubeconfigWriter].
type WriterOpt func(*ClusterAPIKubeconfigSecretWriter)

// WithTimeout sets the optional timeout for a KubeconfigWriter.
func WithTimeout(timeout time.Duration) WriterOpt {
	_ = "STUB: not implemented"
	return *new(WriterOpt)
}

// WithBackoff sets the optional backoff duration for a KubeconfigWriter.
func WithBackoff(backoff time.Duration) WriterOpt {
	_ = "STUB: not implemented"
	return *new(WriterOpt)
}

// NewClusterAPIKubeconfigSecretWriter creates a ClusterAPIKubeconfigSecretWriter.
func NewClusterAPIKubeconfigSecretWriter(unauthClient ClientFactory, opts ...WriterOpt) ClusterAPIKubeconfigSecretWriter {
	_ = "STUB: not implemented"
	return *new(ClusterAPIKubeconfigSecretWriter)
}

// WriteKubeconfig retrieves the contents of the specified cluster's kubeconfig from a secret and copies it to an io.Writer.
func (kr ClusterAPIKubeconfigSecretWriter) WriteKubeconfig(ctx context.Context, clusterName, kubeconfigPath string, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteKubeconfigContent copies a raw kubeconfig to an io.Writer.
func (kr ClusterAPIKubeconfigSecretWriter) WriteKubeconfigContent(ctx context.Context, clusterName string, content []byte, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// GetClusterKubeconfig gets the cluster's kubeconfig from the secret.
func (kr ClusterAPIKubeconfigSecretWriter) GetClusterKubeconfig(ctx context.Context, clusterName, kubeconfigPath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
