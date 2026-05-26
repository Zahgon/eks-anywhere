package workload

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces"
)

// Delete is the workflow that deletes a workload clusters.
type Delete struct {
	provider       providers.Provider
	writer         filewriter.FileWriter
	clusterManager interfaces.ClusterManager
	clusterDeleter interfaces.ClusterDeleter
	gitopsManager  interfaces.GitOpsManager
}

// NewDelete builds a new delete construct.
func NewDelete(provider providers.Provider,
	writer filewriter.FileWriter,
	clusterManager interfaces.ClusterManager,
	clusterDeleter interfaces.ClusterDeleter,
	gitopsManager interfaces.GitOpsManager,
) *Delete {
	_ = "STUB: not implemented"
	return nil
}

// Run executes the tasks to delete a workload cluster.
func (c *Delete) Run(ctx context.Context, workload *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}
