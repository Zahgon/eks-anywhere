package tinkerbell

import (
	"context"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/types"
)

func (p *Provider) BootstrapClusterOpts(_ *cluster.Spec) ([]bootstrapper.BootstrapClusterOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Provider) PreCAPIInstallOnBootstrap(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// We add this annotation to pass the admin machine URL to the controller for cluster creation.

// enable host network on bootstrap cluster

func (p *Provider) PostBootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// ApplyHardwareToCluster adds all the hardwares to the cluster.
func (p *Provider) applyHardware(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Provider) PostWorkloadInit(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// disable host network on workload cluster
// use stack service on workload cluster

// load balancer is handled by kube-vip in control plane nodes
// configure load balancer based on datacenterConfig.Spec.SkipLoadBalancerDeployment

func (p *Provider) SetupAndValidateCreateCluster(ctx context.Context, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(chrisdoherty4) Extract to a defaulting construct and add associated validations to ensure
// there is always a user with ssh key configured.

// for workload cluster use tinkerbell IP of the management cluster

// Checking for empty first as that returns a different error in the datacenter config validate method below

// TODO(chrisdoherty4) Look to inject the validator. Possibly look to use a builder for
// constructing the validations rather than injecting flags into the provider.

// Validate must happen last beacuse we depend on the catalogue entries for some checks.

func (p *Provider) getHardwareFromManagementCluster(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	// Retrieve all unprovisioned hardware from the management cluster and populate the catalogue so
	// it can be considered for the workload creation.
	return nil
}

// Retrieve all provisioned hardware from the management cluster and populate diskExtractors's
// disksProvisionedHardware map for use during workload creation

// Remove all the provisioned hardware from the existing cluster if repeated from the hardware csv input.

func (p *Provider) readCSVToCatalogue() error {
	_ = "STUB: not implemented"
	// Create a catalogue writer used to write hardware to the catalogue.
	return nil
}

// Translate all Machine instances from the p.machines source into Kubernetes object types.
// The PostBootstrapSetup() call invoked elsewhere in the program serializes the catalogue
// and submits it to the clsuter.
