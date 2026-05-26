package clusters

import (
	"context"

	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	"github.com/go-logr/logr"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/pkg/controller"
)

// ControlPlane represents a CAPI spec for a kubernetes cluster.
type ControlPlane struct {
	Cluster *clusterv1beta2.Cluster

	// ProviderCluster is the provider-specific resource that holds the details
	// for provisioning the infrastructure, referenced in Cluster.Spec.InfrastructureRef
	ProviderCluster client.Object

	KubeadmControlPlane *controlplanev1beta2.KubeadmControlPlane

	// ControlPlaneMachineTemplate is the provider-specific machine template referenced
	// in KubeadmControlPlane.Spec.MachineTemplate.InfrastructureRef
	ControlPlaneMachineTemplate client.Object

	EtcdCluster *etcdv1.EtcdadmCluster

	// EtcdMachineTemplate is the provider-specific machine template referenced
	// in EtcdCluster.Spec.InfrastructureTemplate
	EtcdMachineTemplate client.Object

	// Other includes any other provider-specific objects that need to be reconciled
	// as part of the control plane.
	Other []client.Object
}

// AllObjects returns all the control plane objects.
func (cp *ControlPlane) AllObjects() []client.Object { _ = "STUB: not implemented"; return nil }

func (cp *ControlPlane) etcdObjects() []client.Object { _ = "STUB: not implemented"; return nil }

func (cp *ControlPlane) nonEtcdObjects() []client.Object { _ = "STUB: not implemented"; return nil }

// skipCAPIAutoPauseKCPForExternalEtcdAnnotation instructs the CAPI cluster controller to not pause or unpause
// the KCP to wait for etcd endpoints to be ready. When this annotation is present, is left to the user (us)
// to orchestrate this operation if double kcp rollouts are undesirable.
const skipCAPIAutoPauseKCPForExternalEtcdAnnotation = "cluster.x-k8s.io/skip-pause-cp-managed-etcd"

// ReconcileControlPlane orchestrates the ControlPlane reconciliation logic.
func ReconcileControlPlane(ctx context.Context, log logr.Logger, c client.Client, cp *ControlPlane) (controller.Result, error) {
	_ = "STUB: not implemented"
	return *new(controller.Result), nil
}

// always add skip pause annotation since we want to have full control over the kcp-etcd orchestration

// If the CAPI cluster doesn't exist, this is a new cluster, create all objects, no need for extra orchestration.

// If the control plane endpoint is not set in the desired cluster, we want to keep the current one.
// In practice, this condition will always be hit because:
// * We don't set the endpoint in the cluster object in our code, we let CAPI do that
// * The endpoint never changes once the cluster has been created

// For stacked etcd, we don't need orchestration, apply directly

// If there are changes for etcd, we only apply those changes for now and we wait.

// If etcd is not ready yet, we requeue and wait before making any other change to the control plane.

// We need to inject a logger in this method or extract from context

func readCurrentControlPlane(ctx context.Context, c client.Client, cp *ControlPlane) (*clusterv1beta2.Cluster, *controlplanev1beta2.KubeadmControlPlane, *etcdv1.EtcdadmCluster, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// If the CAPI cluster doesn't exist, this is a new cluster, no need to read the rest of the objects.

func applyAllControlPlaneObjects(ctx context.Context, c client.Client, cp *ControlPlane) error {
	_ = "STUB: not implemented"
	return nil
}

func reconcileEtcdChanges(ctx context.Context, log logr.Logger, c client.Client, desiredCP *ControlPlane, currentKCP *controlplanev1beta2.KubeadmControlPlane, currentEtcdadmCluster *etcdv1.EtcdadmCluster) (controller.Result, error) {
	_ = "STUB: not implemented"
	// Before making any changes to etcd, pause the KCP so it doesn't rollout new nodes as the
	// etcd endpoints change.
	return *new(controller.Result), nil
}

// If the etcdadm cluster has changes, this will require a rolling upgrade
// Mark the etcdadm cluster as upgrading
// The etcdadm cluster controller will take care of removing
// this annotation at the right time to orchestrate the kcp upgrade.

// After applying etcd changes, just requeue to wait until etcd finishes updating and is ready
// We use a short wait here just in case etcdadm controller decides that not new machines are
// needed.

func etcdadmClusterReady(etcdadmCluster *etcdv1.EtcdadmCluster) bool {
	_ = "STUB: not implemented"
	// It's important to use status.Ready and not the Ready condition, since the Ready condition
	// only becomes true after the old etcd members have been deleted, which only happens after the
	// kcp finishes its own upgrade.
	return false
}

func reconcileControlPlaneNodeChanges(ctx context.Context, log logr.Logger, c client.Client, desiredCP *ControlPlane, currentKCP *controlplanev1beta2.KubeadmControlPlane) (controller.Result, error) {
	_ = "STUB: not implemented"
	// When the controller reconciles the control plane for a cluster with an external etcd configuration
	// the KubeadmControlPlane.Spec.KubeadmConfigSpec.ClusterConfiguration.Etcd.External.Endpoints field is
	// defaulted to a placeholder value. At some point that field in KubeadmControlPlane object is filled
	// and updated by the kcp controller with real etcd endpoints.
	//
	// We do not want to overwrite real endpoints with the placeholder again, so here we check if the endpoints
	// for the external etcd have already been populated with real values on the KubeadmControlPlane object
	// and preserve them in the desired state before applying.
	return *new(controller.Result), nil
}

// If the KCP is paused, we read the last version (in case we just updated it) and unpause it
// so the cp nodes are reconciled.

// isPlaceholderEndpoint checks if endpoints are placeholder values that we set by default.
func isPlaceholderEndpoint(endpoints []string) bool { _ = "STUB: not implemented"; return false }

func getEtcdadmCluster(ctx context.Context, c client.Client, cluster *clusterv1beta2.Cluster) (*etcdv1.EtcdadmCluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
