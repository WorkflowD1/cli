package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "workflow",
	Short: "This is a CLI to help Workflow's API users.",
	// Long: "Long description... work in progress...",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
