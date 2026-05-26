package framework

import (
	"context"
	_ "embed"
	"testing"
	"time"

	"github.com/bmc-toolbox/bmclib/v2"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/git"
	"github.com/aws/eks-anywhere/pkg/semver"
	"github.com/aws/eks-anywhere/pkg/types"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	clusterf "github.com/aws/eks-anywhere/test/framework/cluster"
)

const (
	defaultClusterConfigFile               = "cluster.yaml"
	defaultBundleReleaseManifestFile       = "bin/local-bundle-release.yaml"
	defaultEksaBinaryLocation              = "eksctl anywhere"
	defaultClusterName                     = "eksa-test"
	defaultDownloadArtifactsOutputLocation = "eks-anywhere-downloads.tar.gz"
	defaultDownloadImagesOutputLocation    = "images.tar"
	eksctlVersionEnvVar                    = "EKSCTL_VERSION"
	eksctlVersionEnvVarDummyVal            = "ham sandwich"
	ClusterPrefixVar                       = "T_CLUSTER_PREFIX"
	JobIdVar                               = "T_JOB_ID"
	BundlesOverrideVar                     = "T_BUNDLES_OVERRIDE"
	ClusterIPPoolEnvVar                    = "T_CLUSTER_IP_POOL"
	ClusterIPEnvVar                        = "T_CLUSTER_IP"
	CleanupResourcesVar                    = "T_CLEANUP_RESOURCES"
	LicenseTokenEnvVar                     = "LICENSE_TOKEN"
	LicenseToken2EnvVar                    = "LICENSE_TOKEN2"
	StagingLicenseTokenEnvVar              = "STAGING_LICENSE_TOKEN"
	StagingLicenseToken2EnvVar             = "STAGING_LICENSE_TOKEN2"
	hardwareYamlPath                       = "hardware.yaml"
	hardwareCsvPath                        = "hardware.csv"
	EksaPackagesInstallation               = "eks-anywhere-packages"
	bundleReleasePathFromArtifacts         = "./eks-anywhere-downloads/bundle-release.yaml"
	releaseV022                            = "v0.22.0"
)

//go:embed testdata/oidc-roles.yaml
var oidcRoles []byte

//go:embed testdata/autoscaler_load.yaml
var autoscalerLoad []byte

//go:embed testdata/local-path-storage.yaml
var localPathProvisioner []byte

type ClusterE2ETest struct {
	T                            T
	ClusterConfigLocation        string
	ClusterConfigFolder          string
	HardwareConfigLocation       string
	HardwareCsvLocation          string
	TestHardware                 map[string]*api.Hardware
	HardwarePool                 map[string]*api.Hardware
	WithNoPowerActions           bool
	WithOOBConfiguration         bool
	ClusterName                  string
	ClusterConfig                *cluster.Config
	clusterStateValidationConfig *clusterf.StateValidationConfig
	Provider                     Provider
	// TODO(g-gaston): migrate uses of clusterFillers to clusterConfigFillers
	clusterFillers       []api.ClusterFiller
	clusterConfigFillers []api.ClusterConfigFiller
	KubectlClient        *executables.Kubectl
	GitProvider          git.ProviderClient
	GitClient            git.Client
	GitWriter            filewriter.FileWriter
	eksaBinaryLocation   string
	OSFamily             v1alpha1.OSFamily
	ExpectFailure        bool
	// PersistentCluster avoids creating the clusters if it finds a kubeconfig
	// in the corresponding cluster folder. Useful for local development of tests.
	// When generating a new base cluster config, it will read from disk instead of
	// using the CLI generate command and will preserve the previous CP endpoint.
	PersistentCluster bool
}

type ClusterE2ETestOpt func(e *ClusterE2ETest)

// NewClusterE2ETest is a support structure for defining an end-to-end test.
func NewClusterE2ETest(t T, provider Provider, opts ...ClusterE2ETestOpt) *ClusterE2ETest {
	_ = "STUB: not implemented"
	return nil
}

