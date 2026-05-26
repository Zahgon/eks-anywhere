package reconciler

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
)

type CiliumReconciler interface {
	Reconcile(ctx context.Context, logger logr.Logger, client client.Client, spec *cluster.Spec) (controller.Result, error)
}

type Reconciler struct {
	ciliumReconciler CiliumReconciler
}

func New(ciliumReconciler CiliumReconciler) *Reconciler { _ = "STUB: not implemented"; return nil }

// Reconcile takes the specified CNI in a cluster to the desired state defined in a cluster Spec
// It uses a controller.Result to indicate when requeues are needed
// Intended to be used in a kubernetes controller
// Only Cilium CNI is supported for now.
func (r *Reconciler) Reconcile(ctx context.Context, logger logr.Logger, client client.Client, spec *cluster.Spec) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}
