package dependencies

import (
	"context"
	"time"

	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/awsiamauth"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cli"
	"github.com/aws/eks-anywhere/pkg/clients/kubernetes"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	"github.com/aws/eks-anywhere/pkg/clustermanager"
	cliconfig "github.com/aws/eks-anywhere/pkg/config"
	"github.com/aws/eks-anywhere/pkg/curatedpackages"
	"github.com/aws/eks-anywhere/pkg/diagnostics"
	"github.com/aws/eks-anywhere/pkg/eksd"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/files"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	gitfactory "github.com/aws/eks-anywhere/pkg/git/factory"
	"github.com/aws/eks-anywhere/pkg/gitops/flux"
	"github.com/aws/eks-anywhere/pkg/helm"
	"github.com/aws/eks-anywhere/pkg/kubeconfig"
	"github.com/aws/eks-anywhere/pkg/manifests"
	"github.com/aws/eks-anywhere/pkg/networking/cilium"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack"
	"github.com/aws/eks-anywhere/pkg/providers/nutanix"
	"github.com/aws/eks-anywhere/pkg/providers/snow"
	"github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
	"github.com/aws/eks-anywhere/pkg/providers/validator"
	"github.com/aws/eks-anywhere/pkg/providers/vsphere"
	"github.com/aws/eks-anywhere/pkg/registrymirror"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/version"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces"
)

type Dependencies struct {
	Logger                      logr.Logger
	Provider                    providers.Provider
	ClusterAwsCli               *executables.Clusterawsadm
	DockerClient                *executables.Docker
	Kubectl                     *executables.Kubectl
	Govc                        *executables.Govc
	CloudStackValidatorRegistry cloudstack.ValidatorRegistry
	SnowAwsClientRegistry       *snow.AwsClientRegistry
	SnowConfigManager           *snow.ConfigManager
	Writer                      filewriter.FileWriter
	Kind                        *executables.Kind
	Clusterctl                  *executables.Clusterctl
	Flux                        *executables.Flux
	Troubleshoot                *executables.Troubleshoot
	Helm                        *executables.Helm
	UnAuthKubeClient            *kubernetes.UnAuthClient
	CiliumTemplater             *cilium.Templater
	AwsIamAuth                  *awsiamauth.Installer
	ClusterManager              *clustermanager.ClusterManager
	KubernetesRetrierClient     *clustermanager.KubernetesRetrierClient
	Bootstrapper                *bootstrapper.Bootstrapper
	GitOpsFlux                  *flux.Flux
	Git                         *gitfactory.GitTools
	EksdInstaller               *eksd.Installer
	EksdUpgrader                *eksd.Upgrader
	ClusterApplier              clustermanager.Applier
	AnalyzerFactory             diagnostics.AnalyzerFactory
	CollectorFactory            diagnostics.CollectorFactory
	DignosticCollectorFactory   diagnostics.DiagnosticBundleFactory
	CAPIManager                 *clusterapi.Manager
	FileReader                  *files.Reader
	ManifestReader              *manifests.Reader
	closers                     []types.Closer
	CliConfig                   *cliconfig.CliConfig
	CreateCliConfig             *cliconfig.CreateClusterCLIConfig
	PackageManager              interfaces.PackageManager
	BundleRegistry              curatedpackages.BundleRegistry
	PackageControllerClient     *curatedpackages.PackageControllerClient
	PackageClient               curatedpackages.PackageHandler
	VSphereValidator            *vsphere.Validator
	VSphereDefaulter            *vsphere.Defaulter
	NutanixClientCache          *nutanix.ClientCache
	NutanixDefaulter            *nutanix.Defaulter
	NutanixValidator            *nutanix.Validator
	SnowValidator               *snow.Validator
	IPValidator                 *validator.IPValidator
	UnAuthKubectlClient         KubeClients
	HelmEnvClientFactory        *helm.EnvClientFactory
	ExecutableBuilder           *executables.ExecutablesBuilder
	CreateClusterDefaulter      cli.CreateClusterDefaulter
	UpgradeClusterDefaulter     cli.UpgradeClusterDefaulter
	KubeconfigWriter            kubeconfig.Writer
	ClusterCreator              *clustermanager.ClusterCreator
	EksaInstaller               *clustermanager.EKSAInstaller
	DeleteClusterDefaulter      cli.DeleteClusterDefaulter
	ClusterDeleter              clustermanager.Deleter
	ClusterMover                *clustermanager.Mover
}

