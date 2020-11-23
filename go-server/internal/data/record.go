package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// A Record consists of an email that unique identifies a user and
// a pair of times that the user clocked in and out at respectively
type Record struct {
	ID       uuid.UUID
	Email    string
	ClockIn  time.Time
	ClockOut time.Time
}

func (r *Record) String() string {
	return fmt.Sprintf("Record<%d, %s, %s, %s>", r.ID, r.Email, r.ClockIn, r.ClockOut)
}

// InsertRecord enters a Record in the database.
// If called, it's assumed to be clocking in, so it will insert a new
// record with the nil equivalent of time.Time for Record.ClockedOut.
func InsertRecord(email string, clockedAt time.Time) {
	queryStr := `
		insert into clock_records (id, email, clockIn, clockOut)
		values ($1, $2, $3, $4)
	`
	stmt := prepSQLStmt(db, queryStr)
	if stmt != nil {
		defer stmt.Close()
	}
	defer stmt.Close()
	id := randomUUID()
	execSQLStmt(stmt, id, email, clockedAt, time.Time{})
}

// FinishRecord updates the clockedOut column of the entry with id.
// If called, it's assumed to be clocking out (hence "finishing" the record).
func FinishRecord(id uuid.UUID, clockedAt time.Time) {
	queryStr := `
		update clock_records 
		set clockOut = $2 
		where id = $1
	`
	stmt := prepSQLStmt(db, queryStr)
	defer stmt.Close()
	execSQLStmt(stmt, id, clockedAt)
}

// FindUnfinishedRecord finds an entry that is "unfinished", or
// have its clockOut column to the nil equivalent for time.Time.
//
// There should only be one such entry, as every clockIn should
// be completed by a clockOut, but if multiple entries match the criteria
// it will return the id of the first row scanned.
func FindUnfinishedRecord(email string) uuid.UUID {
	queryStr := `
		select id
		from clock_records
		where email = $1 and clockOut = $2
	`
	rows := query(db, queryStr, email, time.Time{})
	id := uuid.Nil
	if rows != nil && rows.Next() {
		rows.Scan(&id)
		rows.Close()
	}
	return id
}

// FindRecordsInTimeFrame retrieves all time records in the specified range
// accessToken is gained from an email?
func FindRecordsInTimeFrame(email, accessToken, dateRange string) []*Record {
	queryStr := "select * from clock_records where email = $1"
	records := []*Record{}
	rows := query(db, queryStr, email)
	if rows != nil {
		for rows.Next() {
			record := new(Record)
			err := rows.Scan(&record.ID, &record.Email, &record.ClockIn, &record.ClockOut)
			if err != nil {
				records = append(records, record)
			}
		}
		rows.Close()
	}
	return records
}
