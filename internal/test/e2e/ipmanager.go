package e2e

import (
	"github.com/go-logr/logr"

	"github.com/aws/eks-anywhere/pkg/networkutils"
)

type E2EIPManager struct {
	networkCidr string
	networkIPs  map[string]bool
	logger      logr.Logger
}

func newE2EIPManager(logger logr.Logger, networkCidr string) *E2EIPManager {
	_ = "STUB: not implemented"
	return nil
}

func (ipman *E2EIPManager) reserveIP() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (ipman *E2EIPManager) reserveIPPool(count int) (networkutils.IPPool, error) {
	_ = "STUB: not implemented"
	return *new(networkutils.IPPool), nil
}

func (ipman *E2EIPManager) getUniqueIP(cidr string, usedIPs map[string]bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
