package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // psql driver
)

// InitRecordTable opens a connection to a PostgreSQL instance
// with configuration loaded in stakes/internal/config.
// Then that connection is returned.
func InitRecordTable() RecordTable {
	// connStr := fmt.Sprintf(
	// 	"postgres://%s:%s@%s/%s?sslmode=verify-full",
	// 	config.Get("psql.user"),
	// 	config.Get("psql.password"),
	// 	config.Get("psql.host"),
	// 	config.Get("psql.dbName"),
	// )

	db, err := sql.Open("postgres", "postgres://username:password@localhost/stakes?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	return RecordTable{db: db}
}
