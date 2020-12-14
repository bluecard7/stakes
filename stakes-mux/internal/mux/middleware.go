package mux

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// http://peter.bourgon.org/go-best-practices-2016/#logging-and-instrumentation
// According to this, should explicitly use a logger (I'll put it in ctx again)
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
			// get secret from file, env, etc
			return []byte("secret"), nil
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
		http.Error(w, "Who are you?", 401)
	}
}
