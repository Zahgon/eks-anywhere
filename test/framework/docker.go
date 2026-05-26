package framework

import (
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/executables"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

// Docker is a Provider for running end-to-end tests.
type Docker struct {
	t *testing.T
	executables.Docker
}

const dockerPodCidrVar = "T_DOCKER_POD_CIDR"

// NewDocker creates a new Docker object implementing the Provider interface
// for testing.
func NewDocker(t *testing.T) *Docker { _ = "STUB: not implemented"; return nil }

// Name implements the Provider interface.
func (d *Docker) Name() string {
	_ = "STUB: not implemented"

	// Setup implements the Provider interface.
	return ""
}

func (d *Docker) Setup() {
	_ = "STUB: not implemented"

	// CleanupResources implements the Provider interface.
	return
}

func (d *Docker) CleanupResources(_ string) error {
	_ = "STUB: not implemented"

	// UpdateKubeConfig customizes generated kubeconfig by replacing the server value with correct host
	// and the docker LB port. This is required for the docker provider.
	return nil
}

func (d *Docker) UpdateKubeConfig(content *[]byte, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Docker) WithProviderUpgradeGit() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// There is no config for docker api objects, no-op

// ClusterConfigUpdates satisfies the test framework Provider.
func (d *Docker) ClusterConfigUpdates() []api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return nil
}

// WithNewWorkerNodeGroup returns an api.ClusterFiller that adds a new workerNodeGroupConfiguration and
// a corresponding DockerMachineConfig to the cluster config.
func (d *Docker) WithNewWorkerNodeGroup(machineConfig string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}

// ClusterStateValidations returns a list of provider specific validations.
func (d *Docker) ClusterStateValidations() []clusterf.StateValidation {
	_ = "STUB: not implemented"
	return nil
}

// WithKubeVersionAndOS returns a cluster config filler that sets the cluster kube version.
func (d *Docker) WithKubeVersionAndOS(kubeVersion anywherev1.KubernetesVersion, _ OS, _ *releasev1.EksARelease, _ ...string) api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return *new(api.ClusterConfigFiller)
}
