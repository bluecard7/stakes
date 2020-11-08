package datasource

import (
	"database/sql"
	"stakes/internal/config"
	_ "github.com/lib/pq"
)

func PrepDB() {
	db, err := sql.Open("postgres", "user:password@tcp(127.0.0.1:3306)/hello")
}