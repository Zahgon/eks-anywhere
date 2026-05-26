package flux

import (
	_ "embed"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/providers"
)

const (
	eksaSystemDirName     = "eksa-system"
	kustomizeFileName     = "kustomization.yaml"
	clusterConfigFileName = "eksa-cluster.yaml"
	fluxSyncFileName      = "gotk-sync.yaml"
	fluxPatchFileName     = "gotk-patches.yaml"
)

//go:embed manifests/eksa-system/kustomization.yaml
var eksaKustomizeContent string

//go:embed manifests/flux-system/kustomization.yaml
var fluxKustomizeContent string

//go:embed manifests/flux-system/gotk-sync.yaml
var fluxSyncContent string

type Templater interface {
	WriteToFile(templateContent string, data interface{}, fileName string, f ...filewriter.FileOptionsFunc) (filePath string, err error)
}

type FileGenerator struct {
	fluxWriter, eksaWriter       filewriter.FileWriter
	fluxTemplater, eksaTemplater Templater
}

func NewFileGenerator() *FileGenerator { _ = "STUB: not implemented"; return nil }

// NewFileGeneratorWithWriterTemplater takes flux and eksa writer and templater interface to build the generator.
// This is only for testing.
func NewFileGeneratorWithWriterTemplater(fluxWriter, eksaWriter filewriter.FileWriter, fluxTemplater, eksaTemplater Templater) *FileGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (g *FileGenerator) Init(writer filewriter.FileWriter, eksaSystemDir, fluxSystemDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *FileGenerator) WriteEksaFiles(clusterSpec *cluster.Spec, datacenterConfig providers.DatacenterConfig, machineConfigs []providers.MachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteFluxSystemFiles writes the flux-system files into the flux system git directory.
func (g *FileGenerator) WriteFluxSystemFiles(managementComponents *cluster.ManagementComponents, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *FileGenerator) WriteClusterConfig(clusterSpec *cluster.Spec, datacenterConfig providers.DatacenterConfig, machineConfigs []providers.MachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *FileGenerator) WriteEksaKustomization() error { _ = "STUB: not implemented"; return nil }

// WriteFluxKustomization writes the flux-system kustomization file into the flux system git directory.
func (g *FileGenerator) WriteFluxKustomization(managementComponents *cluster.ManagementComponents, clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *FileGenerator) WriteFluxSync() error { _ = "STUB: not implemented"; return nil }
