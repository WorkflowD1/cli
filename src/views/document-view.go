package views

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"workflow-cli/src/client"
	"workflow-cli/src/config"
	"workflow-cli/src/model"
)

func CreateDocumentMenu() model.Document {
	var document model.Document
	var statusIdStr, productIdStr, modalityIdentifier, cpf, name, phone, email string
	// var filledColumns map[string]interface{}

	baseUrl := config.LoadConfigs().Instance.BaseUrl
	reader := bufio.NewReader(os.Stdin)

	products := client.LoadAllProducts(baseUrl)
	ShowProducts(products)

	fmt.Println("Qual o ID do produto que deseja criar o documento?")
	productIdStr, _ = reader.ReadString('\n')
	productIdStr = strings.TrimSuffix(productIdStr, "\n")
	productId, _ := strconv.ParseInt(productIdStr, 36, 64)

	fmt.Println("Qual o ID do status que será atribuído ao documento?")
	statusIdStr, _ = reader.ReadString('\n')
	statusIdStr = strings.TrimSuffix(statusIdStr, "\n")
	statusId, _ := strconv.ParseInt(statusIdStr, 36, 64)

	fmt.Println("Qual a esteira que será atribuída ao documento?")
	modalityIdentifier, _ = reader.ReadString('\n')
	modalityIdentifier = strings.TrimSuffix(modalityIdentifier, "\n")

	fmt.Println("Qual o CPF do cliente?")
	cpf, _ = reader.ReadString('\n')
	cpf = strings.TrimSuffix(cpf, "\n")

	fmt.Println("Qual o nome do cliente?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Println("Qual o telefone do cliente?")
	phone, _ = reader.ReadString('\n')
	phone = strings.TrimSuffix(phone, "\n")

	fmt.Println("Qual o email do cliente?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSuffix(email, "\n")

	document = model.Document{
		ProductID:          productId,
		StatusID:           statusId,
		ModalityIdentifier: modalityIdentifier,
		CPF:                cpf,
		Name:               name,
		Phone:              phone,
		Email:              email,
	}

	return document
}
