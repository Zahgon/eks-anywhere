package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type Validator interface {
	ValidateEC2SshKeyNameExists(ctx context.Context, m *anywherev1.SnowMachineConfig) error
	ValidateEC2ImageExistsOnDevice(ctx context.Context, m *anywherev1.SnowMachineConfig) error
}

// SnowMachineConfigReconciler reconciles a SnowMachineConfig object.
type SnowMachineConfigReconciler struct {
	client    client.Client
	validator Validator
}

// NewSnowMachineConfigReconciler constructs a new SnowMachineConfigReconciler.
func NewSnowMachineConfigReconciler(client client.Client, validator Validator) *SnowMachineConfigReconciler {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SnowMachineConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: add here kubebuilder permissions as needed.
// Reconcile implements the reconcile.Reconciler interface.
func (r *SnowMachineConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Fetch the SnowMachineConfig object

// Initialize the patch helper

// Always attempt to patch the object and status after each reconciliation.

// There's no need to go any further if the SnowMachineConfig is marked for deletion.

func (r *SnowMachineConfigReconciler) reconcile(ctx context.Context, snowMachineConfig *anywherev1.SnowMachineConfig) (_ ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}
