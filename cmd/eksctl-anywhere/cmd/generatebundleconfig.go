package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/diagnostics"
)

type generateSupportBundleOptions struct {
	fileName              string
	hardwareFileName      string
	tinkerbellBootstrapIP string
}

var gsbo = &generateSupportBundleOptions{}

var generateBundleConfigCmd = &cobra.Command{
	Use:     "support-bundle-config",
	Short:   "Generate support bundle config",
	Long:    "This command is used to generate a default support bundle config yaml",
	PreRunE: bindFlagsToViper,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gsbo.validateCmdInput()
		if err != nil {
			return fmt.Errorf("command input validation failed: %v", err)
		}
		bundle, err := gsbo.generateBundleConfig(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to generate bunlde config: %v", err)
		}
		err = bundle.PrintBundleConfig()
		if err != nil {
			return fmt.Errorf("failed to print bundle config: %v", err)
		}
		return nil
	},
}

func init() {
	generateCmd.AddCommand(generateBundleConfigCmd)
	generateBundleConfigCmd.Flags().StringVarP(&gsbo.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
}

func (gsbo *generateSupportBundleOptions) validateCmdInput() error {
	_ = "STUB: not implemented"
	return nil
}

func (gsbo *generateSupportBundleOptions) generateBundleConfig(ctx context.Context) (diagnostics.DiagnosticBundle, error) {
	_ = "STUB: not implemented"
	return *new(diagnostics.DiagnosticBundle), nil
}

func (gsbo *generateSupportBundleOptions) generateDefaultBundleConfig(ctx context.Context) (diagnostics.DiagnosticBundle, error) {
	_ = "STUB: not implemented"
	return *new(diagnostics.DiagnosticBundle), nil
}
