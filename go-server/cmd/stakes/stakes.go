package main

import (
	"net/http"
	"stakes/internal/routes"
)

func main() {
	// dbConn := db.Open()
	// db.Close(dbConn)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/clock", routes.Clock)

	// TODO:: ListenAndServeTLS later on with user auth
	server.ListenAndServe()
}
