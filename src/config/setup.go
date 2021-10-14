package config

import (
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
