package controllers

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta1"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	v1beta1patch "sigs.k8s.io/cluster-api/util/deprecated/v1beta1/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	controlPlaneLabel = "node-role.kubernetes.io/control-plane"
	podDNEMessage     = "Upgrader pod does not exist"

	// nodeUpgradeFinalizerName is the finalizer added to NodeUpgrade objects to handle deletion.
	nodeUpgradeFinalizerName = "nodeupgrades.anywhere.eks.amazonaws.com/finalizer"
)

// RemoteClientRegistry defines methods for remote cluster controller clients.
type RemoteClientRegistry interface {
	GetClient(ctx context.Context, cluster client.ObjectKey) (client.Client, error)
}

// NodeUpgradeReconciler reconciles a NodeUpgrade object.
type NodeUpgradeReconciler struct {
	client               client.Client
	log                  logr.Logger
	remoteClientRegistry RemoteClientRegistry
}

// NewNodeUpgradeReconciler returns a new instance of NodeUpgradeReconciler.
func NewNodeUpgradeReconciler(client client.Client, remoteClientRegistry RemoteClientRegistry) *NodeUpgradeReconciler {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeUpgradeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=nodeupgrades,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=nodeupgrades/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=nodeupgrades/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;delete
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch
//+kubebuilder:rbac:groups="cluster.x-k8s.io",resources=machines,verbs=list;watch;get;patch;update

// Reconcile reconciles a NodeUpgrade object.
// nolint:gocyclo
func (r *NodeUpgradeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	// TODO(in-place): Add validating webhook to block updating the nodeUpgrade object.
	// It should be immutable. If it needs to be changed, a new spec should be applied.
	return *new(ctrl.Result), nil
}

// Initialize the patch helper

// Always attempt to patch the object and status after each reconciliation.

// We want the observedGeneration to indicate, that the status shown is up-to-date given the desired spec of the same generation.
// However, if there is an error while updating the status, we may get a partial status update, In this case,
// a partially updated status is not considered up to date, so we should not update the observedGeneration

// Patch ObservedGeneration only if the reconciliation completed without error

// Only requeue if we are not already re-queueing and the NodeUpgrade ready condition is false.
// We do this to be able to update the status continuously until the NodeUpgrade becomes ready,
// since there might be changes in state of the world that don't trigger reconciliation requests

// Reconcile the NodeUpgrade deletion if the DeletionTimestamp is set.

func (r *NodeUpgradeReconciler) reconcile(ctx context.Context, log logr.Logger, machineToBeUpgraded *clusterv1beta2.Machine, nodeUpgrade *anywherev1.NodeUpgrade, remoteClient client.Client) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// return early if node upgrade is already complete.

// namespaceOrCreate creates a namespace if it doesn't already exist.
func namespaceOrCreate(ctx context.Context, client client.Client, log logr.Logger, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeUpgradeReconciler) reconcileDelete(ctx context.Context, log logr.Logger, nodeUpgrade *anywherev1.NodeUpgrade, nodeName string, remoteClient client.Client) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// TODO(in-place): Make pod deletion logic more robust by checking if the pod is still running.
// If it is still running and not errored out, then wait before deleting the pod.

// Remove the finalizer from NodeUpgrade object

func (r *NodeUpgradeReconciler) updateStatus(ctx context.Context, log logr.Logger, remoteClient client.Client, nodeUpgrade *anywherev1.NodeUpgrade, nodeName string) error {
	_ = "STUB: not implemented"
	// When NodeUpgrade is fully deleted, we do not need to update the status. Without this check
	// the subsequent patch operations would fail if the status is updated after it is fully deleted.
	return nil
}

// Always update the readyCondition by summarizing the state of other conditions.

func updateComponentsConditions(pod *corev1.Pod, nodeUpgrade *anywherev1.NodeUpgrade) {
	_ = "STUB: not implemented"
	return
}

// this should not happen

func getContainerStatus(pod *corev1.Pod, containerName string) (*corev1.ContainerStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func markAllConditionsFalse(nodeUpgrade *anywherev1.NodeUpgrade, message string, severity clusterv1.ConditionSeverity) {
	_ = "STUB: not implemented"
	return
}

func isControlPlane(node *corev1.Node) bool { _ = "STUB: not implemented"; return false }

// GetNamespacedNameType takes name and namespace and returns NamespacedName in namespace/name format.
func GetNamespacedNameType(name, namespace string) types.NamespacedName {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName)
}

func patchNodeUpgrade(ctx context.Context, patchHelper *v1beta1patch.Helper, nodeUpgrade anywherev1.NodeUpgrade, patchOpts ...v1beta1patch.Option) error {
	_ = "STUB: not implemented"
	// Patch the object, ignoring conflicts on the conditions owned by this controller.
	return nil
}

// Add each condition her that the controller should ignored conflicts for.

// Always attempt to patch the object and status after each reconciliation.

func upgraderPodExists(ctx context.Context, remoteClient client.Client, nodeName string) bool {
	_ = "STUB: not implemented"
	return false
}

func getUpgraderPod(ctx context.Context, remoteClient client.Client, nodeName string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNodeUpgrade(ctx context.Context, remoteClient client.Client, nodeUpgradeName string) (*anywherev1.NodeUpgrade, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nodeUpgradeName returns the name of the node upgrade object based on the machine reference.
func nodeUpgraderName(machineRefName string) string { _ = "STUB: not implemented"; return "" }
