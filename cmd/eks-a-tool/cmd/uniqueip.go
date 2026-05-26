package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var uniqueIpCmd = &cobra.Command{
	Use:    "unique-ip",
	Short:  "Unique IP",
	Long:   "Generate a random unique IP to be used for control plane endpoint ip",
	PreRun: preRunUniqueIp,
	RunE: func(cmd *cobra.Command, args []string) error {
		uniqueIp, err := generateUniqueIP(cmd.Context())
		if err != nil {
			log.Fatalf("Error generating unique ip: %v", err)
		}
		fmt.Println(uniqueIp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uniqueIpCmd)
	uniqueIpCmd.Flags().StringP("cidr", "c", "", "CIDR range for the unique IP")
	err := uniqueIpCmd.MarkFlagRequired("cidr")
	if err != nil {
		log.Fatalf("Error marking flag as required: %v", err)
	}
}

func preRunUniqueIp(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

func generateUniqueIP(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
