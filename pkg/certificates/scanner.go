package certificates

import (
	"context"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

const (
	clusterNameLabel  = "cluster.x-k8s.io/cluster-name"
	controlPlaneLabel = "cluster.x-k8s.io/control-plane"
	externalEtcdLabel = "cluster.x-k8s.io/etcd-cluster"
)

// CertificateScanner defines the interface for checking certificate expiration.
type CertificateScanner interface {
	CheckCertificateExpiry(ctx context.Context, cluster *anywherev1.Cluster) ([]anywherev1.ClusterCertificateInfo, error)
	UpdateClusterCertificateStatus(ctx context.Context, cluster *anywherev1.Cluster) error
}

// Scanner implements the CertificateScanner interface and provides certificate checking functionality.
type Scanner struct {
	client client.Client
	logger logr.Logger
}

// NewCertificateScanner creates a new certificate service.
func NewCertificateScanner(client client.Client, logger logr.Logger) *Scanner {
	_ = "STUB: not implemented"
	return nil
}

// MachineInfo holds machine name and IP information.
type MachineInfo struct {
	Name string
	IP   string
}

// CheckCertificateExpiry checks the certificate expiration for control plane and etcd machines.
func (s *Scanner) CheckCertificateExpiry(ctx context.Context, cluster *anywherev1.Cluster) ([]anywherev1.ClusterCertificateInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scanner) getControlPlaneMachines(ctx context.Context, cluster *anywherev1.Cluster) ([]MachineInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scanner) getEtcdMachines(ctx context.Context, cluster *anywherev1.Cluster) ([]MachineInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Scanner) getMachinesCertificateInfo(machines []MachineInfo, port string) []anywherev1.ClusterCertificateInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) checkMachineCertificateExpiry(ip, port string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We just want to get the certificate, not verify it

// Use the first certificate (leaf certificate)

// UpdateClusterCertificateStatus updates the cluster status with certificate information.
func (s *Scanner) UpdateClusterCertificateStatus(ctx context.Context, cluster *anywherev1.Cluster) error {
	_ = "STUB: not implemented"
	return nil
}
