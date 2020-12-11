package handler

import (
	"net/http"
	"stakes/internal/user"

	"github.com/dgrijalva/jwt-go"
)

var expectedClaims = jwt.MapClaims{}

func authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tokenStr := req.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			// get secret from file, env, etc
			return []byte("secret"), nil
		})
		if err != nil {

		}
		// TODO:: verify expireAt, and maybe other attrs like Issuers etc.
		if claims, ok := token.Claims.(jwt.MapClaims); ok { // && claims.VerifyExpiresAt() {
			email := claims["email"].(string)
			newCtx := user.NewContext(req.Context(), email)
			req.WithContext(newCtx)
			h(w, req)
		} else {
			http.Error(w, "Who are you?", 400)
		}
	}
}
