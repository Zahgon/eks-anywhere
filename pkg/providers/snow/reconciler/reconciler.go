package reconciler

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
	"github.com/aws/eks-anywhere/pkg/controller/clusters"
	"github.com/aws/eks-anywhere/pkg/controller/serverside"
	"github.com/aws/eks-anywhere/pkg/providers/snow"
)

type CNIReconciler interface {
	Reconcile(ctx context.Context, logger logr.Logger, client client.Client, spec *cluster.Spec) (controller.Result, error)
}

type RemoteClientRegistry interface {
	GetClient(ctx context.Context, cluster client.ObjectKey) (client.Client, error)
}

// IPValidator defines an interface for the methods to validate the control plane IP.
type IPValidator interface {
	ValidateControlPlaneIP(ctx context.Context, log logr.Logger, spec *cluster.Spec) (controller.Result, error)
}

type Reconciler struct {
	client               client.Client
	cniReconciler        CNIReconciler
	remoteClientRegistry RemoteClientRegistry
	ipValidator          IPValidator
	*serverside.ObjectApplier
}

// New initializes a new reconciler for the Snow provider.
func New(client client.Client, cniReconciler CNIReconciler, remoteClientRegistry RemoteClientRegistry, ipValidator IPValidator) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) Reconcile(ctx context.Context, log logr.Logger, c *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func (r *Reconciler) ValidateMachineConfigs(ctx context.Context, log logr.Logger, clusterSpec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func (s *Reconciler) ReconcileControlPlane(ctx context.Context, log logr.Logger, clusterSpec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func (r *Reconciler) CheckControlPlaneReady(ctx context.Context, log logr.Logger, clusterSpec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func (s *Reconciler) ReconcileCNI(ctx context.Context, log logr.Logger, clusterSpec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func (s *Reconciler) ReconcileWorkers(ctx context.Context, log logr.Logger, clusterSpec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func toClientControlPlane(cp *snow.ControlPlane) *clusters.ControlPlane {
	_ = "STUB: not implemented"
	return nil
}

func toClientWorkers(workers *snow.Workers) *clusters.Workers {
	_ = "STUB: not implemented"
	return nil
}
