//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

type MetalLBSuite struct {
	suite.Suite
	cluster           *framework.ClusterE2ETest
	kubernetesVersion v1alpha1.KubernetesVersion
	provider          framework.Provider
}

func RunMetalLBDockerTestsForKubeVersion(t *testing.T, kubeVersion v1alpha1.KubernetesVersion) {
	_ = "STUB: not implemented"
	return
}

func (suite *MetalLBSuite) SetupSuite() { _ = "STUB: not implemented"; return }

func kubeVersionNameDiscriminator(version v1alpha1.KubernetesVersion) framework.ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(framework.ClusterE2ETestOpt)
}

func getIPAddressPoolSpec(addresses []string, autoAssign bool) string {
	_ = "STUB: not implemented"
	return ""
}

func getL2AdvertisementSpec(ipPoolNames []string) string { _ = "STUB: not implemented"; return "" }

func getBGPAdvertisementSpec(ipPoolNames []string) string { _ = "STUB: not implemented"; return "" }

func (suite *MetalLBSuite) TestPackagesMetalLB() {
	_ = "STUB: not implemented"
	// This should be split into multiple tests with a cluster setup in `SetupSuite`.
	// This however requires the creation of utilites managing cluster creation.
	return
}
