package diagnostics

import (
	v1 "k8s.io/api/core/v1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
)

// FileReader reads files from local disk or http urls.
type FileReader interface {
	ReadFile(url string) ([]byte, error)
}

// EKSACollectorFactory generates support-bundle collectors for eks-a clusters.
type EKSACollectorFactory struct {
	DiagnosticCollectorImage string
	reader                   FileReader
}

// NewCollectorFactory builds a collector factory.
func NewCollectorFactory(diagnosticCollectorImage string, reader FileReader) *EKSACollectorFactory {
	_ = "STUB: not implemented"
	return nil
}

// NewDefaultCollectorFactory builds a collector factory that will use the default
// diagnostic collector image.
func NewDefaultCollectorFactory(reader FileReader) *EKSACollectorFactory {
	_ = "STUB: not implemented"
	return nil
}

// DefaultCollectors returns the collectors that apply to all clusters.
func (c *EKSACollectorFactory) DefaultCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

// EksaHostCollectors returns the collectors that interact with the kubernetes node machine hosts.
func (c *EKSACollectorFactory) EksaHostCollectors(machineConfigs []providers.MachineConfig) []*Collect {
	_ = "STUB: not implemented"
	return nil
}

// we don't want to duplicate the collectors if multiple machine configs have the same OS family

// HostCollectors returns the collectors that run on host machines.
func (c *EKSACollectorFactory) HostCollectors(datacenter v1alpha1.Ref) []*Collect {
	_ = "STUB: not implemented"
	// Only Tinkerbell needs this right now to collect boots/smee logs in docker container.
	return nil
}

// AuditLogCollectors returns the audit log collectors that run on all control plane nodes.
func (c *EKSACollectorFactory) AuditLogCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

// DataCenterConfigCollectors returns the collectors for the provider datacenter config in the cluster spec.
func (c *EKSACollectorFactory) DataCenterConfigCollectors(datacenter v1alpha1.Ref, spec *cluster.Spec) []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) eksaNutanixCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) eksaSnowCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) eksaTinkerbellCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) eksaVsphereCollectors(spec *cluster.Spec) []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) eksaCloudstackCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) eksaDockerCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) hostTinkerbellCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

// ManagementClusterCollectors returns the collectors that only apply to management clusters.
func (c *EKSACollectorFactory) ManagementClusterCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

// PackagesCollectors returns the collectors that read information for curated packages.
func (c *EKSACollectorFactory) PackagesCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

// FileCollectors returns the collectors that interact with files.
func (c *EKSACollectorFactory) FileCollectors(paths []string) []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) getCollectorsMap() map[v1alpha1.OSFamily][]*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) bottleRocketHostCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) ubuntuHostCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) defaultLogCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) packagesLogCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) managementClusterLogCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

// apiServerCollectors collect connection info when running a pod on an existing cluster.
func (c *EKSACollectorFactory) apiServerCollectors(controlPlaneIP string) []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) controlPlaneNetworkPathCollector(controlPlaneIP string) []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) hostPortCollector(ports []string, hostIP string) *Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) pingHostCollector(hostIP string) *Collect {
	_ = "STUB: not implemented"
	return nil
}

// vmsAccessCollector will connect to API server first, then collect vsphere-cloud-controller-manager logs
// on control plane node.
func (c *EKSACollectorFactory) vmsAccessCollector(controlPlaneConfiguration v1alpha1.ControlPlaneConfiguration) *Collect {
	_ = "STUB: not implemented"
	return nil
}

func makeTolerations(taints []v1.Taint) []v1.Toleration { _ = "STUB: not implemented"; return nil }

func logpath(namespace string) string { _ = "STUB: not implemented"; return "" }

func hostlogPath(logType string) string { _ = "STUB: not implemented"; return "" }

// webhookConfigCollectors returns collectors for admission webhook configurations.
func (c *EKSACollectorFactory) webhookConfigCollectors() []*Collect {
	_ = "STUB: not implemented"
	return nil
}

func (c *EKSACollectorFactory) webhookConfigCollector(resourceType string) *Collect {
	_ = "STUB: not implemented"
	return nil
}

// It's possible for networking to not be working on the cluster or the nodes
// not being ready, so adding tolerations and running the pod on host networking
// to be able to pull the resources from the cluster
