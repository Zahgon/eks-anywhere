package e2e

import (
	_ "embed"
	"regexp"
)

//go:embed IP_POOL_SIZE.yaml
var ipPoolSizeFile []byte

const defaultIPPoolSize = 1

// IPPoolRequirement maps a compiled regex pattern to the number of IPs needed.
type IPPoolRequirement struct {
	re         *regexp.Regexp
	ipPoolSize int
}

// ipPoolEntry is used for YAML unmarshalling.
type ipPoolEntry struct {
	Pattern    string `json:"pattern"`
	IPPoolSize int    `json:"ipPoolSize"`
}

// LoadIPPoolRequirements parses and compiles the embedded IP pool size YAML config.
func LoadIPPoolRequirements() ([]IPPoolRequirement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetIPPoolSize returns the IP pool size for a given test name.
// It checks the test name against each pattern in order and returns
// the IP pool size for the first match. If no pattern matches, it returns
// defaultIPPoolSize (1).
func GetIPPoolSize(testName string, requirements []IPPoolRequirement) int {
	_ = "STUB: not implemented"
	return 0
}
