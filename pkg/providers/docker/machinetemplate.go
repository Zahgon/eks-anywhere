package docker

import (
	"context"

	dockerv1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1beta2"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// GetMachineTemplate gets a DockerMachineTemplate object using the provided client
// If the object doesn't exist, it returns a NotFound error.
func GetMachineTemplate(ctx context.Context, client kubernetes.Client, name, namespace string) (*dockerv1.DockerMachineTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MachineTemplateEqual returns a boolean indicating whether or not the provided DockerMachineTemplates are equal.
func MachineTemplateEqual(new, old *dockerv1.DockerMachineTemplate) bool {
	_ = "STUB: not implemented"
	return false
}
