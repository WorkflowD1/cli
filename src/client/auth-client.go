package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"workflow-cli/src/model"
)

func Authenticate(auth model.AuthRequest) string {
	// Convert AuthRequest struct to JSON
	data, err := json.Marshal(auth)
	if err != nil {
		log.Fatal(err)
	}

	// HTTP POST Request to AWS Cognito
	resp, err := http.Post("https://example.instance.workflows.d1.cx/cognito/user/signIn", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	// Read response body bytes
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the response to AuthRequest struct again
	var res model.AuthResponse
	json.Unmarshal(bodyBytes, &res)

	return res.Token
}
