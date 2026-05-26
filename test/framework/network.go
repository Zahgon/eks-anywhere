package framework

func PopIPFromEnv(ipPoolEnvVar string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// PopIPFromEnv will remove the ip from the pool.
// Therefore, we rewrite the envvar to the system so the next caller can pick from remaining ips in the pool

func GenerateUniqueIp(cidr string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetIP(cidr, ipEnvVar string) (string, error) { _ = "STUB: not implemented"; return "", nil }
