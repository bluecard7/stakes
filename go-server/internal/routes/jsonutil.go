package routes

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(res http.ResponseWriter, v interface{}) {
	res.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(v)
	res.Write(output)
}
