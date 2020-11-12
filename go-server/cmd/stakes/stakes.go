package main

import (
	"stakes/internal/db"
)

func main() {
	dbConn := db.Open()

	db.Close(dbConn)
}
