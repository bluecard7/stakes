package db

import (
	"database/sql"
	"fmt"
	"stakes/internal/config"
	"time"

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

// Clock inserts record
func Clock(accessToken string, dbConn *sql.DB) {
	clockedAt := time.Now()
	year, month, date := clockedAt.Date()
	hour, min, _ := clockedAt.Clock()

	// learn how to format ^ to what's below
	// 2010-01-01 15:30

	sqlCmd := fmt.Sprintf("INSERT INTO %s VALUES (%s, '[)', %s)")
	stmt, err := dbConn.Prepare(sqlCmd)
	// dbConn.Query()
}

// GetRecords retrieves all time records in the specified range
// accessToken is gained
func GetRecords(accessToken, dateRange string) {

}
