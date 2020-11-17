package data

import (
	"database/sql"
	"fmt"
	"time"
)

// A Record represents a time range
type Record struct {
	Id       int
	Email    string
	ClockIn  time.Time
	ClockOut time.Time
}

func (r *Record) String() string {
	return fmt.Sprintf("Record<%d, %s, %s, %s>", r.Id, r.Email, r.ClockIn, r.ClockOut)
}

func InsertRecord(email string, clockedAt time.Time) {
	sqlcmd := `
		insert into clock_records (email, clockIn)
		values ($1, $2)
	`
	stmt, err := db.Prepare(sqlcmd)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, clockedAt)
	if err != nil {
		panic(err)
	}
}

func FindUnfinishedRecord(email string) *Record {
	record := Record{}
	sqlcmd := `
		select email, clockIn, clockOut 
		from clock_records
		where email = $1 and clockOut = NULL
	`
	err := db.QueryRow(sqlcmd, email).Scan(
		&record.Email,
		&record.ClockIn,
		&record.ClockOut,
	)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return &record
}

func UpdateRecord(record *Record) {
	// check if clockIn matches too?
	sqlcmd := `
		update clock_records 
		set clockOut = $2 
		where email = $1
	`
	stmt, err := db.Prepare(sqlcmd)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(record.Email, record.ClockOut)
	if err != nil {
		panic(err)
	}
}

func PrintRecords(email string) {
	record := Record{}
	sqlcmd := "select * from clock_records where email = $1"
	rows, err := db.Query(sqlcmd, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&record.Id, &record.Email, &record.ClockIn, &record.ClockOut)
		if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}

// GetRecords retrieves all time records in the specified range
// accessToken is gained
func GetRecordsInTimeFrame(accessToken, dateRange string) {

}
