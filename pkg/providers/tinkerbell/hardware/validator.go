package hardware

import (
	"regexp"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
)

// MachineAssertion defines a condition that Machine must meet.
type MachineAssertion func(Machine) error

// DefaultMachineValidator validated Machine instances.
type DefaultMachineValidator struct {
	assertions []MachineAssertion
}

var _ MachineValidator = &DefaultMachineValidator{}

// NewDefaultMachineValidator creates a machineValidator instance with default assertions registered.
func NewDefaultMachineValidator() *DefaultMachineValidator { _ = "STUB: not implemented"; return nil }

// Validate validates machine by executing its Validate() method and passing it to all registered MachineAssertions.
func (mv *DefaultMachineValidator) Validate(machine Machine) error {
	_ = "STUB: not implemented"
	return nil
}

// Register registers v MachineAssertions with m.
func (mv *DefaultMachineValidator) Register(v ...MachineAssertion) {
	_ = "STUB: not implemented"
	return
}

var (
	linuxPathRegex      = `^(/dev/[\w-]+)+$`
	linuxPathValidation = regexp.MustCompile(linuxPathRegex)
)

// StaticMachineAssertions defines all static data assertions performed on a Machine.
func StaticMachineAssertions() MachineAssertion {
	_ = "STUB: not implemented"
	return *new(MachineAssertion)
}

// valid VLAN IDs are between 1 and 4094 - https://en.m.wikipedia.org/wiki/VLAN#IEEE_802.1Q

// UniqueIPAddress asserts a given Machine instance has a unique IPAddress field relative to previously seen Machine
// instances. It is not thread safe. It has a 1 time use.
func UniqueIPAddress() MachineAssertion { _ = "STUB: not implemented"; return *new(MachineAssertion) }

// UniqueMACAddress asserts a given Machine instance has a unique MACAddress field relative to previously seen Machine
// instances. It is not thread safe. It has a 1 time use.
func UniqueMACAddress() MachineAssertion { _ = "STUB: not implemented"; return *new(MachineAssertion) }

// UniqueHostnames asserts a given Machine instance has a unique Hostname field relative to previously seen Machine
// instances. It is not thread safe. It has a 1 time use.
func UniqueHostnames() MachineAssertion { _ = "STUB: not implemented"; return *new(MachineAssertion) }

// UniqueBMCIPAddress asserts a given Machine instance has a unique BMCIPAddress field relative to previously seen
// Machine instances. If there is no BMC configuration as defined by machine.HasBMC() the check is a noop. It is
// not thread safe. It has a 1 time use.
func UniqueBMCIPAddress() MachineAssertion {
	_ = "STUB: not implemented"
	return *new(MachineAssertion)
}

// RegisterDefaultAssertions applies a set of default assertions to validator. The default assertions
// include UniqueHostnames and UniqueIDs.
func RegisterDefaultAssertions(validator *DefaultMachineValidator) {
	_ = "STUB: not implemented"
	return
}

func validateLabelKey(k string) error { _ = "STUB: not implemented"; return nil }

func validateLabelValue(v string) error { _ = "STUB: not implemented"; return nil }

// LabelsMatchSelector ensures all selector key-value pairs can be found in labels.
// If selector is empty true is always returned.
func LabelsMatchSelector(selector v1alpha1.HardwareSelector, labels Labels) bool {
	_ = "STUB: not implemented"
	return false
}
