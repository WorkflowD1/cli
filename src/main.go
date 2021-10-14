package main

import (
	"fmt"
	"workflow-cli/src/client"
	"workflow-cli/src/config"
	"workflow-cli/src/model"
)

func main() {
	config := config.LoadConfigs()
	auth := model.AuthRequest{
		Email:    config.User.Email,
		Password: config.User.Password,
	}
	token := client.Authenticate(config.Instance.BaseUrl, auth)
	fmt.Println("Token: ", token)
}
