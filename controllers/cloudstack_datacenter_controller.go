package controllers

import (
	"context"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack"
)

// CloudStackDatacenterReconciler reconciles a CloudStackDatacenterConfig object.
type CloudStackDatacenterReconciler struct {
	client            client.Client
	validatorRegistry cloudstack.ValidatorRegistry
}

// NewCloudStackDatacenterReconciler creates a new instance of the CloudStackDatacenterReconciler struct.
func NewCloudStackDatacenterReconciler(client client.Client, validatorRegistry cloudstack.ValidatorRegistry) *CloudStackDatacenterReconciler {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CloudStackDatacenterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile implements the reconcile.Reconciler interface.
func (r *CloudStackDatacenterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Fetch the CloudStackDatacenter object

// Initialize the patch helper

// Always attempt to patch the object and status after each reconciliation.

// There's no need to go any further if the cloudstackDatacenter is marked for deletion.

func (r *CloudStackDatacenterReconciler) reconcile(ctx context.Context, cloudstackDatacenterConfig *anywherev1.CloudStackDatacenterConfig, log logr.Logger) (_ ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Run validations with validator as Get will construct CMK each time
