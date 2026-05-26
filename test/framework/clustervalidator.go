package framework

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/cluster"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

func validationsForExpectedObjects() []clusterf.StateValidation {
	_ = "STUB: not implemented"
	return nil
}

// This should be checked last as the Cluster should only be ready after all the other validations pass.

func validationsForClusterDoesNotExist() []clusterf.StateValidation {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) buildClusterStateValidationConfig(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func newClusterStateValidator(config *clusterf.StateValidationConfig) *clusterf.StateValidator {
	_ = "STUB: not implemented"
	return nil
}

func buildClusterClient(kubeconfigFileName string) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

// Adding the retry logic here because the connection to the client does not always
// succedd on the first try due to connection failure after the kubeconfig becomes
// available in the cluster.

func buildClusterSpec(ctx context.Context, client client.Client, config *cluster.Config) (*cluster.Spec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The cluster config built by the test does not have certain defaults like the bundle reference,
// so fetch that information from the cluster if missing. This is needed inorder to build the cluster spec.
