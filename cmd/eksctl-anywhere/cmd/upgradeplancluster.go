package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/types"
)

const (
	outputFlagName = "output"
	outputDefault  = outputText
	outputText     = "text"
	outputJson     = "json"
)

var output string

var upgradePlanClusterCmd = &cobra.Command{
	Use:          "cluster",
	Short:        "Provides new release versions for the next cluster upgrade",
	Long:         "Provides a list of target versions for upgrading the core components in the workload cluster",
	PreRunE:      bindFlagsToViper,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := uc.upgradePlanCluster(cmd.Context()); err != nil {
			return fmt.Errorf("failed to display upgrade plan: %v", err)
		}
		return nil
	},
}

func init() {
	upgradePlanCmd.AddCommand(upgradePlanClusterCmd)
	upgradePlanClusterCmd.Flags().StringVarP(&uc.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	upgradePlanClusterCmd.Flags().StringVar(&uc.bundlesOverride, "bundles-override", "", "Override default Bundles manifest (not recommended)")
	upgradePlanClusterCmd.Flags().StringVarP(&output, outputFlagName, "o", outputDefault, "Output format: text|json")
	upgradePlanClusterCmd.Flags().StringVar(&uc.managementKubeconfig, "kubeconfig", "", "Management cluster kubeconfig file")
	err := upgradePlanClusterCmd.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
}

func (uc *upgradeClusterOptions) upgradePlanCluster(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func serialize(componentChangeDiffs *types.ChangeDiff, outputFormat string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func serializeToText(componentChangeDiffs *types.ChangeDiff) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func serializeToJson(componentChangeDiffs *types.ChangeDiff) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
