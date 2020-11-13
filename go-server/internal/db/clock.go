package db

import (
	"database/sql"
	"fmt"
	"time"
)

// Clock inserts record
func Clock(accessToken string, dbConn *sql.DB) {
	psqlTimeFmt := "2006-01-02 03:04"
	incompleteClockRecordQueryFmt := "SELECT clock_interval " +
		"FROM clock_records " +
		"WHERE clock_interval <@ [%s, %s]" +
		"AND upper_inf(clock_interval)"

	clockedAt := time.Now()
	today, dayAfter := clockedAt.Format("2006-01-02"), clockedAt.AddDate(0, 0, 1).Format("2006-01-02")
	incompleteClockRecordQuery := fmt.Sprintf(incompleteClockRecordQueryFmt, today, dayAfter)
	// how to use params here?
	rows, err := dbConn.Query(incompleteClockRecordQuery)
	if rows.Next() {
		// set ClockedAt as upper bound
	} else {
		// if the query returns no rows, its clocking in.
		// sqlCmd := fmt.Sprintf("INSERT INTO %s VALUES (%s, '[%s, )', %s)")
		// stmt, err := dbConn.Prepare(sqlCmd)
	}
}

// GetRecords retrieves all time records in the specified range
// accessToken is gained
func GetRecords(accessToken, dateRange string) {

}
