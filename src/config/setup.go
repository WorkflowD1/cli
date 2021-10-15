package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"workflow-cli/src/model"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func LoadConfigs() model.Configuration {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(homedir + "/workflow")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	var config model.Configuration
	viper.Unmarshal(&config)

	return config
}

func SetConfigs(baseUrl string, email string, password string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath := homedir + "/workflow/config.yml"

	_, err = ioutil.ReadFile(configPath)
	if err != nil {
		_ = os.Mkdir(homedir+"/workflow", 0777)

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

		err = ioutil.WriteFile(configPath, data, 0777)
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

	err = ioutil.WriteFile(configPath, ymlData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Your actual configuration is:")
	// fmt.Println("URL: ", config.Instance.BaseUrl, "\nEmail: ", config.User.Email, "\nSenha: ", config.User.Password)
	fmt.Println(string(ymlData))
}

func ShowConfigs() {
	config := LoadConfigs()
	yml, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(yml))
}
