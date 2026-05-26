// Package certificates provides functionality for managing and renewing certificates in EKS Anywhere clusters.
package certificates

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/types"
)

// VerbosityLevel controls the detail level of logging output.
var VerbosityLevel int

// SSHConfig holds the SSH credential information.
type SSHConfig struct {
	User     string `yaml:"sshUser"`
	KeyPath  string `yaml:"sshKey"`
	Password string `yaml:"-"` // enviroment vairables
}

// NodeConfig holds configuration for a group of nodes.
type NodeConfig struct {
	Nodes []string  `yaml:"nodes"`
	SSH   SSHConfig `yaml:"ssh"`
}

// RenewalConfig defines the configuration for certificate renewal operations.
type RenewalConfig struct {
	ClusterName           string     `yaml:"clusterName"`
	ManagementClusterName string     `yaml:"managementCluster"`
	OS                    string     `yaml:"os"`
	ControlPlane          NodeConfig `yaml:"controlPlane"`
	Etcd                  NodeConfig `yaml:"etcd"`
}

// ParseConfig reads and parses a certificate renewal configuration file.
func ParseConfig(path string) (*RenewalConfig, error) { _ = "STUB: not implemented"; return nil, nil }

// ValidateConfig validates the certificate renewal configuration and ensures all required fields are present.
func ValidateConfig(config *RenewalConfig, component string) error {
	_ = "STUB: not implemented"
	return nil
}

// Etcd nodes are only required if using external etcd.

func validateNodeConfig(config *NodeConfig) error { _ = "STUB: not implemented"; return nil }

// ValidateComponentWithConfig validates that the specified component is compatible with the configuration.
func ValidateComponentWithConfig(component string, config *RenewalConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// PopulateConfig fills in the configuration with control plane and etcd node IPs from the Kubernetes cluster.
func PopulateConfig(ctx context.Context, cfg *RenewalConfig, kubeClient kubernetes.Client, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func getControlPlaneIPs(ctx context.Context, kubeClient kubernetes.Client, cluster *types.Cluster) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getEtcdIPs(ctx context.Context, kubeClient kubernetes.Client, cluster *types.Cluster) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
