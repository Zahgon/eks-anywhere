package cluster

import (
	"k8s.io/apimachinery/pkg/runtime"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// APIObject represents a kubernetes API object.
type APIObject interface {
	runtime.Object
	GetName() string
}

type ObjectLookup map[string]APIObject

// GetFromRef searches in a ObjectLookup for an APIObject referenced by a anywherev1.Ref.
func (o ObjectLookup) GetFromRef(apiVersion string, ref anywherev1.Ref) APIObject {
	_ = "STUB: not implemented"
	return *new(APIObject)
}

func (o ObjectLookup) add(obj APIObject) { _ = "STUB: not implemented"; return }

func keyForRef(apiVersion string, ref anywherev1.Ref) string { _ = "STUB: not implemented"; return "" }

func key(apiVersion, kind, name string) string {
	_ = "STUB: not implemented"
	// this assumes we don't allow to have objects in multiple namespaces
	return ""
}

func keyForObject(o APIObject) string { _ = "STUB: not implemented"; return "" }
