package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/aws/eks-anywhere/pkg/logger"
)

const (
	storageBucketFlagName      = "storage-bucket"
	jobIdFlagName              = "job-id"
	instanceProfileFlagName    = "instance-profile-name"
	regexFlagName              = "regex"
	maxConcurrentTestsFlagName = "max-concurrent-tests"
	skipFlagName               = "skip"
	bundlesOverrideFlagName    = "bundles-override"
	cleanupResourcesFlagName   = "cleanup-resources"
	testReportFolderFlagName   = "test-report-folder"
	branchNameFlagName         = "branch-name"
	instanceConfigFlagName     = "instance-config"
	stageFlagName              = "stage"
)

var runE2ECmd = &cobra.Command{
	Use:          "run",
	Short:        "Run E2E",
	Long:         "Run end to end tests",
	SilenceUsage: true,
	PreRun:       preRunSetup,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runE2E()
		if err != nil {
			logger.Fatal(err, "Failed to run e2e test")
		}
		return nil
	},
}

var requiredFlags = []string{instanceConfigFlagName, storageBucketFlagName, jobIdFlagName, instanceProfileFlagName}

func preRunSetup(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

func init() {
	integrationTestCmd.AddCommand(runE2ECmd)
	runE2ECmd.Flags().StringP(instanceConfigFlagName, "c", "", "File path to the instance-config.yml config")
	runE2ECmd.Flags().StringP(storageBucketFlagName, "s", "", "S3 bucket name to store eks-a binary")
	runE2ECmd.Flags().StringP(jobIdFlagName, "j", "", "Id of the job being run")
	runE2ECmd.Flags().StringP(instanceProfileFlagName, "i", "", "IAM instance profile name to attach to ssm instances")
	runE2ECmd.Flags().StringP(regexFlagName, "r", "", "Run only those tests and examples matching the regular expression. Equivalent to go test -run")
	runE2ECmd.Flags().IntP(maxConcurrentTestsFlagName, "p", 1, "Maximum number of parallel tests that can be run at a time")
	runE2ECmd.Flags().StringSlice(skipFlagName, nil, "List of tests to skip")
	runE2ECmd.Flags().Bool(bundlesOverrideFlagName, false, "Flag to indicate if the tests should run with a bundles override")
	runE2ECmd.Flags().Bool(cleanupResourcesFlagName, false, "Flag to indicate if test resources should be cleaned up automatically as tests complete")
	runE2ECmd.Flags().String(testReportFolderFlagName, "", "Folder destination for JUnit tests reports")
	runE2ECmd.Flags().String(branchNameFlagName, "main", "EKS-A origin branch from where the tests are being run")
	runE2ECmd.Flags().String(stageFlagName, "dev", "Flag to indicate the stage the pipeline from where the tests are being triggered")

	for _, flag := range requiredFlags {
		if err := runE2ECmd.MarkFlagRequired(flag); err != nil {
			log.Fatalf("Error marking flag %s as required: %v", flag, err)
		}
	}
}

func runE2E() error { _ = "STUB: not implemented"; return nil }
