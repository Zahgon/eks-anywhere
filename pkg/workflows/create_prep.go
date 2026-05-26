package workflows

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// CreateNamespaceIfNotPresent creates the namespace on the cluster if it does not already exist.
func CreateNamespaceIfNotPresent(ctx context.Context, namespace string, client kubernetes.Client) error {
	_ = "STUB: not implemented"
	return nil
}
