package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // psql driver
)

// TableConfig is a struct type used to configure connection to database
type TableConfig struct {
	Username string
	Password string
	Host     string
	DBName   string
}

// InitRecordTable opens a connection to a PostgreSQL instance
// with configuration loaded in stakes/internal/config.
// Then that connection is returned.
func InitRecordTable(cfg *TableConfig) RecordTableImpl {
	// TODO:: enable sslmode
	connectURL := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.DBName,
	)
	db, err := sql.Open("postgres", connectURL)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	// return cleanup func?
	return RecordTableImpl{db: db}
}
