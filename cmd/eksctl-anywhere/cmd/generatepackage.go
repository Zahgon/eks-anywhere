package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type generatePackageOptions struct {
	kubeVersion string
	clusterName string
	registry    string
	// kubeConfig is an optional kubeconfig file to use when querying an
	// existing cluster.
	kubeConfig      string
	bundlesOverride string
}

var gpOptions = &generatePackageOptions{}

func init() {
	generateCmd.AddCommand(generatePackageCommand)
	generatePackageCommand.Flags().StringVar(&gpOptions.clusterName, "cluster", "", "Name of cluster for package generation")
	generatePackageCommand.Flags().StringVar(&gpOptions.kubeVersion, "kube-version", "", "Kubernetes Version of the cluster to be used. Format <major>.<minor>")
	generatePackageCommand.Flags().StringVar(&gpOptions.registry, "registry", "", "Used to specify an alternative registry for package generation")
	generatePackageCommand.Flags().StringVar(&gpOptions.kubeConfig, "kubeconfig", "",
		"Path to an optional kubeconfig file to use.")
	generatePackageCommand.Flags().StringVar(&gpOptions.bundlesOverride, "bundles-override", "", "Override default Bundles manifest (not recommended)")
	if err := generatePackageCommand.MarkFlagRequired("cluster"); err != nil {
		log.Fatalf("marking cluster flag as required: %s", err)
	}
}

var generatePackageCommand = &cobra.Command{
	Use:          "package [flags] <package>",
	Aliases:      []string{"packages"},
	Short:        "Generate package(s) configuration",
	Long:         "Generates Kubernetes configuration files for curated packages",
	PreRunE:      preRunPackages,
	SilenceUsage: true,
	RunE:         runGeneratePackages,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err == nil {
			return nil
		}
		return fmt.Errorf("The name of the package to install must be specified as an argument")
	},
}

func runGeneratePackages(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func generatePackages(ctx context.Context, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
