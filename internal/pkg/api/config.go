package api

import "github.com/aws/eks-anywhere/pkg/cluster"

// ClusterConfigFiller updates a cluster.Config.
type ClusterConfigFiller func(*cluster.Config)

// UpdateClusterConfig updates the given cluster.Config by applying the fillers.
func UpdateClusterConfig(config *cluster.Config, fillers ...ClusterConfigFiller) {
	_ = "STUB: not implemented"
	return
}
