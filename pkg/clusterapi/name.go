package clusterapi

import (
	"context"
	"regexp"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

var nameRegex = regexp.MustCompile(`(.*?)(-)(\d+)$`)

// Object represents a kubernetes API object.
type Object[O kubernetes.Object] interface {
	kubernetes.Object
	DeepCopy() O
}

// ObjectComparator returns true only if only both kubernetes Object's are identical
// Most of the time, this only requires comparing the Spec field, but that can variate
// from object to object.
type ObjectComparator[O Object[O]] func(current, new O) bool

// ObjectRetriever gets a kubernetes API object using the provided client
// If the object doesn't exist, it returns a NotFound error.
type ObjectRetriever[O Object[O]] func(ctx context.Context, client kubernetes.Client, name, namespace string) (O, error)

// IncrementName takes an object name and increments the suffix number by one.
// This method is used for updating objects (e.g. machinetemplate, kubeadmconfigtemplate) that are either immutable
// or require recreation to trigger machine rollout. The original object name should follow the name convention of
// alphanumeric followed by dash digits, e.g. abc-1, md-0, kct-2. An error will be raised if the original name does not follow
// this pattern.
func IncrementName(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// IncrementNameWithFallbackDefault calls the IncrementName and fallbacks to use the default name if IncrementName
// returns an error. This method is used to accommodate for any objects with name breaking changes from a previous version.
// For example, in beta capi snowmachinetemplate is named after the eks-a snowmachineconfig name, without the '-1' suffix.
// We set the object name to the default new machinetemplate name after detecting the invalid old name.
func IncrementNameWithFallbackDefault(name, defaultName string) string {
	_ = "STUB: not implemented"
	return ""
}

func ObjectName(baseName string, version int) string { _ = "STUB: not implemented"; return "" }

func DefaultObjectName(baseName string) string { _ = "STUB: not implemented"; return "" }

// KubeadmControlPlaneName generates the kubeadmControlPlane name for an EKSA Cluster.
func KubeadmControlPlaneName(cluster *v1alpha1.Cluster) string {
	_ = "STUB: not implemented"
	return ""

	// EtcdClusterName sets the default EtcdCluster object name.
}

func EtcdClusterName(clusterName string) string { _ = "STUB: not implemented"; return "" }

// MachineDeploymentName returns the name for the corresponding MachineDeployment to an EKS-A worker node group.
func MachineDeploymentName(cluster *v1alpha1.Cluster, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration) string {
	_ = "STUB: not implemented"
	// Adding cluster name prefix guarantees the machine deployment name uniqueness
	// among clusters under the same management cluster setting.
	return ""
}

func DefaultKubeadmConfigTemplateName(clusterSpec *cluster.Spec, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration) string {
	_ = "STUB: not implemented"
	return ""
}

func clusterWorkerNodeGroupName(cluster *v1alpha1.Cluster, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration) string {
	_ = "STUB: not implemented"
	return ""
}

// ControlPlaneMachineTemplateName sets the default object name on the control plane machine template.
func ControlPlaneMachineTemplateName(cluster *v1alpha1.Cluster) string {
	_ = "STUB: not implemented"
	return ""
}

// EtcdMachineTemplateName sets the default object name on the etcd machine template.
func EtcdMachineTemplateName(cluster *v1alpha1.Cluster) string {
	_ = "STUB: not implemented"
	return ""
}

func WorkerMachineTemplateName(clusterSpec *cluster.Spec, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration) string {
	_ = "STUB: not implemented"
	return ""
}

// ControlPlaneMachineHealthCheckName returns a name for a kcp machine health check.
func ControlPlaneMachineHealthCheckName(cluster *v1alpha1.Cluster) string {
	_ = "STUB: not implemented"
	return ""
}

// WorkerMachineHealthCheckName returns a name for a worker machine health check.
func WorkerMachineHealthCheckName(cluster *v1alpha1.Cluster, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration) string {
	_ = "STUB: not implemented"
	return ""
}

// InitialTemplateNamesForWorkers returns the default initial names for workers machine templates and kubeadm config templates.
func InitialTemplateNamesForWorkers(clusterSpec *cluster.Spec) (machineTemplateNames, kubeadmConfigTemplateNames map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EnsureNewNameIfChanged updates an object's name if such object is different from its current state in the cluster.
func EnsureNewNameIfChanged[M Object[M]](ctx context.Context,
	client kubernetes.Client,
	retrieve ObjectRetriever[M],
	equal ObjectComparator[M],
	new M,
) error {
	_ = "STUB: not implemented"
	return nil
}

// if object doesn't exist with same name in same namespace, no need to compare, there won't be a conflict

// ClusterCASecretName returns the name of the cluster CA secret for the cluster.
func ClusterCASecretName(clusterName string) string { _ = "STUB: not implemented"; return "" }

// ClusterKubeconfigSecretName returns the name of the kubeconfig secret for the cluster.
func ClusterKubeconfigSecretName(clusterName string) string { _ = "STUB: not implemented"; return "" }
