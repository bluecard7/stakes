package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// A Record represents a time range
type Record struct {
	Id       uuid.UUID
	Email    string
	ClockIn  time.Time
	ClockOut time.Time
}

func (r *Record) String() string {
	return fmt.Sprintf("Record<%d, %s, %s, %s>", r.Id, r.Email, r.ClockIn, r.ClockOut)
}

func FindUnfinishedRecord(email string) (id uuid.UUID) {
	sqlcmd := `
		select id
		from clock_records
		where email = $1 and clockOut = $2
	`
	err := db.QueryRow(sqlcmd, email, time.Time{}).Scan(&id)
	if err == sql.ErrNoRows {
		return uuid.Nil
	} else if err != nil {
		panic(err)
	}
	return
}

func InsertRecord(email string, clockedAt time.Time) {
	sqlcmd := `
		insert into clock_records (id, email, clockIn, clockOut)
		values ($1, $2, $3, $4)
	`
	stmt, err := db.Prepare(sqlcmd)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var id uuid.UUID
	id, err = uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id, email, clockedAt, time.Time{})
	if err != nil {
		panic(err)
	}
}

func FinishRecord(id uuid.UUID, clockedAt time.Time) {
	sqlcmd := `
		update clock_records 
		set clockOut = $2 
		where id = $1
	`
	stmt, err := db.Prepare(sqlcmd)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, clockedAt)
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
