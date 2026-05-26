package kubernetes

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClientFactory builds clients from a kubeconfig file by
// wrapping around NewRuntimeClientFromFileName to facilitate mocking.
type ClientFactory struct{}

// BuildClientFromKubeconfig builds a K8s client from a kubeconfig file.
func (f ClientFactory) BuildClientFromKubeconfig(kubeconfigPath string) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

// NewRuntimeClientFromFileName creates a new controller runtime client given a kubeconfig filename.
func NewRuntimeClientFromFileName(kubeConfigFilename string) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func initScheme(scheme *runtime.Scheme) error { _ = "STUB: not implemented"; return nil }

// InitScheme initializes a runtime scheme with all required types.
func InitScheme(scheme *runtime.Scheme) error { _ = "STUB: not implemented"; return nil }

func newRuntimeClient(data []byte, rc restConfigurator, scheme *runtime.Scheme) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

// restConfigurator abstracts the creation of a controller-runtime *rest.Config.
//
// This abstraction improves testing, as all known methods of instantiating a
// *rest.Config try to make network calls, and that's something we'd like to
// keep out of our unit tests as much as possible. In addition, where we do
// use them in unit tests, we need to be prepared with a controller-runtime
// EnvTest environment.
//
// For normal, non-test use, this can safely be ignored.
type restConfigurator func([]byte) (*rest.Config, error)

// Config generates and returns a rest.Config from a kubeconfig in bytes.
func (c restConfigurator) Config(data []byte) (*rest.Config, error) {
	_ = "STUB: not implemented"

	// ObjectsToRuntimeObjects converts objects of another type to runtime.Object's.
	return nil, nil
}

func ObjectsToRuntimeObjects[T runtime.Object](objs []T) []runtime.Object {
	_ = "STUB: not implemented"
	return nil
}
