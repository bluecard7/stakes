package data

import (
	"database/sql"
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

// RecordTableInterface defines operations used to interact with Record table
// for routes
type RecordTableInterface interface {
	InsertRecord(email string, clockedAt time.Time)
	FinishRecord(id uuid.UUID, clockedAt time.Time)
	FindUnfinishedRecord(email string) uuid.UUID
	FindRecordsInTimeFrame(email, fromISO, toISO string) []Record
}

// RecordTable implements RecordTableInterface
type RecordTable struct {
	db *sql.DB
}

// InsertRecord enters a Record in the database.
// If called, it's assumed to be clocking in, so it will insert a new
// record with the nil equivalent of time.Time for Record.ClockedOut.
func (table RecordTable) InsertRecord(email string, clockedAt time.Time) {
	queryStr := `
		insert into clock_records (id, email, clockIn, clockOut)
		values ($1, $2, $3, $4)
	`
	stmt, err := table.db.Prepare(queryStr)
	if err == nil {
		defer stmt.Close()
	}
	id, err := uuid.NewRandom()
	if err == nil {
		stmt.Exec(id, email, clockedAt, time.Time{})
	}
	// success indicator?
}

// FinishRecord updates the clockedOut column of the entry with id.
// If called, it's assumed to be clocking out (hence "finishing" the record).
func (table RecordTable) FinishRecord(id uuid.UUID, clockedAt time.Time) {
	queryStr := `
		update clock_records 
		set clockOut = $2 
		where id = $1
	`
	stmt, err := table.db.Prepare(queryStr)
	if err == nil {
		defer stmt.Close()
		stmt.Exec(id, clockedAt)
	}
}

// FindUnfinishedRecord finds an entry that is "unfinished", or
// have its clockOut column to the nil equivalent for time.Time.
//
// There should only be one such entry, as every clockIn should
// be completed by a clockOut, but if multiple entries match the criteria
// it will return the id of the first row scanned.
func (table RecordTable) FindUnfinishedRecord(email string) uuid.UUID {
	queryStr := `
		select id
		from clock_records
		where email = $1 and clockOut = $2
	`
	rows, err := table.db.Query(queryStr, email, time.Time{})
	id := uuid.Nil
	if err == nil && rows != nil && rows.Next() {
		rows.Scan(&id)
		rows.Close()
	}
	return id
}

// FindRecordsInTimeFrame retrieves all time records whose
// whose clockIn is in [fromISO, toISO].
// fromISO and toISO are ISO8601 strings that represent dates.
// toISO needs to be one day after the intended end of range, or it will not be included.
func (table RecordTable) FindRecordsInTimeFrame(email, fromISO, toISO string) []Record { // map result from * to literals?
	queryStr := `
		select * from clock_records 
		where email = $1 and clockIn <@ tsrange($2, $3, '[]')
	`
	fromTime, _ := time.Parse("2006-01-02", fromISO)
	toTime, _ := time.Parse("2006-01-02", toISO)
	records := []Record{}
	rows, err := table.db.Query(queryStr, email, fromTime, toTime)
	if err == nil && rows != nil {
		for rows.Next() {
			record := Record{}
			err := rows.Scan(&record.ID, &record.Email, &record.ClockIn, &record.ClockOut)
			if err == nil {
				records = append(records, record)
			}
		}
		rows.Close()
	}
	return records
}
