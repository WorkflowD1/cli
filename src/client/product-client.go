package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"workflow-cli/src/model"
)

func LoadAllProducts(baseUrl string) []model.Product {
	url := baseUrl + "product/loadAll"

	resp, err := Post(url, nil)
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

func CreateProduct(baseUrl string, product model.Product) model.Product {
	url := baseUrl + "product/create"

	resp, err := Post(url, product)
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
