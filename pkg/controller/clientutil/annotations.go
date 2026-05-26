package clientutil

import "sigs.k8s.io/controller-runtime/pkg/client"

// AddAnnotation adds an annotation to the given object.
// If the annotation already exists, it overwrites its value.
func AddAnnotation(o client.Object, key, value string) { _ = "STUB: not implemented"; return }

// AddLabel adds a label to the given object.
// If the label already exists, it overwrites its value.
func AddLabel(o client.Object, key, value string) { _ = "STUB: not implemented"; return }

// RemoveAnnotation removes an annotation from the given object.
func RemoveAnnotation(o client.Object, key string) { _ = "STUB: not implemented"; return }
