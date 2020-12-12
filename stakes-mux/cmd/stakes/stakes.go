package main

import (
	"log"
	"net/http"
	"os"
	"stakes/internal/config"
	"stakes/internal/data"
	"stakes/internal/mux"

	"github.com/spf13/viper"
)

func main() {
	config.ConfigureApp()

	stakesSrv := mux.StakesServer{
		Table: data.InitRecordTable(&data.TableConfig{
			Username: viper.GetString("psql.username"),
			Password: viper.GetString("psql.password"),
			Host:     viper.GetString("psql.host"),
			DBName:   viper.GetString("psql.dbName"),
		}),
		Router: http.NewServeMux(),
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
	stakesSrv.MapRoutes()

	srv := http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: stakesSrv.Router,
	}
	// TODO:: ListenAndServeTLS...?
	log.Println("Server listening...")
	srv.ListenAndServe()
}
