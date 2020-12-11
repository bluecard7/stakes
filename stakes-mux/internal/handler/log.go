package handler

import (
	stdLog "log"
	"net/http"
)

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		stdLog.Println(req.Method, req.RequestURI)
		h(w, req)
	}
}
