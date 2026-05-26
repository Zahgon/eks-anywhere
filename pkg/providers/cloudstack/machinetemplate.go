package cloudstack

import (
	"context"

	cloudstackv1 "sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta3"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// GetMachineTemplate gets a CloudStackMachineTemplate object using the provided client
// If the object doesn't exist, it returns a NotFound error.
func GetMachineTemplate(ctx context.Context, client kubernetes.Client, name, namespace string) (*cloudstackv1.CloudStackMachineTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// machineTemplateEqual returns a boolean indicating whether the provided CloudStackMachineTemplates are equal.
func machineTemplateEqual(new, old *cloudstackv1.CloudStackMachineTemplate) bool {
	_ = "STUB: not implemented"
	// Compare new -> old and old -> new because DeepDerivative ignores fields in the first param
	// that are default values. This is important for cases where an optional field is removed
	// from the spec.
	return false
}
