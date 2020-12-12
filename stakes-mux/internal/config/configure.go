package config

import (
	"log"

	"github.com/spf13/viper"
)

// AppConfig is used to pass options to ConfigureApp.
type AppConfig struct {
	EnvName string
	Dir     string
}

// ConfigureApp attempts to find configuration file and load it in
// based on options passed through cfg.
func ConfigureApp(cfg *AppConfig) {
	viper.SetConfigName("properties." + cfg.EnvName)
	viper.SetConfigType("yml")
	viper.AddConfigPath(cfg.Dir)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to configure service:", err)
	}
}
