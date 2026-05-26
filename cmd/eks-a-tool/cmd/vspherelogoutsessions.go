package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var sessions []string

var vsphereSessionRmCommand = &cobra.Command{
	Use:    "sessions",
	Short:  "vsphere logout sessions command",
	Long:   "This command logs out all of the provided VSphere user sessions ",
	PreRun: prerunCmdBindFlags,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := vsphereLogoutSessions(cmd.Context(), sessions)
		if err != nil {
			log.Fatalf("Error removing sessions: %v", err)
		}
		return nil
	},
}

const (
	sessionTokensFlag      = "sessionTokens"
	tlsInsecureFlag        = "tlsInsecure"
	vsphereApiEndpointFlag = "vsphereApiEndpoint"
)

func init() {
	vsphereRmCmd.AddCommand(vsphereSessionRmCommand)
	vsphereSessionRmCommand.Flags().StringSliceVarP(&sessions, sessionTokensFlag, "s", []string{}, "sessions to logout")
	vsphereSessionRmCommand.Flags().Bool(tlsInsecureFlag, false, "if endpoint is tls secure or not")
	vsphereSessionRmCommand.Flags().StringP(vsphereApiEndpointFlag, "e", "", "the URL of the vsphere API endpoint")

	err := vsphereSessionRmCommand.MarkFlagRequired(vsphereApiEndpointFlag)
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}

	err = vsphereSessionRmCommand.MarkFlagRequired("sessionTokens")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
}

func vsphereLogoutSessions(_ context.Context, sessions []string) error {
	_ = "STUB: not implemented"
	return nil
}

func logoutSession(session string) error { _ = "STUB: not implemented"; return nil }
