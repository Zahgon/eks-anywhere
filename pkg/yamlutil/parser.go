package yamlutil

import (
	"io"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	// APIObjectGenerator returns an implementor of the APIObject interface.
	APIObjectGenerator func() APIObject
	// ParsedProcessor fills the struct of type T with the parsed API objects in ObjectLookup.
	ParsedProcessor[T any] func(*T, ObjectLookup)

	// Parser allows to parse from yaml with kubernetes style objects and
	// store them in a type implementing Builder
	// It allows to dynamically register configuration for mappings between kind and concrete types.
	Parser struct {
		apiObjectMapping   map[string]APIObjectGenerator
		generateObjAnyKind APIObjectGenerator
		logger             logr.Logger
	}
)

func NewParser(logger logr.Logger) *Parser { _ = "STUB: not implemented"; return nil }

// RegisterMapping records the mapping between a kubernetes Kind and an API concrete type.
func (c *Parser) RegisterMapping(kind string, generator APIObjectGenerator) error {
	_ = "STUB: not implemented"
	return nil
}

// Mapping mapping between a kubernetes Kind and an API concrete type of type T.
type Mapping[T APIObject] struct {
	New  func() T
	Kind string
}

func NewMapping[T APIObject](kind string, new func() T) Mapping[T] {
	_ = "STUB: not implemented"
	return nil
}

// ToAPIObjectMapping is helper to convert from other concrete types of Mapping
// to a APIObject Mapping
// This is mostly to help pass Mappings to RegisterMappings.
func (m Mapping[T]) ToAPIObjectMapping() Mapping[APIObject] { _ = "STUB: not implemented"; return nil }

// RegisterMappings records a collection of mappings.
func (c *Parser) RegisterMappings(mappings ...Mapping[APIObject]) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterMappingForAnyKind records an object generator that will be used
// as fallback when there is not a specific APIObjectGenerator registered for that particular kind.
func (c *Parser) RegisterMappingForAnyKind(generator APIObjectGenerator) {
	_ = "STUB: not implemented"
	return
}

// Builder processes the parsed API objects contained in a lookup.
type Builder interface {
	BuildFromParsed(ObjectLookup) error
}

// Parse reads yaml manifest content with the registered mappings and passes
// the result to the Builder for further processing.
func (p *Parser) Parse(yamlManifest []byte, b Builder) error { _ = "STUB: not implemented"; return nil }

// Read reads yaml manifest content with the registered mappings and passes
// the result to the Builder for further processing.
func (p *Parser) Read(reader io.Reader, b Builder) error { _ = "STUB: not implemented"; return nil }

type basicAPIObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

func (k *basicAPIObject) empty() bool { _ = "STUB: not implemented"; return false }

type parsed struct {
	objects ObjectLookup
}

func (p *Parser) unmarshal(reader io.Reader) (*parsed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read one YAML document at a time, until io.EOF is returned

// Ignore empty objects.
// Empty objects are generated if there are weird things in manifest files like e.g. two --- in a row without a yaml doc in the middle

func (p *Parser) buildConfigFromParsed(parsed *parsed, b Builder) error {
	_ = "STUB: not implemented"
	return nil
}
