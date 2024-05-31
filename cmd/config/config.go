package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string

	DbUser    string `mapstructure:"DB_USER"`
	DbPass    string `mapstructure:"DB_PASS"`
	DbHost    string `mapstructure:"DB_HOST"`
	DbPort    string `mapstructure:"DB_PORT"`
	DbName    string `mapstructure:"DB_NAME"`
	DbSslMode string `mapstructure:"DB_SSLMODE"`
}

func NewConfig() *Config {
	config := loadEnvVariables()

	return config
}

func loadEnvVariables() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
