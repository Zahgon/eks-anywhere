package e2e

import (
	e2etests "github.com/aws/eks-anywhere/test/framework"
)

var proxyVarsByProvider = map[string]e2etests.ProxyRequiredEnvVars{
	"CloudStack": e2etests.CloudstackProxyRequiredEnvVars,
	"VSphere":    e2etests.VsphereProxyRequiredEnvVars,
	"Tinkerbell": e2etests.TinkerbellProxyRequiredEnvVars,
}

func (e *E2ESession) setupProxyEnv(testRegex string) error { _ = "STUB: not implemented"; return nil }
