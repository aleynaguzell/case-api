package config

import (
	"case-api/pkg/logger"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Mongo  MongoConfiguration
	Server ServerConfiguration
}

type MongoConfiguration struct {
	Url      string
	Database string
}
type ServerConfiguration struct {
	Port string
}

// Setup initialize configuration
func Setup(path string, env string) error {
	var configuration *Configuration

	viper.SetConfigName("config" + "." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Error("Unable to decode into struct, %v", err)
		return err
	}

	Config = configuration

	return nil
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
