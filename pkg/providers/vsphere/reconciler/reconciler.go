package reconciler

import (
	"context"

	"github.com/go-logr/logr"
	apiv1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	c "github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
	"github.com/aws/eks-anywhere/pkg/controller/clusters"
	"github.com/aws/eks-anywhere/pkg/controller/serverside"
	"github.com/aws/eks-anywhere/pkg/providers/vsphere"
)

// CNIReconciler is an interface for reconciling CNI in the VSphere cluster reconciler.
type CNIReconciler interface {
	Reconcile(ctx context.Context, logger logr.Logger, client client.Client, spec *c.Spec) (controller.Result, error)
}

// RemoteClientRegistry is an interface that defines methods for remote clients.
type RemoteClientRegistry interface {
	GetClient(ctx context.Context, cluster client.ObjectKey) (client.Client, error)
}

// IPValidator is an interface that defines methods to validate the control plane IP.
type IPValidator interface {
	ValidateControlPlaneIP(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error)
}

type Reconciler struct {
	client               client.Client
	validator            *vsphere.Validator
	defaulter            *vsphere.Defaulter
	cniReconciler        CNIReconciler
	remoteClientRegistry RemoteClientRegistry
	ipValidator          IPValidator
	*serverside.ObjectApplier
}

// New defines a new VSphere reconciler.
func New(client client.Client, validator *vsphere.Validator, defaulter *vsphere.Defaulter, cniReconciler CNIReconciler, remoteClientRegistry RemoteClientRegistry, ipValidator IPValidator) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

func VsphereCredentials(ctx context.Context, cli client.Client) (*apiv1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetupEnvVars(ctx context.Context, vsphereDatacenter *anywherev1.VSphereDatacenterConfig, cli client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) Reconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ValidateDatacenterConfig updates the cluster status if the VSphereDatacenter status indicates that the spec is invalid.
func (r *Reconciler) ValidateDatacenterConfig(ctx context.Context, log logr.Logger, clusterSpec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ValidateMachineConfigs performs additional, context-aware validations on the machine configs.
func (r *Reconciler) ValidateMachineConfigs(ctx context.Context, log logr.Logger, clusterSpec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// Set up env vars for executing Govc cmd

// ValidateFailureDomains performs validations for the provided failure domains and the assigned failure domains in worker node group.
func (r *Reconciler) ValidateFailureDomains(ctx context.Context, log logr.Logger, clusterSpec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ReconcileFailureDomains applies the Vsphere FailureDomain objects to the cluster.
// It also takes care of  deleting the old Vsphere FailureDomains that are not in in the cluster spec anymore.
func (r *Reconciler) ReconcileFailureDomains(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// CAPV by default deletes VSphereFailureDomain if its corresponding VSphereDeploymentZone is deleted
// Hence we are only deleting the VsphereDeploymentZones
// https://github.com/kubernetes-sigs/cluster-api-provider-vsphere/blob/d7af9d7199f21b442402714ca00d0a9801237f5d/controllers/vspheredeploymentzone_controller.go#L274.

// ReconcileControlPlane applies the control plane CAPI objects to the cluster.
func (r *Reconciler) ReconcileControlPlane(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// CheckControlPlaneReady checks whether the control plane for an eks-a cluster is ready or not.
// Requeues with the appropriate wait times whenever the cluster is not ready yet.
func (r *Reconciler) CheckControlPlaneReady(ctx context.Context, log logr.Logger, clusterSpec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ReconcileCNI takes the Cilium CNI in a cluster to the desired state defined in a cluster spec.
func (r *Reconciler) ReconcileCNI(ctx context.Context, log logr.Logger, clusterSpec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ReconcileWorkers applies the worker CAPI objects to the cluster.
func (r *Reconciler) ReconcileWorkers(ctx context.Context, log logr.Logger, spec *c.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func toClientControlPlane(cp *vsphere.ControlPlane) *clusters.ControlPlane {
	_ = "STUB: not implemented"
	return nil
}
