package envtest

import (
	"context"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CreateObjs creates Objects using the provided kube client and waits until its cache
// has been updated with those objects.
func CreateObjs(ctx context.Context, t testing.TB, c client.Client, objs ...client.Object) {
	_ = "STUB: not implemented"
	return
}

// we copy objects because the client modifies them while making creating/updating calls

// namespaces can't be deleted
// assuming most tests just want the namespace to exist, since it already does
// we ignore the error
// for more advance usecases, handle namespaces manually outside of this helper

// If the status doesn't need to be updated, just wait for the object to
// to be available.

type updatedStatus struct {
	obj       client.Object
	newStatus map[string]interface{}
}

// UpdateStatusAndWait updates an objects status subresource and waits until the cache refreshes
// and reflects the new status.
func UpdateStatusAndWait(ctx context.Context, t testing.TB, c client.Client, o client.Object) {
	_ = "STUB: not implemented"
	return
}

func updateStatus(ctx context.Context, t testing.TB, c client.Client, o client.Object) (newStatus map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// We need to update the status independently, kubernetes doesn't allow to create the main objects and
// its subresources all at once

// Some objects without a subresource will fail here,
// so we just try and if it fails with a 404, we ignore the error

func waitForStatusUpdated(ctx context.Context, t testing.TB, c client.Client, o client.Object, newStatus map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

// Filter out CAPI v1beta1 fields that are deprecated in v1beta2 for CAPI objects
// These fields (infrastructureReady, controlPlaneReady, ready) are removed in CAPI v1beta2
// and may not be preserved during API version conversions
// TODO: Remove these conditions once we move to using CAPI v1beta2 objects inside EKS-Anywhere

// isCAPIObject checks if the object is any CAPI object that might have deprecated fields.
func isCAPIObject(obj client.Object) bool { _ = "STUB: not implemented"; return false }

// Check for CAPI core objects and infrastructure objects

// isDockerInfraObject checks if the object is a Docker infrastructure object used in tests.
func isDockerInfraObject(obj client.Object) bool { _ = "STUB: not implemented"; return false }

// Check for Docker infrastructure objects used in CAPI tests

// filterDeprecatedCAPIFields removes deprecated CAPI v1beta1 fields that don't exist in v1beta2.
func filterDeprecatedCAPIFields(status map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Skip deprecated fields that were removed in CAPI v1beta2

func waitForObjectAvailable(ctx context.Context, t testing.TB, c client.Client, obj client.Object) *unstructured.Unstructured {
	_ = "STUB: not implemented"
	return nil
}

func isNamespace(obj client.Object) bool { _ = "STUB: not implemented"; return false }

func copyObject(t testing.TB, obj client.Object) client.Object {
	_ = "STUB: not implemented"
	return *new(client.Object)
}

// APIExpecter is a helper to define eventual expectations over API resources in tests.
// It's useful when working with clients that maintain a cache, since changes might not be
// reflected immediately, causing tests to flake.
type APIExpecter struct {
	t       testing.TB
	client  client.Client
	g       gomega.Gomega
	timeout time.Duration
}

// NewAPIExpecter constructs a new APIExpecter.
func NewAPIExpecter(t testing.TB, client client.Client) *APIExpecter {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAndWait sends delete requests for a collection of objects and waits until
// the client cache reflects the changes.
func (a *APIExpecter) DeleteAndWait(ctx context.Context, objs ...client.Object) {
	_ = "STUB: not implemented"
	return
}

// namespaces can't be deleted with envtest

// DeleteAllOfAndWait deletes all objects of the given type and waits until the client's
// cache reflects those changes.
func (a *APIExpecter) DeleteAllOfAndWait(ctx context.Context, obj client.Object) {
	_ = "STUB: not implemented"
	return
}

// ShouldEventuallyExist defines an eventual expectation that succeeds if the provided object
// becomes readable by the client before the timeout expires.
func (a *APIExpecter) ShouldEventuallyExist(ctx context.Context, obj client.Object) {
	_ = "STUB: not implemented"
	return
}

// ShouldEventuallyMatch defines an eventual expectation that succeeds if the provided object
// becomes readable by the client and matches the provider expectation before the timeout expires.
func (a *APIExpecter) ShouldEventuallyMatch(ctx context.Context, obj client.Object, match func(g gomega.Gomega)) {
	_ = "STUB: not implemented"
	return
}

// CloneNameNamespace returns an empty client object of the same type
// with the same and namespace. This is a helper to pass a new object to the "Eventually"
// methods while preserving the original object's data.
func CloneNameNamespace[T any, PT interface {
	*T
	client.Object
}](obj PT,
) PT {
	_ = "STUB: not implemented"
	return *new(PT)
}

// ShouldEventuallyNotExist defines an eventual expectation that succeeds if the provided object
// becomes not found by the client before the timeout expires.
func (a *APIExpecter) ShouldEventuallyNotExist(ctx context.Context, obj client.Object) {
	_ = "STUB: not implemented"
	return
}
