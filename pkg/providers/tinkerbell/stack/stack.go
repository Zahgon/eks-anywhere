package stack

import (
	"context"
	"net/url"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/helm"
	"github.com/aws/eks-anywhere/pkg/registrymirror"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	overridesFileName = "tinkerbell-chart-overrides.yaml"

	smee         = "smee"
	smeeHTTPPort = "7171"
	grpcPort     = "42113"

	// localTinkWorkerImage is the path to the tink-worker image embedded in HookOS.
	localTinkWorkerImage         = "127.0.0.1/embedded/tink-worker"
	rufioMaxConcurrentReconciles = 10
	tinkMaxConcurrentReconciles  = 5
)

type Docker interface {
	CheckContainerExistence(ctx context.Context, name string) (bool, error)
	ForceRemove(ctx context.Context, name string) error
	Run(ctx context.Context, image string, name string, cmd []string, flags ...string) error
}

type Helm interface {
	RegistryLogin(ctx context.Context, endpoint, username, password string) error
	UpgradeInstallChartWithValuesFile(ctx context.Context, chart, ociURI, version, kubeconfigFilePath, namespace, valuesFilePath string, opts ...helm.Opt) error
	Uninstall(ctx context.Context, chart, kubeconfigFilePath, namespace string, opts ...helm.Opt) error
	ListCharts(ctx context.Context, kubeconfigFilePath, filter, namespace string) ([]string, error)
}

// StackInstaller deploys a Tinkerbell stack.
//
//nolint:revive // Stutter and the interface shouldn't exist. Will clean up (chrisdoherty4)
type StackInstaller interface {
	CleanupLocalBoots(ctx context.Context, forceCleanup bool) error
	Install(ctx context.Context, bundle releasev1alpha1.TinkerbellBundle, tinkerbellIP, kubeconfig, hookOverride string, opts ...InstallOption) error
	UninstallLocal(ctx context.Context) error
	UninstallChart(ctx context.Context, chartName, kubeconfig, namespace string) error
	Upgrade(_ context.Context, _ releasev1alpha1.TinkerbellBundle, tinkerbellIP, kubeconfig, hookOverride string, opts ...InstallOption) error
	UpgradeInstallCRDs(ctx context.Context, bundle releasev1alpha1.TinkerbellBundle, kubeconfig string, opts ...InstallOption) error
	UpgradeLegacy(ctx context.Context, bundle releasev1alpha1.TinkerbellBundle, kubeconfig string, opts ...InstallOption) error
	AddNoProxyIP(IP string)
	GetNamespace() string
	HasChart(ctx context.Context, chartName, kubeconfig, namespace string) (bool, error)
}

type Installer struct {
	docker                Docker
	filewriter            filewriter.FileWriter
	helm                  Helm
	podCidrRange          string
	registryMirror        *registrymirror.RegistryMirror
	proxyConfig           *v1alpha1.ProxyConfiguration
	namespace             string
	loadBalancerInterface string
	hookIsoURL            string
	smeeOnDocker          bool
	hostNetwork           bool
	loadBalancer          bool
	stackService          bool
	dhcpRelay             bool
}

type InstallOption func(s *Installer)

// WithLoadBalancerInterface is an InstallOption that allows you to configure load balancer interface for the tinkerbell stack.
func WithLoadBalancerInterface(loadBalancerInterface string) InstallOption {
	_ = "STUB: not implemented"
	return *new(InstallOption)
}

// WithSmeeOnDocker is an InstallOption to run Boots as a Docker container.
func WithSmeeOnDocker() InstallOption { _ = "STUB: not implemented"; return *new(InstallOption) }

// WithSmeeOnKubernetes is an InstallOption to run Boots as a Kubernetes deployment.
func WithSmeeOnKubernetes() InstallOption { _ = "STUB: not implemented"; return *new(InstallOption) }

// WithHostNetworkEnabled is an InstallOption that allows you to enable/disable host network for Tinkerbell deployments.
func WithHostNetworkEnabled(enabled bool) InstallOption {
	_ = "STUB: not implemented"
	return *new(InstallOption)
}

// WithLoadBalancerEnabled is an InstallOption that allows you to setup a LoadBalancer to expose hegel and tink-server.
func WithLoadBalancerEnabled(enabled bool) InstallOption {
	_ = "STUB: not implemented"
	return *new(InstallOption)
}

// WithStackServiceEnabled is an InstallOption that allows you to enable the nginx service as a reverse proxy.
func WithStackServiceEnabled(enabled bool) InstallOption {
	_ = "STUB: not implemented"
	return *new(InstallOption)
}

// WithDHCPRelayEnabled is an InstallOption that allows you to enable DHCP Relay.
func WithDHCPRelayEnabled(enabled bool) InstallOption {
	_ = "STUB: not implemented"
	return *new(InstallOption)
}

