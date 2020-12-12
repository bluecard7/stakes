package mux

import (
	"encoding/json"
	"net/http"
	"time"

	"stakes/internal/data"
	"stakes/internal/user"

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
func (s *StakesServer) getRecords(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	email, _ := user.FromContext(req.Context())
	from, to := query.Get("from"), query.Get("to")
	if from == "" || to == "" {
		http.Error(w, "Need to specify from and to dates in yyyy-mm-dd", http.StatusBadRequest)
		return
	}
	// how to handle err here?
	records := s.Table.FindRecordsInTimeFrame(email, from, to)
	respondJSON(w, struct {
		Records []data.Record `json:"records"`
	}{Records: records})
}

// curl -X POST -d '{"email":"my@email.com"}' -H 'Content-Type: application/json' http://localhost:8000/clock
func (s *StakesServer) clock(w http.ResponseWriter, req *http.Request) {
	var record *data.Record
	clockedAt := time.Now()
	email, _ := user.FromContext(req.Context())
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
