package reconciler

import (
	"context"

	"github.com/go-logr/logr"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	anywhereCluster "github.com/aws/eks-anywhere/pkg/cluster"
)

// Reconciler allows to reconcile machine health checks.
type Reconciler struct {
	client    client.Client
	defaulter anywhereCluster.MachineHealthCheckDefaulter
}

// New returns a new Reconciler.
func New(client client.Client, defaulter anywhereCluster.MachineHealthCheckDefaulter) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile installs machine health checks for a given cluster.
func (r *Reconciler) Reconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func copyNodeStartupTimeoutToCluster(cluster *anywherev1.Cluster, capiMHC *clusterv1beta2.MachineHealthCheck) {
	_ = "STUB: not implemented"
	return
}

func copyUnhealthyMachineTimeoutToCluster(cluster *anywherev1.Cluster, capiMHC *clusterv1beta2.MachineHealthCheck) {
	_ = "STUB: not implemented"
	return
}
