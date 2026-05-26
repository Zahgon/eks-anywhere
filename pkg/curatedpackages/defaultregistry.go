package curatedpackages

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/version"
)

type DefaultRegistry struct {
	releaseManifestReader Reader
	kubeVersion           string
	cliVersion            version.Info
}

func NewDefaultRegistry(rmr Reader, kv string, cv version.Info) *DefaultRegistry {
	_ = "STUB: not implemented"
	return nil
}

func (dr *DefaultRegistry) GetRegistryBaseRef(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Use package controller registry to fetch packageBundles.
// Format of controller image is: <uri>/<env_type>/<repository_name>
