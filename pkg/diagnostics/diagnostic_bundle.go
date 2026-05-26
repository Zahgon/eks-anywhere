package diagnostics

import (
	"context"
	_ "embed"
	"time"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/retrier"
)

//go:embed config/diagnostic-collector-rbac.yaml
var diagnosticCollectorRbac []byte

const (
	troubleshootApiVersion      = "troubleshoot.sh/v1beta2"
	generatedBundleNameFormat   = "%s-%s-bundle.yaml"
	generatedAnalysisNameFormat = "%s-%s-analysis.yaml"
	maxRetries                  = 5
	backOffPeriod               = 5 * time.Second
	defaultClusterName          = "eksa-cluster"
)

type EksaDiagnosticBundle struct {
	bundle           *supportBundle
	hostBundle       *supportBundle
	bundlePath       string
	client           BundleClient
	collectorFactory CollectorFactory
	clusterSpec      *cluster.Spec
	analyzerFactory  AnalyzerFactory
	kubeconfig       string
	kubectl          *executables.Kubectl
	retrier          *retrier.Retrier
	writer           filewriter.FileWriter
	analysis         []*executables.SupportBundleAnalysis
}

func newDiagnosticBundleManagementCluster(af AnalyzerFactory, cf CollectorFactory, spec *cluster.Spec, client BundleClient,
	kubectl *executables.Kubectl, kubeconfig string, writer filewriter.FileWriter,
) (*EksaDiagnosticBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDiagnosticBundleFromSpec(af AnalyzerFactory, cf CollectorFactory, spec *cluster.Spec, provider providers.Provider,
	client BundleClient, kubectl *executables.Kubectl, kubeconfig string, writer filewriter.FileWriter, auditLogs bool,
) (*EksaDiagnosticBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDiagnosticBundleDefault(af AnalyzerFactory, cf CollectorFactory) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func newDiagnosticBundleCustom(af AnalyzerFactory, cf CollectorFactory, client BundleClient, kubectl *executables.Kubectl, bundlePath string, kubeconfig string, writer filewriter.FileWriter) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) CollectAndAnalyze(ctx context.Context, sinceTimeValue *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) PrintBundleConfig() error { _ = "STUB: not implemented"; return nil }

func (e *EksaDiagnosticBundle) WriteBundleConfig() error { _ = "STUB: not implemented"; return nil }

func (e *EksaDiagnosticBundle) PrintAnalysis() error { _ = "STUB: not implemented"; return nil }

func (e *EksaDiagnosticBundle) WriteAnalysisToFile() (path string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// WithHostCollectors configures host bundle with collectors that run on host machines.
func (e *EksaDiagnosticBundle) WithHostCollectors(config v1alpha1.Ref) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

// WithAuditLogs configures bundle to collect audit logs from control plane nodes.
func (e *EksaDiagnosticBundle) WithAuditLogs() *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

// WithDefaultHostCollectors collects the default collectors that run on the host machine.
func (e *EksaDiagnosticBundle) WithDefaultHostCollectors(config v1alpha1.Ref) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithDefaultCollectors() *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithDefaultAnalyzers() *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithManagementCluster(isSelfManaged bool) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

// WithFileCollectors appends collectors that collect static data from the specified paths to the bundle.
func (e *EksaDiagnosticBundle) WithFileCollectors(paths []string) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithPackagesCollectors() *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithDatacenterConfig(config v1alpha1.Ref, spec *cluster.Spec) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithOidcConfig(config *v1alpha1.OIDCConfig) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithExternalEtcd(config *v1alpha1.ExternalEtcdConfiguration) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithGitOpsConfig(config *v1alpha1.GitOpsConfig) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithMachineConfigs(configs []providers.MachineConfig) *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

func (e *EksaDiagnosticBundle) WithLogTextAnalyzers() *EksaDiagnosticBundle {
	_ = "STUB: not implemented"
	return nil
}

// hasHostCollectors checks if the host bundle has any collectors added.
func (e *EksaDiagnosticBundle) hasHostCollectors() bool { _ = "STUB: not implemented"; return false }

// combineWithHostBundle adds host bundle YAML to the main bundle YAML if host collectors exist.
func (e *EksaDiagnosticBundle) combineWithHostBundle(bundleYaml []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add a separator between the two YAML documents

// Combine the original bundle YAML with the host collector YAML

// createDiagnosticNamespace attempts to create the namespace eksa-diagnostics and associated RBAC objects.
// collector pods, for example host log collectors or run command collectors, will be launched in this namespace with the default service account.
// this method intentionally does not return an error
// a cluster in need of diagnosis may be unable to create new API objects and we should not stop our collection/analysis just because the namespace fails to create.
func (e *EksaDiagnosticBundle) createDiagnosticNamespaceAndRoles(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (e *EksaDiagnosticBundle) deleteDiagnosticNamespaceAndRoles(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func ParseTimeFromDuration(since string) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseTimeOptions(since string, sinceTime string) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// returning an uninitialized (zero) Time value here results in a
// sinceTimeValue of "0001-01-01 00:00:00 +0000 UTC"
// so all pod logs will be collected from the very beginning

func (e *EksaDiagnosticBundle) clusterName() string { _ = "STUB: not implemented"; return "" }
