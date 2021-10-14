package model

type Configuration struct {
	User     UserConfig
	Instance InstanceConfig
}

type UserConfig struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

type InstanceConfig struct {
	BaseUrl string `yaml:"baseUrl"`
}
