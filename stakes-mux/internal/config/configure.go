package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

func ConfigureApp() {
	var (
		envName string
		cfgDir  string
	)
	flag.StringVar(&envName, "e", "local", "environment name")
	flag.StringVar(&cfgDir, "c", "config", "location of config files")
	flag.Parse()

	viper.SetConfigName("properties." + envName)
	viper.SetConfigType("yml")
	viper.AddConfigPath(cfgDir)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to configure service:", err)
	}
}
