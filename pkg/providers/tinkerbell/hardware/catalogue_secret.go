package hardware

import (
	corev1 "k8s.io/api/core/v1"
)

// IndexSecret indexes Secret instances on index by extracfting the key using fn.
func (c *Catalogue) IndexSecret(index string, fn KeyExtractorFunc) {
	_ = "STUB: not implemented"
	return
}

// InsertSecret inserts Secrets into the catalogue. If any indexes exist, the Secret is indexed.
func (c *Catalogue) InsertSecret(secret *corev1.Secret) error {
	_ = "STUB: not implemented"
	return nil
}

// AllSecrets retrieves a copy of the catalogued Secret instances.
func (c *Catalogue) AllSecrets() []*corev1.Secret { _ = "STUB: not implemented"; return nil }

// LookupSecret retrieves Secret instances on index with a key of key. Multiple Secrets _may_
// have the same key hence it can return multiple Secrets.
func (c *Catalogue) LookupSecret(index, key string) ([]*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalSecrets returns the total Secrets registered in the catalogue.
func (c *Catalogue) TotalSecrets() int { _ = "STUB: not implemented"; return 0 }

const SecretNameIndex = ".ObjectMeta.Name"

// WithSecretNameIndex creates a Secret index using SecretNameIndex on Secret.ObjectMeta.Name.
func WithSecretNameIndex() CatalogueOption { _ = "STUB: not implemented"; return *new(CatalogueOption) }

// SecretCatalogueWriter converts Machine instances to Tinkerbell BaseboardManagement and inserts them
// in a catalogue.
type SecretCatalogueWriter struct {
	catalogue *Catalogue
}

var _ MachineWriter = &SecretCatalogueWriter{}

// NewSecretCatalogueWriter creates a new SecretCatalogueWriter instance.
func NewSecretCatalogueWriter(catalogue *Catalogue) *SecretCatalogueWriter {
	_ = "STUB: not implemented"
	return nil
}

// Write converts m to a Tinkerbell BaseboardManagement and inserts it into w's Catalogue.
func (w *SecretCatalogueWriter) Write(m Machine) error { _ = "STUB: not implemented"; return nil }

func baseboardManagementSecretFromMachine(m Machine) []*corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}
