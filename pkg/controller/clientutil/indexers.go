package clientutil

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClusterNameIndexer is an indexer for controller to list Cluster objects based on name.
func ClusterNameIndexer(obj client.Object) []string { _ = "STUB: not implemented"; return nil }
