package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources",
	Long:  "Use eksctl anywhere get to display one or many resources",
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func preRunPackages(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func getResources(ctx context.Context, resourceType, output, kubeConfig, clusterName, bundlesOverride string, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
