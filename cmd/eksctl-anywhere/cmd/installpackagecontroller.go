package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

type installControllerOptions struct {
	fileName        string
	kubeConfig      string
	bundlesOverride string
}

var ico = &installControllerOptions{}

func init() {
	installCmd.AddCommand(installPackageControllerCommand)
	installPackageControllerCommand.Flags().StringVarP(&ico.fileName, "filename", "f", "", "Filename that contains EKS-A cluster configuration")
	installPackageControllerCommand.Flags().StringVar(&ico.kubeConfig, "kubeConfig", "", "Management cluster kubeconfig file")
	installPackageControllerCommand.Flags().StringVar(&ico.bundlesOverride, "bundles-override", "",
		"Override default Bundles manifest (not recommended)")
	if err := installPackageControllerCommand.MarkFlagRequired("filename"); err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
}

var installPackageControllerCommand = &cobra.Command{
	Use:          "packagecontroller",
	Aliases:      []string{"pc"},
	Short:        "Install packagecontroller on the cluster",
	Long:         "This command is used to Install the packagecontroller on to an existing cluster",
	PreRunE:      preRunPackages,
	SilenceUsage: true,
	RunE:         runInstallPackageController,
}

func runInstallPackageController(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func installPackageController(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
