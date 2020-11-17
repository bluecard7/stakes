package data

import (
	"database/sql"

	_ "github.com/lib/pq" // psql driver
)

var db *sql.DB

// Open opens a connection to a PostgreSQL instance
// with configuration loaded in stakes/internal/config.
// Then that connection is returned.
func init() {
	// connStr := fmt.Sprintf(
	// 	"postgres://%s:%s@%s/%s?sslmode=verify-full",
	// 	config.Get("psql.user"),
	// 	config.Get("psql.password"),
	// 	config.Get("psql.host"),
	// 	config.Get("psql.dbName"),
	// )

	var err error
	// db, err = sql.Open("postgres", "user=username password=password dbname=stakes sslmode=disable")
	db, err = sql.Open("postgres", "postgres://username:password@localhost/stakes?sslmode=disable")
	if err != nil {
		// log? panic?
		panic(err)
	}
}

func Close() {
	db.Close()
}
