package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/logger"
)

const (
	maxAgeFlagName = "max-age"
	tagFlagName    = "tag"
)

var cleanUpAwsCmd = &cobra.Command{
	Use:          "aws",
	Short:        "Clean up e2e resources on aws",
	Long:         "Clean up resources created for e2e testing on aws",
	SilenceUsage: true,
	PreRun:       preRunCleanUpAwsSetup,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cleanUpAwsTestResources(cmd.Context())
		if err != nil {
			logger.Fatal(err, "Failed to cleanup e2e resources on aws")
		}
		return nil
	},
}

func preRunCleanUpAwsSetup(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

var requiredAwsCleanUpFlags = []string{storageBucketFlagName, maxAgeFlagName, tagFlagName}

func init() {
	cleanUpInstancesCmd.AddCommand(cleanUpAwsCmd)
	cleanUpAwsCmd.Flags().StringP(storageBucketFlagName, "s", "", "Name of s3 bucket used for e2e testing")
	cleanUpAwsCmd.Flags().StringP(maxAgeFlagName, "a", "0", "Instance age in seconds after which it should be deleted")
	cleanUpAwsCmd.Flags().StringP(tagFlagName, "t", "", "EC2 instance tag")

	for _, flag := range requiredAwsCleanUpFlags {
		if err := cleanUpAwsCmd.MarkFlagRequired(flag); err != nil {
			log.Fatalf("Error marking flag %s as required: %v", flag, err)
		}
	}
}

func cleanUpAwsTestResources(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
