package views

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"workflow-cli/src/model"
)

func CreateProductMenu() model.Product {
	var product model.Product
	var name, description string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Qual o nome do produto?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Println("E qual a descrição do produto?")
	description, _ = reader.ReadString('\n')
	description = strings.TrimSuffix(description, "\n")

	product = model.Product{
		Name:        name,
		Description: description,
	}

	return product
}
