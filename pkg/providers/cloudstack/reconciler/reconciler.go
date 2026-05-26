package reconciler

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	c "github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack"
)

// IPValidator is an interface that defines methods to validate the control plane IP.
type IPValidator interface {
	ValidateControlPlaneIP(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error)
}

// CNIReconciler is an interface for reconciling CNI in the CloudStack cluster reconciler.
type CNIReconciler interface {
	Reconcile(ctx context.Context, logger logr.Logger, client client.Client, spec *c.Spec) (controller.Result, error)
}

// RemoteClientRegistry is an interface that defines methods for remote clients.
type RemoteClientRegistry interface {
	GetClient(ctx context.Context, cluster client.ObjectKey) (client.Client, error)
}

// Reconciler for CloudStack.
type Reconciler struct {
	client               client.Client
	ipValidator          IPValidator
	cniReconciler        CNIReconciler
	remoteClientRegistry RemoteClientRegistry
	validatorRegistry    cloudstack.ValidatorRegistry
}

// New defines a new CloudStack reconciler.
func New(client client.Client, ipValidator IPValidator, cniReconciler CNIReconciler, remoteClientRegistry RemoteClientRegistry, validatorRegistry cloudstack.ValidatorRegistry) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile reconciles cluster to desired state.
func (r *Reconciler) Reconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ValidateDatacenterConfig updates the cluster status if the CloudStackDatacenter status indicates that the spec is invalid.
func (r *Reconciler) ValidateDatacenterConfig(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ValidateMachineConfig performs additional, context-aware validations on the machine configs.
func (r *Reconciler) ValidateMachineConfig(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ReconcileControlPlane applies the control plane CAPI objects to the cluster.
func (r *Reconciler) ReconcileControlPlane(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// CheckControlPlaneReady checks whether the control plane for an eks-a cluster is ready or not.
// Requeues with the appropriate wait times whenever the control plane is not ready yet.
func (r *Reconciler) CheckControlPlaneReady(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ReconcileWorkers applies the worker CAPI objects to the cluster.
func (r *Reconciler) ReconcileWorkers(ctx context.Context, log logr.Logger, clusterSpec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ReconcileCNI reconciles the CNI to the desired state.
func (r *Reconciler) ReconcileCNI(ctx context.Context, log logr.Logger, clusterSpec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}
