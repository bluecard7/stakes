package mux

import (
	"log"
	"net/http"
	"os"
	"stakes/internal/config"
	"stakes/internal/data"
	"testing"

	"github.com/spf13/viper"
)

// I think integrations in context of this package is chaining the handlers and
// executing the endpoints on the test database
func TestIntegrations(t *testing.T) {
	if !*integration {
		t.Skip("Integration tests skipped - running unit tests instead")
	}

	// parse flags in main and pass in to this to make configureApp configurable
	config.ConfigureApp()
	stakesSrv := &StakesServer{
		Table: data.InitRecordTable(&data.TableConfig{
			Username: viper.GetString("psql.username"),
			Password: viper.GetString("psql.password"),
			Host:     viper.GetString("psql.host"),
			DBName:   viper.GetString("psql.dbName"),
		}),
		Router: http.NewServeMux(),
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}
