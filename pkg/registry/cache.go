package registry

// Cache storage client for an OCI registry.
type Cache struct {
	registries map[string]StorageClient
}

// NewCache creates an OCI registry client.
func NewCache() *Cache { _ = "STUB: not implemented"; return nil }

// Get cached registry client or make it.
func (cache *Cache) Get(context StorageContext) (StorageClient, error) {
	_ = "STUB: not implemented"
	return *new(StorageClient), nil
}

// Set a client in the cache.
func (cache *Cache) Set(registryName string, client StorageClient) {
	_ = "STUB: not implemented"
	return
}