// KubeClients defines super struct that exposes all behavior.
type KubeClients struct {
	*executables.Kubectl
	*kubernetes.UnAuthClient
}

func (d *Dependencies) Close(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Reverse the loop so we close like LIFO
	return nil
}

// ForSpec constructs a Factory using the bundle referenced by clusterSpec.
func ForSpec(clusterSpec *cluster.Spec) *Factory { _ = "STUB: not implemented"; return nil }

// Factory helps initialization.
type Factory struct {
	executablesConfig        *executablesConfig
	config                   config
	registryMirror           *registrymirror.RegistryMirror
	proxyConfiguration       map[string]string
	writerFolder             string
	diagnosticCollectorImage string
	buildSteps               []buildStep
	dependencies             Dependencies
}

type executablesConfig struct {
	builder            *executables.ExecutablesBuilder
	image              string
	useDockerContainer bool
	dockerClient       executables.DockerClient
	mountDirs          []string
}

type config struct {
	bundlesOverride string
	noTimeouts      bool
}

type buildStep func(ctx context.Context) error

func NewFactory() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) Build(ctx context.Context) (*Dependencies, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// clean up stack

// Make copy of dependencies since its attributes are public

func (f *Factory) WithWriterFolder(folder string) *Factory { _ = "STUB: not implemented"; return nil }

// WithRegistryMirror configures the factory to use registry mirror wherever applicable.
func (f *Factory) WithRegistryMirror(registryMirror *registrymirror.RegistryMirror) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) UseProxyConfiguration(proxyConfig map[string]string) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) GetProxyConfiguration() map[string]string { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithProxyConfiguration() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) UseExecutableImage(image string) *Factory { _ = "STUB: not implemented"; return nil }

// WithExecutableImage sets the right cli tools image for the executable builder, reading
// from the Bundle and using the first VersionsBundle
// This is just the default for when there is not an specific kubernetes version available
// For commands that receive a cluster config file or a kubernetes version directly as input,
// use UseExecutableImage to specify the image directly.
func (f *Factory) WithExecutableImage() *Factory { _ = "STUB: not implemented"; return nil }

// selectImageFromBundleOverride retrieves an image from a bundles override.
//
// Handles cases where the bundle is configured with an override.
func (f *Factory) selectImageFromBundleOverride(bundlesOverride string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Note: Currently using the first available version of the cli tools
// This is because the binaries bundled are all the same version hence no compatibility concerns
// In case, there is a change to this behavior, there might be a need to reassess this item

// WithCustomBundles allows configuring a bundle override.
func (f *Factory) WithCustomBundles(bundlesOverride string) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) WithExecutableMountDirs(mountDirs ...string) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) WithLocalExecutables() *Factory { _ = "STUB: not implemented"; return nil }

// UseExecutablesDockerClient forces a specific DockerClient to build
// Executables as opposed to follow the normal building flow
// This is only for testing.
func (f *Factory) UseExecutablesDockerClient(client executables.DockerClient) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// dockerLogin performs a docker login with the ENV VARS.
func dockerLogin(ctx context.Context, registry string, docker executables.DockerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// WithDockerLogin adds a docker login to the build steps.
func (f *Factory) WithDockerLogin() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithExecutableBuilder() *Factory {
	_ = "STUB: not implemented"
	// Ensure the file writer is created before the tools container is launched. This is necessary
	// because we bind mount the cluster directory into the tools container. If the directory
	// doesn't exist, dockerd (running as root) creates the hostpath for the bind mount with root
	// ownership. This prevents further files from being written to the cluster directory.
	return nil
}

// WithHelmExecutableBuilder adds a build step to initializes the helm.ExecutableBuilder dependency.
func (f *Factory) WithHelmExecutableBuilder() *Factory { _ = "STUB: not implemented"; return nil }

// ProviderOptions contains per provider options.
type ProviderOptions struct {
	// Tinkerbell contains Tinkerbell specific options.
	Tinkerbell *TinkerbellOptions
}

// TinkerbellOptions contains Tinkerbell specific options.
type TinkerbellOptions struct {
	// BMCOptions contains options for configuring BMC interactions.
	BMCOptions *hardware.BMCOptions
}

// WithProvider initializes the provider dependency and adds to the build steps.
func (f *Factory) WithProvider(clusterConfigFile string, clusterConfig *v1alpha1.Cluster, skipIPCheck bool, hardwareCSVPath string, force bool, tinkerbellBootstrapIP string, skippedValidations map[string]bool, opts *ProviderOptions) *Factory {
	_ = "STUB: not implemented" // nolint:gocyclo
	return nil
}

// Set BMC timeout based on noTimeouts flag
// Default 5 minutes

// 1 week, effectively no timeout

// WithKubeconfigWriter adds the KubeconfigReader dependency depending on the provider.
func (f *Factory) WithKubeconfigWriter(clusterConfig *v1alpha1.Cluster) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithClusterCreator adds the ClusterCreator dependency.
func (f *Factory) WithClusterCreator(clusterConfig *v1alpha1.Cluster) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) WithDocker() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithKubectl() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithGovc() *Factory { _ = "STUB: not implemented"; return nil }

