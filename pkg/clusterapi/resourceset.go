package clusterapi

import (
	addons "sigs.k8s.io/cluster-api/api/addons/v1beta2"
)

type ClusterResourceSet struct {
	resources   map[string][]byte
	clusterName string
	namespace   string
}

func NewClusterResourceSet(clusterName string) *ClusterResourceSet {
	_ = "STUB: not implemented"
	return nil
}

func (c ClusterResourceSet) AddResource(name string, content []byte) {
	_ = "STUB: not implemented"
	return
}

func (c ClusterResourceSet) ToYaml() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c ClusterResourceSet) buildSet() *addons.ClusterResourceSet {
	_ = "STUB: not implemented"
	return nil
}

func (c ClusterResourceSet) resourceRefs() []addons.ResourceRef {
	_ = "STUB: not implemented"
	return nil
}

func (c ClusterResourceSet) buildResourceConfigMaps() []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func marshall(objects ...interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
