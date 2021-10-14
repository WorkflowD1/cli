package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"workflow-cli/src/model"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
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
		blankYml := model.Configuration{
			User: model.UserConfig{
				Email:    "",
				Password: "",
			},
			Instance: model.InstanceConfig{
				BaseUrl: "",
			},
		}
		data, err := yaml.Marshal(blankYml)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("./src/config/config.yml", data, 0644)
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

	ymlData, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./src/config/config.yml", ymlData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your actual configuration is:")
	// fmt.Println("URL: ", config.Instance.BaseUrl, "\nEmail: ", config.User.Email, "\nSenha: ", config.User.Password)
	fmt.Println(string(ymlData))
}
