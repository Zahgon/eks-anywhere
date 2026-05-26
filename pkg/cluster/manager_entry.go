package cluster

type (
	// APIObjectGenerator returns an implementor of the APIObject interface.
	APIObjectGenerator func() APIObject
	// ParsedProcessor fills the Config struct from the parsed API objects in ObjectLookup.
	ParsedProcessor func(*Config, ObjectLookup)
	// Validation performs a validation over the Config object.
	Validation func(*Config) error
	// Defaulter sets defaults in a Config object.
	Defaulter func(*Config) error
)

// ConfigManagerEntry allows to declare the necessary configuration to parse
// from yaml, set defaults and validate a Cluster struct for one or more types.
// It is semantically equivalent to use the individual register methods and its
// only purpose is convenience.
type ConfigManagerEntry struct {
	APIObjectMapping map[string]APIObjectGenerator
	Processors       []ParsedProcessor
	Validations      []Validation
	Defaulters       []Defaulter
}

// NewConfigManagerEntry builds a ConfigManagerEntry with empty configuration.
func NewConfigManagerEntry() *ConfigManagerEntry { _ = "STUB: not implemented"; return nil }

// Merge combines the configuration declared in multiple ConfigManagerEntry.
func (c *ConfigManagerEntry) Merge(entries ...*ConfigManagerEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterMapping records the mapping between a kubernetes Kind and an API concrete type.
func (c *ConfigManagerEntry) RegisterMapping(kind string, generator APIObjectGenerator) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterProcessors records setters to fill the Config struct from the parsed API objects.
func (c *ConfigManagerEntry) RegisterProcessors(processors ...ParsedProcessor) {
	_ = "STUB: not implemented"
	return
}

// RegisterValidations records validations for a Config struct.
func (c *ConfigManagerEntry) RegisterValidations(validations ...Validation) {
	_ = "STUB: not implemented"
	return
}

// RegisterDefaulters records defaults for a Config struct.
func (c *ConfigManagerEntry) RegisterDefaulters(defaulters ...Defaulter) {
	_ = "STUB: not implemented"
	return
}
