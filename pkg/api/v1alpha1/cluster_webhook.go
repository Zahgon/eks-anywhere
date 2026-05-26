// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const supportedMinorVersionIncrement int64 = 1

// log is for logging in this package.
var clusterlog = logf.Log.WithName("cluster-resource")

func (r *Cluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:webhook:path=/mutate-anywhere-eks-amazonaws-com-v1alpha1-cluster,mutating=true,failurePolicy=fail,sideEffects=None,groups=anywhere.eks.amazonaws.com,resources=clusters,verbs=create;update,versions=v1alpha1,name=mutation.cluster.anywhere.amazonaws.com,admissionReviewVersions={v1,v1beta1}

var _ webhook.CustomDefaulter = &Cluster{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the type.
func (r *Cluster) Default(_ context.Context, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// Change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-anywhere-eks-amazonaws-com-v1alpha1-cluster,mutating=false,failurePolicy=fail,sideEffects=None,groups=anywhere.eks.amazonaws.com,resources=clusters,verbs=create;update,versions=v1alpha1,name=validation.cluster.anywhere.amazonaws.com,admissionReviewVersions={v1,v1beta1}

var _ webhook.CustomValidator = &Cluster{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type.
func (r *Cluster) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type.
func (r *Cluster) ValidateUpdate(_ context.Context, old, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type.
func (r *Cluster) ValidateDelete(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func supportsEtcdEncryption(cluster *Cluster) bool { _ = "STUB: not implemented"; return false }

// ValidateEksaVersionSkew ensures that upgrades are sequential by CLI minor versions.
func ValidateEksaVersionSkew(new, old *Cluster) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// allow users to update cluster if old cluster is invalid

// if major different or upgrade difference greater than one minor version

// Allow "downgrades" if old version is greater than managment cluster. We can't check if a cluster's EksaVersion
// is less than or equal to the mgmt cluster in the webhook. Instead we check it in the controller where the
// EksaVersion will already be applied. However, the cluster will never begin to reconcile due to the validation.
// We should not block users from changing EksaVersion to a lower semver in this scenario.

// don't allow downgrades if old version was valid

func validateBundlesRefCluster(new, old *Cluster) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateEksaVersionCluster(new, old *Cluster) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateUpgradeRequestTinkerbell(new, old *Cluster) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateImmutableFieldsCluster(new, old *Cluster) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// ValidateKubernetesVersionSkew validates Kubernetes version skew between upgrades.
func ValidateKubernetesVersionSkew(new, old *Cluster) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateKubeVersionSkew(newVersion, oldVersion KubernetesVersion, path *field.Path) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// ValidateWorkerKubernetesVersionSkew validates worker node group Kubernetes version skew between upgrades.
func ValidateWorkerKubernetesVersionSkew(new, old *Cluster) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func performWorkerKubernetesValidationsNewNodeGroup(newVersion *KubernetesVersion, newClusterVersion KubernetesVersion) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func performWorkerKubernetesValidations(oldVersion, newVersion *KubernetesVersion, oldClusterVersion, newClusterVersion KubernetesVersion) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateRemoveWorkerKubernetesVersion(newCPVersion, oldCPVersion KubernetesVersion, oldWorkerVersion *KubernetesVersion) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validKubeMinorVersionDiff(old, new KubernetesVersion) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateCPWorkerKubeSkew(cpVersion, workerVersion KubernetesVersion) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateEtcdEncryptionSupport(cluster *Cluster) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}
