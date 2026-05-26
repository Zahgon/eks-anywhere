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

var nutanixmachineconfiglog = logf.Log.WithName("nutanixmachineconfig-resource")

// SetupWebhookWithManager sets up and registers the webhook with the manager.
func (in *NutanixMachineConfig) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomValidator = &NutanixMachineConfig{}

//+kubebuilder:webhook:path=/validate-anywhere-eks-amazonaws-com-v1alpha1-nutanixmachineconfig,mutating=false,failurePolicy=fail,sideEffects=None,groups=anywhere.eks.amazonaws.com,resources=nutanixmachineconfigs,verbs=create;update,versions=v1alpha1,name=validation.nutanixmachineconfig.anywhere.amazonaws.com,admissionReviewVersions={v1,v1beta1}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type.
func (in *NutanixMachineConfig) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type.
func (in *NutanixMachineConfig) ValidateUpdate(_ context.Context, old, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type.
func (in *NutanixMachineConfig) ValidateDelete(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// TODO(user): fill in your validation logic upon object deletion.

func validateImmutableFieldsNutantixMachineConfig(new, old *NutanixMachineConfig) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateImmutableFieldsControlPlane(new, old *NutanixMachineConfig) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}
