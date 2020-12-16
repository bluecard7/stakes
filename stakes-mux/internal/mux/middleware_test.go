package mux

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

// place async token in properties..
func testJWT(t *testing.T, claims jwt.MapClaims) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		t.Fatal(err)
	}
	return signedToken
}

func Test_authenticate(t *testing.T) {

	stakesSrv := &StakesServer{}

	tests := []struct {
		Claims jwt.MapClaims
	}{}

	nopHandlerFunc := func(w http.ResponseWriter, r *http.Request) {}
	for _, test := range tests {

		req, err := http.NewRequest("", "", nil)
		if err != nil {
			t.Fatal("test request could not be created")
		}
		req.Header.Set("Authorization", testJWT(t, test.Claims))
		w := httptest.NewRecorder()
		stakesSrv.authenticate(nopHandlerFunc)(w, req)

		// check if email gained correctly, if incalidated, etc
	}
}
