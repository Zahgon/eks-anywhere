package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/aws/eks-anywhere/pkg/logger"
)

var rootCmd = &cobra.Command{
	Use:              "anywhere",
	Short:            "Amazon EKS Anywhere",
	Long:             `Use eksctl anywhere to build your own self-managing cluster on your hardware with the best of Amazon EKS`,
	PersistentPreRun: rootPersistentPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		outputFilePath := logger.GetOutputFilePath()
		if outputFilePath == "" {
			return
		}
		if err := os.Remove(outputFilePath); err != nil {
			fmt.Printf("Failed to cleanup log file %s: %s", outputFilePath, err)
		}
	},
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

// RootCmd returns the eksctl-anywhere root cmd.
func RootCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }
