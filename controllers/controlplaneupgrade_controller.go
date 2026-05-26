/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	bootstrapv1beta2 "sigs.k8s.io/cluster-api/api/bootstrap/kubeadm/v1beta2"
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// controlPlaneUpgradeFinalizerName is the finalizer added to NodeUpgrade objects to handle deletion.
const (
	controlPlaneUpgradeFinalizerName      = "controlplaneupgrades.anywhere.eks.amazonaws.com/finalizer"
	kubeadmClusterConfigurationAnnotation = "controlplane.cluster.x-k8s.io/kubeadm-cluster-configuration"
	cloneFromNameAnnotationInfraMachine   = "cluster.x-k8s.io/cloned-from-name"
	kubeVipStaticPodPath                  = "/etc/kubernetes/manifests/kube-vip.yaml"
)

// ControlPlaneUpgradeReconciler reconciles a ControlPlaneUpgradeReconciler object.
type ControlPlaneUpgradeReconciler struct {
	client               client.Client
	remoteClientRegistry RemoteClientRegistry
	log                  logr.Logger
}

// NewControlPlaneUpgradeReconciler returns a new instance of ControlPlaneUpgradeReconciler.
func NewControlPlaneUpgradeReconciler(client client.Client, remoteClientRegistry RemoteClientRegistry) *ControlPlaneUpgradeReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=controlplaneupgrades,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=controlplaneupgrades/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=controlplaneupgrades/finalizers,verbs=update
//+kubebuilder:rbac:groups=bootstrap.cluster.x-k8s.io,resources=kubeadmconfigs,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=tinkerbellmachines;vspheremachines,verbs=get;list;update;patch
//+kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch

// Reconcile reconciles a ControlPlaneUpgrade object.
// nolint:gocyclo
func (r *ControlPlaneUpgradeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Always attempt to patch the object and status after each reconciliation.

// We want the observedGeneration to indicate, that the status shown is up-to-date given the desired spec of the same generation.
// However, if there is an error while updating the status, we may get a partial status update, In this case,
// a partially updated status is not considered up to date, so we should not update the observedGeneration

// Patch ObservedGeneration only if the reconciliation completed without error

// Always attempt to patch the object and status after each reconciliation.

// Only requeue if we are not already re-queueing and the Cluster ready condition is false.
// We do this to be able to update the status continuously until the cluster becomes ready,
// since there might be changes in state of the world that don't trigger reconciliation requests

// Reconcile the NodeUpgrade deletion if the DeletionTimestamp is set.

// SetupWithManager sets up the controller with the Manager.
func (r *ControlPlaneUpgradeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ControlPlaneUpgradeReconciler) reconcile(ctx context.Context, log logr.Logger, cpUpgrade *anywherev1.ControlPlaneUpgrade) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// return early if controlplane upgrade is already complete

// check if kube-vip config map exists and clean it up

func nodeUpgrader(machineRef corev1.ObjectReference, kubernetesVersion, etcdVersion string, firstControlPlane bool) *anywherev1.NodeUpgrade {
	_ = "STUB: not implemented"
	return nil
}

func (r *ControlPlaneUpgradeReconciler) reconcileDelete(ctx context.Context, log logr.Logger, cpUpgrade *anywherev1.ControlPlaneUpgrade) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Remove the finalizer on ControlPlaneUpgrade objext

func (r *ControlPlaneUpgradeReconciler) updateStatus(ctx context.Context, log logr.Logger, cpUpgrade *anywherev1.ControlPlaneUpgrade) error {
	_ = "STUB: not implemented"
	// When ControlPlaneUpgrade is fully deleted, we do not need to update the status. Without this check
	// the subsequent patch operations would fail if the status is updated after it is fully deleted.
	return nil
}

func (r *ControlPlaneUpgradeReconciler) updateResources(ctx context.Context, log logr.Logger, cpUpgrade *anywherev1.ControlPlaneUpgrade, nodeUpgrade *anywherev1.NodeUpgrade) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the machine kubeadmClusterConfiguration annotation

// Update the machine k8s version and update the KubeadmClusterConfiguration annotation

func (r *ControlPlaneUpgradeReconciler) updateKubeadmConfig(ctx context.Context, log logr.Logger, kcpSpec *controlplanev1beta2.KubeadmControlPlaneSpec, machine *clusterv1beta2.Machine) error {
	_ = "STUB: not implemented"
	return nil
}

// Apply CAPI's runtime-computed default FeatureGates for the target Kubernetes version.
// CAPI's KCP controller adds ControlPlaneKubeletLocalMode for K8s 1.31-1.35 via
// DefaultFeatureGates() during KubeadmConfig creation, but this gate is not stored in the
// KCP spec. Since CAPI v1.12+ compares the full KubeadmConfig (including ClusterConfiguration),
// the gate must be present to avoid the UpToDate condition remaining False.

// applyCAPIDefaultFeatureGates applies the same default FeatureGates that CAPI's KCP controller
// adds at runtime via DefaultFeatureGates().
// FeatureGates already present in the KubeadmConfig (from the KCP spec) take precedence.
func applyCAPIDefaultFeatureGates(kc *bootstrapv1beta2.KubeadmConfig, kubernetesVersion string) {
	_ = "STUB: not implemented"
	return
}

func (r *ControlPlaneUpgradeReconciler) updateInfraMachine(ctx context.Context, log logr.Logger, kcpSpec *controlplanev1beta2.KubeadmControlPlaneSpec, machine *clusterv1beta2.Machine) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the cloned-from-name annotation to match the updated infra machine template name in KubeadmControlPlane

func decodeAndUnmarshalKcpSpecData(kcpSpecData string) (*controlplanev1beta2.KubeadmControlPlaneSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCapiMachine(ctx context.Context, client client.Client, nodeUpgrade *anywherev1.NodeUpgrade) (*clusterv1beta2.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cleanupKubeVipCM(ctx context.Context, log logr.Logger, remoteClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func createKubeVipCMIfNotExist(ctx context.Context, remoteClient client.Client, cpUpgrade *anywherev1.ControlPlaneUpgrade) error {
	_ = "STUB: not implemented"
	return nil
}

func kubeVipConfigMap(cpUpgrade *anywherev1.ControlPlaneUpgrade) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
