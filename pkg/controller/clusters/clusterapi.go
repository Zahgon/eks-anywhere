package clusters

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/controller"
)

// CheckControlPlaneReady is a controller helper to check whether KCP object for
// the cluster is ready or not. This is intended to be used from cluster reconcilers
// due its signature and that it returns controller results with appropriate wait times whenever
// the cluster is not ready.
func CheckControlPlaneReady(ctx context.Context, client client.Client, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// We make sure to check that the status is up to date before using it

// Checking for version as well to avoid race condition of status not being updated in time at least for Kubernetes version upgrades
