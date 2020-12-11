package main

import (
	"net/http"
	"stakes/internal/config"
	"stakes/internal/data"
	"stakes/internal/handler"

	"github.com/spf13/viper"
)

func main() {
	config.ConfigureApp()
	recordTable := data.InitRecordTable(&data.TableConfig{
		Username: viper.GetString("psql.username"),
		Password: viper.GetString("psql.password"),
		Host:     viper.GetString("psql.host"),
		DBName:   viper.GetString("psql.dbName"),
	})
	srv := http.Server{
		Addr: viper.GetString("server.addr"),
	}
	http.HandleFunc("/clock", handler.ClockHandler(recordTable))
	// TODO:: ListenAndServeTLS...?
	srv.ListenAndServe()
}
