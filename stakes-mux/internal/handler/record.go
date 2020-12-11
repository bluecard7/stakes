package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"stakes/internal/data"
	"stakes/internal/user"

	"github.com/google/uuid"
)

// ClockHandler returns handler function
func ClockHandler(table data.RecordTable) http.HandlerFunc {
	return log(authenticate(func(w http.ResponseWriter, req *http.Request) {
		newCtx := data.NewContext(req.Context(), table)
		req.WithContext(newCtx)
		switch req.Method {
		case "GET":
			getRecords(w, req)
		case "POST":
			clock(w, req)
		}
	}))
}

// curl -H 'Content-Type: application/json' 'http://localhost:8000/clock?from=2020-11-22&to=2020-11-23'
func getRecords(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	email, _ := user.FromContext(req.Context())
	table, _ := data.FromContext(req.Context())
	records := table.FindRecordsInTimeFrame(
		email,
		query.Get("from"),
		query.Get("to"),
	)
	respondJSON(w, struct {
		Records []data.Record `json:"records"`
	}{Records: records})
}

// curl -X POST -d '{"email":"my@email.com"}' -H 'Content-Type: application/json' http://localhost:8000/clock
func clock(w http.ResponseWriter, req *http.Request) {
	var record data.Record
	clockedAt := time.Now()
	email, _ := user.FromContext(req.Context())
	table, _ := data.FromContext(req.Context())
	if id := table.FindUnfinishedRecord(email); id == uuid.Nil {
		table.InsertRecord(email, clockedAt)
	} else {
		table.FinishRecord(id, clockedAt)
	}
	respondJSON(w, record)
}

func respondJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "Sorry! Couldn't send my response.", 404)
	} else {
		w.Write(data)
	}
}
