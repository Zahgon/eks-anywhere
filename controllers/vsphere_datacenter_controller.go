package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/providers/vsphere"
)

// VSphereDatacenterReconciler reconciles a VSphereDatacenterConfig object.
type VSphereDatacenterReconciler struct {
	client    client.Client
	defaulter *vsphere.Defaulter
	validator *vsphere.Validator
}

// NewVSphereDatacenterReconciler constructs a new VSphereDatacenterReconciler.
func NewVSphereDatacenterReconciler(client client.Client, validator *vsphere.Validator, defaulter *vsphere.Defaulter) *VSphereDatacenterReconciler {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VSphereDatacenterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: add here kubebuilder permissions as neeeded.
// Reconcile implements the reconcile.Reconciler interface.
func (r *VSphereDatacenterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Fetch the VsphereDatacenter object

// Initialize the patch helper

// Always attempt to patch the object and status after each reconciliation.

// There's no need to go any further if the VsphereDatacenterConfig is marked for deletion.

func (r *VSphereDatacenterReconciler) reconcile(ctx context.Context, vsphereDatacenter *anywherev1.VSphereDatacenterConfig, log logr.Logger) (_ ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	// Set up envs for executing Govc cmd and default values for datacenter config
	return *new(ctrl.Result), nil
}

// Determine if VsphereDatacenterConfig is valid

func (r *VSphereDatacenterReconciler) reconcileDelete(ctx context.Context, vsphereDatacenter *anywherev1.VSphereDatacenterConfig, log logr.Logger) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}
