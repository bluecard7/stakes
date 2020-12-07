package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"stakes/internal/data"

	"github.com/google/uuid"
)

// Clock wraps handlers for methods on "/clock" route
func Clock(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getRecords(res, req)
	case "POST":
		clock(res, req)
	}
}

// curl -H 'Content-Type: application/json' 'http://localhost:8000/clock?from=2020-11-22T08:00:00.000Z&to=2020-11-22T08:00:00.000Z'
func getRecords(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	email := "my@email.com"
	fromISO := query["from"][0]
	toISO := query["to"][0]
	records := data.FindRecordsInTimeFrame(email, fromISO, toISO)
	respondWithJSON(res, struct {
		Records []data.Record `json:"records"`
	}{Records: records})
}

// curl -X POST -d '{"email":"my@email.com"}' -H 'Content-Type: application/json' http://localhost:8000/clock
func clock(res http.ResponseWriter, req *http.Request) {
	var clockReq struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(req.Body).Decode(&clockReq); err != nil {
		// probably replace with own error struct response
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var clockType string
	if id := data.FindUnfinishedRecord(clockReq.Email); id == uuid.Nil {
		data.InsertRecord(clockReq.Email, time.Now())
		clockType = "IN"
	} else {
		data.FinishRecord(id, time.Now())
		clockType = "OUT"
	}
	respondWithJSON(res, struct {
		Clocked string `json:"clocked"`
	}{Clocked: clockType})
}
