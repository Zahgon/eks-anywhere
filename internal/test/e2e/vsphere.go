package e2e

const (
	vsphereCidrVar               = "T_VSPHERE_CIDR"
	vspherePrivateNetworkCidrVar = "T_VSPHERE_PRIVATE_NETWORK_CIDR"
	vsphereRegex                 = `^.*VSphere.*$`
)

func (e *E2ESession) setupVSphereEnv(testRegex string) error { _ = "STUB: not implemented"; return nil }

// This algorithm is not very efficient with two nested loops
// Making the assumption that VSphereExtraEnvVarPrefixes() returns a very small number of prefixes
// this should be ok and probably not worth the complexity of building a more complex data structure.
// If in the future we see the need to have a bigger number of prefixes, we will need
// to change this to avoid the n*m complexity
