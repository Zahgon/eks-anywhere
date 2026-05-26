package cluster

import (
	"regexp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// ConfigManager allows to parse from yaml, set defaults and validate a Cluster struct
// It allows to dynamically register configuration for all those operations.
type ConfigManager struct { // TODO: find a better name
	entry *ConfigManagerEntry
}

// NewConfigManager builds a ConfigManager with empty configuration.
func NewConfigManager() *ConfigManager { _ = "STUB: not implemented"; return nil }

// Register records the configuration defined in a ConfigManagerEntry into the ConfigManager
// This is equivalent to the individual register methods.
func (c *ConfigManager) Register(entries ...*ConfigManagerEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterMapping records the mapping between a kubernetes Kind and an API concrete type.
func (c *ConfigManager) RegisterMapping(kind string, generator APIObjectGenerator) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterProcessors records setters to fill the Config struct from the parsed API objects.
func (c *ConfigManager) RegisterProcessors(processors ...ParsedProcessor) {
	_ = "STUB: not implemented"
	return
}

// RegisterValidations records validations for a Config struct.
func (c *ConfigManager) RegisterValidations(validations ...Validation) {
	_ = "STUB: not implemented"
	return
}

// RegisterDefaulters records defaults for a Config struct.
func (c *ConfigManager) RegisterDefaulters(defaulters ...Defaulter) {
	_ = "STUB: not implemented"
	return
}

// Parse reads yaml manifest with at least one cluster object and generates the corresponding Config.
func (c *ConfigManager) Parse(yamlManifest []byte) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse set the registered defaults in a Config struct.
func (c *ConfigManager) SetDefaults(config *Config) error { _ = "STUB: not implemented"; return nil }

// Validate performs the registered validations in a Config struct.
func (c *ConfigManager) Validate(config *Config) error { _ = "STUB: not implemented"; return nil }

type basicAPIObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

func (k *basicAPIObject) empty() bool { _ = "STUB: not implemented"; return false }

type parsed struct {
	objects ObjectLookup
	cluster *anywherev1.Cluster
}

var separatorRegex = regexp.MustCompile(`(?m)^---$`)

func (c *ConfigManager) unmarshal(yamlManifest []byte) (*parsed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignore empty objects.
// Empty objects are generated if there are weird things in manifest files like e.g. two --- in a row without a yaml doc in the middle

func (c *ConfigManager) buildConfigFromParsed(p *parsed) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// machineConfigsProcessor is a helper to generate a ParsedProcessor for all machine configs in a Cluster.
func machineConfigsProcessor(processMachineRef func(c *Config, o ObjectLookup, machineRef *anywherev1.Ref)) ParsedProcessor {
	_ = "STUB: not implemented"
	return *new(ParsedProcessor)
}
