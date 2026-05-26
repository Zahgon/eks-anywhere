package certificates

import (
	"context"
)

const (
	linuxEtcdCertDir           = "/etc/etcd"
	linuxControlPlaneCertDir   = "/etc/kubernetes/pki"
	linuxControlPlaneManifests = "/etc/kubernetes/manifests"
	linuxTempDir               = "/tmp"
)

// LinuxRenewer implements OSRenewer for Linux-based systems (Ubuntu / RHEL).
type LinuxRenewer struct {
	osType OSType
	backup string
}

// NewLinuxRenewer creates a new renewer for Linux-based operating systems.
func NewLinuxRenewer(backupDir string) *LinuxRenewer { _ = "STUB: not implemented"; return nil }

// RenewControlPlaneCerts renews certificates for control plane nodes.
func (l *LinuxRenewer) RenewControlPlaneCerts(
	ctx context.Context,
	node string,
	cfg *RenewalConfig,
	component string,
	ssh SSHRunner,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RenewEtcdCerts renews certificates for etcd nodes.
func (l *LinuxRenewer) RenewEtcdCerts(
	ctx context.Context,
	node string,
	ssh SSHRunner,
) error {
	_ = "STUB: not implemented"
	return nil
}

// CopyEtcdCertsToLocal copies the etcd certificates from the specified node to the local machine.
func (l *LinuxRenewer) CopyEtcdCertsToLocal(
	ctx context.Context,
	node string,
	ssh SSHRunner,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *LinuxRenewer) backupControlPlaneCerts(_ string, hasExternalEtcd bool, backup string) string {
	_ = "STUB: not implemented"
	return ""
}

func (l *LinuxRenewer) renewControlPlaneCerts(_ string, hasExternalEtcd bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (l *LinuxRenewer) restartControlPlaneStaticPods() string { _ = "STUB: not implemented"; return "" }

func (l *LinuxRenewer) backupEtcdCerts() string { _ = "STUB: not implemented"; return "" }

func (l *LinuxRenewer) validateEtcdCerts() string { _ = "STUB: not implemented"; return "" }

// TransferCertsToControlPlaneFromLocal transfers etcd client certificates to a control plane node.
func (l *LinuxRenewer) TransferCertsToControlPlaneFromLocal(
	ctx context.Context, node string, ssh SSHRunner,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *LinuxRenewer) copyExternalEtcdCerts() string { _ = "STUB: not implemented"; return "" }
