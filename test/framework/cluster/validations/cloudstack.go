package validations

import (
	"context"

	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

// ValidateAvailabilityZones checks each availability zones defined cloudstackdatacenterconfig in the cluster.Spec
// have corresponding cloudstackfailuredomains objects within the cluster.
func ValidateAvailabilityZones(ctx context.Context, vc clusterf.StateValidationConfig) error {
	_ = "STUB: not implemented"
	return nil
}
