package mux

import (
	"log"
	"net/http"
	"stakes/internal/data"
)

// StakesServer holds relevant data structures
type StakesServer struct {
	Table  data.RecordTable
	Router *http.ServeMux
	Logger *log.Logger
}

// MapRoutes links handler functions to routes
func (s *StakesServer) MapRoutes() {
	s.Router.HandleFunc("/clock", s.logRequest(s.authenticate(s.handleClock)))
}
