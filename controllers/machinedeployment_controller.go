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
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	mdInPlaceUpgradeNeededAnnotation = "machinedeployment.clusters.x-k8s.io/in-place-upgrade-needed"
	workerMachineLabel               = "cluster.x-k8s.io/deployment-name"
	machineDeploymentKind            = "MachineDeployment"
)

// MachineDeploymentReconciler reconciles a MachineDeploymentReconciler object.
type MachineDeploymentReconciler struct {
	// client reads from a cache and is not a fully direct client.
	client client.Client
	// uncachedClient reads directly from the API server and is slightly slower.
	uncachedClient client.Reader
	log            logr.Logger
}

// NewMachineDeploymentReconciler returns a new instance of MachineDeploymentReconciler.
func NewMachineDeploymentReconciler(client client.Client, uncachedClient client.Reader) *MachineDeploymentReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=cluster.x-k8s.io,resources=machinedeployment,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=cluster.x-k8s.io,resources=machinedeployment/status,verbs=get

// Reconcile reconciles a MachineDeployment object for in place upgrades.
func (r *MachineDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, reterr error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Always attempt to patch after each reconciliation in case annotation is removed.

// Only requeue if we are not already re-queueing and the "in-place-upgrade-needed" annotation is not set.
// We do this to be able to update the status continuously until it becomes ready,
// since there might be changes in state of the world that don't trigger reconciliation requests

// SetupWithManager sets up the controller with the Manager.
func (r *MachineDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *MachineDeploymentReconciler) reconcile(ctx context.Context, log logr.Logger, md *clusterv1beta2.MachineDeployment) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Remove the in-place-upgrade-needed annotation only after the MachineDeploymentUpgrade object is deleted

func (r *MachineDeploymentReconciler) inPlaceUpgradeNeeded(md *clusterv1beta2.MachineDeployment) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *MachineDeploymentReconciler) machinesToUpgrade(ctx context.Context, md *clusterv1beta2.MachineDeployment) ([]*clusterv1beta2.Machine, []corev1.ObjectReference, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func sortMachinesByCreationTimestamp(list *clusterv1beta2.MachineList) []*clusterv1beta2.Machine {
	_ = "STUB: not implemented"
	return nil
}

func machineDeploymentUpgrade(md *clusterv1beta2.MachineDeployment, machines []corev1.ObjectReference) (*anywherev1.MachineDeploymentUpgrade, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mdUpgradeName(mdName string) string { _ = "STUB: not implemented"; return "" }

func mdMachineHealthCheckName(mdName string) string { _ = "STUB: not implemented"; return "" }
