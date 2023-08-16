package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

//see: https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d

type Config struct {
	Name                              string
	Gin_Cors_AllowAllOrigins          bool
	Gin_Cors_AdditionalAllowedHeaders []string
	Gin_Web_Path                      string
	Gin_WebServerAddress              string
	Gin_Api_BaseUrl                   string
	Gin_Web_BaseUrl                   string
	Data_Import_Glob                  string
	Db_Type                           string
	Db_Sqlite_Connection              string
}

var Current *Config = nil

func loadConfig(path string) (err error) {

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

func Configure() {
	//load config path from command line or use "." (current application path)
	var configPath string = "."
	if len(os.Args) > 1 {
		configPath = os.Args[1:][0]
	}

	err := loadConfig(configPath)
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
}
