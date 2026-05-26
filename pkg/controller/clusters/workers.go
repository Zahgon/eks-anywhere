package clusters

import (
	"context"

	"github.com/go-logr/logr"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	"github.com/aws/eks-anywhere/pkg/controller"
)

// Workers represents the CAPI spec for an eks-a cluster's workers.
type Workers struct {
	Groups []WorkerGroup

	// Other includes any other provider-specific objects that need to be reconciled
	// as part of the worker groups.
	Other []client.Object
}

// objects returns a list of API objects for a collection of worker groups.
func (w *Workers) objects() []client.Object { _ = "STUB: not implemented"; return nil }

// WorkerGroup represents the CAPI spec for an eks-a worker group.
type WorkerGroup struct {
	KubeadmConfigTemplate   *bootstrapv1beta2.KubeadmConfigTemplate
	MachineDeployment       *clusterv1beta2.MachineDeployment
	ProviderMachineTemplate client.Object
}

func (g *WorkerGroup) objects() []client.Object { _ = "STUB: not implemented"; return nil }

// ToWorkers converts the generic clusterapi Workers definition to the concrete one defined
// here. It's just a helper for callers generating workers spec using the clusterapi package.
func ToWorkers[M clusterapi.Object[M]](capiWorkers *clusterapi.Workers[M]) *Workers {
	_ = "STUB: not implemented"
	return nil
}

// ReconcileWorkersForEKSA orchestrates the worker node reconciliation logic for a particular EKS-A cluster.
// It takes care of applying all desired objects in the Workers spec and deleting the
// old MachineDeployments that are not in it.
func ReconcileWorkersForEKSA(ctx context.Context, log logr.Logger, c client.Client, cluster *anywherev1.Cluster, w *Workers) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// cluster doesn't exist, this might be transient, requeuing

// ReconcileWorkers orchestrates the worker node reconciliation logic.
// It takes care of applying all desired objects in the Workers spec and deleting the
// old MachineDeployments that are not in it.
func ReconcileWorkers(ctx context.Context, c client.Client, cluster *clusterv1beta2.Cluster, w *Workers) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}
