//go:build e2e
// +build e2e

package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/test/framework"
)

func runCuratedPackageInstall(test *framework.ClusterE2ETest) { _ = "STUB: not implemented"; return }

func runCuratedPackageInstallSimpleFlow(test *framework.ClusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

func runDisabledCuratedPackage(test *framework.ClusterE2ETest) { _ = "STUB: not implemented"; return }

func runDisabledCuratedPackageInstallSimpleFlow(test *framework.ClusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

// addDefaultOCINamespacesFromEnv adds default OCI namespaces from environment variables if they're not already set.
func addDefaultOCINamespacesFromEnv(test *framework.ClusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

func runCuratedPackageInstallSimpleFlowRegistryMirror(test *framework.ClusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

// Determine the correct alias from the bundle after download

// Add default OCI namespaces from env vars, then append the curated packages one

func WaitForPackageNamespace(test *framework.ClusterE2ETest, ctx context.Context, mgmtKubeconfig string, workloadClusterName string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func runCuratedPackageRemoteClusterInstallSimpleFlow(test *framework.MulticlusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

func runCuratedPackageInstallTinkerbellSingleNodeFlow(test *framework.ClusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

type resourcePredicate func(string, error) bool

func NoErrorPredicate(_ string, err error) bool {
	_ = "STUB: not implemented"

	// TODO turn them into generics using comparable once 1.18 is allowed
	return false
}

func StringMatchPredicate(s string) resourcePredicate {
	_ = "STUB: not implemented"
	return *new(resourcePredicate)
}

func IntEqualPredicate(i int) resourcePredicate {
	_ = "STUB: not implemented"
	return *new(resourcePredicate)
}

func WaitForResource(
	test *framework.ClusterE2ETest,
	ctx context.Context,
	resource string,
	namespace string,
	jsonpath string,
	timeout time.Duration,
	predicates ...resourcePredicate,
) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitForDaemonset(
	test *framework.ClusterE2ETest,
	ctx context.Context,
	daemonsetName string,
	namespace string,
	numberOfNodes int,
	timeout time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Hackish way to get the latest bundle. This assumes no bundle is created outside of the normal PBC bundle fetch timer.
// This should be modified to get the bundle from the previous build step and use that only.
func WaitForLatestBundleToBeAvailable(
	test *framework.ClusterE2ETest,
	ctx context.Context,
	timeout time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitForPackageToBeInstalled(
	test *framework.ClusterE2ETest,
	ctx context.Context,
	packageName string,
	timeout time.Duration,
) error {
	_ = "STUB: not implemented"
	//--for=jsonpath isn't supported in v1.22. Update once it's supported
	//_, err = test.KubectlClient.Execute(
	//    ctx, "wait", "--timeout", "1m",
	//    "--for", "jsonpath='{.status.state}'=installed",
	//    "package", packagePrefix, "--kubeconfig", kubeconfig,
	//    "-n", "eksa-packages",
	//)
	return nil
}

func GetLatestBundleFromCluster(test *framework.ClusterE2ETest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func withCluster(cluster *framework.ClusterE2ETest) *types.Cluster {
	_ = "STUB: not implemented"
	return nil
}

func SetupSimpleMultiCluster(t *testing.T, provider framework.Provider, kubeVersion v1alpha1.KubernetesVersion) *framework.MulticlusterE2ETest {
	_ = "STUB: not implemented"
	return nil
}
