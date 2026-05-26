package tinkerbell

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const defaultUsername = "ec2-user"

func ensureMachineConfigsHaveAtLeast1User(machines map[string]*v1alpha1.TinkerbellMachineConfig) {
	_ = "STUB: not implemented"
	return
}

func extractUserConfigurationsWithoutSshKeys(machines map[string]*v1alpha1.TinkerbellMachineConfig) []*v1alpha1.UserConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func applySshKeyToUsers(users []*v1alpha1.UserConfiguration, key string) {
	_ = "STUB: not implemented"
	return
}

func stripCommentsFromSshKeys(machines map[string]*v1alpha1.TinkerbellMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) configureSshKeys() error { _ = "STUB: not implemented"; return nil }
