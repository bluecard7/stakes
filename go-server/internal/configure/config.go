package configure

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
)

const (
	PROJECT_BASE = fmt.Sprintf("%s/../..",os.Getwd())
)

func init() {
	fmt.Println("Project base dir", PROJECT_BASE)
	viper.SetConfigName("properties.yml")
	viper.AddConfigPath(PROJECT_BASE + "/configs")
	
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func Get(key string) string{
	value, ok := viper.Get(key).(string)
	if !ok {
		panic(fmt.Errorf("couldn't get value from config: %s \n", key))
	}
	return value
}