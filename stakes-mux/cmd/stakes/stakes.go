package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"stakes/internal/config"
	"stakes/internal/data"
	"stakes/internal/mux"

	"github.com/spf13/viper"
)

func main() {
	envName := flag.String("e", "local", "environment name")
	cfgDir := flag.String("c", "config", "location of config files")
	flag.Parse()
	config.ConfigureApp(&config.AppConfig{
		EnvName: *envName,
		Dir:     *cfgDir,
	})

	stakesSrv := &mux.StakesServer{
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

	httpSrv := http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: stakesSrv.Router,
	}
	// TODO:: ListenAndServeTLS...?
	stakesSrv.Logger.Println("Server listening...")
	httpSrv.ListenAndServe()
}
