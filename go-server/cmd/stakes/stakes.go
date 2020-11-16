package main

import (
	"fmt"
	"stakes/internal/data"
)

func main() {

	record := data.ClockIn()
	fmt.Println(record)
	record = data.GetRecordById(record.Id)
	fmt.Println(record)
	data.Close()

	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }

	// http.HandleFunc("/clock", routes.Clock)

	// // TODO:: ListenAndServeTLS later on with user auth
	// server.ListenAndServe()
}
