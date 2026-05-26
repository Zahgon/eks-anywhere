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

// log is for logging in this package.
var vspheremachineconfiglog = logf.Log.WithName("vspheremachineconfig-resource")

func (r *VSphereMachineConfig) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-anywhere-eks-amazonaws-com-v1alpha1-vspheremachineconfig,mutating=true,failurePolicy=fail,sideEffects=None,groups=anywhere.eks.amazonaws.com,resources=vspheremachineconfigs,verbs=create;update,versions=v1alpha1,name=mutation.vspheremachineconfig.anywhere.amazonaws.com,admissionReviewVersions={v1,v1beta1}

var _ webhook.CustomDefaulter = &VSphereMachineConfig{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the type.
func (r *VSphereMachineConfig) Default(_ context.Context, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-anywhere-eks-amazonaws-com-v1alpha1-vspheremachineconfig,mutating=false,failurePolicy=fail,sideEffects=None,groups=anywhere.eks.amazonaws.com,resources=vspheremachineconfigs,verbs=create;update,versions=v1alpha1,name=validation.vspheremachineconfig.anywhere.amazonaws.com,admissionReviewVersions={v1,v1beta1}

var _ webhook.CustomValidator = &VSphereMachineConfig{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type.
func (r *VSphereMachineConfig) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type.
func (r *VSphereMachineConfig) ValidateUpdate(_ context.Context, old, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func validateImmutableFieldsVSphereMachineConfig(new, old *VSphereMachineConfig) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type.
func (r *VSphereMachineConfig) ValidateDelete(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
