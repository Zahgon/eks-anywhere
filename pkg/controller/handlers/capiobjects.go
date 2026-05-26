package handlers

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/handler"
)

// CAPIObjectToCluster returns a request handler that enqueues an EKS-A Cluster
// reconcile request for CAPI objects that contain the cluster name and namespace labels.
func CAPIObjectToCluster(log logr.Logger) handler.MapFunc {
	_ = "STUB: not implemented"
	return *new(handler.MapFunc)
}

// Object not managed by an eks-a Cluster, don't enqueue
