package e2e

import (
	_ "embed"
)

//go:embed TINKERBELL_HARDWARE_COUNT.yaml
var tinkerbellHardwareCountFile []byte

// GetTinkerbellTestsHardwareRequirements returns a map of Tinkerbell test name to required hardware.
func GetTinkerbellTestsHardwareRequirements() (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
