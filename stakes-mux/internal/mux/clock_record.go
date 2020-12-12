package mux

import (
	"encoding/json"
	"net/http"
	"time"

	"stakes/internal/data"
	"stakes/internal/user"

	"github.com/google/uuid"
)

func (s *StakesServer) handleClock() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			getRecords(w, req, s.Table)
		case "POST":
			clock(w, req, s.Table)
		}
	}
}

// curl -H 'Content-Type: application/json' 'http://localhost:8000/clock?from=2020-11-22&to=2020-11-23'
func getRecords(w http.ResponseWriter, req *http.Request, table data.RecordTable) {
	query := req.URL.Query()
	email, _ := user.FromContext(req.Context())
	records := table.FindRecordsInTimeFrame(email, query.Get("from"), query.Get("to"))
	respondJSON(w, struct {
		Records []data.Record `json:"records"`
	}{Records: records})
}

// curl -X POST -d '{"email":"my@email.com"}' -H 'Content-Type: application/json' http://localhost:8000/clock
func clock(w http.ResponseWriter, req *http.Request, table data.RecordTable) {
	var record *data.Record
	clockedAt := time.Now()
	email, _ := user.FromContext(req.Context())
	if id := table.FindUnfinishedRecord(email); id == uuid.Nil {
		record = table.InsertRecord(email, clockedAt)
	} else {
		record = table.FinishRecord(id, clockedAt)
	}
	if record == nil {
		http.Error(w, "Sorry! Something failed on our end when you clocked in/out.", 404)
		return
	}
	respondJSON(w, *record)
}

func respondJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "Sorry! Couldn't send back the response.", 404)
	} else {
		w.Write(data)
	}
}
