package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const TinkerbellMachineConfigKind = "TinkerbellMachineConfig"

// +kubebuilder:object:generate=false
type TinkerbellMachineConfigGenerateOpt func(config *TinkerbellMachineConfigGenerate)

// Used for generating yaml for generate clusterconfig command.
func NewTinkerbellMachineConfigGenerate(name string, opts ...TinkerbellMachineConfigGenerateOpt) *TinkerbellMachineConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *TinkerbellMachineConfigGenerate) APIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *TinkerbellMachineConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *TinkerbellMachineConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func WithTemplateRef(ref ProviderRefAccessor) TinkerbellMachineConfigGenerateOpt {
	_ = "STUB: not implemented"
	return *new(TinkerbellMachineConfigGenerateOpt)
}

func validateTinkerbellMachineConfig(config *TinkerbellMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate hardware selection (HardwareSelector vs HardwareAffinity)

// validateHardwareSelection validates the hardware selection configuration.
// HardwareSelector and HardwareAffinity are mutually exclusive.
func validateHardwareSelection(config *TinkerbellMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Check mutual exclusivity

// At least one must be specified

// Validate HardwareSelector if present

// Validate HardwareAffinity if present

// validateHardwareAffinity validates the HardwareAffinity configuration.
func validateHardwareAffinity(affinity *HardwareAffinity, configName string) error {
	_ = "STUB: not implemented"
	// Required terms must have at least one entry
	return nil
}

// Validate each required term

// Validate each preferred term

// Validate weight range (1-100)

// validateLabelSelector validates a Kubernetes LabelSelector.
func validateLabelSelector(selector *metav1.LabelSelector, configName, path string) error {
	_ = "STUB: not implemented"
	// Validate matchExpressions
	return nil
}

// Validate operator

// Validate values for In/NotIn operators

// Validate that Exists/DoesNotExist don't have values

// isValidLabelSelectorOperator checks if the operator is a valid LabelSelector operator.
func isValidLabelSelectorOperator(op metav1.LabelSelectorOperator) bool {
	_ = "STUB: not implemented"
	return false
}

// ValidateHardwareAffinityOperator validates a single operator string.
func ValidateHardwareAffinityOperator(op string) bool { _ = "STUB: not implemented"; return false }

// ValidateHardwareAffinityWeight validates a weight value.
func ValidateHardwareAffinityWeight(weight int32) bool { _ = "STUB: not implemented"; return false }

// ValidateLabelSelectorRequirement validates a single LabelSelectorRequirement.
func ValidateLabelSelectorRequirement(req metav1.LabelSelectorRequirement) error {
	_ = "STUB: not implemented"
	return nil
}

func setTinkerbellMachineConfigDefaults(machineConfig *TinkerbellMachineConfig) {
	_ = "STUB: not implemented"
	return
}

func normalizeSSHKeys(machineConfig *TinkerbellMachineConfig) { _ = "STUB: not implemented"; return }

func stripCommentsFromSSHKeys(machine *TinkerbellMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}
