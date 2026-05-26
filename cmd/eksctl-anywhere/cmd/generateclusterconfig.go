package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/validations"
)

var removeFromDefaultConfig = []string{"spec.clusterNetwork.dns"}

var generateClusterConfigCmd = &cobra.Command{
	Use:    "clusterconfig <cluster-name> (max 80 chars)",
	Short:  "Generate cluster config",
	Long:   "This command is used to generate a cluster config yaml for the create cluster command",
	PreRun: preRunGenerateClusterConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		clusterName, err := validations.ValidateClusterNameArg(args)
		if err != nil {
			return err
		}
		err = generateClusterConfig(clusterName)
		if err != nil {
			return fmt.Errorf("generating eks-a cluster config: %v", err) // need to have better error handling here in own func
		}
		return nil
	},
}

func preRunGenerateClusterConfig(cmd *cobra.Command, args []string) {
	_ = "STUB: not implemented"
	return
}

func init() {
	generateCmd.AddCommand(generateClusterConfigCmd)
	generateClusterConfigCmd.Flags().StringP("provider", "p", "", fmt.Sprintf("Provider to use (%s)", strings.Join(constants.SupportedProviders, " or ")))
	err := generateClusterConfigCmd.MarkFlagRequired("provider")
	if err != nil {
		log.Fatalf("marking flag as required: %v", err)
	}
}

func generateClusterConfig(clusterName string) error { _ = "STUB: not implemented"; return nil }

// need to default control plane config name to something different from the cluster name based on assumption
// in controller code

// need to default control plane config name to something different from the cluster name based on assumption
// in controller code
