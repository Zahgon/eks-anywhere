package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/validations"
)

var validateClusterCmd = &cobra.Command{
	Use:   "validate-cluster <cluster-name> <kubeconfig>",
	Short: "Validate eks-a cluster command",
	Long:  "Use eks-a-tool validate eks-anywhere cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			log.Fatalf("Some args are missing. See usage for required arguments")
		}
		clusterName, err := validations.ValidateClusterNameArg(args)
		if err != nil {
			log.Fatalf("Error validating the cluster: %v", err)
		}
		kubeconfig := args[1]
		if !validations.FileExists(kubeconfig) {
			log.Fatalf("Error validating the cluster: kubeconfig file %s not found", kubeconfig)
		}
		cluster := &types.Cluster{
			Name:           clusterName,
			KubeconfigFile: kubeconfig,
		}
		err = validateCluster(cmd.Context(), cluster, clusterName)
		if err != nil {
			log.Fatalf("Error validating the cluster: %v", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateClusterCmd)
}

func validateCluster(ctx context.Context, cluster *types.Cluster, clusterName string) error {
	_ = "STUB: not implemented"
	return nil
}
