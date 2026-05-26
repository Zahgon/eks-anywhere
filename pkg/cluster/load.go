package cluster

import (
	"github.com/aws/eks-anywhere/pkg/types"
)

type kubeConfigCluster struct {
	Name string `json:"name"`
}

type kubeConfigYAML struct {
	Clusters []*kubeConfigCluster `json:"clusters"`
}

func LoadManagement(kubeconfig string) (*types.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
