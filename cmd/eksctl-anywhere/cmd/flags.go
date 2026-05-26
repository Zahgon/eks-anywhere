package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	TinkerbellHardwareCSVFlagName        = "hardware-csv"
	TinkerbellHardwareCSVFlagAlias       = "z"
	TinkerbellHardwareCSVFlagDescription = "Path to a CSV file containing hardware data."
	KubeconfigFile                       = "kubeconfig"

	forceCleanupDeprecationMessageForUpgrade = `The flag --force-cleanup has been removed. For more information on how to troubleshoot existing bootstrap clusters, please refer to the documentation:
https://anywhere.eks.amazonaws.com/docs/troubleshooting/troubleshooting/#cluster-upgrade-fails-with-management-components-on-bootstrap-cluster`
	forceCleanupDeprecationMessageForCreateDelete = `The flag --force-cleanup has been removed. For more information on how to troubleshoot existing bootstrap clusters, please refer to the documentation:
https://anywhere.eks.amazonaws.com/docs/troubleshooting/troubleshooting/#bootstrap-cluster-fails-to-come-up-nodes-already-exist-for-a-cluster-with-the-name`
)

func bindFlagsToViper(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Environment variables can't have dashes in them, so bind them to their equivalent
// keys with underscores, e.g. --hardware-csv to HARDWARE_CSV

// viper.AutomaticEnv() needs help with dashes in flag names.

func applyClusterOptionFlags(flagSet *pflag.FlagSet, clusterOpt *clusterOptions) {
	_ = "STUB: not implemented"
	return
}

func applyTinkerbellHardwareFlag(flagSet *pflag.FlagSet, pathOut *string) {
	_ = "STUB: not implemented"
	return
}

func checkTinkerbellFlags(flagSet *pflag.FlagSet, hardwareCSVPath string, operationType Operation) error {
	_ = "STUB: not implemented"
	return nil
}

// If no flag was returned there is a developer error as the flag has been removed
// from the program rendering it invalid.

// For upgrade and workload cluster create, hardware-csv is an optional flag

func hideForceCleanup(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }
