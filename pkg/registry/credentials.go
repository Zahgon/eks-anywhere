package registry

import (
	"github.com/docker/cli/cli/config/configfile"
	"oras.land/oras-go/v2/registry/remote/auth"
)

// CredentialStore for registry credentials such as ~/.docker/config.json.
type CredentialStore struct {
	directory  string
	configFile *configfile.ConfigFile
}

// NewCredentialStore create a credential store.
func NewCredentialStore() *CredentialStore { _ = "STUB: not implemented"; return nil }

// SetDirectory override default directory.
func (cs *CredentialStore) SetDirectory(directory string) { _ = "STUB: not implemented"; return }

// Init initialize a credential store.
func (cs *CredentialStore) Init() (err error) { _ = "STUB: not implemented"; return nil }

// Credential get an authentication credential for a given registry.
func (cs *CredentialStore) Credential(registry string) (auth.Credential, error) {
	_ = "STUB: not implemented"
	return *new(auth.Credential), nil
}
