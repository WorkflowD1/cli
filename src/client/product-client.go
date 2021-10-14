package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"workflow-cli/src/model"
)

func LoadAllProducts(baseUrl string, auth model.AuthRequest) []model.Product {
	url := baseUrl + "product/loadAll"
	contentType := "application/json"
	body := bytes.NewBuffer([]byte(""))

	bearerToken := "Bearer " + Authenticate(baseUrl, auth)

	httpclient := &http.Client{}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", bearerToken)

	resp, err := httpclient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res []model.Product
	json.Unmarshal(bodyBytes, &res)

	for _, product := range res {
		fmt.Println("#", product.ID, " - ", product.Name)
	}

	return res
}

func CreateProduct(baseUrl string, auth model.AuthRequest, product model.Product) model.Product {
	url := baseUrl + "product/create"
	contentType := "application/json"

	data, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewBuffer(data)

	bearerToken := "Bearer " + Authenticate(baseUrl, auth)

	httpclient := &http.Client{}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", bearerToken)

	resp, err := httpclient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res model.Product
	json.Unmarshal(bodyBytes, &res)

	return res
}
