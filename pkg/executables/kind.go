package executables

import (
	"context"
	_ "embed"

	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/registrymirror"
	"github.com/aws/eks-anywhere/pkg/types"
)

const kindPath = "kind"

//go:embed config/kind.yaml
var kindConfigTemplate string

//go:embed config/hosts.toml
var hostsTomlTemplate string

const configFileName = "kind_tmp.yaml"

type Kind struct {
	writer filewriter.FileWriter
	Executable
	execConfig *kindExecConfig
}

// kindExecConfig contains transient information for the execution of kind commands
// It's used by BootstrapClusterClientOption's to store/change information prior to a command execution
// It must be cleaned after each execution to prevent side effects from past executions options.
type kindExecConfig struct {
	env                  map[string]string
	ConfigFile           string
	KindImage            string
	KubernetesRepository string
	EtcdRepository       string
	EtcdVersion          string
	CorednsRepository    string
	CorednsVersion       string
	KubernetesVersion    string
	RegistryConfigDir    string
	ExtraPortMappings    []int
	DockerExtraMounts    bool
	DisableDefaultCNI    bool
	PodSubnet            string
	ServiceSubnet        string
	AuditPolicyPath      string
}

func NewKind(executable Executable, writer filewriter.FileWriter) *Kind {
	_ = "STUB: not implemented"
	return nil
}

// CreateAuditPolicy creates an audit policy file to be used by the bootstrap cluster's api server.
func (k *Kind) CreateAuditPolicy(clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kind) CreateBootstrapCluster(ctx context.Context, clusterSpec *cluster.Spec, opts ...bootstrapper.BootstrapClusterClientOption) (kubeconfig string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *Kind) ClusterExists(ctx context.Context, clusterName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (k *Kind) GetKubeconfig(ctx context.Context, clusterName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *Kind) WithExtraDockerMounts() bootstrapper.BootstrapClusterClientOption {
	_ = "STUB: not implemented"
	return *new(bootstrapper.BootstrapClusterClientOption)
}

func (k *Kind) WithExtraPortMappings(ports []int) bootstrapper.BootstrapClusterClientOption {
	_ = "STUB: not implemented"
	return *new(bootstrapper.BootstrapClusterClientOption)
}

func (k *Kind) WithEnv(env map[string]string) bootstrapper.BootstrapClusterClientOption {
	_ = "STUB: not implemented"
	return *new(bootstrapper.BootstrapClusterClientOption)
}

func (k *Kind) DeleteBootstrapCluster(ctx context.Context, cluster *types.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}

// RegistryConfig contains configuration for setting up a registry with containerd.
type RegistryConfig struct {
	Server     string // The server URL for the registry
	Host       string // The host URL to redirect to
	CACertPath string // CA certificate path
	AuthHeader string
	OutputDir  string // Directory where to write hosts.toml
}

// setupRegistryConfig creates a registry configuration hosts.toml file.
func setupRegistryConfig(config RegistryConfig) error { _ = "STUB: not implemented"; return nil }

// Generate hosts.toml content using template

// Write hosts.toml file

func (k *Kind) setupExecConfig(clusterSpec *cluster.Spec) error {
	_ = "STUB: not implemented"
	return nil
}

// setupRegistryMirror handles the complete setup of registry mirror configuration.
func (k *Kind) setupRegistryMirror(clusterSpec *cluster.Spec, registryMirror *registrymirror.RegistryMirror) error {
	_ = "STUB: not implemented"
	return nil
}

// Create the base certs.d directory

// Generate authorization header if authentication is required

// Create CA certificate only once for the base registry and determine shared CA path

// Create the base registry directory and CA certificate file

// Write the CA certificate file

// Setup configuration for the mirror registry

// Setup configuration for each original registry that should be mirrored

func (k *Kind) cleanExecConfig() { _ = "STUB: not implemented"; return }

func (k *Kind) buildConfigFile() error { _ = "STUB: not implemented"; return nil }

func (k *Kind) execArguments(clusterName string, kubeconfigName string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (k *Kind) createKubeConfig(clusterName string, content []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func processOpts(opts []bootstrapper.BootstrapClusterClientOption) error {
	_ = "STUB: not implemented"
	return nil
}

func getInternalName(clusterName string) string { _ = "STUB: not implemented"; return "" }
