package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"stakes/internal/data"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func StakesRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/clock").HandlerFunc(ClockHandler)
	return r
}

func ClockHandler(table data.RecordTableInterface) http.HandleFunc {
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
func getRecords(res http.ResponseWriter, req *http.Request, table data.RecordTableInterface) {
	query := req.URL.Query()
	email := "my@email.com"
	fromISO := query["from"][0]
	toISO := query["to"][0]
	records := table.FindRecordsInTimeFrame(email, fromISO, toISO)
	respondWithJSON(res, struct {
		Records []data.Record `json:"records"`
	}{Records: records})
}

// curl -X POST -d '{"email":"my@email.com"}' -H 'Content-Type: application/json' http://localhost:8000/clock
func clock(res http.ResponseWriter, req *http.Request, table data.RecordTableInterface) {
	var clockReq struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(req.Body).Decode(&clockReq); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var clockType string
	if id := table.FindUnfinishedRecord(clockReq.Email); id == uuid.Nil {
		table.InsertRecord(clockReq.Email, time.Now())
		clockType = "IN"
	} else {
		table.FinishRecord(id, time.Now())
		clockType = "OUT"
	}
	respondWithJSON(res, struct {
		Clocked string `json:"clocked"`
	}{Clocked: clockType})
}
