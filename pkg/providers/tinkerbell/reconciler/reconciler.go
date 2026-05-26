package reconciler

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	rufiov1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell/rufio"
	c "github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/controller"
	"github.com/aws/eks-anywhere/pkg/controller/clusters"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell"
)

const (
	// NewClusterOperation indicates to create a new cluster.
	NewClusterOperation Operation = "NewCluster"
	// K8sVersionUpgradeOperation indicates to upgrade all nodes to a new Kubernetes version.
	K8sVersionUpgradeOperation Operation = "K8sVersionUpgrade"
	// NoChange indicates no change made to cluster during periodical sync.
	NoChange Operation = "NoChange"
)

// Operation indicates the desired change on a cluster.
type Operation string

// CNIReconciler is an interface for reconciling CNI in the Tinkerbell cluster reconciler.
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

// Scope object for Tinkerbell reconciler.
type Scope struct {
	ClusterSpec  *c.Spec
	ControlPlane *tinkerbell.ControlPlane
	Workers      *tinkerbell.Workers
}

// NewScope creates a new Tinkerbell Reconciler Scope.
func NewScope(clusterSpec *c.Spec) *Scope { _ = "STUB: not implemented"; return nil }

// Reconciler for Tinkerbell.
type Reconciler struct {
	client               client.Client
	cniReconciler        CNIReconciler
	remoteClientRegistry RemoteClientRegistry
	ipValidator          IPValidator
}

// New defines a new Tinkerbell reconciler.
func New(client client.Client, cniReconciler CNIReconciler, remoteClientRegistry RemoteClientRegistry, ipValidator IPValidator) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile reconciles cluster to desired state.
func (r *Reconciler) Reconcile(ctx context.Context, log logr.Logger, cluster *anywherev1.Cluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	// Implement reconcile all here.
	// This would include validating machine and datacenter configs
	// and reconciling cp and worker nodes.
	return *new(controller.Result), nil
}

// ValidateControlPlaneIP passes the cluster spec from tinkerbellScope to the IP Validator.
func (r *Reconciler) ValidateControlPlaneIP(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// CleanupStatusAfterValidate removes errors from the cluster status with the tinkerbellScope.
func (r *Reconciler) CleanupStatusAfterValidate(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ValidateClusterSpec performs a set of assertions on a cluster spec.
func (r *Reconciler) ValidateClusterSpec(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// GenerateSpec generates Tinkerbell control plane and workers spec.
func (r *Reconciler) GenerateSpec(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// DetectOperation detects change type.
func (r *Reconciler) DetectOperation(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

// The restriction that not allowing scaling and rolling is covered in webhook.

// ReconcileControlPlane applies the control plane CAPI objects to the cluster.
func (r *Reconciler) ReconcileControlPlane(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// CheckControlPlaneReady checks whether the control plane for an eks-a cluster is ready or not.
// Requeues with the appropriate wait times whenever the cluster is not ready yet.
func (r *Reconciler) CheckControlPlaneReady(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ReconcileWorkers applies the worker CAPI objects to the cluster.
func (r *Reconciler) ReconcileWorkers(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ValidateDatacenterConfig updates the cluster status if the TinkerbellDatacenter status indicates that the spec is invalid.
func (r *Reconciler) ValidateDatacenterConfig(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// ReconcileCNI reconciles the CNI to the desired state.
func (r *Reconciler) ReconcileCNI(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

func (r *Reconciler) validateTinkerbellIPMatch(ctx context.Context, clusterSpec *c.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// for workload cluster tinkerbell IP must match management cluster tinkerbell IP

func toClientControlPlane(cp *tinkerbell.ControlPlane) *clusters.ControlPlane {
	_ = "STUB: not implemented"
	return nil
}

// ValidateHardware performs a set of validations on the tinkerbell hardware read from the cluster.
func (r *Reconciler) ValidateHardware(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// We need a new reader each time so that the catalogue gets recreated.

// skip extra hardware validation for InPlace upgrades

// eksa version upgrade cannot be triggered from controller, so set it to false.

// Hardware selectors for controlPlane and worker nodes are mutually exclusive, so its safe to copy
// as no keys are going to be overwritten

func (r *Reconciler) getValidatableCAPI(ctx context.Context, cluster *anywherev1.Cluster) (*tinkerbell.ValidatableTinkerbellCAPI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateHardwareReqForKCP returns minium hardware requirements for the KCP to rollout new control plane nodes
// CAPI rolls out a new control-plane node whenever the associated MachineTemplate changes in the kcp object
// There will be no rollout if the template stays the same.
func (r *Reconciler) validateHardwareReqForKCP(validatableCAPI *tinkerbell.ValidatableTinkerbellCAPI, tinkerbellScope *Scope) (tinkerbell.MinimumHardwareRequirements, error) {
	_ = "STUB: not implemented"
	return *new(tinkerbell.MinimumHardwareRequirements), nil
}

// validateHardwareReqForMachineDeployments returns minium hardware requirements for the md's to rollout new worker nodes
// CAPI rolls out a new worker node only whenever the associated MachineTemplate changes in the md object
// A single cluster can have multiple MachineDeployment objects and in case of modular upgrades
// only few of those worker groups might need a rollout.
func (r *Reconciler) validateHardwareReqForMachineDeployments(ctx context.Context, tinkerbellScope *Scope) (requirements tinkerbell.MinimumHardwareRequirements, err error) {
	_ = "STUB: not implemented"
	return *new(tinkerbell.MinimumHardwareRequirements), nil
}

// EKS-A names MachineDeployment with the clusterName prefix followed by the WorkerNodeGroup name provider concatenated by '-'
// We just need the workerNodeGroup name to fetch the corresponding workerNodeGroup config from the spec

// ValidateRufioMachines checks to ensure all the Rufio machines condition contactable is True.
func (r *Reconciler) ValidateRufioMachines(ctx context.Context, log logr.Logger, tinkerbellScope *Scope) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// Skip contactability check if the machine has the skip label set to "true"

func (r *Reconciler) checkContactable(rm *rufiov1alpha1.Machine) error {
	_ = "STUB: not implemented"
	return nil
}
