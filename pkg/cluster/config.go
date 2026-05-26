package cluster

import (
	v1 "k8s.io/api/core/v1"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
)

type Config struct {
	Cluster                   *anywherev1.Cluster
	CloudStackDatacenter      *anywherev1.CloudStackDatacenterConfig
	VSphereDatacenter         *anywherev1.VSphereDatacenterConfig
	DockerDatacenter          *anywherev1.DockerDatacenterConfig
	SnowDatacenter            *anywherev1.SnowDatacenterConfig
	NutanixDatacenter         *anywherev1.NutanixDatacenterConfig
	TinkerbellDatacenter      *anywherev1.TinkerbellDatacenterConfig
	VSphereMachineConfigs     map[string]*anywherev1.VSphereMachineConfig
	CloudStackMachineConfigs  map[string]*anywherev1.CloudStackMachineConfig
	SnowMachineConfigs        map[string]*anywherev1.SnowMachineConfig
	NutanixMachineConfigs     map[string]*anywherev1.NutanixMachineConfig
	TinkerbellMachineConfigs  map[string]*anywherev1.TinkerbellMachineConfig
	TinkerbellTemplateConfigs map[string]*anywherev1.TinkerbellTemplateConfig
	OIDCConfigs               map[string]*anywherev1.OIDCConfig
	AWSIAMConfigs             map[string]*anywherev1.AWSIamConfig
	GitOpsConfig              *anywherev1.GitOpsConfig
	FluxConfig                *anywherev1.FluxConfig
	SnowCredentialsSecret     *v1.Secret
	SnowIPPools               map[string]*anywherev1.SnowIPPool
}

func (c *Config) VsphereMachineConfig(name string) *anywherev1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) CloudStackMachineConfig(name string) *anywherev1.CloudStackMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) SnowMachineConfig(name string) *anywherev1.SnowMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

// SnowIPPool returns a SnowIPPool based on a name.
func (c *Config) SnowIPPool(name string) *anywherev1.SnowIPPool {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) OIDCConfig(name string) *anywherev1.OIDCConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) AWSIamConfig(name string) *anywherev1.AWSIamConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) NutanixMachineConfig(name string) *anywherev1.NutanixMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) DeepCopy() *Config { _ = "STUB: not implemented"; return nil }

// ChildObjects returns all API objects in Config except the Cluster.
func (c *Config) ChildObjects() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

// machine configs length + datacenter + OIDC + IAM + gitops

// ClusterAndChildren returns all kubernetes objects in the cluster Config.
// It's equivalent to appending the Cluster to the result of ChildObjects.
func (c *Config) ClusterAndChildren() []kubernetes.Object { _ = "STUB: not implemented"; return nil }

func appendIfNotNil(objs []kubernetes.Object, elems ...kubernetes.Object) []kubernetes.Object {
	_ = "STUB: not implemented"
	return nil

	// Since we receive interfaces, these will never be nil since they contain
	// the type of the original implementing struct
	// I can't find another clean option of doing this
}
