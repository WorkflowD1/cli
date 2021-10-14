package cmd

import (
	"workflow-cli/src/config"

	"github.com/spf13/cobra"
)

var baseUrl, email, password string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setup authentication and instance stuff.",
	// Long:  "Long description... Work in progress...",
	Run: func(cmd *cobra.Command, args []string) {
		config.SetConfigs(baseUrl, email, password)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&baseUrl, "url", "u", "", "Base URL to access workflow API.")
	configCmd.Flags().StringVarP(&email, "email", "e", "", "User e-mail to authenticate.")
	configCmd.Flags().StringVarP(&password, "password", "p", "", "User password to authenticate.")
}
