package hardware

import (
	"reflect"
)

// FieldIndexer indexes collection of objects for a single type against one of its fields.
// FieldIndexer is not thread safe.
type FieldIndexer struct {
	expectedType reflect.Type
	indexes      map[string]*fieldIndex
}

// NewFieldIndexer creates a new FieldIndexer instance. object is the object to be indexed and will
// be checked during Insert() calls. NewFieldIndexer will panic if object is nil.
func NewFieldIndexer(object interface{}) *FieldIndexer { _ = "STUB: not implemented"; return nil }

// KeyExtractorFunc returns a key from object that can be used to look up the object.
type KeyExtractorFunc func(object interface{}) string

// IndexField registers a new index with i. field is the index name and should represent a path
// to the field such as `.Spec.ID`. fn is used to extract the lookup key on Insert() from the object
// to be inserted.
func (i *FieldIndexer) IndexField(field string, fn KeyExtractorFunc) {
	_ = "STUB: not implemented"
	return
}

// Insert inserts v into i on all indexed fields registered with IndexField. If v is not of the
// expected type defined by NewFieldIndexer() ErrIncorrectType is returned. Multiple objects
// with the same index value may be inserted.
func (i *FieldIndexer) Insert(v interface{}) error { _ = "STUB: not implemented"; return nil }

// Lookup uses the index associated with field to find and return all objects associated with key.
// If field has no associated index created by IndexField ErrUnknownIndex is returned.
func (i *FieldIndexer) Lookup(field string, key string) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove removes v from all indexes if present. If v is not present Remove is a no-op. If v is of
// an incorrect type ErrUnknownType is returned.
func (i *FieldIndexer) Remove(v interface{}) error { _ = "STUB: not implemented"; return nil }

// fieldIndex represents a single index on a particular object. When inserting into the fieldIndex
// the key is extracted from the object using the KeyExtractorFunc.
type fieldIndex struct {
	index            map[string][]interface{}
	keyExtractorFunc KeyExtractorFunc
}

func (i *fieldIndex) Insert(v interface{}) { _ = "STUB: not implemented"; return }

func (i *fieldIndex) Lookup(key string) []interface{} { _ = "STUB: not implemented"; return nil }

func (i *fieldIndex) Remove(v interface{}) { _ = "STUB: not implemented"; return }

// ErrIncorrectType indicates an incorrect type was used with a FieldIndexer.
type ErrIncorrectType struct {
	Expected reflect.Type
	Received reflect.Type
}

func (e ErrIncorrectType) Error() string { _ = "STUB: not implemented"; return "" }

type ErrUnknownIndex struct {
	Field string
}

func (e ErrUnknownIndex) Error() string { _ = "STUB: not implemented"; return "" }
