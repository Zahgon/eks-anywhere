package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var autofillCmd = &cobra.Command{
	Use:    "autofill",
	Short:  "Autofill provider config",
	Long:   "Fills provider config with values set in environment variables",
	PreRun: preRunAutofill,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := autofill(cmd.Context())
		if err != nil {
			log.Fatalf("Error filling the provider config: %v", err)
		}
		return nil
	},
}

func init() {
	vsphereCmd.AddCommand(autofillCmd)
	autofillCmd.Flags().StringP("filename", "f", "", "Cluster config yaml filepath")
	err := autofillCmd.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
}

func preRunAutofill(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

func autofill(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
