package snow

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	snowv1 "github.com/aws/eks-anywhere/pkg/providers/snow/api/v1beta1"
)

func getMachineTemplate(ctx context.Context, client kubernetes.Client, name, namespace string) (*snowv1.AWSSnowMachineTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
