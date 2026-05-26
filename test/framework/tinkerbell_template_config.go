package framework

import (
	_ "embed"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

//go:embed testdata/tinkerbell/custom_config.yaml
var customTinkerbellConfigYAML []byte

// GetCustomTinkerbellConfig returns a custom TinkerbellTemplateConfig.
func GetCustomTinkerbellConfig(tinkerbellLBIP string, osImage string) (*anywherev1.TinkerbellTemplateConfig, error) {
	_ = "STUB: not implemented"
	// Replace placeholders with actual values using string replacement
	return nil, nil
}

// Parse the YAML into TinkerbellTemplateConfig
