package clustermanager

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/kubeconfig"
	"github.com/aws/eks-anywhere/pkg/types"
)

// ClusterApplier is responsible for applying the cluster spec to the cluster.
type ClusterApplier interface {
	Run(ctx context.Context, spec *cluster.Spec, managementCluster types.Cluster) error
}

// ClusterCreator is responsible for applying the cluster config and writing the kubeconfig file.
type ClusterCreator struct {
	ClusterApplier
	kubeconfigWriter kubeconfig.Writer
	fs               filewriter.FileWriter
}

// NewClusterCreator creates a ClusterCreator.
func NewClusterCreator(applier ClusterApplier, kubeconfigWriter kubeconfig.Writer, fs filewriter.FileWriter) *ClusterCreator {
	_ = "STUB: not implemented"
	return nil
}

// CreateSync creates a workload cluster using the EKS-A controller and returns the types.Cluster object for that cluster.
func (cc ClusterCreator) CreateSync(ctx context.Context, spec *cluster.Spec, managementCluster *types.Cluster) (*types.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc ClusterCreator) buildClusterAccess(ctx context.Context, clusterName string, management *types.Cluster) (*types.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
