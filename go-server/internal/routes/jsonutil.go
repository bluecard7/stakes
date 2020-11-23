package routes

import (
	"encoding/json"
	"net/http"
)

func decodeRequestBody(req *http.Request, v interface{}) error {
	err := json.NewDecoder(req.Body).Decode(v)
	return err
}

func respondWithJSON(res http.ResponseWriter, v interface{}) {
	res.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(v)
	res.Write(output)
}
