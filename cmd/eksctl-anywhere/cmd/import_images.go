package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

// imagesCmd represents the images command.
var importImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Import images and charts to a registry from a tarball",
	Long: `Import all the images and helm charts necessary for EKS Anywhere clusters into a registry.
Use this command in conjunction with download images, passing it output tarball as input to this command.`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		return importImagesCommand.Call(ctx)
	},
}

func init() {
	importCmd.AddCommand(importImagesCmd)

	importImagesCmd.Flags().StringVarP(&importImagesCommand.InputFile, "input", "i", "", "Input tarball containing all images and charts to import")
	if err := importImagesCmd.MarkFlagRequired("input"); err != nil {
		log.Fatalf("Cannot mark 'input' as required: %s", err)
	}
	importImagesCmd.Flags().StringVarP(&importImagesCommand.RegistryEndpoint, "registry", "r", "", "Registry where to import images and charts")
	if err := importImagesCmd.MarkFlagRequired("registry"); err != nil {
		log.Fatalf("Cannot mark 'registry' as required: %s", err)
	}
	importImagesCmd.Flags().StringVarP(&importImagesCommand.BundlesFile, "bundles", "b", "", "Bundles file to read artifact dependencies from")
	if err := importImagesCmd.MarkFlagRequired("bundles"); err != nil {
		log.Fatalf("Cannot mark 'bundles' as required: %s", err)
	}
	importImagesCmd.Flags().BoolVar(&importImagesCommand.includePackages, "include-packages", false, "Flag to indicate inclusion of curated packages in imported images")
	importImagesCmd.Flag("include-packages").Deprecated = "use copy packages command"
	importImagesCmd.Flags().BoolVar(&importImagesCommand.insecure, "insecure", false, "Flag to indicate skipping TLS verification while pushing helm charts and bundles")
}

var importImagesCommand = ImportImagesCommand{}

type ImportImagesCommand struct {
	InputFile        string
	RegistryEndpoint string
	BundlesFile      string
	includePackages  bool
	insecure         bool
}

func (c ImportImagesCommand) Call(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Import the eksa tools image into the registry first, so it can be used immediately
// after to build the helm executable
