package cmd

import (
	// "errors"
	"os"
	"path/filepath"

	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "the root command to config",
	Run: func(cmd *cobra.Command, args []string) {

		prompt := promptui.Prompt{
			Label: "VAULT ADDR",
		}
		vaultAddr, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		prompt = promptui.Prompt{
			Label: "GitHub token",
		}
		githubToken, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		homeDir, _ := os.UserHomeDir()
		viper.SetConfigFile(filepath.Join(homeDir, ".config/gitops-ctl.yaml"))
		viper.SetConfigType("yaml")
		viper.Set("vault_addr", vaultAddr)
		viper.Set("github_token", githubToken)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Printf("write config failed %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
