package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"stakes/internal/data"

	"github.com/google/uuid"
)

// TODO:: middleware to validate JWT token and extract email

// ClockHandler returns handler function
func ClockHandler(table data.RecordTable) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			getRecords(res, req, table)
		case "POST":
			clock(res, req, table)
		}
	}
}

// curl -H 'Content-Type: application/json' 'http://localhost:8000/clock?from=2020-11-22&to=2020-11-23'
func getRecords(res http.ResponseWriter, req *http.Request, table data.RecordTable) {
	query := req.URL.Query()
	// would email get from JWT in header
	email := "my@email.com"
	from := query["from"][0]
	to := query["to"][0]
	records := table.FindRecordsInTimeFrame(email, from, to)
	// respondWithJSON(res, struct {
	// 	Records []data.Record `json:"records"`
	// }{Records: records})

	res.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(struct {
		Records []data.Record `json:"records"`
	}{Records: records})
	res.Write(output)
}

// curl -X POST -d '{"email":"my@email.com"}' -H 'Content-Type: application/json' http://localhost:8000/clock
func clock(res http.ResponseWriter, req *http.Request, table data.RecordTable) {
	// would email get from JWT in header
	var record data.Record
	clockedAt := time.Now()
	if id := table.FindUnfinishedRecord("my@email.com"); id == uuid.Nil {
		table.InsertRecord("my@email.com", clockedAt)
	} else {
		table.FinishRecord(id, clockedAt)
	}
	res.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(record)
	res.Write(output)
}
