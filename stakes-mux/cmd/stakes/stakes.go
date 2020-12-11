package main

import (
	"net/http"
	"stakes/internal/data"
	"stakes/internal/handler"
)

func main() {
	// Handle configuration with flags
	/*
		viper.SetConfigName("properties")
		viper.SetConfigType("yml")
		viper.AddConfigPath("./configs")

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	*/

	recordTable := data.InitRecordTable(&data.TableConfig{
		User:     "",
		Password: "",
		Host:     "",
		DBName:   "",
	})
	srv := http.Server{
		Addr: ":8000",
	}
	http.HandleFunc("/clock", handler.ClockHandler(recordTable))
	// TODO:: ListenAndServeTLS later on with user auth
	srv.ListenAndServe()
}
