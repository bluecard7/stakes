package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // psql driver
)

type TableConfig struct {
	User     string
	Password string
	Host     string
	DBName   string
}

// InitRecordTable opens a connection to a PostgreSQL instance
// with configuration loaded in stakes/internal/config.
// Then that connection is returned.
func InitRecordTable(cfg *TableConfig) RecordTableImpl {
	db, err := sql.Open("postgres", "postgres://username:password@localhost/stakes?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	return RecordTableImpl{db: db}
}