// WithCloudStackValidatorRegistry initializes the CloudStack validator for the object being constructed to make it available in the constructor.
func (f *Factory) WithCloudStackValidatorRegistry(skipIPCheck bool) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) WithSnowConfigManager() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithAwsSnow() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithWriter() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithKind() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithClusterctl() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithFlux() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithTroubleshoot() *Factory { _ = "STUB: not implemented"; return nil }

// WithHelm initializes a new Helm executable as a factory dependency.
func (f *Factory) WithHelm(opts ...helm.Opt) *Factory { _ = "STUB: not implemented"; return nil }

// WithHelmEnvClientFactory configures the HelmEnvClientFactory dependency with a helm.EnvClientFactory.
func (f *Factory) WithHelmEnvClientFactory(opts ...helm.Opt) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) WithCiliumTemplater() *Factory { _ = "STUB: not implemented"; return nil }

// WithAwsIamAuth builds dependencies for AWS IAM Auth.
func (f *Factory) WithAwsIamAuth(clusterConfig *v1alpha1.Cluster) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithIPValidator builds the IPValidator for the given cluster.
func (f *Factory) WithIPValidator() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithBootstrapper() *Factory { _ = "STUB: not implemented"; return nil }

type clusterManagerClient struct {
	*executables.Clusterctl
	*clustermanager.KubernetesRetrierClient
}

// ClusterManagerTimeoutOptions maintains the timeout options for cluster manager.
type ClusterManagerTimeoutOptions struct {
	NoTimeouts bool

	ControlPlaneWait, ExternalEtcdWait, MachineWait time.Duration
}

func (f *Factory) eksaInstallerOpts() []clustermanager.EKSAInstallerOpt {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) clusterManagerOpts(timeoutOpts *ClusterManagerTimeoutOptions) []clustermanager.ClusterManagerOpt {
	_ = "STUB: not implemented"
	return nil
}

// WithClusterManager builds a cluster manager based on the cluster config and timeout options.
func (f *Factory) WithClusterManager(clusterConfig *v1alpha1.Cluster, timeoutOpts *ClusterManagerTimeoutOptions) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithEKSAInstaller builds a cluster manager based on the cluster config and timeout options.
func (f *Factory) WithEKSAInstaller() *Factory { _ = "STUB: not implemented"; return nil }

// WithKubernetesRetrierClient builds a cluster manager based on the cluster config and timeout options.
func (f *Factory) WithKubernetesRetrierClient() *Factory { _ = "STUB: not implemented"; return nil }

// WithNoTimeouts injects no timeouts to all the dependencies with configurable timeout.
// Calling this method sets no timeout for the waits and retries in all the
// cluster operations, i.e. cluster manager, eksa installer, networking installer.
// Instead of passing the option to each dependency's constructor, use this
// method to pass no timeouts to new dependency.
func (f *Factory) WithNoTimeouts() *Factory { _ = "STUB: not implemented"; return nil }

