package diagnostics

import (
	_ "embed"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
)

type EksaDiagnosticBundleFactoryOpts struct {
	AnalyzerFactory  AnalyzerFactory
	Client           BundleClient
	CollectorFactory CollectorFactory
	Kubectl          *executables.Kubectl
	Writer           filewriter.FileWriter
}

type eksaDiagnosticBundleFactory struct {
	analyzerFactory  AnalyzerFactory
	client           BundleClient
	collectorFactory CollectorFactory
	kubectl          *executables.Kubectl
	writer           filewriter.FileWriter
}

func NewFactory(opts EksaDiagnosticBundleFactoryOpts) *eksaDiagnosticBundleFactory {
	_ = "STUB: not implemented"
	return nil
}

func (f *eksaDiagnosticBundleFactory) DiagnosticBundle(spec *cluster.Spec, provider providers.Provider, kubeconfig string, bundlePath string, auditLogs bool) (DiagnosticBundle, error) {
	_ = "STUB: not implemented"
	return *new(DiagnosticBundle), nil
}

func (f *eksaDiagnosticBundleFactory) DiagnosticBundleManagementCluster(spec *cluster.Spec, kubeconfig string) (DiagnosticBundle, error) {
	_ = "STUB: not implemented"
	return *new(DiagnosticBundle), nil
}

func (f *eksaDiagnosticBundleFactory) DiagnosticBundleWorkloadCluster(spec *cluster.Spec, provider providers.Provider, kubeconfig string, auditLogs bool) (DiagnosticBundle, error) {
	_ = "STUB: not implemented"
	return *new(DiagnosticBundle), nil
}

func (f *eksaDiagnosticBundleFactory) DiagnosticBundleDefault() DiagnosticBundle {
	_ = "STUB: not implemented"
	return *new(DiagnosticBundle)
}

func (f *eksaDiagnosticBundleFactory) DiagnosticBundleCustom(kubeconfig string, bundlePath string) DiagnosticBundle {
	_ = "STUB: not implemented"
	return *new(DiagnosticBundle)
}
