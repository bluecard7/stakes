package db

import (
	"database/sql"
	"fmt"
	"stakes/internal/config"

	_ "github.com/lib/pq" // psql driver
)

// Open opens a connection to a PostgreSQL instance
// with configuration loaded in stakes/internal/config.
// Then that connection is returned.
func Open() *sql.DB {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=verify-full",
		config.Get("psql.user"),
		config.Get("psql.password"),
		config.Get("psql.host"),
		config.Get("psql.dbName"),
	)
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		// log? panic?
	}
	return dbConn
}

// Close closes a given sql.DB instance
func Close(dbConn *sql.DB) {
	dbConn.Close()
}
