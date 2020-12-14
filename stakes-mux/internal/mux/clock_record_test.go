package mux

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"stakes/internal/data"
	"testing"
	"time"

	"github.com/google/uuid"
)

// could test handleClock by running requests through it
// targetting the switch cases and expecting some output

type MockRecordTable struct {
}

func (table MockRecordTable) InsertRecord(email string, clockedAt time.Time) *data.Record {
	return nil
}

func (table MockRecordTable) FinishRecord(id uuid.UUID, clockedAt time.Time) *data.Record {
	return nil
}

func (table MockRecordTable) FindUnfinishedRecord(email string) uuid.UUID {
	return uuid.Nil
}

func (table MockRecordTable) FindRecordsInTimeFrame(email string, from, to time.Time) []data.Record {
	return []data.Record{
		{
			ID:       uuid.Nil,
			Email:    "test@email.com",
			ClockIn:  time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
			ClockOut: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}
}

var (
	stakesSrv   *StakesServer
	integration = flag.Bool("integration", false, "runs integration tests instead of unit tests")
)

func Test_getRecords(t *testing.T) {
	if *integration {
		t.Skip("Unit tests skipped - running integration tests instead")
	}

	stakesSrv := &StakesServer{
		Table: MockRecordTable{},
	}

	expectError := func(w *httptest.ResponseRecorder) error {
		if w.Code != http.StatusBadRequest {
			msg := fmt.Sprintf("Expected code to be %d, got %d", http.StatusBadRequest, w.Code)
			return errors.New(msg)
		}
		errMsg := string(w.Body.Bytes())
		expectedMsg := "Need to specify from and to dates in yyyy-mm-dd format as query params.\n"
		if errMsg != expectedMsg {
			msg := fmt.Sprintf("Expected response to be \"%s\", got \"%s\"", expectedMsg, errMsg)
			return errors.New(msg)
		}
		return nil
	}

	expectSuccess := func(w *httptest.ResponseRecorder) error {
		return errors.New("not impl")
	}

	tests := []struct {
		Scenario       string
		URL            string
		VerifyResponse func(*httptest.ResponseRecorder) error
	}{
		{
			Scenario:       "url has no query params",
			URL:            "/clock",
			VerifyResponse: expectError,
		},
		{
			Scenario:       "url missing to query param",
			URL:            "/clock?from=2006-01-02",
			VerifyResponse: expectError,
		},
		{
			Scenario:       "url missing from query param",
			URL:            "/clock?to=2006-01-03",
			VerifyResponse: expectError,
		},
		{
			Scenario:       "url has query params instead of to or from",
			URL:            "/clock?foo=2006-01-02&bar=2006-01-03",
			VerifyResponse: expectError,
		},
		{
			Scenario:       "url uses wrong format for times in query params",
			URL:            "/clock?from=01-02-2016&to=01-03-2016",
			VerifyResponse: expectError,
		},
		{
			Scenario:       "url ignores query params other than from and to",
			URL:            "/clock?from=2006-01-02&foo=2006-01-02&to=2006-01-03",
			VerifyResponse: expectSuccess,
		},
		{
			Scenario:       "request has both from and to",
			URL:            "/clock?from=2006-01-02&to=2006-01-03",
			VerifyResponse: expectSuccess,
		},
	}

	for _, test := range tests {
		t.Run(test.Scenario, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.URL, nil)
			if err != nil {
				t.Fatal("test request could not be created")
			}
			req = req.WithContext(newContextWithUserID(req.Context(), "test@email.com"))
			w := httptest.NewRecorder()
			stakesSrv.handleClock(w, req)

			if err = test.VerifyResponse(w); err != nil {
				t.Error(err)
			}
		})
	}
}

func Test_clock(t *testing.T) {
	tests := []struct {
		Scenario       string
		VerifyResponse func(*httptest.ResponseRecorder) error
	}{
		{
			Scenario:       "",
			VerifyResponse: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Scenario, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/clock", nil)
			if err != nil {
				t.Fatal("test request could not be created")
			}
			req = req.WithContext(newContextWithUserID(req.Context(), "test@email.com"))
			w := httptest.NewRecorder()
			stakesSrv.handleClock(w, req)

			if err = test.VerifyResponse(w); err != nil {
				t.Error(err)
			}
		})
	}
}
