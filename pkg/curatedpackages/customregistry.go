package curatedpackages

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/helm"
)

type CustomRegistry struct {
	helm.Client
	registry string
}

// NewCustomRegistry returns a new CustomRegistry.
func NewCustomRegistry(helm helm.Client, registry string) *CustomRegistry {
	_ = "STUB: not implemented"
	return nil
}

func (cm *CustomRegistry) GetRegistryBaseRef(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
