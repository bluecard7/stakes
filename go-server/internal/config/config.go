package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("properties")
	viper.SetConfigType("yml")
	// assumes binary is ran from go-server/
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func Get(key string) interface{} {
	value := viper.Get(key)
	return value
}
