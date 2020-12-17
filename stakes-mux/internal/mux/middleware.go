package mux

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// http://peter.bourgon.org/go-best-practices-2016/#logging-and-instrumentation
// USE or RED?
func (s *StakesServer) logRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		s.Logger.Println(req.Method, req.RequestURI)
		h(w, req)
	}
}

// Authenticate is a middleware handler that decodes the JWT from the request
// header. If JWT is decoded successfully, then request is considered
// authenticated.
func (s *StakesServer) authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tokenStr := req.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("JWT_SECRET")), nil
		})
		if err == nil {
			// TODO:: verify expireAt, and maybe other attrs like Issuers etc.
			if claims, ok := token.Claims.(jwt.MapClaims); ok { // && claims.VerifyExpiresAt() {
				email := claims["email"].(string)
				newCtx := newContextWithUserID(req.Context(), email)
				h(w, req.WithContext(newCtx))
				return
			}
		}
		http.Error(w, "Your JWT token is wack", 401)
	}
}
