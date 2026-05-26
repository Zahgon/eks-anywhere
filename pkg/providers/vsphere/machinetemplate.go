package vsphere

import (
	"context"

	vspherev1 "sigs.k8s.io/cluster-api-provider-vsphere/apis/v1beta1"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

func getMachineTemplate(ctx context.Context, client kubernetes.Client, name, namespace string) (*vspherev1.VSphereMachineTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func machineTemplateEqual(new, old *vspherev1.VSphereMachineTemplate) bool {
	_ = "STUB: not implemented"
	return false
}
