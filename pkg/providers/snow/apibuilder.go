package snow

import (
	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	snowv1 "github.com/aws/eks-anywhere/pkg/providers/snow/api/v1beta1"
)

const (
	// SnowClusterKind is the kubernetes object kind for CAPAS Cluster.
	SnowClusterKind = "AWSSnowCluster"
	// SnowMachineTemplateKind is the kubernetes object kind for CAPAS machine template.
	SnowMachineTemplateKind = "AWSSnowMachineTemplate"
	// SnowIPPoolKind is the kubernetes object kind for CAPAS IP pool.
	SnowIPPoolKind                                   = "AWSSnowIPPool"
	ignoreEtcdKubernetesManifestFolderPreflightError = "DirAvailable--etc-kubernetes-manifests"
)

// CAPICluster generates the CAPICluster object for snow provider.
func CAPICluster(clusterSpec *cluster.Spec, snowCluster *snowv1.AWSSnowCluster, kubeadmControlPlane *controlplanev1beta2.KubeadmControlPlane, etcdCluster *etcdv1.EtcdadmCluster) *clusterv1beta2.Cluster {
	_ = "STUB: not implemented"
	return nil
}

// KubeadmControlPlane generates the kubeadmControlPlane object for snow provider from clusterSpec and snowMachineTemplate.
func KubeadmControlPlane(log logr.Logger, clusterSpec *cluster.Spec, snowMachineTemplate *snowv1.AWSSnowMachineTemplate) (*controlplanev1beta2.KubeadmControlPlane, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KubeadmConfigTemplate generates the kubeadmConfigTemplate object for snow provider from clusterSpec and workerNodeGroupConfig.
func KubeadmConfigTemplate(log logr.Logger, clusterSpec *cluster.Spec, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration) (*bootstrapv1beta2.KubeadmConfigTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func machineDeployment(clusterSpec *cluster.Spec, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration, kubeadmConfigTemplate *bootstrapv1beta2.KubeadmConfigTemplate, snowMachineTemplate *snowv1.AWSSnowMachineTemplate) *clusterv1beta2.MachineDeployment {
	_ = "STUB: not implemented"
	return nil
}

// EtcdadmCluster builds an etcdadmCluster based on an eks-a cluster spec and snowMachineTemplate.
func EtcdadmCluster(log logr.Logger, clusterSpec *cluster.Spec, snowMachineTemplate *snowv1.AWSSnowMachineTemplate) *etcdv1.EtcdadmCluster {
	_ = "STUB: not implemented"
	return nil
}

func SnowCluster(clusterSpec *cluster.Spec, credentialsSecret *v1.Secret) *snowv1.AWSSnowCluster {
	_ = "STUB: not implemented"
	return nil
}

func CredentialsSecret(name, namespace string, credsB64, certsB64 []byte) *v1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func CAPASCredentialsSecret(clusterSpec *cluster.Spec, credsB64, certsB64 []byte) *v1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func EksaCredentialsSecret(datacenter *v1alpha1.SnowDatacenterConfig, credsB64, certsB64 []byte) *v1.Secret {
	_ = "STUB: not implemented"
	return nil
}

// CAPASIPPools defines a set of CAPAS AWSSnowPool objects.
type CAPASIPPools map[string]*snowv1.AWSSnowIPPool

func (p CAPASIPPools) addPools(dnis []v1alpha1.SnowDirectNetworkInterface, m map[string]*v1alpha1.SnowIPPool) {
	_ = "STUB: not implemented"
	return
}

func buildSnowIPPool(pool v1alpha1.IPPool) snowv1.IPPool {
	_ = "STUB: not implemented"
	return *new(snowv1.IPPool)
}

func toAWSSnowIPPool(pool *v1alpha1.SnowIPPool) *snowv1.AWSSnowIPPool {
	_ = "STUB: not implemented"
	return nil
}

func buildDNI(dni v1alpha1.SnowDirectNetworkInterface, capasPools CAPASIPPools) snowv1.AWSSnowDirectNetworkInterface {
	_ = "STUB: not implemented"
	return *new(snowv1.AWSSnowDirectNetworkInterface)
}

// MachineTemplate builds a snowMachineTemplate based on an eks-a snowMachineConfig and a capasIPPool.
func MachineTemplate(name string, machineConfig *v1alpha1.SnowMachineConfig, capasPools CAPASIPPools) *snowv1.AWSSnowMachineTemplate {
	_ = "STUB: not implemented"
	return nil
}
