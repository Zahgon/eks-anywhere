package cmd

import (
	"github.com/spf13/cobra"
)

const fmTemplate = `---
title: "%s"
linkTitle: "%s"
---

`

var cmdDocPath string

var docgenCmd = &cobra.Command{
	Use:    "docgen",
	Short:  "Generate the documentation for the CLI commands",
	Long:   "Use eks-a-tool docgen to auto generate CLI commands documentation",
	Hidden: true,
	RunE:   docgenCmdRun,
}

func init() {
	docgenCmd.Flags().StringVar(&cmdDocPath, "path", "./docs/content/en/docs/reference/eksctl", "Path to write the generated documentation to")
	rootCmd.AddCommand(docgenCmd)
}

func docgenCmdRun(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func filePrepender(filename string) string { _ = "STUB: not implemented"; return "" }

func linkHandler(name string) string { _ = "STUB: not implemented"; return "" }
