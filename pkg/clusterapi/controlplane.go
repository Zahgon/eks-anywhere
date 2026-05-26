package clusterapi

import (
	"context"

	etcdv1 "github.com/aws/etcdadm-controller/api/v1beta1"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

// ControlPlane represents the provider-specific spec for a CAPI control plane using the kubeadm CP provider.
type ControlPlane[C Object[C], M Object[M]] struct {
	Cluster *clusterv1beta2.Cluster

	// ProviderCluster is the provider-specific resource that holds the details
	// for provisioning the infrastructure, referenced in Cluster.Spec.InfrastructureRef
	ProviderCluster C

	KubeadmControlPlane *controlplanev1beta2.KubeadmControlPlane

	// ControlPlaneMachineTemplate is the provider-specific machine template referenced
	// in KubeadmControlPlane.Spec.MachineTemplate.Spec.InfrastructureRef
	ControlPlaneMachineTemplate M

	EtcdCluster *etcdv1.EtcdadmCluster

	// EtcdMachineTemplate is the provider-specific machine template referenced
	// in EtcdCluster.Spec.InfrastructureTemplate
	EtcdMachineTemplate M
}

// Objects returns all API objects that form a concrete provider-specific control plane.
func (cp *ControlPlane[C, M]) Objects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// UpdateImmutableObjectNames checks if any control plane immutable objects have changed by comparing the new definition
// with the current state of the cluster. If they had, it generates a new name for them by increasing a monotonic number
// at the end of the name
// This is applied to all provider machine templates.
func (cp *ControlPlane[C, M]) UpdateImmutableObjectNames(
	ctx context.Context,
	client kubernetes.Client,
	machineTemplateRetriever ObjectRetriever[M],
	machineTemplateComparator ObjectComparator[M],
) error {
	_ = "STUB: not implemented"
	return nil
}

// KubeadmControlPlane doesn't exist, this is a new cluster so machine templates should use their default name

// Ensure we don't set an empty name which would cause "resource name may not be empty" errors

// If the name is empty, we keep the default name that was set during template generation

// EtcdadmCluster doesn't exist, this is a new cluster so machine templates should use their default name

// Ensure we don't set an empty name which would cause "resource name may not be empty" errors

// If the name is empty, we keep the default name that was set during template generation
