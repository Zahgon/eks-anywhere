package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

// imagesCmd represents the images command.
var downloadImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Download all eks-a images to disk",
	Long: `Creates a tarball containing all necessary images
to create an eks-a cluster for any of the supported
Kubernetes versions.`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		return downloadImagesRunner.Run(ctx)
	},
}

func init() {
	downloadCmd.AddCommand(downloadImagesCmd)

	downloadImagesCmd.Flags().StringVarP(&downloadImagesRunner.outputFile, "output", "o", "", "Output tarball containing all downloaded images")
	if err := downloadImagesCmd.MarkFlagRequired("output"); err != nil {
		log.Fatalf("Cannot mark 'output' flag as required: %s", err)
	}

	downloadImagesCmd.Flags().BoolVar(&downloadImagesRunner.includePackages, "include-packages", false, "this flag no longer works, use copy packages instead")
	downloadImagesCmd.Flag("include-packages").Deprecated = "use copy packages command"
	downloadImagesCmd.Flags().StringVarP(&downloadImagesRunner.bundlesOverride, "bundles-override", "", "", "Override default Bundles manifest (not recommended)")
	downloadImagesCmd.Flags().BoolVar(&downloadImagesRunner.insecure, "insecure", false, "Flag to indicate skipping TLS verification while downloading helm charts")
}

var downloadImagesRunner = downloadImagesCommand{}

type downloadImagesCommand struct {
	outputFile      string
	bundlesOverride string
	includePackages bool
	insecure        bool
}

func (c downloadImagesCommand) Run(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type packager interface {
	UnPackage(orgFile, dstFolder string) error
	Package(sourceFolder, dstFile string) error
}

func packagerForFile(file string) packager { _ = "STUB: not implemented"; return *new(packager) }
