package vsphere

import (
	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

func SetupEnvVars(datacenterConfig *anywherev1.VSphereDatacenterConfig) error {
	_ = "STUB: not implemented"
	// TODO(cxbrowne): We set environment variables here in response to existing of other environment
	// variables. Investigate why this is done, and possible remove the need for this.
	// https://github.com/aws/eks-anywhere-internal/issues/2192
	return nil
}

// TODO: move this somewhere else since it's not vSphere specific
