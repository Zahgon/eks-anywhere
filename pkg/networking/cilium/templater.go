package cilium

import (
	"context"
	_ "embed"
	"time"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/helm"
	"github.com/aws/eks-anywhere/pkg/retrier"
	"github.com/aws/eks-anywhere/pkg/semver"
)

//go:embed network_policy.yaml
var networkPolicyAllowAll string

const (
	maxRetries           = 10
	defaultBackOffPeriod = 5 * time.Second
	namespace            = constants.KubeSystemNamespace
)

// HelmClientFactory provides a helm client for a cluster.
type HelmClientFactory interface {
	Get(ctx context.Context, clus *anywherev1.Cluster) (helm.Client, error)
}

type Templater struct {
	helmFactory HelmClientFactory
}

// NewTemplater returns a new Templater.
func NewTemplater(helmFactory HelmClientFactory) *Templater { _ = "STUB: not implemented"; return nil }

func (t *Templater) GenerateUpgradePreflightManifest(ctx context.Context, spec *cluster.Spec) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ManifestOpt allows to modify options for a cilium manifest.
type ManifestOpt func(*ManifestConfig)

type ManifestConfig struct {
	values      values
	retrier     *retrier.Retrier
	kubeVersion string
	namespaces  []string
}

// WithKubeVersion allows to generate the Cilium manifest for a different kubernetes version
// than the one specified in the cluster spec. Useful for upgrades scenarios where Cilium is upgraded before
// the kubernetes components.
func WithKubeVersion(kubeVersion string) ManifestOpt {
	_ = "STUB: not implemented"
	return *new(ManifestOpt)
}

// WithRetrier introduced for optimizing unit tests.
func WithRetrier(retrier *retrier.Retrier) ManifestOpt {
	_ = "STUB: not implemented"
	return *new(ManifestOpt)
}

// WithUpgradeFromVersion allows to specify the compatibility Cilium version to use in the manifest.
// This is necessary for Cilium upgrades.
func WithUpgradeFromVersion(version semver.Version) ManifestOpt {
	_ = "STUB: not implemented"
	return *new(ManifestOpt)
}

// WithPolicyAllowedNamespaces allows to specify which namespaces traffic should be allowed when using
// and "Always" policy enforcement mode.
func WithPolicyAllowedNamespaces(namespaces []string) ManifestOpt {
	_ = "STUB: not implemented"
	return *new(ManifestOpt)
}

func (t *Templater) GenerateManifest(ctx context.Context, spec *cluster.Spec, opts ...ManifestOpt) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if policy enforcement mode is "always" to append network policy

// Check helmValues first (takes precedence)

// Fall back to deprecated field if helmValues not provided or empty

func (t *Templater) GenerateNetworkPolicyManifest(spec *cluster.Spec, namespaces []string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type values map[string]interface{}

func (c values) set(value interface{}, path ...string) { _ = "STUB: not implemented"; return }

// convertToValues recursively converts map[string]interface{} to values type.
func convertToValues(input map[string]interface{}) values {
	_ = "STUB: not implemented"
	return *new(values)
}

// convertSliceToValues converts slices that may contain maps to use values type.
func convertSliceToValues(input []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func templateValues(spec *cluster.Spec, versionsBundle *cluster.VersionsBundle) values {
	_ = "STUB: not implemented"
	// If HelmValues are configured, use them instead of templated values
	return *new(values)
}

// Unmarshal JSON to map

// The chart expects an "incomplete" repository
// and will add the necessary suffix ("-generic" in our case)

func getChartURIAndVersion(versionsBundle *cluster.VersionsBundle) (uri, version string) {
	_ = "STUB: not implemented"
	return "", ""
}

func getKubeVersion(versionsBundle *cluster.VersionsBundle) (*semver.Version, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getKubeVersionString(spec *cluster.Spec, versionsBundle *cluster.VersionsBundle) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
