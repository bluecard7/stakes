package internal

import (
	"net/http"
	"net/http/httptest"
	"stakes/internal/config"
	"stakes/internal/data"
	"stakes/internal/mux"
	"testing"

	"github.com/spf13/viper"
)

// I think integrations in context of this package is chaining the handlers and
// executing the endpoints on the test database
func TestIntegrations(t *testing.T) {
	// test malformed urls

	config.ConfigureApp(&config.AppConfig{
		EnvName: "test",
		Dir:     "../config",
	})

	stakesSrv := &mux.StakesServer{
		Table: data.InitRecordTable(&data.TableConfig{
			Username: viper.GetString("psql.username"),
			Password: viper.GetString("psql.password"),
			Host:     viper.GetString("psql.host"),
			DBName:   viper.GetString("psql.dbName"),
		}),
		Router: http.NewServeMux(),
	}
	stakesSrv.MapRoutes()

	tests := []struct {
		Scenario string
		Method   string
		URL      string
		// Want... or golden file?(against golden file since I want to use the actual uuid)
	}{}

	// TODO:: Create test binary and place that in an image from scratch in a docker network.
	// It depends on the postgres test instance, which will be prepopulated programatically before this.
	for _, test := range tests {
		t.Run(test.Scenario, func(t *testing.T) {
			req, err := http.NewRequest(test.Method, test.URL, nil)
			if err != nil {
				t.Fatal("test request could not be created")
			}
			req.Header.Set("Authorization", "JWT token")
			w := httptest.NewRecorder()
			stakesSrv.Router.ServeHTTP(w, req)
		})
	}
}
