package cmd

import (
	"fmt"
	"os"
	"strings"
	"workflow-cli/src/client"
	"workflow-cli/src/config"
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
				res := client.CreateProduct(config.Instance.BaseUrl, product)
				fmt.Println("Produto criado:\nID: ", res.ID, "\nNome: ", res.Name, "\nDescrição: ", res.Description)
				os.Exit(0)
			}

			fmt.Println("Entidade inválida.")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&entity, "entity", "e", "", "Any available Workflow entity.")
}
