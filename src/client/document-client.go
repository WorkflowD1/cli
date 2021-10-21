package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"workflow-cli/src/model"
)

func CreateDocument(baseUrl string, document model.Document) model.DocumentResponse {
	url := baseUrl + "document/create"

	resp, err := Post(url, document)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res model.DocumentResponse
	json.Unmarshal(bodyBytes, &res)

	return res
}
