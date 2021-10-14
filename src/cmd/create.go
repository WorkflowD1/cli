package cmd

import (
	"fmt"
	"strings"
	"workflow-cli/src/client"
	"workflow-cli/src/config"
	"workflow-cli/src/model"
	"workflow-cli/src/views"

	"github.com/spf13/cobra"
)

var entity string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an entity.",
	// Long:  "Long description... Work in progress...",
	Run: func(cmd *cobra.Command, args []string) {
		entity = strings.ToLower(entity)

		config := config.LoadConfigs()

		if entity != "" {
			if entity == "product" {
				product := views.CreateProductMenu()
				res := client.CreateProduct(config.Instance.BaseUrl, model.AuthRequest(config.User), product)
				fmt.Println("Produto criado:\nID: ", res.ID, "\nNome: ", res.Name, "\nDescrição: ", res.Description)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&entity, "entity", "e", "", "Any available Workflow entity.")
}
