package cmd

import (
	"github.com/spf13/cobra"
)

var codebuildCmd = &cobra.Command{
	Use:   "cb",
	Short: "the root command when working with AWS Codebuild",
}

func init() {
	rootCmd.AddCommand(codebuildCmd)
}
