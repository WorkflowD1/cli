package views

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"workflow-cli/src/model"
)

func ConfigurationMenu() model.Configuration {
	var config model.Configuration
	var baseUrl, email, password string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Qual a URL da instância? (exemplo: http://instancia.api.workflows.d1.cx/)")
	fmt.Println("Atenção: é importante que a URL esteja no formato apresentado no exemplo, incluindo a barra no final.")
	baseUrl, _ = reader.ReadString('\n')
	baseUrl = strings.TrimSuffix(baseUrl, "\n")

	fmt.Println("Qual o seu e-mail?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSuffix(email, "\n")

	fmt.Println("Qual a sua senha?")
	password, _ = reader.ReadString('\n')
	password = strings.TrimSuffix(password, "\n")

	config = model.Configuration{
		User: model.UserConfig{
			Email:    email,
			Password: password,
		},
		Instance: model.InstanceConfig{
			BaseUrl: baseUrl,
		},
	}

	return config
}
