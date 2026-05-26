package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces"
)

var upgradePlanManagementComponentsCmd = &cobra.Command{
	Use:          "management-components",
	Short:        "Lists the current and target versions for upgrading the management components in a management cluster",
	Long:         "Provides a list of current and target versions for upgrading the management components in a management cluster. The term 'management components' encompasses all Kubernetes controllers and their CRDs present in the management cluster that are responsible for reconciling your EKS Anywhere (EKS-A) cluster.",
	PreRunE:      bindFlagsToViper,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := uc.upgradePlanManagementComponents(cmd.Context()); err != nil {
			return fmt.Errorf("failed to display upgrade plan: %v", err)
		}
		return nil
	},
}

func init() {
	upgradePlanCmd.AddCommand(upgradePlanManagementComponentsCmd)
	upgradePlanManagementComponentsCmd.Flags().StringVarP(&uc.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	upgradePlanManagementComponentsCmd.Flags().StringVar(&uc.bundlesOverride, "bundles-override", "", "Override default Bundles manifest (not recommended)")
	upgradePlanManagementComponentsCmd.Flags().StringVarP(&output, outputFlagName, "o", outputDefault, "Output format: text|json")
	upgradePlanManagementComponentsCmd.Flags().StringVar(&uc.managementKubeconfig, "kubeconfig", "", "Management cluster kubeconfig file")
	err := upgradePlanManagementComponentsCmd.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
}

func (uc *upgradeClusterOptions) upgradePlanManagementComponents(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func getManagementComponentsChangeDiffs(ctx context.Context, clientFactory interfaces.ClientFactory, managementCluster *types.Cluster, currentSpec *cluster.Spec, newClusterSpec *cluster.Spec, provider providers.Provider) (*types.ChangeDiff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
