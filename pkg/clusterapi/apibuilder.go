package clusterapi

import (
	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

const (
	clusterKind               = "Cluster"
	kubeadmControlPlaneKind   = "KubeadmControlPlane"
	etcdadmClusterKind        = "EtcdadmCluster"
	kubeadmConfigTemplateKind = "KubeadmConfigTemplate"
	machineDeploymentKind     = "MachineDeployment"
	EKSAClusterLabelName      = "cluster.anywhere.eks.amazonaws.com/cluster-name"
	EKSAClusterLabelNamespace = "cluster.anywhere.eks.amazonaws.com/cluster-namespace"
)

var (
	clusterAPIVersion             = clusterv1beta2.GroupVersion.String()
	machineDeploymentAPIVersion   = clusterv1beta2.GroupVersion.String()
	machineHealthCheckAPIVersion  = clusterv1beta2.GroupVersion.String()
	kubeadmControlPlaneAPIVersion = controlplanev1beta2.GroupVersion.String()
	bootstrapAPIVersion           = bootstrapv1beta2.GroupVersion.String()
	etcdAPIVersion                = etcdv1.GroupVersion.String()
)

type APIObject interface {
	runtime.Object
	GetName() string
}

func InfrastructureAPIVersion() string { _ = "STUB: not implemented"; return "" }

func eksaClusterLabels(clusterSpec *cluster.Spec) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func capiClusterLabel(clusterSpec *cluster.Spec) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func capiObjectLabels(clusterSpec *cluster.Spec) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func mergeLabels(labels ...map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// ClusterName generates the CAPI cluster name for an EKSA Cluster.
func ClusterName(cluster *anywherev1.Cluster) string { _ = "STUB: not implemented"; return "" }

// Cluster builds a CAPI Cluster based on an eks-a cluster spec, infrastructureObject, controlPlaneObject and unstackedEtcdObject.
func Cluster(clusterSpec *cluster.Spec, infrastructureObject, controlPlaneObject, unstackedEtcdObject APIObject) *clusterv1beta2.Cluster {
	_ = "STUB: not implemented"
	return nil
}

func KubeadmControlPlane(clusterSpec *cluster.Spec, infrastructureObject APIObject) (*controlplanev1beta2.KubeadmControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KubeadmConfigTemplate(clusterSpec *cluster.Spec, workerNodeGroupConfig anywherev1.WorkerNodeGroupConfiguration) (*bootstrapv1beta2.KubeadmConfigTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MachineDeployment builds a machineDeployment based on an eks-a cluster spec, workerNodeGroupConfig, bootstrapObject and infrastructureObject.
func MachineDeployment(clusterSpec *cluster.Spec, workerNodeGroupConfig anywherev1.WorkerNodeGroupConfiguration, bootstrapObject, infrastructureObject APIObject) *clusterv1beta2.MachineDeployment {
	_ = "STUB: not implemented"
	return nil
}

// EtcdadmCluster builds a etcdadmCluster based on an eks-a cluster spec and infrastructureTemplate.
func EtcdadmCluster(clusterSpec *cluster.Spec, infrastructureTemplate APIObject) *etcdv1.EtcdadmCluster {
	_ = "STUB: not implemented"
	return nil
}
