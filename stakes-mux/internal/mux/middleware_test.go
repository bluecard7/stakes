package mux

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

func testJWT(t *testing.T, claims jwt.MapClaims) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		t.Fatal(err)
	}
	return signedToken
}

func verifyResponse(want, got []byte) error {
	if !bytes.Equal(want, got) {
		return fmt.Errorf("Expected %s, got %s", want, got)
	}
	return nil
}

func Test_authenticate(t *testing.T) {
	viper.AutomaticEnv()
	os.Setenv("JWT_SECRET", "decoding-secret")
	defer os.Unsetenv("JWT_SECRET")
	os.Setenv("JWT_ISSUER", "stakes-jwt-issuer-id")
	defer os.Unsetenv("JWT_ISSUER")

	stakesSrv := &StakesServer{}
	tests := []struct {
		Scenario string
		Claims   jwt.MapClaims
		Want     []byte
	}{
		{
			Scenario: "Invalid JWT",
			Claims: jwt.MapClaims{
				"exp":   time.Now().Add(-time.Minute * 15).Unix(),
				"iss":   viper.GetString("JWT_ISSUER"),
				"email": "doesn't matter",
			},
			Want: []byte("Your JWT token is wack\n"),
		},
		{
			Scenario: "Valid JWT",
			Claims: jwt.MapClaims{
				"exp":   time.Now().Add(time.Minute * 15).Unix(),
				"iss":   viper.GetString("JWT_ISSUER"),
				"email": "test@email.com",
			},
			Want: []byte("test@email.com"),
		},
	}

	recordRequest := func(w http.ResponseWriter, req *http.Request) {
		email, ok := userIDFromContext(req.Context())
		if !ok {
			email = ""
		}
		w.Write([]byte(email))
	}
	for _, test := range tests {
		t.Run(test.Scenario, func(t *testing.T) {
			req, err := http.NewRequest("", "", nil)
			if err != nil {
				t.Fatal("test request could not be created")
			}
			req.Header.Set("Authorization", testJWT(t, test.Claims))
			w := httptest.NewRecorder()
			stakesSrv.authenticate(recordRequest)(w, req)

			if err = verifyResponse(test.Want, w.Body.Bytes()); err != nil {
				t.Error(err)
			}
		})
	}
}
