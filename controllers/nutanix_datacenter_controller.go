package controllers

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/providers/nutanix"
)

// NutanixDatacenterReconciler reconciles a NutanixDatacenterConfig object.
type NutanixDatacenterReconciler struct {
	client    client.Client
	defaulter *nutanix.Defaulter
}

// Reconcile reconciles a NutanixDatacenterConfig object.
func (r *NutanixDatacenterReconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (r *NutanixDatacenterReconciler) reconcileDelete(ctx context.Context, dc *anywherev1.NutanixDatacenterConfig) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// NewNutanixDatacenterReconciler constructs a new NutanixDatacenterReconciler.
func NewNutanixDatacenterReconciler(client client.Client, defaulter *nutanix.Defaulter) *NutanixDatacenterReconciler {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NutanixDatacenterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
