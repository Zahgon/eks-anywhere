package controller

import (
	"context"

	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// GetCAPICluster reads a cluster-api Cluster for an eks-a cluster using a kube client
// If the CAPI cluster is not found, the method returns (nil, nil).
func GetCAPICluster(ctx context.Context, client client.Client, cluster *anywherev1.Cluster) (*clusterv1beta2.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CapiClusterObjectKey generates an ObjectKey for the CAPI cluster owned by
// the provided eks-a cluster.
func CapiClusterObjectKey(cluster *anywherev1.Cluster) client.ObjectKey {
	_ = "STUB: not implemented"
	// TODO: we should consider storing a reference to the CAPI cluster in the eksa cluster status
	return *new(client.ObjectKey)
}

// GetKubeadmControlPlane reads a cluster-api KubeadmControlPlane for an eks-a cluster using a kube client
// If the KubeadmControlPlane is not found, the method returns (nil, nil).
func GetKubeadmControlPlane(ctx context.Context, client client.Client, cluster *anywherev1.Cluster) (*controlplanev1beta2.KubeadmControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KubeadmControlPlane reads a cluster-api KubeadmControlPlane for an eks-a cluster using a kube client.
func KubeadmControlPlane(ctx context.Context, client client.Client, cluster *anywherev1.Cluster) (*controlplanev1beta2.KubeadmControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CAPIKubeadmControlPlaneKey generates an ObjectKey for the CAPI Kubeadm control plane owned by
// the provided eks-a cluster.
func CAPIKubeadmControlPlaneKey(cluster *anywherev1.Cluster) client.ObjectKey {
	_ = "STUB: not implemented"
	return *new(client.ObjectKey)
}

// GetMachineDeployment reads a cluster-api MachineDeployment for an eks-a cluster using a kube client.
// If the MachineDeployment is not found, the method returns (nil, nil).
func GetMachineDeployment(ctx context.Context, client client.Client, machineDeploymentName string) (*clusterv1beta2.MachineDeployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMachineDeployments reads all of cluster-api MachineDeployment for an eks-a cluster using a kube client.
func GetMachineDeployments(ctx context.Context, c client.Client, cluster *anywherev1.Cluster) ([]clusterv1beta2.MachineDeployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
