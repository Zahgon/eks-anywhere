package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type vSphereSetupUserOptions struct {
	fileName string
	force    bool
	password string
}

var setupUserOptions = &vSphereSetupUserOptions{}

var setupUserCmd = &cobra.Command{
	Use:          "user -f <config-file> [flags]",
	Short:        "Setup vSphere user",
	Long:         "Use eksctl anywhere vsphere setup user to configure EKS Anywhere vSphere user",
	PreRunE:      bindFlagsToViper,
	SilenceUsage: false,
	RunE:         setupUserOptions.setupUser,
}

func init() {
	vsphereSetupCmd.AddCommand(setupUserCmd)

	setupUserCmd.Flags().StringVarP(&setupUserOptions.fileName, "filename", "f", "", "Filename containing vsphere setup configuration")
	setupUserCmd.Flags().StringVarP(&setupUserOptions.password, "password", "p", "", "Password for creating new user")
	setupUserCmd.Flags().BoolVarP(&setupUserOptions.force, "force", "", false, "Force flag. When set, setup user will proceed even if the group and role objects already exist. Mutually exclusive with --password flag, as it expects the user to already exist. default: false")

	if err := setupUserCmd.MarkFlagRequired("filename"); err != nil {
		log.Fatalf("error marking flag as required: %v", err)
	}
}

func (setupUserOptions *vSphereSetupUserOptions) setupUser(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

// when using the force flag we assume the user already exists
