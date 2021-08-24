package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Revision string
	Branch   string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("gitops-ctl. (branch=%s, revision=%s)", Branch, Revision))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
