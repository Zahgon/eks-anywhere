package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type installPackageOptions struct {
	kubeVersion   string
	clusterName   string
	packageName   string
	registry      string
	customConfigs []string
	// kubeConfig is an optional kubeconfig file to use when querying an
	// existing cluster.
	kubeConfig      string
	bundlesOverride string
}

var ipo = &installPackageOptions{}

func init() {
	installCmd.AddCommand(installPackageCommand)

	installPackageCommand.Flags().StringVar(&ipo.kubeVersion, "kube-version", "",
		"Kubernetes Version of the cluster to be used. Format <major>.<minor>")
	installPackageCommand.Flags().StringVarP(&ipo.packageName, "package-name", "n",
		"", "Custom name of the curated package to install")
	installPackageCommand.Flags().StringVar(&ipo.registry, "registry",
		"", "Used to specify an alternative registry for discovery")
	installPackageCommand.Flags().StringArrayVar(&ipo.customConfigs, "set",
		[]string{}, "Provide custom configurations for curated packages. Format key:value")
	installPackageCommand.Flags().StringVar(&ipo.kubeConfig, "kubeconfig", "",
		"Path to an optional kubeconfig file to use.")
	installPackageCommand.Flags().StringVar(&ipo.clusterName, "cluster", "",
		"Target cluster for installation.")
	installPackageCommand.Flags().StringVar(&ipo.bundlesOverride, "bundles-override", "",
		"Override default Bundles manifest (not recommended)")

	if err := installPackageCommand.MarkFlagRequired("package-name"); err != nil {
		log.Fatalf("marking package-name flag as required: %s", err)
	}
	if err := installPackageCommand.MarkFlagRequired("cluster"); err != nil {
		log.Fatalf("marking cluster flag as required: %s", err)
	}
}

var installPackageCommand = &cobra.Command{
	Use:          "package [flags] package",
	Short:        "Install package",
	Long:         "This command is used to Install a curated package. Use list to discover curated packages",
	PreRunE:      preRunPackages,
	SilenceUsage: true,
	RunE:         runInstallPackages,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err == nil {
			return nil
		}
		return fmt.Errorf("The name of the package to install must be specified as an argument")
	},
	Deprecated: "use `kubectl apply` instead",
}

func runInstallPackages(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func installPackages(ctx context.Context, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
