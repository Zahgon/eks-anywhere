package clusters

import (
	"context"

	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	"github.com/go-logr/logr"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// UpdateClusterStatusForControlPlane checks the current state of the Cluster's control plane and updates the
// Cluster status information.
// There is a possibility that UpdateClusterStatusForControlPlane does not update the
// controlplane status specially in case where it is still waiting for cluster objects to be created.
func UpdateClusterStatusForControlPlane(ctx context.Context, client client.Client, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateClusterStatusForWorkers checks the current state of the Cluster's workers and updates the
// Cluster status information.
func UpdateClusterStatusForWorkers(ctx context.Context, client client.Client, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateClusterStatusForCNI updates the Cluster status for the default cni before the control plane is ready. The CNI reconciler
// handles the rest of the logic for determining the condition and updating the status based on the current state of the cluster.
func UpdateClusterStatusForCNI(ctx context.Context, cluster *anywherev1.Cluster) {
	_ = "STUB: not implemented"
	// Here, we want to initialize the DefaultCNIConfigured condition only when the condition does not exist,
	// such as in the event of cluster creation. In this case, when the control plane is not ready, we can assume
	// the CNI is not ready yet.
	return
}

// UpdateClusterCertificateStatus updates the cluster status with the certificate information
// about cluster machines such as control plane and external etcd machines. It will only update
// if the cluster is ready to avoid unncessary TLS connections.
func UpdateClusterCertificateStatus(ctx context.Context, client client.Client, log logr.Logger, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// updateConditionsForEtcdAndControlPlane updates the ControlPlaneReady condition if etcdadm cluster is not ready.
func updateConditionsForEtcdAndControlPlane(cluster *anywherev1.Cluster, kcp *controlplanev1beta2.KubeadmControlPlane, etcdadmCluster *etcdv1.EtcdadmCluster) {
	_ = "STUB: not implemented"
	// Make sure etcd cluster is ready before marking ControlPlaneReady status to true
	// This condition happens while creating a workload cluster from the management cluster using controller
	// where it tries to get the etcdadm cluster for the first time before it generates the resources.
	return
}

// Make sure etcd machine is ready before marking ControlPlaneReady status to true

// updateControlPlaneReadyCondition updates the ControlPlaneReady condition, after checking the state of the control plane
// in the cluster.
func updateControlPlaneReadyCondition(cluster *anywherev1.Cluster, kcp *controlplanev1beta2.KubeadmControlPlane) {
	_ = "STUB: not implemented"
	return
}

// We make sure to check that the status is up to date before using it

// The control plane should be marked ready when the count specified in the spec is
// equal to the ready number of nodes in the cluster and they're all of the right version specified.

// First, in the case of a rolling upgrade, we get the number of outdated nodes, and as long as there are some,
// we want to reflect in the message that the Cluster is in progress updating the old nodes with the
// new machine spec.

// Then, we check that the number of nodes in the cluster match the expected amount. If not, we
// mark that the Cluster is scaling up or scale down the control plane replicas to the expected amount.

// We check the condition signifying the overall health of the control plane components. Usually, the control plane should be healthy
// at this point but if that is not the case, we report it as an error.

// We check for the Available condition on the kubeadm control plane as a final validation. Usually, the kcp objects
// should be available at this point but if that is not the case, we report it as an error.

// updateControlPlaneInitializedCondition updates the ControlPlaneInitialized condition if it hasn't already been set.
// This condition should be set only once.
func updateControlPlaneInitializedCondition(cluster *anywherev1.Cluster, kcp *controlplanev1beta2.KubeadmControlPlane) {
	_ = "STUB: not implemented"
	// Return early if the ControlPlaneInitializedCondition is already "True"
	return
}

// We make sure to check that the status is up to date before using it

// Then, we'll check explicitly for that the control plane is available. This way, we do not rely on CAPI
// to implicitly to fill out our v1beta1conditions reasons, and we can have custom messages.

// updateWorkersReadyCondition updates the WorkersReadyCondition condition after checking the state of the worker node groups
// in the cluster.
func updateWorkersReadyCondition(cluster *anywherev1.Cluster, machineDeployments []clusterv1beta2.MachineDeployment) {
	_ = "STUB: not implemented"
	return
}

// We want to consider only the worker node groups which don't have autoscaling configuration for expected worker nodes count.

// First, we need to aggregate the number of nodes across worker node groups to be able to assess the condition of the workers
// as a whole.

// We make sure to check that the status is up to date before using the information from the machine deployment status.

// Skip updating the replicas for the machine deployments which have autoscaling configuration annotation

// There may be worker nodes that are not up to date yet in the case of a rolling upgrade,
// so reflect that on the condition with an appropriate message.

// We are checking the control plane configuration here because we already validate that all the machines
// have the same upgrade strategy.

// If the number of worker nodes replicas need to be scaled up.

// If the number of worker nodes replicas need to be scaled down.

// Iterating through the machine deployments which have autoscaling configured to check if the number of worker nodes replicas
// are between min count and max count specified in the cluster spec.

// We check for the Ready condition on the machine deployments as a final validation. Usually, the md objects
// should be ready at this point but if that is not the case, we report it as an error.

// controlPlaneInitializationInProgressCondition returns a new "False" condition for the ControlPlaneInitializationInProgress reason.
func controlPlaneInitializationInProgressCondition() *anywherev1.Condition {
	_ = "STUB: not implemented"
	return nil
}
