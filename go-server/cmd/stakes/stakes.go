package main

import (
	"net/http"
	_ "stakes/internal/data"
	"stakes/internal/routes"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/clock", routes.Clock)

	// TODO:: ListenAndServeTLS later on with user auth
	server.ListenAndServe()
}
