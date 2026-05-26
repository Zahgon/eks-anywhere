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
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	// mdUpgradeFinalizerName is the finalizer added to MachineDeploymentUpgrade objects to handle deletion.
	mdUpgradeFinalizerName = "machinedeploymentupgrades.anywhere.eks.amazonaws.com/finalizer"
	mdLabelIdentifier      = "cluster.x-k8s.io/deployment-name"
)

// MachineDeploymentUpgradeReconciler reconciles a MachineDeploymentUpgrade object.
type MachineDeploymentUpgradeReconciler struct {
	client client.Client
	log    logr.Logger
}

// NewMachineDeploymentUpgradeReconciler returns a new instance of MachineDeploymentUpgradeReconciler.
func NewMachineDeploymentUpgradeReconciler(client client.Client) *MachineDeploymentUpgradeReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=machinedeploymentupgrades,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=machinedeploymentupgrades/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=anywhere.eks.amazonaws.com,resources=machinedeploymentupgrades/finalizers,verbs=update
//+kubebuilder:rbac:groups=cluster.x-k8s.io,resources=machinesets,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=cluster.x-k8s.io,resources=machines,verbs=get;list;watch;update;patch

// Reconcile reconciles a MachineDeploymentUpgrade object.
// nolint:gocyclo
// TODO: Reduce high cyclomatic complexity: https://github.com/aws/eks-anywhere-internal/issues/2119
func (r *MachineDeploymentUpgradeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Always attempt to patch the object and status after each reconciliation.

// We want the observedGeneration to indicate, that the status shown is up-to-date given the desired spec of the same generation.
// However, if there is an error while updating the status, we may get a partial status update, In this case,
// a partially updated status is not considered up to date, so we should not update the observedGeneration

// Patch ObservedGeneration only if the reconciliation completed without error

// Only requeue if we are not already re-queueing and the Ready condition is false.
// We do this to be able to update the status continuously until it becomes ready,
// since there might be changes in state of the world that don't trigger reconciliation requests

// Reconcile the MachineDeploymentUpgrade deletion if the DeletionTimestamp is set.

// AddFinalizer	is idempotent

// SetupWithManager sets up the controller with the Manager.
func (r *MachineDeploymentUpgradeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *MachineDeploymentUpgradeReconciler) reconcile(ctx context.Context, log logr.Logger, mdUpgrade *anywherev1.MachineDeploymentUpgrade) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (r *MachineDeploymentUpgradeReconciler) reconcileDelete(ctx context.Context, log logr.Logger, mdUpgrade *anywherev1.MachineDeploymentUpgrade) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Remove the finalizer on MachineDeploymentUpgrade object

func patchMachineDeploymentUpgrade(ctx context.Context, patchHelper *patch.Helper, mdUpgrade *anywherev1.MachineDeploymentUpgrade, patchOpts ...patch.Option) error {
	_ = "STUB: not implemented"
	// Always attempt to patch the object and status after each reconciliation.
	return nil
}

func (r *MachineDeploymentUpgradeReconciler) updateStatus(ctx context.Context, log logr.Logger, mdUpgrade *anywherev1.MachineDeploymentUpgrade, ms *clusterv1beta2.MachineSet) error {
	_ = "STUB: not implemented"
	// When MachineDeploymentUpgrade is fully deleted, we do not need to update the status. Without this check
	// the subsequent patch operations would fail if the status is updated after it is fully deleted.
	return nil
}

func (r *MachineDeploymentUpgradeReconciler) getCurrentMachineSet(ctx context.Context, md *clusterv1beta2.MachineDeployment) (*clusterv1beta2.MachineSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *MachineDeploymentUpgradeReconciler) updateMachineSet(ctx context.Context, ms *clusterv1beta2.MachineSet, spec clusterv1beta2.MachineSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *MachineDeploymentUpgradeReconciler) updateMachineVersion(ctx context.Context, log logr.Logger, machineRef corev1.ObjectReference, kubernetesVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if update is needed

func mdNodeUpgrader(machineRef corev1.ObjectReference, kubernetesVersion string) *anywherev1.NodeUpgrade {
	_ = "STUB: not implemented"
	return nil
}
