package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Get cluster versions",
	Long:  "Get the versions of images in cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := versions(cmd.Context())
		if err != nil {
			log.Fatalf("Error getting image versions: %v", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}

func versions(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
