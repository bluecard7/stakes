package handler

import "net/http"

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// log something
		h(w, req)
	}
}
