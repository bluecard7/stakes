// +build integration

package mux

import (
	"stakes/internal/config"
	"testing"
)

// I think integrations in context of this package is chaining the handlers and
// executing the endpoints on the test database
func TestIntegrations(t *testing.T) {
	// if !*integration {
	// 	t.Skip("Integration tests skipped - running unit tests instead")
	// }
	// test malformed urls

	config.ConfigureApp(&config.AppConfig{
		EnvName: "test",
		Dir:     "../../config",
	})
	// stakesSrv := &StakesServer{
	// 	Table: data.InitRecordTable(&data.TableConfig{
	// 		Username: viper.GetString("psql.username"),
	// 		Password: viper.GetString("psql.password"),
	// 		Host:     viper.GetString("psql.host"),
	// 		DBName:   viper.GetString("psql.dbName"),
	// 	}),
	// 	Router: http.NewServeMux(),
	// 	Logger: log.New(os.Stdout, "", log.LstdFlags),
	// }

	t.Run("Should not run", func(t *testing.T) {
		t.FailNow()
	})
}
