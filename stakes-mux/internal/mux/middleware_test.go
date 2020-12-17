package mux

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

// place async token in properties..
func testJWT(t *testing.T, claims jwt.MapClaims) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		t.Fatal(err)
	}
	return signedToken
}

func Test_authenticate(t *testing.T) {
	os.Setenv("JWT_SECRET", "decoding-secret")
	defer os.Unsetenv("JWT_SECRET")

	stakesSrv := &StakesServer{}
	tests := []struct {
		Scenario string
		Claims   jwt.MapClaims
	}{
		{
			Scenario: "Invalid JWT",
			Claims: jwt.MapClaims{
				"exp":   time.Now().Add(-time.Minute * 15).Unix(),
				"email": "test@email.com",
			},
		},
		{
			Scenario: "Valid JWT",
			Claims: jwt.MapClaims{
				"exp":   time.Now().Add(time.Minute * 15).Unix(),
				"email": "test@email.com",
			},
		},
	}

	nopHandlerFunc := func(w http.ResponseWriter, r *http.Request) {}
	for _, test := range tests {
		t.Run(test.Scenario, func(t *testing.T) {
			req, err := http.NewRequest("", "", nil)
			if err != nil {
				t.Fatal("test request could not be created")
			}
			req.Header.Set("Authorization", testJWT(t, test.Claims))
			w := httptest.NewRecorder()
			stakesSrv.authenticate(nopHandlerFunc)(w, req)

			// check if email gained correctly, if incalidated, etc
		})
	}
}
