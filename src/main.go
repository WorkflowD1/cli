package main

import (
	"workflow-cli/src/client"
	"workflow-cli/src/model"
)

func main() {
	auth := model.AuthRequest{
		Email:    "example.email@d1.cx",
		Password: "mystrongpassword",
	}
	client.Authenticate(auth)
}
