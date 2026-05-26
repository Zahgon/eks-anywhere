package vsphere

import (
	"context"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

const minDiskGib int = 20

type Defaulter struct {
	govc ProviderGovcClient
}

func NewDefaulter(govc ProviderGovcClient) *Defaulter { _ = "STUB: not implemented"; return nil }

func (d *Defaulter) setDefaultsForMachineConfig(ctx context.Context, spec *Spec) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Defaulter) SetDefaultsForDatacenterConfig(ctx context.Context, datacenterConfig *anywherev1.VSphereDatacenterConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func setDefaultsForEtcdMachineConfig(machineConfig *anywherev1.VSphereMachineConfig) {
	_ = "STUB: not implemented"
	return
}

func (d *Defaulter) setWorkerDefaultTemplateIfMissing(ctx context.Context, spec *Spec, workerNodeGroup anywherev1.WorkerNodeGroupConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Defaulter) setDefaultTemplateIfMissing(ctx context.Context, spec *Spec, m *anywherev1.VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Defaulter) setupDefaultTemplate(ctx context.Context, spec *Spec, machineConfig *anywherev1.VSphereMachineConfig, versionsBundle *cluster.VersionsBundle) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: figure out if it's worth refactoring the factory to be able to reuse across machine configs.

// TODO: remove the factory's dependency on a machineConfig

func max(a, b int) int { _ = "STUB: not implemented"; return 0 }

func (d *Defaulter) setCloneModeAndDiskSizeDefaults(ctx context.Context, machineConfig *anywherev1.VSphereMachineConfig, datacenter string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMachineWithNoCloneMode(templateHasSnapshot bool, templateDiskSize int, machineConfig *anywherev1.VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMachineWithLinkedCloneMode(templateHasSnapshot bool, templateDiskSize int, machineConfig *anywherev1.VSphereMachineConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Defaulter) setTemplateFullPath(ctx context.Context,
	datacenterConfig *anywherev1.VSphereDatacenterConfig,
	machine *anywherev1.VSphereMachineConfig,
) error {
	_ = "STUB: not implemented"
	return nil
}
