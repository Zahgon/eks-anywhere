package v1alpha1

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1/thirdparty/tinkerbell"
)

const TinkerbellTemplateConfigKind = "TinkerbellTemplateConfig"

// +kubebuilder:object:generate=false
type ActionOpt func(action *[]tinkerbell.Action)

// NewDefaultTinkerbellTemplateConfigCreate returns a default TinkerbellTemplateConfig with the required Tasks and Actions.
func NewDefaultTinkerbellTemplateConfigCreate(clusterSpec *Cluster, osImageOverride, tinkerbellLocalIP, tinkerbellLBIP string, osFamily OSFamily) *TinkerbellTemplateConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *TinkerbellTemplateConfigGenerate) APIVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *TinkerbellTemplateConfigGenerate) Kind() string { _ = "STUB: not implemented"; return "" }

func (c *TinkerbellTemplateConfigGenerate) Name() string { _ = "STUB: not implemented"; return "" }
