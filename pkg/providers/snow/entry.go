package snow

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/cluster"
)

type ConfigManager struct {
	validator  *Validator
	defaulters *Defaulters
}

// NewConfigManager returns a new snow config manager.
func NewConfigManager(defaulters *Defaulters, validators *Validator) *ConfigManager {
	_ = "STUB: not implemented"
	return nil
}

func (cm *ConfigManager) SetDefaultsAndValidate(ctx context.Context, config *cluster.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (cm *ConfigManager) snowEntry(ctx context.Context) *cluster.ConfigManagerEntry {
	_ = "STUB: not implemented"
	return nil
}
