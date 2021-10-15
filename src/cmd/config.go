package cmd

import (
	"os"
	"workflow-cli/src/config"
	"workflow-cli/src/views"

	"github.com/spf13/cobra"
)

var baseUrl, email, password string
var show bool

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setup authentication and instance stuff.",
	// Long:  "Long description... Work in progress...",
	Run: func(cmd *cobra.Command, args []string) {
		if show {
			config.ShowConfigs()
			os.Exit(0)
		}

		if baseUrl == "" && email == "" && password == "" {
			userConfig := views.ConfigurationMenu()
			config.SetConfigs(userConfig.Instance.BaseUrl, userConfig.User.Email, userConfig.User.Password)
		} else {
			config.SetConfigs(baseUrl, email, password)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&baseUrl, "url", "u", "", "Base URL to access workflow API.")
	configCmd.Flags().StringVarP(&email, "email", "e", "", "User e-mail to authenticate.")
	configCmd.Flags().StringVarP(&password, "password", "p", "", "User password to authenticate.")
	configCmd.Flags().BoolVarP(&show, "show", "s", false, "Show current configuration.")
}
