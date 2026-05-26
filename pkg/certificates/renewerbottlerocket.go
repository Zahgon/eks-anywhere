package certificates

import (
	"context"
)

const (
	brEtcdCertDir           = "/var/lib/etcd"
	brEtcdPkiDir            = "/var/lib/etcd/pki"
	brControlPlaneCertDir   = "/var/lib/kubeadm/pki"
	brControlPlaneManifests = "/var/lib/kubeadm/manifests"
	brTempDir               = "/run/host-containerd/io.containerd.runtime.v2.task/default/admin/rootfs/tmp"
)

// BottlerocketRenewer implements OSRenewer for Bottlerocket systems.
type BottlerocketRenewer struct {
	osType OSType
	backup string
}

// NewBottlerocketRenewer creates a new BottlerocketRenewer.
func NewBottlerocketRenewer(backupDir string) *BottlerocketRenewer {
	_ = "STUB: not implemented"
	return nil
}

// RenewControlPlaneCerts renews certificates for control plane nodes.
func (b *BottlerocketRenewer) RenewControlPlaneCerts(
	ctx context.Context,
	node string,
	cfg *RenewalConfig,
	component string,
	ssh SSHRunner,
) error {
	_ = "STUB: not implemented"
	return nil
}

// TransferCertsToControlPlaneFromLocal transfers etcd client certificates to a control plane node.
func (b *BottlerocketRenewer) TransferCertsToControlPlaneFromLocal(
	ctx context.Context, node string, ssh SSHRunner,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RenewEtcdCerts renews etcd certificates on a Bottlerocket node.
func (b *BottlerocketRenewer) RenewEtcdCerts(ctx context.Context, node string, ssh SSHRunner) error {
	_ = "STUB: not implemented"
	return nil
}

// CopyEtcdCertsToLocal copies the etcd certificates from the specified node to the local machine.
func (b *BottlerocketRenewer) CopyEtcdCertsToLocal(ctx context.Context, node string, ssh SSHRunner) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *BottlerocketRenewer) sheltie(commands ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BottlerocketRenewer) pullContainerImage() string { _ = "STUB: not implemented"; return "" }

func (b *BottlerocketRenewer) backupControlPlaneCerts(_ string, hasExternalEtcd bool, certDir string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BottlerocketRenewer) renewControlPlaneCerts() string { _ = "STUB: not implemented"; return "" }

func (b *BottlerocketRenewer) checkControlPlaneCerts() string { _ = "STUB: not implemented"; return "" }

func (b *BottlerocketRenewer) copyExternalEtcdCerts() string { _ = "STUB: not implemented"; return "" }

func (b *BottlerocketRenewer) restartControlPlaneStaticPods() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BottlerocketRenewer) backupEtcdCerts() string { _ = "STUB: not implemented"; return "" }

func (b *BottlerocketRenewer) renewEtcdCerts() string { _ = "STUB: not implemented"; return "" }

func (b *BottlerocketRenewer) validateEtcdCerts() string { _ = "STUB: not implemented"; return "" }

func (b *BottlerocketRenewer) copyEtcdCertsToTemp(tempDir string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BottlerocketRenewer) cleanupEtcdTempFiles(tempDir string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BottlerocketRenewer) createTempDirectory(tempLocalEtcdCertsDir string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BottlerocketRenewer) writeCertToTemp(certificateBytes64 string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BottlerocketRenewer) writeKeyToTemp(keyBytes64 string) string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BottlerocketRenewer) readTempFile(filePath string) string {
	_ = "STUB: not implemented"
	return ""
}
