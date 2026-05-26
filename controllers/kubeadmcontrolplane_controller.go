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
	controlplanev1beta2 "sigs.k8s.io/cluster-api/api/controlplane/kubeadm/v1beta2"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	kcpInPlaceUpgradeNeededAnnotation = "controlplane.clusters.x-k8s.io/in-place-upgrade-needed"
	controlPlaneMachineLabel          = "cluster.x-k8s.io/control-plane-name"
	kubeadmControlPlaneKind           = "KubeadmControlPlane"
)

// KubeadmControlPlaneReconciler reconciles a KubeadmControlPlaneReconciler object.
type KubeadmControlPlaneReconciler struct {
	// client reads from a cache and is not a fully direct client.
	client client.Client
	// uncachedClient reads directly from the API server and is slightly slower.
	uncachedClient client.Reader
	log            logr.Logger
}

// NewKubeadmControlPlaneReconciler returns a new instance of KubeadmControlPlaneReconciler.
func NewKubeadmControlPlaneReconciler(client client.Client, uncachedClient client.Reader) *KubeadmControlPlaneReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=controlplane.cluster.x-k8s.io,resources=kubeadmcontrolplane,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=controlplane.cluster.x-k8s.io,resources=kubeadmcontrolplane/status,verbs=get

// Reconcile reconciles a KubeadmControlPlane object for in place upgrades.
func (r *KubeadmControlPlaneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Always attempt to patch after each reconciliation in case annotation is removed.

// Only requeue if we are not already re-queueing and the "in-place-upgrade-needed" annotation is not set.
// We do this to be able to update the status continuously until it becomes ready,
// since there might be changes in state of the world that don't trigger reconciliation requests

// SetupWithManager sets up the controller with the Manager.
func (r *KubeadmControlPlaneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *KubeadmControlPlaneReconciler) reconcile(ctx context.Context, log logr.Logger, kcp *controlplanev1beta2.KubeadmControlPlane) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Remove the in-place-upgrade-needed annotation only after the ControlPlaneUpgrade object is deleted

func (r *KubeadmControlPlaneReconciler) inPlaceUpgradeNeeded(kcp *controlplanev1beta2.KubeadmControlPlane) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *KubeadmControlPlaneReconciler) machinesToUpgrade(ctx context.Context, kcp *controlplanev1beta2.KubeadmControlPlane) ([]corev1.ObjectReference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *KubeadmControlPlaneReconciler) validateStackedEtcd(kcp *controlplanev1beta2.KubeadmControlPlane) error {
	_ = "STUB: not implemented"
	return nil
}

func pauseMachineHealthCheck(ctx context.Context, mhc *clusterv1beta2.MachineHealthCheck, mhcPatchHelper *patch.Helper) error {
	_ = "STUB: not implemented"
	return nil
}

func resumeMachineHealthCheck(ctx context.Context, mhc *clusterv1beta2.MachineHealthCheck, mhcPatchHelper *patch.Helper) error {
	_ = "STUB: not implemented"
	return nil
}

func controlPlaneUpgrade(kcp *controlplanev1beta2.KubeadmControlPlane, machines []corev1.ObjectReference) (*anywherev1.ControlPlaneUpgrade, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cpUpgradeName(kcpName string) string { _ = "STUB: not implemented"; return "" }

func cpMachineHealthCheckName(kcpName string) string { _ = "STUB: not implemented"; return "" }