// WithHookIsoOverride is an InstallOption allows you to set a URL of the HookOS ISO image.
func WithHookIsoOverride(url string) InstallOption {
	_ = "STUB: not implemented"
	return *new(InstallOption)
}

// AddNoProxyIP is for workload cluster upgrade, we have to pass
// controlPlaneEndpoint IP of managemement cluster if proxy is configured.
func (s *Installer) AddNoProxyIP(IP string) { _ = "STUB: not implemented"; return }

// NewInstaller returns a Tinkerbell StackInstaller which can be used to install or uninstall the Tinkerbell stack.
func NewInstaller(docker Docker, filewriter filewriter.FileWriter, helm Helm, hookIsoURL, namespace, podCidrRange string, registryMirror *registrymirror.RegistryMirror, proxyConfig *v1alpha1.ProxyConfiguration) StackInstaller {
	_ = "STUB: not implemented"
	return *new(StackInstaller)
}

// Install installs the Tinkerbell stack on a target cluster using a helm chart and providing the necessary values overrides.
func (s *Installer) Install(ctx context.Context, bundle releasev1alpha1.TinkerbellBundle, tinkerbellIP, kubeconfig, hookOverride string, opts ...InstallOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Installer) installSmeeOnDocker(ctx context.Context, bundle releasev1alpha1.TinkerbellStackBundle, tinkServerIP, kubeconfig, hookOverride, isoOverride string) error {
	_ = "STUB: not implemented"
	return nil
}

// Mono-repo tinkerbell binary uses TINKERBELL_ prefix for env vars

// Mono-repo binary uses env vars only, no command line args

func (s *Installer) getSmeeKernelArgs(_ releasev1alpha1.TinkerbellStackBundle) []string {
	_ = "STUB: not implemented"
	return nil
}

// UninstallLocal currently removes local docker container running Boots.
func (s *Installer) UninstallLocal(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Installer) uninstallBootsFromDocker(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func getURIDir(uri string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// CleanupLocalBoots determines whether Boots is already running locally
// and either cleans it up or errors out depending on the `remove` flag.
func (s *Installer) CleanupLocalBoots(ctx context.Context, remove bool) error {
	_ = "STUB: not implemented"
	return nil
}

// return error if the docker call failed

// return nil if boots container doesn't exist

// if remove is set, try to delete boots

// finally, return an "already exists" error if boots exists and forceCleanup is not set

func (s *Installer) localRegistryURL(originalURL string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *Installer) authenticateHelmRegistry(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Upgrade the Tinkerbell stack using images specified in bundle.
func (s *Installer) Upgrade(ctx context.Context, bundle releasev1alpha1.TinkerbellBundle, tinkerbellIP, kubeconfig, hookOverride string, opts ...InstallOption) error {
	_ = "STUB: not implemented"
	return nil
}

// UpgradeInstallCRDs the Tinkerbell CRDs using images specified in bundle.
func (s *Installer) UpgradeInstallCRDs(ctx context.Context, bundle releasev1alpha1.TinkerbellBundle, kubeconfig string, opts ...InstallOption) error {
	_ = "STUB: not implemented"
	return nil
}

// UninstallChart uninstalls a tinkerbell helm chart by name.
func (s *Installer) UninstallChart(ctx context.Context, chartName, kubeconfig, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetNamespace retrieves the namespace the installer is using for stack deployment.
func (s *Installer) GetNamespace() string {
	_ = "STUB: not implemented"

	// UpgradeLegacy upgrades the legacy Tinkerbell stack using images specified in bundle.
	return ""
}

func (s *Installer) UpgradeLegacy(ctx context.Context, bundle releasev1alpha1.TinkerbellBundle, kubeconfig string, opts ...InstallOption) error {
	_ = "STUB: not implemented"
	return nil
}

// HasChart returns whether or not a helm chart with the given name exists on the cluster.
func (s *Installer) HasChart(ctx context.Context, chartName, kubeconfig, namespace string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// createValuesOverride generates the values override file for the mono-repo tinkerbell helm chart.
// The mono-repo chart uses a single deployment with all services (smee, tootles, tink-server,
// tink-controller, rufio) running in one pod, controlled via enable flags.
func (s *Installer) createValuesOverride(bundle releasev1alpha1.TinkerbellBundle, bootEnv []string, tinkerbellIP, loadBalancerInterface, isoURL string, osiePath *url.URL) map[string]any {
	_ = "STUB: not implemented"
	// Build OSIE URL from parsed path
	return nil
}

// Set load balancer interface if specified

// parseImageURI splits an image URI into image and tag components.
// Example: "public.ecr.aws/eks-anywhere/tinkerbell:v0.1.0" returns ("public.ecr.aws/eks-anywhere/tinkerbell", "v0.1.0").
func parseImageURI(uri string) (string, string) {
	_ = "STUB: not implemented"
	// Handle both tag (:) and digest (@) separators
	return "", ""
}

// Make sure we're not splitting on the port in the registry URL
// by checking if there's a / after the :
