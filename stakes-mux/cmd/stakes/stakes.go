package main

import (
	"net/http"
	_ "stakes/internal/data"

	"github.com/gorilla/mux"
)

func main() {
	srv := http.Server{
		Addr:    ":8000",
		Handler: mux.NewRouter(),
	}
	// TODO:: ListenAndServeTLS later on with user auth
	srv.ListenAndServe()
}
