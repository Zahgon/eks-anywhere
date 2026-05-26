package framework

const (
	kubeConformanceImage                  = "registry.k8s.io/conformance"
	minKubernetesVersionRequiringTestSkip = "v1.29.0"
	skippedTestName                       = "Services should serve endpoints on same port and different protocols"
)

func (e *ClusterE2ETest) RunConformanceTests() { _ = "STUB: not implemented"; return }

// If running conformance tests for Kubernetes 1.29 or above, skip this particular test
// because it will not pass with our deployment of Cilium.
// References:
// 1. https://github.com/kubernetes/kubernetes/pull/120069
// 2. https://github.com/cilium/cilium/issues/29913
// 3. https://github.com/cncf/k8s-conformance/pull/3049

// Only mode or --e2e-skip can be used at a time. Because we are using --e2e-skip
// for k8s 1.29 and higher, we need to skip --mode=certified-conformance for those versions.
// Once we stop skipping e2e, we can add the mode back to k8s 1.29 and higher.

func (e *ClusterE2ETest) getEksdReleaseKubeVersion() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Function to parse the conformace test results and look for any failed tests.
// By default we run 2 plugins so we check for failed tests in twice.
func hasFailed(results string) bool { _ = "STUB: not implemented"; return false }
