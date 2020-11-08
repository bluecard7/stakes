package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	currDir, _ := os.Getwd()
	projectBase := fmt.Sprintf("%s", currDir)
	fmt.Println("Project base dir", projectBase)
	viper.SetConfigName("properties")
	viper.SetConfigType("yml")
	viper.AddConfigPath(projectBase + "/configs")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("fatal error config file: %s", err))
		} else {
			panic(fmt.Errorf("Another err: %s", err))
		}
	}
}

func Get(key string) interface{} {
	value := viper.Get(key)
	return value
}
