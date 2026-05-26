package yamlutil

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clusterv1beta2 "sigs.k8s.io/cluster-api/api/core/v1beta2"
)

// APIObject represents a kubernetes API object.
type APIObject interface {
	runtime.Object
	GetName() string
}

// ObjectLookup allows to search APIObjects by a unique key composed of apiVersion, kind, and name.
type ObjectLookup map[string]APIObject

// GetFromRef searches in a ObjectLookup for an APIObject referenced by a corev1.ObjectReference.
func (o ObjectLookup) GetFromRef(ref corev1.ObjectReference) APIObject {
	_ = "STUB: not implemented"
	return *

	// GetFromContractVersionedRef searches for an APIObject referenced by a v1beta2 ContractVersionedObjectReference.
	new(APIObject)
}

func (o ObjectLookup) GetFromContractVersionedRef(ref clusterv1beta2.ContractVersionedObjectReference) APIObject {
	_ = "STUB: not implemented"
	return *new(APIObject)
}

func (o ObjectLookup) add(obj APIObject) { _ = "STUB: not implemented"; return }

func NewObjectLookupBuilder() *ObjectLookupBuilder { _ = "STUB: not implemented"; return nil }

// ObjectLookupBuilder allows to construct an ObjectLookup and add APIObjects to it.
type ObjectLookupBuilder struct {
	lookup ObjectLookup
}

// Add acumulates an API object that will be included in the built ObjectLookup.
func (o *ObjectLookupBuilder) Add(objs ...APIObject) *ObjectLookupBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Build constructs and returns an ObjectLookup
// After this method is called, the builder is reset and loses track
// of all previously added objects.
func (o *ObjectLookupBuilder) Build() ObjectLookup {
	_ = "STUB: not implemented"
	return *new(ObjectLookup)
}

// Key builds the yaml object key.
func Key(apiVersion, kind, name string) string {
	_ = "STUB: not implemented"
	// this assumes we don't allow to have objects in multiple namespaces
	return ""
}

func keyForRef(ref corev1.ObjectReference) string { _ = "STUB: not implemented"; return "" }

func keyForObject(o APIObject) string { _ = "STUB: not implemented"; return "" }