// UpdateClusterName updates the cluster name for the test. This will drive both the name of the eks-a
// cluster config objects as well as the cluster config file name and the folder where the cluster config
// file is stored.
// The cluster config folder will be updated to the new cluster name only if it was using the default value.
func (e *ClusterE2ETest) UpdateClusterName(name string) { _ = "STUB: not implemented"; return }

// Only update the folder if it was using the old name. This is the default value.

func withHardware(requiredCount int, hardareType string, labels map[string]string) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithNoPowerActions() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func ExpectFailure(expected bool) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithControlPlaneHardware(requiredCount int) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithWorkerHardware(requiredCount int) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithCustomLabelHardware(requiredCount int, label string) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithExternalEtcdHardware(requiredCount int) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithClusterName sets the name that will be used for the cluster. This will drive both the name of the eks-a
// cluster config objects as well as the cluster config file name.
func WithClusterName(name string) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// PersistentCluster  avoids creating the clusters if it finds a kubeconfig
// in the corresponding cluster folder. Useful for local development of tests.
func PersistentCluster() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func (e *ClusterE2ETest) GetHardwarePool() map[string]*api.Hardware {
	_ = "STUB: not implemented"
	return nil
}

func (e *ClusterE2ETest) RunClusterFlowWithGitOps(clusterOpts ...ClusterE2ETestOpt) {
	_ = "STUB: not implemented"
	return
}

