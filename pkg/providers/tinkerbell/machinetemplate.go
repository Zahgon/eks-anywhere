package tinkerbell

import (
	"context"

	tinkerbellv1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/capt/v1beta1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// GetMachineTemplate gets a TinkerbellMachineTemplate object using the provided client
// If the object doesn't exist, it returns a NotFound error.
func GetMachineTemplate(ctx context.Context, client kubernetes.Client, name, namespace string) (*tinkerbellv1.TinkerbellMachineTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// machineTemplateEqual returns a boolean indicating whether the provided TinkerbellMachineTemplates are equal.
func machineTemplateEqual(new, old *tinkerbellv1.TinkerbellMachineTemplate) bool {
	_ = "STUB: not implemented"
	return false
}
