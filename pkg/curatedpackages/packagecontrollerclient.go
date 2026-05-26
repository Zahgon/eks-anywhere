package curatedpackages

import (
	"context"
	_ "embed"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/registrymirror"
	releasev1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

//go:embed config/secrets.yaml
var secretsValueYaml string

const (
	eksaDefaultRegion             = "us-west-2"
	valueFileName                 = "values.yaml"
	defaultRegistryMirrorUsername = "username"
	defaultRegistryMirrorPassword = "password"
)

type PackageControllerClientOpt func(client *PackageControllerClient)

type PackageControllerClient struct {
	kubeConfig string
	chart      *releasev1.Image
	// chartManager installs and deletes helm charts.
	chartManager          ChartManager
	clusterName           string
	clusterSpec           *anywherev1.ClusterSpec
	managementClusterName string
	kubectl               KubectlRunner
	eksaAccessKeyID       string
	eksaSecretAccessKey   string
	eksaSessionToken      string
	eksaRegion            string
	eksaAwsConfig         string
	httpProxy             string
	httpsProxy            string
	noProxy               []string
	registryMirror        *registrymirror.RegistryMirror
	// activeBundleTimeout is the timeout to activate a bundle on installation.
	activeBundleTimeout time.Duration
	valuesFileWriter    filewriter.FileWriter
	// skipWaitForPackageBundle indicates whether the installer should wait
	// until a package bundle is activated.
	//
	// Skipping the wait is desirable for full cluster lifecycle use cases,
	// where resource creation and error reporting are asynchronous in nature.
	skipWaitForPackageBundle bool
	// tracker creates k8s clients for workload clusters managed via full
	// cluster lifecycle API.
	clientBuilder ClientBuilder

	// mu provides some thread-safety.
	mu sync.Mutex

	// registryAccessTester test if the aws credential has access to registry
	registryAccessTester RegistryAccessTester
}

// ClientBuilder returns a k8s client for the specified cluster.
type ClientBuilder interface {
	GetClient(context.Context, types.NamespacedName) (client.Client, error)
}

type ChartInstaller interface {
	InstallChart(ctx context.Context, chart, ociURI, version, kubeconfigFilePath, namespace, valueFilePath string, skipCRDs bool, values []string) error
}

// ChartUninstaller handles deleting chart installations.
type ChartUninstaller interface {
	Delete(ctx context.Context, kubeconfigFilePath, installName, namespace string) error
}

// ChartManager installs and uninstalls helm charts.
type ChartManager interface {
	ChartInstaller
	ChartUninstaller
	RegistryLogin(ctx context.Context, registry, username, password string) error
}

// NewPackageControllerClientFullLifecycle creates a PackageControllerClient
// for the Full Cluster Lifecycle controller.
//
// It differs because the CLI use case has far more information available at
// instantiation, while the FCL use case has less information at
// instantiation, and the rest when cluster creation is triggered.
func NewPackageControllerClientFullLifecycle(logger logr.Logger, chartManager ChartManager, kubectl KubectlRunner, clientBuilder ClientBuilder) *PackageControllerClient {
	_ = "STUB: not implemented"
	return nil
}

// EnableFullLifecycle wraps Enable to handle run-time arguments.
//
// This method fills in the gaps between the original CLI use case, where all
// information is known at PackageControllerClient initialization, and the
// Full Cluster Lifecycle use case, where there's limited information at
// initialization. Basically any parameter here isn't known at instantiation
// of the PackageControllerClient during full cluster lifecycle usage, hence
// why this method exists.
func (pc *PackageControllerClient) EnableFullLifecycle(ctx context.Context, log logr.Logger, clusterName, kubeConfig string, chart *releasev1.Image, registryMirror *registrymirror.RegistryMirror, options ...PackageControllerClientOpt) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// This anonymous function ensures that the pc.mu is unlocked before
// Enable is called, preventing deadlocks in the event that Enable tries
// to acquire pc.mu.

// NewPackageControllerClient instantiates a new instance of PackageControllerClient.
func NewPackageControllerClient(chartManager ChartManager, kubectl KubectlRunner, clusterName, kubeConfig string, chart *releasev1.Image, registryMirror *registrymirror.RegistryMirror, options ...PackageControllerClientOpt) *PackageControllerClient {
	_ = "STUB: not implemented"
	return nil
}

// Enable curated packages in a cluster
//
// In case the cluster is management cluster, it performs the following actions:
//   - Installation of Package Controller through helm chart installation
//   - Creation of secret credentials
//   - Creation of a single run of a cron job refresher
//   - Activation of a curated packages bundle
//
// In case the cluster is a workload cluster, it performs the following actions:
//   - Creation of package bundle controller (PBC) custom resource in management cluster
func (pc *PackageControllerClient) Enable(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Provide proxy details for curated packages helm chart when proxy details provided

// Helm requires commas to be escaped: https://github.com/rancher/rancher/issues/16195

// GetCuratedPackagesRegistries gets value for configurable registries from PBC.
func (pc *PackageControllerClient) GetCuratedPackagesRegistries(ctx context.Context) (sourceRegistry, defaultRegistry, defaultImageRegistry string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// If registry mirror is configured, admin machine may be airgapped.
// We will use registry mirror configuration as source of truth to decide package registries

// registry name is added as part of sourceRegistry field in package controller helm chart
// https://github.com/aws/eks-anywhere-packages/blob/main/charts/eks-anywhere-packages/values.yaml#L15-L18

// use regional registry when the above credential is good

// CreateHelmOverrideValuesYaml creates a temp file to override certain values in package controller helm install.
func (pc *PackageControllerClient) CreateHelmOverrideValuesYaml() (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (pc *PackageControllerClient) generateHelmOverrideValues() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// packageBundleControllerResource is the name of the package bundle controller
// resource in the API.
const packageBundleControllerResource string = "packageBundleController"

// waitForActiveBundle polls the package bundle controller for its active bundle.
//
// It returns nil on success. Success is defined as receiving a valid package
// bundle controller from the API with a non-empty active bundle.
//
// If no timeout is specified, a default of 3 minutes is used.
func (pc *PackageControllerClient) waitForActiveBundle(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if eks-anywhere-packages Package is installed

// TODO read a polling interval value from the context, falling
// back to this as a default.

// IsInstalled checks if a package controller custom resource exists.
func (pc *PackageControllerClient) IsInstalled(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func formatYamlLine(space, key, value string) string { _ = "STUB: not implemented"; return "" }

func formatImageResource(resource *anywherev1.ImageResource, name string) (result string) {
	_ = "STUB: not implemented"
	return ""
}

func formatCronJob(cronJob *anywherev1.PackageControllerCronJob) (result string) {
	_ = "STUB: not implemented"
	return ""
}

func formatResources(resources *anywherev1.PackageControllerResources) (result string) {
	_ = "STUB: not implemented"
	return ""
}

// GetPackageControllerConfiguration returns the default kubernetes version for a Cluster.
func (pc *PackageControllerClient) GetPackageControllerConfiguration() (result string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Reconcile installs resources when a full cluster lifecycle cluster is created.
func (pc *PackageControllerClient) Reconcile(ctx context.Context, logger logr.Logger, client client.Client, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// No Kubeconfig is passed. This is intentional. The helm executable will
// get that configuration from its environment.

// getBundleFromCluster based on the cluster's k8s version.
func (pc *PackageControllerClient) getBundleFromCluster(ctx context.Context, client client.Client, clusterObj *anywherev1.Cluster) (*releasev1.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KubeDeleter abstracts client.Client so mocks can be substituted in tests.
type KubeDeleter interface {
	Delete(context.Context, client.Object, ...client.DeleteOption) error
}

// ReconcileDelete removes resources after a full cluster lifecycle cluster is
// deleted.
func (pc *PackageControllerClient) ReconcileDelete(ctx context.Context, logger logr.Logger, client KubeDeleter, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

func WithEksaAccessKeyId(eksaAccessKeyId string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithActiveBundleTimeout(timeout time.Duration) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithEksaSecretAccessKey(eksaSecretAccessKey string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

// WithEksaSessionToken set the eksaSessionToken field.
func WithEksaSessionToken(eksaSessionToken string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithEksaRegion(eksaRegion string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

// WithEksaAwsConfig set the eksaAwsConfig field.
func WithEksaAwsConfig(eksaAwsConfig string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithHTTPProxy(httpProxy string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithHTTPSProxy(httpsProxy string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithNoProxy(noProxy []string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

func WithManagementClusterName(managementClusterName string) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

// WithValuesFileWriter sets up a writer to generate temporary values.yaml to
// override some values in package controller helm chart.
func WithValuesFileWriter(writer filewriter.FileWriter) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

// WithClusterSpec sets the cluster spec.
func WithClusterSpec(clusterSpec *cluster.Spec) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

// WithRegistryAccessTester sets the registryTester.
func WithRegistryAccessTester(registryTester RegistryAccessTester) func(client *PackageControllerClient) {
	_ = "STUB: not implemented"
	return nil
}

// WithSkipWait sets skipWaitForPackageBundle.
func WithSkipWait() func(client *PackageControllerClient) { _ = "STUB: not implemented"; return nil }
