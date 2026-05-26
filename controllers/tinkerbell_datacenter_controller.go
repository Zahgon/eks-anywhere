package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TinkerbellDatacenterReconciler reconciles a TinkerbellDatacenterConfig object.
type TinkerbellDatacenterReconciler struct {
	client client.Client
}

// NewTinkerbellDatacenterReconciler creates a new instance of the TinkerbellDatacenterReconciler struct.
func NewTinkerbellDatacenterReconciler(client client.Client) *TinkerbellDatacenterReconciler {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TinkerbellDatacenterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: add here kubebuilder permissions as neeeded.

// Reconcile implements the reconcile.Reconciler interface.
func (r *TinkerbellDatacenterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	// TODO fetch Tinkerbell datacenter object and implement reconcile
	return *new(ctrl.Result), nil
}
