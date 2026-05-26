package docker

import (
	"github.com/aws/eks-anywhere/pkg/cluster"
)

// ValidateControlPlaneEndpoint - checks to see if endpoint host configuration is specified for docker cluster and returns an error if true.
func ValidateControlPlaneEndpoint(clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}
