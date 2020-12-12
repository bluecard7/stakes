package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

func ConfigureApp() {
	envName := flag.String("e", "local", "environment name")
	cfgDir := flag.String("c", "config", "location of config files")
	flag.Parse()

	viper.SetConfigName("properties." + *envName)
	viper.SetConfigType("yml")
	viper.AddConfigPath(*cfgDir)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to configure service:", err)
	}
}
