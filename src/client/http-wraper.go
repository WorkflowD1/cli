package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"workflow-cli/src/config"
	"workflow-cli/src/model"
)

type Empty struct{}

func Post(url string, body interface{}) (res *http.Response, err error) {
	var bodyBytes io.Reader

	if body == nil {
		bodyBytes = bytes.NewBuffer([]byte(""))
	} else {
		bodyJson, err := json.Marshal(body)
		if err != nil {
			log.Fatal(err)
		}

		bodyBytes = bytes.NewBuffer(bodyJson)
	}

	req, err := http.NewRequest("POST", url, bodyBytes)
	if err != nil {
		log.Fatal(err)
	}

	config := config.LoadConfigs()
	token := Authenticate(config.Instance.BaseUrl, model.AuthRequest(config.User))
	bearerToken := "Bearer " + token

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearerToken)

	httpClient := &http.Client{}

	return httpClient.Do(req)
}
