package networkutils

type IPPool []string

func NewIPPool() IPPool { _ = "STUB: not implemented"; return *new(IPPool) }

func NewIPPoolFromString(fromString string) IPPool { _ = "STUB: not implemented"; return *new(IPPool) }

func NewIPPoolFromEnv(ipPoolEnvVar string) (IPPool, error) {
	_ = "STUB: not implemented"
	return *new(IPPool), nil
}

func (ipPool *IPPool) ToString() string { _ = "STUB: not implemented"; return "" }

func (ipPool *IPPool) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (ipPool *IPPool) AddIP(ip string) { _ = "STUB: not implemented"; return }

func (ipPool *IPPool) PopIP() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (ipPool *IPPool) ToEnvVar(envVarName string) error { _ = "STUB: not implemented"; return nil }
