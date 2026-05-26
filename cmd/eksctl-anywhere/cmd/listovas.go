package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/cluster"
)

type listOvasOptions struct {
	fileName        string
	bundlesOverride string
}

type listOvasOutput struct {
	URI    string
	SHA256 string
	SHA512 string
}

var listOvaOpts = &listOvasOptions{}

func init() {
	listCmd.AddCommand(listOvasCmd)
	listOvasCmd.Flags().StringVarP(&listOvaOpts.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	listOvasCmd.Flags().StringVarP(&listOvaOpts.bundlesOverride, "bundles-override", "", "", "Override default Bundles manifest (not recommended)")
	err := listOvasCmd.MarkFlagRequired("filename")
	if err != nil {
		log.Fatalf("Error marking filename flag as required: %v", err)
	}
}

var listOvasCmd = &cobra.Command{
	Use:          "ovas",
	Short:        "List the OVAs that are supported by current version of EKS Anywhere",
	Long:         "This command is used to list the vSphere OVAs from the EKS Anywhere bundle manifest for the current version of the EKS Anywhere CLI",
	PreRunE:      preRunListOvasCmd,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := listOvas(cmd.Context(), listOvaOpts.fileName, listOvaOpts.bundlesOverride); err != nil {
			return err
		}
		return nil
	},
}

func listOvas(context context.Context, clusterSpecPath, bundlesOverride string) error {
	_ = "STUB: not implemented"
	return nil
}

func printOvas(bundle *cluster.VersionsBundle) error { _ = "STUB: not implemented"; return nil }

func preRunListOvasCmd(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func yamlIndent(level int, yamlString string) string { _ = "STUB: not implemented"; return "" }
