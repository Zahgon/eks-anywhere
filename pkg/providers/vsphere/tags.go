package vsphere

import (
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
)

func requiredTemplateTags(machineConfig *v1alpha1.VSphereMachineConfig, versionsBundle *cluster.VersionsBundle) []string {
	_ = "STUB: not implemented"
	return nil
}

func requiredTemplateTagsByCategory(machineConfig *v1alpha1.VSphereMachineConfig, versionsBundle *cluster.VersionsBundle) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}
