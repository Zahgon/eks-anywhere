package snow

import (
	"context"

	"github.com/google/uuid"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/filewriter"
)

type Defaulters struct {
	clientRegistry ClientRegistry
	writer         filewriter.FileWriter
	keyGenerator   SshKeyGenerator
	uuid           uuid.UUID
}

type SshKeyGenerator interface {
	GenerateSSHAuthKey(filewriter.FileWriter) (string, error)
}

type DefaultersOpt func(defaulters *Defaulters)

func NewDefaulters(clientRegistry ClientRegistry, writer filewriter.FileWriter, opts ...DefaultersOpt) *Defaulters {
	_ = "STUB: not implemented"
	return nil
}

// In the future if we need a cluster wide uuid that is shared, we should move this call to the dependency factory for reuse.

func WithKeyGenerator(generator SshKeyGenerator) DefaultersOpt {
	_ = "STUB: not implemented"
	return *new(DefaultersOpt)
}

// WithUUID will set uuid generated outside of constructor.
func WithUUID(uuid uuid.UUID) DefaultersOpt { _ = "STUB: not implemented"; return *new(DefaultersOpt) }

// GenerateDefaultSSHKeys generates ssh key if it doesn't exist already.
func (d *Defaulters) GenerateDefaultSSHKeys(ctx context.Context, machineConfigs map[string]*v1alpha1.SnowMachineConfig, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

type MachineConfigDefaulters struct {
	sshKey     string
	defaulters *Defaulters
}

func NewMachineConfigDefaulters(d *Defaulters) *MachineConfigDefaulters {
	_ = "STUB: not implemented"
	return nil
}

// SetupDefaultSSHKey creates and imports a default ssh key to snow devices listed in the snow machine config.
// If not exist, a ssh auth key is generated locally first. Then we loop through the devices in the machine config,
// and import the key to any device that does not have the key. In the end the default ssh key name is assigned to
// the snow machine config.
func (md *MachineConfigDefaulters) SetupDefaultSSHKey(ctx context.Context, m *v1alpha1.SnowMachineConfig, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (md *MachineConfigDefaulters) defaultSSHKeyName(clusterName string) string {
	_ = "STUB: not implemented"
	return ""
}

func SetupEksaCredentialsSecret(c *cluster.Config) error { _ = "STUB: not implemented"; return nil }
