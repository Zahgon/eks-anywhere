package cmd

import (
	"github.com/spf13/cobra"
)

const shortGenerateTinkerbellTemplateConfigHelp = "Generate TinkerbellTemplateConfig objects"

const longGenerateTinkerbellTemplateConfigHelp = `Generate TinkerbellTemplateConfig objects for your cluster specification.

The TinkerbellTemplateConfig is part of an EKS Anywhere bare metal cluster 
specification. When no template config is specified on TinkerbellMachineConfig
objects, EKS Anywhere generates the template config internally. The template 
config defines the actions for provisioning a bare metal host such as streaming 
an OS image to disk. Actions vary based on the OS - see the EKS Anywhere 
documentation for more details on the individual actions.

The template config include it in your bare metal cluster specification and
reference it in the TinkerbellMachineConfig object using the .spec.templateRef
field.
`

// NewGenerateTinkerbellTemplateConfig creates a command that will generate a TinkerbellTemplateConfig
// using the cluster configuration.
func NewGenerateTinkerbellTemplateConfig() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Configure the flagset. Some of these flags are duplicated from other parts of the cmd code
// for consistency but their descriptions may vary because of the commands use-case.

// When the bootstrap IP is unspecified attempt to derive it from IPs assigned to the
// primary interface.

// Validation logic called by newClusterSpec arbitrarily logs warnings. Until it can
// be refactored we need to redirect logging such that the output is swallowed.

// Reset the logger before evaluating anything in-case logic higher up the call path
// needs the logger.

// Handle the newClusterSpec() error.

// Generating the TinkerbellTemplateConfig requires the OS family to ensure the right
// actions are produced. The OS family is specified per TinkerbellMachineConfig,
// therefore is specified in multiple places (control plane and worker node groups).
// However, we only support single OS family clusters so validation should error out
// earlier if they aren't the same. This means we can use the control plane machine
// configs OS family.

// Configure the commands flags.
