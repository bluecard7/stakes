package routes

import (
	"encoding/json"
	"net/http"
)

// Clock wraps handlers for methods on "/clock" route
func Clock(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		clockIn(res)
	case "PATCH":
		clockOut(res)
	}
}

// ClockResponse models JSON returned in response to rquests on "/clock"
type ClockResponse struct {
	Clocked string `json:"clocked"`
}

func clockIn(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(ClockResponse{Clocked: "IN"})
	res.Write(output)
}

func clockOut(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(ClockResponse{Clocked: "OUT"})
	res.Write(output)
}
