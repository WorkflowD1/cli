package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"workflow-cli/src/model"
)

func Authenticate(baseUrl string, auth model.AuthRequest) string {
	// Convert AuthRequest struct to JSON
	data, err := json.Marshal(auth)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare request variables
	url := baseUrl + "cognito/user/signIn"
	contentType := "application/json"
	body := bytes.NewBuffer(data)

	// HTTP POST Request to AWS Cognito
	resp, err := http.Post(url, contentType, body)
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
