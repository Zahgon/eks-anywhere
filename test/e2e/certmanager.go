//go:build e2e
// +build e2e

package e2e

import (
	"github.com/aws/eks-anywhere/test/framework"
)

const (
	cmPackageName = "cert-manager"
)

func runCertManagerRemoteClusterInstallSimpleFlow(test *framework.MulticlusterE2ETest) {
	_ = "STUB: not implemented"
	return
}

// Ensure cleanup happens even if the test fails
