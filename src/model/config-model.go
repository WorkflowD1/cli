package model

type Configuration struct {
	User     UserConfig
	Instance InstanceConfig
}

type UserConfig struct {
	Email    string
	Password string
}

type InstanceConfig struct {
	BaseUrl string
}
