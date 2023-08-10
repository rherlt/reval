package config

import (
	"github.com/spf13/viper"
)

//see: https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d

type Config struct {
	DataPath                          string
	Name                              string
	Gin_Cors_AllowAllOrigins          bool
	Gin_Cors_AdditionalAllowedHeaders []string
	Gin_Web_Path                      string
	Gin_WebServerAddress              string
	Gin_Api_BaseUrl                   string
	Gin_Web_BaseUrl                   string
}

var Current *Config = nil

func LoadConfig(path string) (err error) {

	viper.AddConfigPath(".")
	viper.AddConfigPath(path)
	viper.SetConfigName("reval")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Current)
	return
}
