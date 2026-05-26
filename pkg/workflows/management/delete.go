package management

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
	bootstrapper   interfaces.Bootstrapper
	provider       providers.Provider
	writer         filewriter.FileWriter
	clusterManager interfaces.ClusterManager
	gitopsManager  interfaces.GitOpsManager
	clusterDeleter interfaces.ClusterDeleter
	eksdInstaller  interfaces.EksdInstaller
	eksaInstaller  interfaces.EksaInstaller
	clientFactory  interfaces.ClientFactory
	clusterMover   interfaces.ClusterMover
}

// NewDelete builds a new delete construct.
func NewDelete(bootstrapper interfaces.Bootstrapper,
	provider providers.Provider,
	writer filewriter.FileWriter,
	clusterManager interfaces.ClusterManager,
	gitopsManager interfaces.GitOpsManager,
	clusterDeleter interfaces.ClusterDeleter,
	eksdInstaller interfaces.EksdInstaller,
	eksaInstaller interfaces.EksaInstaller,
	clientFactory interfaces.ClientFactory,
	mover interfaces.ClusterMover,
) *Delete {
	_ = "STUB: not implemented"
	return nil
}

// Run executes the tasks to delete a management cluster.
func (c *Delete) Run(ctx context.Context, workload *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}
