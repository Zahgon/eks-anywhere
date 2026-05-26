package executables

import (
	"bytes"
	"context"

	"github.com/aws/eks-anywhere/pkg/helm"
)

const (
	helmPath               = "helm"
	insecureSkipVerifyFlag = "--insecure-skip-tls-verify"
)

type Helm struct {
	executable Executable
	helmConfig *helm.Config // Embed HelmOptions in Helm struct
	env        map[string]string
}

// NewHelm returns a new Helm executable client.
func NewHelm(executable Executable, opts ...helm.Opt) *Helm { _ = "STUB: not implemented"; return nil }

// mergeMaps joins the default and the provided maps together, then return the
// new map.
func mergeMaps(defaultEnv, newEnv map[string]string) { _ = "STUB: not implemented"; return }

func (h *Helm) Template(ctx context.Context, ociURI, version, namespace string, values interface{}, kubeVersion string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Helm) PullChart(ctx context.Context, ociURI, version string) error {
	_ = "STUB: not implemented"
	return nil
}

// ShowValues get the values of a chart.
func (h *Helm) ShowValues(ctx context.Context, ociURI, version string) (bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return *new(bytes.Buffer), nil
}

func (h *Helm) PushChart(ctx context.Context, chart, registry string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Helm) RegistryLogin(ctx context.Context, registry, username, password string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Helm) SaveChart(ctx context.Context, ociURI, version, folder string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Helm) InstallChartFromName(ctx context.Context, ociURI, kubeConfig, name, version string) error {
	_ = "STUB: not implemented"
	// Using upgrade --install will install the chart if it doesn't exist, but
	// upgrades it otherwise, making this more idempotent than install, which
	// would error out if the chart is already installed, and has no similar
	// "--upgrade" flag.
	return nil
}

// InstallChart installs a helm chart to the target cluster.
//
// If kubeconfigFilePath is the empty string, it won't be passed at all.
func (h *Helm) InstallChart(ctx context.Context, chart, ociURI, version, kubeconfigFilePath, namespace, valueFilePath string, skipCRDs bool, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete removes an installation.
func (h *Helm) Delete(ctx context.Context, kubeconfigFilePath, installName, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// ListCharts lists helm charts filtered on the given regex filter.
// If namespace is provided, it lists charts in that namespace; otherwise uses helm's default behavior.
func (h *Helm) ListCharts(ctx context.Context, kubeconfigFilePath, filter, namespace string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Helm) addInsecureFlagIfProvided(params []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (h *Helm) url(originalURL string) string { _ = "STUB: not implemented"; return "" }

func GetHelmValueArgs(values []string) []string { _ = "STUB: not implemented"; return nil }

// UpgradeInstallChartWithValuesFile runs a helm upgrade --install with the provided values file and waits for the
// chart deployment to be ready. This will upgrade the chart if it exists and install if it does not.
func (h *Helm) UpgradeInstallChartWithValuesFile(ctx context.Context, chart, ociURI, version, kubeconfigFilePath, namespace, valuesFilePath string, opts ...helm.Opt) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: we should not update the receiver here, so this needs to change.
// This is not thread safe.
// https://github.com/aws/eks-anywhere/issues/7176

// Uninstall runs a helm uninstall for the given chart in given namespace.
func (h *Helm) Uninstall(ctx context.Context, chart, kubeconfigFilePath, namespace string, opts ...helm.Opt) error {
	_ = "STUB: not implemented"
	return nil
}