// WithCliConfig builds a cli config.
func (f *Factory) WithCliConfig(cliConfig *cliconfig.CliConfig) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithCreateClusterDefaulter builds a create cluster defaulter that builds defaulter dependencies specific to the create cluster command. The defaulter is then run once the factory is built in the create cluster command.
func (f *Factory) WithCreateClusterDefaulter(createCliConfig *cliconfig.CreateClusterCLIConfig) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithUpgradeClusterDefaulter builds a create cluster defaulter that builds defaulter dependencies specific to the create cluster command. The defaulter is then run once the factory is built in the create cluster command.
func (f *Factory) WithUpgradeClusterDefaulter(upgradeCliConfig *cliconfig.UpgradeClusterCLIConfig) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithDeleteClusterDefaulter builds a delete cluster defaulter that builds defaulter dependencies specific to the delete cluster command. The defaulter is then run once the factory is built in the delete cluster command.
func (f *Factory) WithDeleteClusterDefaulter(deleteCliConfig *cliconfig.DeleteClusterCLIConfig) *Factory {
	_ = "STUB: not implemented"
	return nil
}

type eksdInstallerClient struct {
	*executables.Kubectl
}

func (f *Factory) WithEksdInstaller() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithEksdUpgrader() *Factory { _ = "STUB: not implemented"; return nil }

// WithClusterApplier builds a cluster applier.
func (f *Factory) WithClusterApplier() *Factory { _ = "STUB: not implemented"; return nil }

// opts = append(opts, clustermanager.ManagementUpgraderRetrier(*retrier.NewWithNoTimeout()))

// WithClusterDeleter builds a cluster deleter.
func (f *Factory) WithClusterDeleter() *Factory { _ = "STUB: not implemented"; return nil }

// WithClusterMover builds a cluster mover.
func (f *Factory) WithClusterMover() *Factory { _ = "STUB: not implemented"; return nil }

// WithValidatorClients builds KubeClients.
func (f *Factory) WithValidatorClients() *Factory { _ = "STUB: not implemented"; return nil }

// WithLogger setups a logger to be injected in constructors. It uses the logger
// package level logger.
func (f *Factory) WithLogger() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithGit(clusterConfig *v1alpha1.Cluster, fluxConfig *v1alpha1.FluxConfig) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithGitOpsFlux builds a gitops flux.
func (f *Factory) WithGitOpsFlux(clusterConfig *v1alpha1.Cluster, fluxConfig *v1alpha1.FluxConfig, cliConfig *cliconfig.CliConfig) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithPackageManager builds a package manager.
func (f *Factory) WithPackageManager(spec *cluster.Spec, packagesLocation, kubeConfig string) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithPackageManagerWithoutWait builds a package manager that doesn't wait for active bundles.
func (f *Factory) WithPackageManagerWithoutWait(spec *cluster.Spec, packagesLocation, kubeConfig string) *Factory {
	_ = "STUB: not implemented"
	return nil
}

// WithPackageControllerClient builds a client for package controller.
func (f *Factory) WithPackageControllerClient(spec *cluster.Spec, kubeConfig string, opts ...curatedpackages.PackageControllerClientOpt) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) WithPackageClient() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithCuratedPackagesRegistry(registryName, kubeVersion string, version version.Info) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) WithDiagnosticBundleFactory() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithAnalyzerFactory() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithDiagnosticCollectorImage(diagnosticCollectorImage string) *Factory {
	_ = "STUB: not implemented"
	return nil
}

func (f *Factory) WithCollectorFactory() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithCAPIManager() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithFileReader() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithManifestReader() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithUnAuthKubeClient() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithVSphereValidator() *Factory { _ = "STUB: not implemented"; return nil }

func (f *Factory) WithVSphereDefaulter() *Factory { _ = "STUB: not implemented"; return nil }

// WithNutanixDefaulter adds a new NutanixDefaulter to the factory.
func (f *Factory) WithNutanixDefaulter() *Factory { _ = "STUB: not implemented"; return nil }

// WithNutanixValidator adds a new NutanixValidator to the factory.
func (f *Factory) WithNutanixValidator() *Factory { _ = "STUB: not implemented"; return nil }

// WithNutanixClientCache adds a new NutanixClientCache to the factory.
func (f *Factory) WithNutanixClientCache() *Factory { _ = "STUB: not implemented"; return nil }

func getProxyConfiguration(clusterSpec *cluster.Spec) (httpProxy, httpsProxy string, noProxy []string) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func getManagementClusterName(clusterSpec *cluster.Spec) string {
	_ = "STUB: not implemented"
	return ""
}
