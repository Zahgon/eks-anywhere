package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:              "eks-a-test-tool",
	Short:            "Amazon EKS Anywhere Test Tooling",
	Long:             `Use eks-a-test-tool to evaluate EKS-A test results`,
	PersistentPreRun: rootPersistentPreRun,
}

func init() {
	rootCmd.PersistentFlags().IntP("verbosity", "v", 0, "Set the log level verbosity")
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatalf("failed to bind flags for root: %v", err)
	}
}

func rootPersistentPreRun(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

func initLogger() error { _ = "STUB: not implemented"; return nil }

func Execute() error { _ = "STUB: not implemented"; return nil }
