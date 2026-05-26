package clusters

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// FetchManagementEksaCluster returns the management cluster object for a given workload Cluster.
// If we are unable to find the management cluster using the same namespace as the current cluster, we will attempt
// to get a list of the clusters with that name across all the namespaces. If we find multiple, which usually should
// not happen as these clusters get mapped to a cluster-api cluster object in the eksa-system namespace, then we
// also error on that because it is not possible to have multiple resources with the same name within a namespace.
func FetchManagementEksaCluster(ctx context.Context, cli client.Client, cluster *v1alpha1.Cluster) (*v1alpha1.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Save error returned from Get if we don't end up finding the cluster through List as it won't return an error
