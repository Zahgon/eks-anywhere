package controllers

import (
	"context"

	"github.com/go-logr/logr"
	clusterctlv1 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
	"sigs.k8s.io/cluster-api/controllers/clustercache"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	awsiamconfigreconciler "github.com/aws/eks-anywhere/pkg/awsiamauth/reconciler"
	mhcreconciler "github.com/aws/eks-anywhere/pkg/clusterapi/machinehealthcheck/reconciler"
	"github.com/aws/eks-anywhere/pkg/controller/clusters"
	"github.com/aws/eks-anywhere/pkg/curatedpackages"
	"github.com/aws/eks-anywhere/pkg/dependencies"
	"github.com/aws/eks-anywhere/pkg/networking/cilium"
	cnireconciler "github.com/aws/eks-anywhere/pkg/networking/reconciler"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack"
	cloudstackreconciler "github.com/aws/eks-anywhere/pkg/providers/cloudstack/reconciler"
	dockerreconciler "github.com/aws/eks-anywhere/pkg/providers/docker/reconciler"
	nutanixreconciler "github.com/aws/eks-anywhere/pkg/providers/nutanix/reconciler"
	snowreconciler "github.com/aws/eks-anywhere/pkg/providers/snow/reconciler"
	tinkerbellreconciler "github.com/aws/eks-anywhere/pkg/providers/tinkerbell/reconciler"
	vspherereconciler "github.com/aws/eks-anywhere/pkg/providers/vsphere/reconciler"
)

type Manager = manager.Manager

type Factory struct {
	buildSteps                   []buildStep
	dependencyFactory            *dependencies.Factory
	manager                      Manager
	registryBuilder              *clusters.ProviderClusterReconcilerRegistryBuilder
	reconcilers                  Reconcilers
	tracker                      clustercache.ClusterCache
	registry                     *clusters.ProviderClusterReconcilerRegistry
	dockerClusterReconciler      *dockerreconciler.Reconciler
	vsphereClusterReconciler     *vspherereconciler.Reconciler
	tinkerbellClusterReconciler  *tinkerbellreconciler.Reconciler
	snowClusterReconciler        *snowreconciler.Reconciler
	cloudstackClusterReconciler  *cloudstackreconciler.Reconciler
	nutanixClusterReconciler     *nutanixreconciler.Reconciler
	cniReconciler                *cnireconciler.Reconciler
	ipValidator                  *clusters.IPValidator
	awsIamConfigReconciler       *awsiamconfigreconciler.Reconciler
	machineHealthCheckReconciler *mhcreconciler.Reconciler
	logger                       logr.Logger
	deps                         *dependencies.Dependencies
	packageControllerClient      *curatedpackages.PackageControllerClient
	cloudStackValidatorRegistry  cloudstack.ValidatorRegistry
	ciliumTemplater              *cilium.Templater
	helmClientFactory            cilium.HelmClientFactory
}

type Reconcilers struct {
	ClusterReconciler                  *ClusterReconciler
	DockerDatacenterReconciler         *DockerDatacenterReconciler
	VSphereDatacenterReconciler        *VSphereDatacenterReconciler
	SnowMachineConfigReconciler        *SnowMachineConfigReconciler
	TinkerbellDatacenterReconciler     *TinkerbellDatacenterReconciler
	CloudStackDatacenterReconciler     *CloudStackDatacenterReconciler
	NutanixDatacenterReconciler        *NutanixDatacenterReconciler
	KubeadmControlPlaneReconciler      *KubeadmControlPlaneReconciler
	MachineDeploymentReconciler        *MachineDeploymentReconciler
	ControlPlaneUpgradeReconciler      *ControlPlaneUpgradeReconciler
	MachineDeploymentUpgradeReconciler *MachineDeploymentUpgradeReconciler
	NodeUpgradeReconciler              *NodeUpgradeReconciler
}

type buildStep func(ctx context.Context) error

func NewFactory(logger logr.Logger, manager Manager) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) Build(ctx context.Context) (*Reconcilers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close cleans up any open resources from the created dependencies.
func (f *Factory) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// WithClusterReconciler builds the cluster reconciler.
func (f *Factory) WithClusterReconciler(capiProviders []clusterctlv1.Provider, opts ...ClusterReconcilerOption) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithDockerDatacenterReconciler adds the DockerDatacenterReconciler to the controller factory.
func (f *Factory) WithDockerDatacenterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithVSphereDatacenterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithSnowMachineConfigReconciler() *Factory { _ = "STUB: not implemented"; return nil }

// WithTinkerbellDatacenterReconciler adds the TinkerbellDatacenterReconciler to the controller factory.
func (f *Factory) WithTinkerbellDatacenterReconciler() *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithCloudStackDatacenterReconciler adds the CloudStackDatacenterReconciler to the controller factory.
func (f *Factory) WithCloudStackDatacenterReconciler() *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithNutanixDatacenterReconciler adds the NutanixDatacenterReconciler to the controller factory.
func (f *Factory) WithNutanixDatacenterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

// withNutanixClusterReconciler adds the NutanixClusterReconciler to the controller factory.
func (f *Factory) withNutanixClusterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withTracker() *Factory { _ = "STUB: not implemented"; return nil }

const (
	dockerProviderName     = "docker"
	snowProviderName       = "snow"
	vSphereProviderName    = "vsphere"
	tinkerbellProviderName = "tinkerbell"
	cloudstackProviderName = "cloudstack"
	nutanixProviderName    = "nutanix"
)

func (f *Factory) WithProviderClusterReconcilerRegistry(capiProviders []clusterctlv1.Provider) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) withDockerClusterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withVSphereClusterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withSnowClusterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withTinkerbellClusterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withCloudStackClusterReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withCloudStackValidatorRegistry() *Factory { _ = "STUB: not implemented"; return nil }

// withHelmClientFactory configures the HelmClientFactory dependency with a helm.ClientFactory.
func (f *Factory) withHelmClientFactory() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withCiliumTemplater() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withCNIReconciler(providerNamespace string) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) withIPValidator() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withAWSIamConfigReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withPackageControllerClient() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) withMachineHealthCheckReconciler() *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithKubeadmControlPlaneReconciler builds the KubeadmControlPlane reconciler.
func (f *Factory) WithKubeadmControlPlaneReconciler() *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithMachineDeploymentReconciler builds the MachineDeployment reconciler.
func (f *Factory) WithMachineDeploymentReconciler() *Factory { _ = "STUB: not implemented"; return nil }

// WithControlPlaneUpgradeReconciler builds the ControlPlaneUpgrade reconciler.
func (f *Factory) WithControlPlaneUpgradeReconciler() *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithMachineDeploymentUpgradeReconciler builds the WithMachineDeploymentUpgrade reconciler.
func (f *Factory) WithMachineDeploymentUpgradeReconciler() *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithNodeUpgradeReconciler builds the WithNodeUpgrade reconciler.
func (f *Factory) WithNodeUpgradeReconciler() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) getProviderNamespace(providerName string) string {
	_ = "STUB: not implemented"
	return ""
}
