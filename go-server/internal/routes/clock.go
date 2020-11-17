package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"stakes/internal/data"
)

// Clock wraps handlers for methods on "/clock" route
func Clock(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		clock(res, req)
	}
}

// ClockRequest models request body expected in requests on "/clock"
type ClockRequest struct {
	Email string `json:"email"`
}

// ClockResponse models JSON returned in response to requests on "/clock"
type ClockResponse struct {
	Clocked string `json:"clocked"`
}

func clock(res http.ResponseWriter, req *http.Request) {
	var clockReq ClockRequest
	err := json.NewDecoder(req.Body).Decode(&clockReq)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var clockType string

	unfinished := data.FindUnfinishedRecord(clockReq.Email)
	if unfinished == nil {
		data.InsertRecord(clockReq.Email, time.Now())
		clockType = "IN"
	} else {
		unfinished.ClockOut = time.Now()
		data.UpdateRecord(unfinished)
		clockType = "OUT"
	}
	// if that doesn't exist, insert new record
	data.PrintRecords(clockReq.Email)

	res.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(ClockResponse{Clocked: clockType})
	res.Write(output)
}
