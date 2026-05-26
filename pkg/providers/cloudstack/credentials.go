package cloudstack

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack/decoder"
)

// GetCloudstackExecConfig gets cloudstack exec config from secrets.
func GetCloudstackExecConfig(ctx context.Context, cli client.Client, datacenterConfig *v1alpha1.CloudStackDatacenterConfig) (*decoder.CloudStackExecConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
