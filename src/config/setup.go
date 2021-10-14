package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"workflow-cli/src/model"

	"github.com/spf13/viper"
)

func LoadConfigs() model.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath("./src/config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var config model.Configuration
	viper.Unmarshal(&config)

	return config
}

func SetConfigs(baseUrl string, email string, password string) {
	_, err := ioutil.ReadFile("./src/config/config.yml")
	if err != nil {
		blankYml := "user:\n    email: " + `""` + "\n    password: " + `""` + "\n\ninstance:\n    baseUrl: " + `""`
		blankYmlBytes := []byte(blankYml)
		err := ioutil.WriteFile("./src/config/config.yml", blankYmlBytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration file was created.")
	}

	config := LoadConfigs()

	if baseUrl != "" {
		config.Instance.BaseUrl = baseUrl
	}

	if email != "" {
		config.User.Email = email
	}

	if password != "" {
		config.User.Password = password
	}

	ymlContent := "user:\n    email: " + config.User.Email + "\n    password: " + config.User.Password + "\n\ninstance:\n    baseUrl: " + config.Instance.BaseUrl
	ymlBytes := []byte(ymlContent)

	err = ioutil.WriteFile("./src/config/config.yml", ymlBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your actual configuration is:")
	fmt.Println("URL: ", config.Instance.BaseUrl, "\nEmail: ", config.User.Email, "\nSenha: ", config.User.Password)
}
