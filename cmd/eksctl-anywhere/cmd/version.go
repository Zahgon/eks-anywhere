package cmd

import (
	"github.com/spf13/cobra"
)

type versionOptions struct {
	output string
}

var vo = &versionOptions{}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the eksctl anywhere version",
	Long:  "This command prints the version of eksctl anywhere",
	RunE: func(cmd *cobra.Command, args []string) error {
		return vo.printVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().StringVarP(&vo.output, "output", "o", "", "specifies the output format (valid option: json)")
}

func (vo *versionOptions) printVersion() error { _ = "STUB: not implemented"; return nil }