func WithClusterFiller(f ...api.ClusterFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithClusterSingleNode helps to create an e2e test option for a single node cluster.
func WithClusterSingleNode(v v1alpha1.KubernetesVersion) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithClusterConfigLocationOverride(path string) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithEksaVersion(version *semver.Version) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithLatestMinorReleaseFromMain() ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

func WithEnvVar(key, val string) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

type Provider interface {
	Name() string
	// ClusterConfigUpdates allows a provider to modify the default cluster config
	// after this one is generated for the first time. This is not reapplied on every CLI operation.
	// Prefer to call UpdateClusterConfig directly from the tests to make it more explicit.
	ClusterConfigUpdates() []api.ClusterConfigFiller
	Setup()
	CleanupResources(clusterName string) error
	UpdateKubeConfig(content *[]byte, clusterName string) error
	ClusterStateValidations() []clusterf.StateValidation
	WithKubeVersionAndOS(kubeVersion v1alpha1.KubernetesVersion, os OS, release *releasev1.EksARelease, kernelVariant ...string) api.ClusterConfigFiller
	WithNewWorkerNodeGroup(name string, workerNodeGroup *WorkerNodeGroup) api.ClusterConfigFiller
}

// GenerateClusterConfig generates a cluster configuration.
func (e *ClusterE2ETest) GenerateClusterConfig(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// GenerateClusterConfigWithLicenseToken generates a cluster configuration while setting a specific license token.
func (e *ClusterE2ETest) GenerateClusterConfigWithLicenseToken(licenseToken string, opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func newBmclibClient(log logr.Logger, hostIP, username, password string) *bmclib.Client {
	_ = "STUB: not implemented"
	return nil
}

// ValidateHardwareDecommissioned checks that the all hardware was powered off during the cluster deletion.
// This function tests that the hardware was powered off during the cluster deletion.
func (e *ClusterE2ETest) ValidateHardwareDecommissioned() { _ = "STUB: not implemented"; return }

// add sleep retries to give the machine time to power off

func (e *ClusterE2ETest) GenerateHardwareConfig(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) generateHardwareConfig(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// create hardware CSV with no bmc username/password

// GenerateClusterConfigForVersion generates cluster configuration for the specified EKS-A version and license token.
func (e *ClusterE2ETest) GenerateClusterConfigForVersion(eksaVersion, licenseToken string, opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return

	// LicenseToken field was introduced in cluster spec only in release-22
	// attempting to populate the field for any prior versions would break the api.
	// We will need the conditional check as long as the latest minor is 'v0.21.*'.
}

func (e *ClusterE2ETest) generateClusterConfigObjects(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// Copy all objects that might be generated by the CLI.
// Don't replace the whole ClusterConfig since some ClusterE2ETestOpt might
// have already set some data in it.

// UpdateClusterConfig applies the cluster Config provided updates to e.ClusterConfig, marshalls its content
// to yaml and writes it to a file on disk configured by e.ClusterConfigLocation. Call this method when you want
// make changes to the eks-a cluster definition before running a CLI command or API operation.
func (e *ClusterE2ETest) UpdateClusterConfig(fillers ...api.ClusterConfigFiller) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) baseClusterConfigUpdates() []api.ClusterConfigFiller {
	_ = "STUB: not implemented"
	return nil
}

// This defaults all tests to a 1:1:1 configuration. Since all the fillers defined on each test are run
// after these defaults, if the tests is explicit about any of these, the defaults will be overwritten

// Disable packages by default for non-package e2e tests
// Package tests have "CuratedPackages" in their test name

// If we are persisting an existing cluster, set the control plane endpoint back to the original, since
// it is immutable

func (e *ClusterE2ETest) generateClusterConfigWithCLI(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) parseClusterConfigFromDisk(file string) { _ = "STUB: not implemented"; return }

// WithClusterConfig generates a base cluster config using the CLI `generate clusterconfig` command
// and updates them with the provided fillers. Helpful for defining the initial Cluster config
// before running a create operation.
func (e *ClusterE2ETest) WithClusterConfig(fillers ...api.ClusterConfigFiller) *ClusterE2ETest {
	_ = "STUB: not implemented"
	return nil
}

// DownloadArtifacts runs the EKS-A `download artifacts` command with appropriate args.
func (e *ClusterE2ETest) DownloadArtifacts(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

// ExtractDownloadedArtifacts extracts the downloaded artifacts.
func (e *ClusterE2ETest) ExtractDownloadedArtifacts(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// CleanupDownloadedArtifactsAndImages cleans up the downloaded artifacts and images.
func (e *ClusterE2ETest) CleanupDownloadedArtifactsAndImages(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func getBundleManifestLocation() string { _ = "STUB: not implemented"; return "" }

// DownloadImages runs the EKS-A `download images` command with appropriate args.
func (e *ClusterE2ETest) DownloadImages(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

// ImportImages runs the EKS-A `import images` command with appropriate args.
// When ociNamespaces are configured, it imports images into each namespace path
// so containerd can find them at the correct namespaced location.
func (e *ClusterE2ETest) ImportImages(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

// GetDiagnosticCollectorImage returns the diagnostic collector image from the bundle file,
// or falls back to a hardcoded URI if reading from disk fails.
func (e *ClusterE2ETest) GetDiagnosticCollectorImage() string { _ = "STUB: not implemented"; return "" }

// Try to read the bundle file

// Try to parse the bundle

// Check if version bundles exist

// GetPackageControllerRepo returns the bundle's package controller chart image.
func (e *ClusterE2ETest) GetPackageControlleChartRepo() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Read the bundle file

// Parse the bundle

// Get the package controller chart image from the first version bundle

// CopyPackages runs the EKS-A `copy packages` command to copy curated packages to the registry mirror.
func (e *ClusterE2ETest) CopyPackages(packageMirrorAlias string, packageChartRegistry string, packageRegistry string, opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// ChangeInstanceSecurityGroup modifies the security group of the instance to the provided value.
func (e *ClusterE2ETest) ChangeInstanceSecurityGroup(securityGroup string) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) CreateCluster(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) createCluster(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) ValidateCluster(kubeVersion v1alpha1.KubernetesVersion) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) WaitForMachineDeploymentReady(machineDeploymentName string) {
	_ = "STUB: not implemented"
	return
}

// GetEKSACluster retrieves the EKSA cluster from the runtime environment using kubectl.
func (e *ClusterE2ETest) GetEKSACluster() *v1alpha1.Cluster { _ = "STUB: not implemented"; return nil }

func (e *ClusterE2ETest) GetCapiMachinesForCluster(clusterName string) map[string]types.Machine {
	_ = "STUB: not implemented"
	return nil
}

// CapiMachinesForCluster reads all the CAPI Machines for a particular cluster and returns them
// index by their name.
func (e *ClusterE2ETest) CapiMachinesForCluster(clusterName string) (map[string]types.Machine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyClusterManifest uses client-side logic to create/update objects defined in a cluster yaml manifest.
func (e *ClusterE2ETest) ApplyClusterManifest() { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) applyClusterManifest(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// WithClusterUpgrade adds a cluster upgrade.
func WithClusterUpgrade(fillers ...api.ClusterFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// WithUpgradeClusterConfig adds a cluster upgrade.
// When we migrate usages of ClusterFiller to ClusterConfigFiller we can rename this to WithClusterUpgrade.
func WithUpgradeClusterConfig(fillers ...api.ClusterConfigFiller) ClusterE2ETestOpt {
	_ = "STUB: not implemented"
	return *new(ClusterE2ETestOpt)
}

// LoadClusterConfigGeneratedByCLI loads the full cluster config from the file generated when a cluster is created using the CLI.
func (e *ClusterE2ETest) LoadClusterConfigGeneratedByCLI(fillers ...api.ClusterConfigFiller) {
	_ = "STUB: not implemented"
	return
}

// UpgradeClusterWithNewConfig applies the test options, re-generates the cluster config file and runs the CLI upgrade command.
func (e *ClusterE2ETest) UpgradeClusterWithNewConfig(clusterOpts []ClusterE2ETestOpt, commandOpts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) upgradeCluster(clusterOpts []ClusterE2ETestOpt, commandOpts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// UpgradeCluster runs the CLI upgrade command.
func (e *ClusterE2ETest) UpgradeCluster(commandOpts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) generateClusterConfigYaml() []byte { _ = "STUB: not implemented"; return nil }

// This is required because Flux requires a namespace be specified for objects
// to be able to reconcile right.

// This is required because Flux requires a namespace be specified for objects
// to be able to reconcile right.

func (e *ClusterE2ETest) buildClusterConfigFile() { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) DeleteCluster(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

// cleanupResources is a helper to clean up test resources. It is a noop if the T_CLEANUP_RESOURCES environment variable
// is false or unset.
func (e *ClusterE2ETest) cleanupResources() { _ = "STUB: not implemented"; return }

func (e *ClusterE2ETest) CleanupDockerEnvironment() { _ = "STUB: not implemented"; return }

func shouldCleanUpResources() bool { _ = "STUB: not implemented"; return false }

func (e *ClusterE2ETest) deleteCluster(opts ...CommandOpt) { _ = "STUB: not implemented"; return }

// GenerateSupportBundleOnCleanupIfTestFailed does what it says on the tin.
//
// It uses testing.T.Cleanup to register a handler that checks if the test
// failed, and generates a support bundle only in the event of a failure.
func (e *ClusterE2ETest) GenerateSupportBundleOnCleanupIfTestFailed(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

// GenerateSupportBundleIfTestFailed runs generates a support bundle if the test failed.
func (e *ClusterE2ETest) GenerateSupportBundleIfTestFailed(opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) Run(name string, args ...string) { _ = "STUB: not implemented"; return }

// Look for the last line of the out put that starts with 'Error:'

func (e *ClusterE2ETest) RunEKSA(args []string, opts ...CommandOpt) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) StopIfFailed() { _ = "STUB: not implemented"; return }

// Cluster builds a cluster obj using the ClusterE2ETest name and kubeconfig.
func (e *ClusterE2ETest) Cluster() *types.Cluster { _ = "STUB: not implemented"; return nil }

func (e *ClusterE2ETest) managementCluster() *types.Cluster { _ = "STUB: not implemented"; return nil }

// KubeconfigFilePath retrieves the Kubeconfig path used for the workload cluster.
func (e *ClusterE2ETest) KubeconfigFilePath() string { _ = "STUB: not implemented"; return "" }

// BuildWorkloadClusterClient creates a client for the workload cluster created by e.
func (e *ClusterE2ETest) BuildWorkloadClusterClient() (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

// Adding the retry logic here because the connection to the client does not always
// succedd on the first try due to connection failure after the kubeconfig becomes
// available in the cluster.

func (e *ClusterE2ETest) managementKubeconfigFilePath() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *ClusterE2ETest) GetEksaVSphereMachineConfigs() []v1alpha1.VSphereMachineConfig {
	_ = "STUB: not implemented"
	return nil
}

func GetTestNameHash(name string) string { _ = "STUB: not implemented"; return "" }

func getClusterName(t T) string { _ = "STUB: not implemented"; return "" }

// Append hash to make each cluster name unique per test. Using the testname will be too long
// and would fail validations

func getBundlesOverride() string { _ = "STUB: not implemented"; return "" }

// GetLicenseToken retrieves the license token from the environment variables.
func GetLicenseToken() string { _ = "STUB: not implemented"; return "" }

// GetLicenseToken2 retrieves the license token 2 from the environment variables.
func GetLicenseToken2() string { _ = "STUB: not implemented"; return "" }

// GetStagingLicenseToken retrieves the staging license token from the environment variables.
func GetStagingLicenseToken() string { _ = "STUB: not implemented"; return "" }

// GetStagingLicenseToken2 retrieves the staging license token 2 from the environment variables.
func GetStagingLicenseToken2() string { _ = "STUB: not implemented"; return "" }

func getCleanupResourcesVar() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func setEksctlVersionEnvVar() error { _ = "STUB: not implemented"; return nil }

// CreateNamespace creates a namespace.
func (e *ClusterE2ETest) CreateNamespace(namespace string) { _ = "STUB: not implemented"; return }

// DeleteNamespace deletes a namespace.
func (e *ClusterE2ETest) DeleteNamespace(namespace string) { _ = "STUB: not implemented"; return }

// SetPackageBundleActive will set the current packagebundle to the active state.
func (e *ClusterE2ETest) SetPackageBundleActive() { _ = "STUB: not implemented"; return }

// ValidatingNoPackageController make sure there is no package controller.
func (e *ClusterE2ETest) ValidatingNoPackageController() { _ = "STUB: not implemented"; return }

// InstallCuratedPackage will install a curated package.
func (e *ClusterE2ETest) InstallCuratedPackage(packageName, packagePrefix, kubeconfig string, opts ...string) {
	_ = "STUB: not implemented"
	return
}

// The package install command doesn't (yet?) have a --kubeconfig flag.

// InstallCuratedPackageFile will install a curated package from a yaml file, this is useful since target namespace isn't supported on the CLI.
func (e *ClusterE2ETest) InstallCuratedPackageFile(packageFile, kubeconfig string, opts ...string) {
	_ = "STUB: not implemented"
	return
}

// ValidatePackageBundleControllerRegistry checks if the registries for helm charts and images match for curated packages tests.
func (e *ClusterE2ETest) ValidatePackageBundleControllerRegistry() {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) generatePackageConfig(ns, targetns, prefix, packageName string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// BuildPackageConfigFile will create the file in the test directory for the curated package.
func (e *ClusterE2ETest) BuildPackageConfigFile(packageName, prefix, ns string) string {
	_ = "STUB: not implemented"
	return ""
}

func (e *ClusterE2ETest) CreateResource(ctx context.Context, resource string) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) UninstallCuratedPackage(packagePrefix string, opts ...string) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) InstallLocalStorageProvisioner() { _ = "STUB: not implemented"; return }

// WithCluster helps with bringing up and tearing down E2E test clusters.
func (e *ClusterE2ETest) WithCluster(f func(e *ClusterE2ETest)) { _ = "STUB: not implemented"; return }

// Like WithCluster but does not delete the cluster. Useful for debugging.
func (e *ClusterE2ETest) WithPersistentCluster(f func(e *ClusterE2ETest)) {
	_ = "STUB: not implemented"
	return
}

// VerifyHarborPackageInstalled is checking if the harbor package gets installed correctly.
func (e *ClusterE2ETest) VerifyHarborPackageInstalled(prefix, namespace string) {
	_ = "STUB: not implemented"
	return
}

// Log Package/Deployment outputs

func (e *ClusterE2ETest) printPackageSpec(ctx context.Context, params []string) {
	_ = "STUB: not implemented"
	return
}

func (e *ClusterE2ETest) printDeploymentSpec(ctx context.Context, ns string) {
	_ = "STUB: not implemented"
	return
}

// VerifyHelloPackageInstalled is checking if the hello eks anywhere package gets installed correctly.
func (e *ClusterE2ETest) VerifyHelloPackageInstalled(packageName string, mgmtCluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// Log Package/Deployment outputs

// VerifyAdotPackageInstalled is checking if the ADOT package gets installed correctly.
func (e *ClusterE2ETest) VerifyAdotPackageInstalled(packageName, targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

// Log Package/Deployment outputs

//go:embed testdata/adot_package_deployment.yaml
var adotPackageDeployment []byte

//go:embed testdata/adot_package_daemonset.yaml
var adotPackageDaemonset []byte

// VerifyAdotPackageDeploymentUpdated is checking if deployment config changes trigger resource reloads correctly.
func (e *ClusterE2ETest) VerifyAdotPackageDeploymentUpdated(packageName, targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

// Deploy ADOT as a deployment and scrape the apiservers

// Add sleep to allow package to change state

// Log Package/Deployment outputs

// VerifyAdotPackageDaemonSetUpdated is checking if daemonset config changes trigger resource reloads correctly.
func (e *ClusterE2ETest) VerifyAdotPackageDaemonSetUpdated(packageName, targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

// Deploy ADOT as a daemonset and scrape the node

// Add sleep to allow package to change state

// Log Package/Deployment outputs

//go:embed testdata/emissary_listener.yaml
var emisarryListener []byte

//go:embed testdata/emissary_package.yaml
var emisarryPackage []byte

// VerifyEmissaryPackageInstalled is checking if emissary package gets installed correctly.
func (e *ClusterE2ETest) VerifyEmissaryPackageInstalled(packageName string, mgmtCluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// Log Package/Deployment outputs

// TestEmissaryPackageRouting is checking if emissary is able to create Ingress, host, and mapping that function correctly.
func (e *ClusterE2ETest) TestEmissaryPackageRouting(packageName, checkName string, mgmtCluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// Log Package/Deployment outputs

// Functional testing of Emissary Ingress

// VerifyPrometheusPackageInstalled is checking if the Prometheus package gets installed correctly.
func (e *ClusterE2ETest) VerifyPrometheusPackageInstalled(packageName, targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

// VerifyCertManagerPackageInstalled is checking if the cert manager package gets installed correctly.
func (e *ClusterE2ETest) VerifyCertManagerPackageInstalled(prefix, namespace, packageName string, mgmtCluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// Log Package/Deployment outputs

//go:embed testdata/certmanager/certmanager_selfsignedissuer.yaml
var certManagerSelfSignedIssuer []byte

//go:embed testdata/certmanager/certmanager_selfsignedcert.yaml
var certManagerSelfSignedCert []byte

func (e *ClusterE2ETest) verifySelfSignedCertificate() error { _ = "STUB: not implemented"; return nil }

//go:embed testdata/certmanager/certmanager_letsencrypt_issuer.yaml
var certManagerLetsEncryptIssuer string

//go:embed testdata/certmanager/certmanager_letsencrypt_cert.yaml
var certManagerLetsEncryptCert []byte

//go:embed testdata/certmanager/certmanager_secret.yaml
var certManagerSecret string

func (e *ClusterE2ETest) verifyLetsEncryptCert() error { _ = "STUB: not implemented"; return nil }

// CleanupCerts cleans up letsencrypt certificates.
func (e *ClusterE2ETest) CleanupCerts(mgmtCluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPrometheusPrometheusServerStates is checking if the Prometheus package prometheus-server component is functioning properly.
func (e *ClusterE2ETest) VerifyPrometheusPrometheusServerStates(packageName, targetNamespace, mode string) {
	_ = "STUB: not implemented"
	return
}

// VerifyPrometheusNodeExporterStates is checking if the Prometheus package node-exporter component is functioning properly.
func (e *ClusterE2ETest) VerifyPrometheusNodeExporterStates(packageName, targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

//go:embed testdata/prometheus_package_deployment.yaml
var prometheusPackageDeployment []byte

//go:embed testdata/prometheus_package_statefulset.yaml
var prometheusPackageStatefulSet []byte

// ApplyPrometheusPackageServerDeploymentFile is checking if deployment config changes trigger resource reloads correctly.
func (e *ClusterE2ETest) ApplyPrometheusPackageServerDeploymentFile(packageName, targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

// ApplyPrometheusPackageServerStatefulSetFile is checking if statefulset config changes trigger resource reloads correctly.
func (e *ClusterE2ETest) ApplyPrometheusPackageServerStatefulSetFile(packageName, targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

// VerifyPackageControllerNotInstalled is verifying that package controller is not installed.
func (e *ClusterE2ETest) VerifyPackageControllerNotInstalled() { _ = "STUB: not implemented"; return }

// VerifyAutoScalerPackageInstalled is verifying that the autoscaler package is installed and deployed.
func (e *ClusterE2ETest) VerifyAutoScalerPackageInstalled(packageName, targetNamespace string, mgmtCluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// Log Package/Deployment outputs

// VerifyMetricServerPackageInstalled is verifying that metrics-server is installed and deployed.
func (e *ClusterE2ETest) VerifyMetricServerPackageInstalled(packageName, targetNamespace string, mgmtCluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// Log Package/Deployment outputs

//go:embed testdata/autoscaler_package.yaml
var autoscalerPackageDeploymentTemplate string

//go:embed testdata/metrics_server_package.yaml
var metricsServerPackageDeploymentTemplate string

// InstallAutoScalerWithMetricServer installs autoscaler and metrics-server with a given target namespace.
func (e *ClusterE2ETest) InstallAutoScalerWithMetricServer(targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

//go:embed testdata/autoscaler_package_workload_cluster.yaml
var autoscalerPackageWorkloadClusterDeploymentTemplate string

// InstallAutoScaler installs autoscaler with a given target namespace.
func (e *ClusterE2ETest) InstallAutoScaler(workloadClusterName, targetNamespace string) {
	_ = "STUB: not implemented"
	return
}

//go:embed testdata/certmanager/certmanager_package.yaml
var certManagerPackageTemplate string

// InstallCertManagerPackageWithAwsCredentials installs cert-manager package by setting aws credentials in the pod.
func (e *ClusterE2ETest) InstallCertManagerPackageWithAwsCredentials(prefix, packageName, namespace, clusterName string) {
	_ = "STUB: not implemented"
	return
}

// CombinedAutoScalerMetricServerTest verifies that new nodes are spun up after using a HPA to scale a deployment.
func (e *ClusterE2ETest) CombinedAutoScalerMetricServerTest(autoscalerName, metricServerName, targetNamespace string, mgmtCluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// DeployTestWorkload deploys the test workload on the cluster.
func (e *ClusterE2ETest) DeployTestWorkload(cluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// VerifyWorkerNodesScaleUp verifies that the worker nodes are scaled up after a test workload is deployed on a cluster with Autoscaler installed.
func (e *ClusterE2ETest) VerifyWorkerNodesScaleUp(mgmtCluster *types.Cluster) {
	_ = "STUB: not implemented"
	return
}

// RestartClusterAutoscaler restarts the cluster autoscaler deployment in the target namespace.
func (e *ClusterE2ETest) RestartClusterAutoscaler(targetNamespace string) {
	_ = "STUB: not implemented"
	// There is a bug in cluster autoscaler currently where it's not able to autoscale the cluster
	// because of missing permissions on infrastructure machine template.
	// Cluster Autoscaler does restart after ~10 min after which it starts functioning normally.
	// We are force triggering a restart so the e2e doesn't have to wait 10 min for the restart.
	// This can be removed once the following issue is resolve upstream.
	// https://github.com/kubernetes/autoscaler/issues/6490
	return
}

// ValidateClusterState runs a set of validations against the cluster to identify an invalid cluster state.
func (e *ClusterE2ETest) ValidateClusterState() { _ = "STUB: not implemented"; return }

// ValidateClusterStateWithT runs a set of validations against the cluster to identify an invalid cluster state and accepts *testing.T as a parameter.
func (e *ClusterE2ETest) ValidateClusterStateWithT(t *testing.T) { _ = "STUB: not implemented"; return }

func validateClusterState(t *testing.T, e *ClusterE2ETest) { _ = "STUB: not implemented"; return }

// ApplyPackageFile is applying a package file in the cluster.
func (e *ClusterE2ETest) ApplyPackageFile(packageName, targetNamespace string, PackageFile []byte) {
	_ = "STUB: not implemented"
	return
}

// Add sleep to allow package to change state

// CurlEndpoint creates a pod with command to curl the target endpoint,
// and returns the created pod name.
func (e *ClusterE2ETest) CurlEndpoint(endpoint, namespace string, extraCurlArgs ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// MatchLogs matches the log from a container to the expected content. Given it
// takes time for logs to be populated, a retrier with configurable timeout duration
// is added.
func (e *ClusterE2ETest) MatchLogs(targetNamespace, targetPodName string,
	targetContainerName, expectedLogs string, timeout time.Duration,
) {
	_ = "STUB: not implemented"
	return
}

// ValidateEndpointContent validates the contents at the target endpoint.
func (e *ClusterE2ETest) ValidateEndpointContent(endpoint, namespace, expectedContent string, extraCurlArgs ...string) {
	_ = "STUB: not implemented"
	return
}

// AirgapDockerContainers airgap docker containers. Outside network should not be reached during airgapped deployment.
func (e *ClusterE2ETest) AirgapDockerContainers(localCIDRs string) {
	_ = "STUB: not implemented"
	return
}

// CreateAirgappedUser create airgapped user and setup the iptables rule. Notice that OUTPUT chain is flushed each time.
func (e *ClusterE2ETest) CreateAirgappedUser(localCIDR string) { _ = "STUB: not implemented"; return }

// Allow the airgap user to access logs folder
// Allow the airgap user to access working dir

// AssertAirgappedNetwork make sure that the admin machine is indeed airgapped.
func (e *ClusterE2ETest) AssertAirgappedNetwork() { _ = "STUB: not implemented"; return }

func dumpFile(description, path string, t T) { _ = "STUB: not implemented"; return }

// CreateCloudStackCredentialsSecretFromEnvVar parses the cloudstack credentials from an environment variable,
// builds a new secret object from the credentials in the provided profile and creates it in the cluster.
func (e *ClusterE2ETest) CreateCloudStackCredentialsSecretFromEnvVar(name, profileName string) {
	_ = "STUB: not implemented"
	return
}

// Create a new secret with the credentials from the profile, but with a new name.

func (e *ClusterE2ETest) addClusterConfigFillers(fillers ...api.ClusterConfigFiller) {
	_ = "STUB: not implemented"
	return
}
