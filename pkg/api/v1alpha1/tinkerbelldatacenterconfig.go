package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const TinkerbellDatacenterKind = "TinkerbellDatacenterConfig"

// Used for generating yaml for generate clusterconfig command.
func NewTinkerbellDatacenterConfigGenerate(clusterName string) *TinkerbellDatacenterConfigGenerate {
	_ = "STUB: not implemented"
	return nil
}

func (c *TinkerbellDatacenterConfigGenerate) APIVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *TinkerbellDatacenterConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *TinkerbellDatacenterConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }

func GetTinkerbellDatacenterConfig(fileName string) (*TinkerbellDatacenterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateDatacenterConfig(config *TinkerbellDatacenterConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateObjectMeta(meta metav1.ObjectMeta) error { _ = "STUB: not implemented"; return nil }
