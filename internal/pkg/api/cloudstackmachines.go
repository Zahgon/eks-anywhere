package api

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

type CloudStackMachineConfigFiller func(m *anywherev1.CloudStackMachineConfig)

func FillCloudStackMachineConfig(m *anywherev1.CloudStackMachineConfig, fillers ...CloudStackMachineConfigFiller) {
	_ = "STUB: not implemented"
	return
}

func WithCloudStackComputeOffering(value string) CloudStackMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackMachineConfigFiller)
}

func WithCloudStackSSHKey(value string) CloudStackMachineConfigFiller {
	_ = "STUB: not implemented"
	return *new(CloudStackMachineConfigFiller)
}

func setCloudStackSSHKeyForFirstUser(m *anywherev1.CloudStackMachineConfig, key string) {
	_ = "STUB: not implemented"
	return
}
