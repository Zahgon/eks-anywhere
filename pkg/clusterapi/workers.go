package clusterapi

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// Workers represents the provider specific CAPI spec for an eks-a cluster's workers.
type Workers[M Object[M]] struct {
	Groups []WorkerGroup[M]
}

// WorkerObjects returns a list of API objects for concrete provider-specific collection of worker groups.
func (w *Workers[M]) WorkerObjects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// UpdateImmutableObjectNames checks if any immutable objects have changed by comparing the new definition
// with the current state of the cluster. If they had, it generates a new name for them by increasing a monotonic number
// at the end of the name.
func (w *Workers[M]) UpdateImmutableObjectNames(
	ctx context.Context,
	client kubernetes.Client,
	machineTemplateRetriever ObjectRetriever[M],
	machineTemplateComparator ObjectComparator[M],
) error {
	_ = "STUB: not implemented"
	return nil
}

// WorkerGroup represents the provider specific CAPI spec for an eks-a worker group.
type WorkerGroup[M Object[M]] struct {
	KubeadmConfigTemplate   *bootstrapv1beta2.KubeadmConfigTemplate
	MachineDeployment       *clusterv1beta2.MachineDeployment
	ProviderMachineTemplate M
}

// Objects returns a list of API objects for a provider-specific of the worker group.
func (g *WorkerGroup[M]) Objects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// UpdateImmutableObjectNames checks if any immutable objects have changed by comparing the new definition
// with the current state of the cluster. If they had, it generates a new name for them by increasing a monotonic number
// at the end of the name.
// This process is performed to the provider machine template and the kubeadmconfigtemplate.
// The kubeadmconfigtemplate is not immutable at the API level but we treat it as such for consistency.
func (g *WorkerGroup[M]) UpdateImmutableObjectNames(
	ctx context.Context,
	client kubernetes.Client,
	machineTemplateRetriever ObjectRetriever[M],
	machineTemplateComparator ObjectComparator[M],
) error {
	_ = "STUB: not implemented"
	return nil
}

// MachineDeployment doesn't exist, this is a new cluster so machine templates should use their default name

// DeepCopy generates a new WorkerGroup copying the contexts of the receiver.
func (g *WorkerGroup[M]) DeepCopy() *WorkerGroup[M] { _ = "STUB: not implemented"; return nil }

// GetKubeadmConfigTemplate retrieves a KubeadmConfigTemplate using a client
// Implements ObjectRetriever.
func GetKubeadmConfigTemplate(ctx context.Context, client kubernetes.Client, name, namespace string) (*bootstrapv1beta2.KubeadmConfigTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KubeadmConfigTemplateEqual returns true only if the new version of a KubeadmConfigTemplate
// involves changes with respect to the old one when applied to the cluster.
// Implements ObjectComparator.
func KubeadmConfigTemplateEqual(new, old *bootstrapv1beta2.KubeadmConfigTemplate) bool {
	_ = "STUB: not implemented"
	// DeepDerivative treats empty map (length == 0) as unset field. We need to manually compare certain fields
	// such as taints, so that setting it to empty will trigger machine recreate
	// The file check with deep equal has been added since the introduction of kubelet configuration in case users
	// want to get rid of the files with that context.
	return false
}

func kubeadmConfigTemplateTaintsEqual(new, old *bootstrapv1beta2.KubeadmConfigTemplate) bool {
	_ = "STUB: not implemented"
	return false
}

func kubeadmConfigTemplateExtraArgsEqual(new, old *bootstrapv1beta2.KubeadmConfigTemplate) bool {
	_ = "STUB: not implemented"
	return false
}

// taintsFromPtr dereferences a *[]Taint to []Taint, returning nil if the pointer is nil.
func taintsFromPtr(t *[]corev1.Taint) []corev1.Taint { _ = "STUB: not implemented"; return nil }

// taintsToPtr returns a pointer to the taints slice, or nil if the slice is nil.
func taintsToPtr(t []corev1.Taint) *[]corev1.Taint { _ = "STUB: not implemented"; return nil }
