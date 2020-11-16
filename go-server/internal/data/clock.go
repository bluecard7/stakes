package data

import (
	"stakes/internal/models"
)

// Clock inserts record
// func Clock(accessToken string, dbConn *sql.DB) {
// 	psqlTimeFmt := "2006-01-02 03:04"
// 	incompleteClockRecordQueryFmt := "select clock_interval " +
// 		"FROM clock_records " +
// 		"WHERE clock_interval <@ [%s, %s]" +
// 		"AND upper_inf(clock_interval)"

// 	clockedAt := time.Now()
// 	today, dayAfter := clockedAt.Format("2006-01-02"), clockedAt.AddDate(0, 0, 1).Format("2006-01-02")
// 	incompleteClockRecordQuery := fmt.Sprintf(incompleteClockRecordQueryFmt, today, dayAfter)
// 	// how to use params here?
// 	rows, err := dbConn.Query(incompleteClockRecordQuery)
// 	if rows.Next() {
// 		// set ClockedAt as upper bound
// 	} else {
// 		// if the query returns no rows, its clocking in.
// 		// sqlCmd := fmt.Sprintf("INSERT INTO %s VALUES (%s, '[%s, )', %s)")
// 		// stmt, err := dbConn.Prepare(sqlCmd)
// 	}
// }

func ClockIn() models.Record {
	command := "insert into clock_records (user_email) values ($1) returning id"
	stmt, err := db.Prepare(command)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	record := models.Record{UserEmail: "here@gmail.com"}
	err = stmt.QueryRow(record.UserEmail).Scan(&record.Id)
	if err != nil {
		panic(err)
	}
	return record
}

func GetRecordById(id int) models.Record {
	record := models.Record{}
	command := "select id, user_email from clock_records where id = $1"
	err := db.QueryRow(command, id).Scan(&record.Id, &record.UserEmail)
	if err != nil {
		panic(err)
	}
	return record
}

// GetRecords retrieves all time records in the specified range
// accessToken is gained
func GetRecords(accessToken, dateRange string) {

}
