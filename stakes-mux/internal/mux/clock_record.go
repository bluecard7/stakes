package mux

import (
	"encoding/json"
	"net/http"
	"time"

	"stakes/internal/data"

	"github.com/google/uuid"
)

func (s *StakesServer) handleClock(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s.getRecords(w, req)
	case "POST":
		s.clock(w, req)
	}
}

// curl -H 'Content-Type: application/json' 'http://localhost:8000/clock?from=2020-11-22&to=2020-11-23'
// getRecords handles requests for GET /clock.
// getRecords returns the recorded clock in/out times of the user between the
// specified from and to dates.
// Dates are expected to be in yyyy-mm-dd format.
func (s *StakesServer) getRecords(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	email, _ := userIDFromContext(req.Context())
	from, failedFrom := time.Parse("2006-01-02", query.Get("from"))
	to, failedTo := time.Parse("2006-01-02", query.Get("to"))
	if failedFrom != nil || failedTo != nil {
		http.Error(w, "Need to specify from and to dates in yyyy-mm-dd format as query params.", http.StatusBadRequest)
		return
	}
	// how to handle err here?
	records := s.Table.FindRecordsInTimeFrame(email, from, to)
	respondJSON(w, struct {
		Records []data.Record `json:"records"`
	}{Records: records})
}

// curl -X POST -d '{"email":"my@email.com"}' -H 'Content-Type: application/json' http://localhost:8000/clock
// clock handles requests for POST /clock.
// clock() clocks in/out the user in the system.
func (s *StakesServer) clock(w http.ResponseWriter, req *http.Request) {
	var record *data.Record
	clockedAt := time.Now()
	email, _ := userIDFromContext(req.Context())
	if id := s.Table.FindUnfinishedRecord(email); id == uuid.Nil {
		record = s.Table.InsertRecord(email, clockedAt)
	} else {
		record = s.Table.FinishRecord(id, clockedAt)
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
