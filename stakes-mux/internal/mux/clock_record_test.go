package mux

import (
	"encoding/json"
	"flag"
	"io"
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

func (table MockRecordTable) FindRecordsInTimeFrame(email, fromISO, toISO string) []data.Record {
	return nil
}

var (
	stakesSrv   *StakesServer
	integration = flag.Bool("integration", false, "runs integration tests instead of unit tests")
)

func TestGetRecords(t *testing.T) {
	if *integration {
		t.Skip("Unit tests skipped - running integration tests instead")
	}

	stakesSrv := &StakesServer{
		Table: MockRecordTable{},
		// Router: http.NewServeMux(),
		// Logger: log.New(ioutil.Discard, "", 0),
	}
	stakesSrv.MapRoutes()

	t.Run("request has no query params", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := testRequest(t, "POST", "/clock", nil)

		if w.Code != 200 {
		}
		record := data.Record{}
		err := json.Unmarshal(w.Body.Bytes(), &record)
	})

	t.Run("request has missing to query param", func(t *testing.T) {
		url := "/clock?from=yyyy-mm-dd"
		w := httptest.NewRecorder()
		req := testRequest(t, "GET", url, nil)
	})

	t.Run("request has missing from query param", func(t *testing.T) {
		url := "/clock?from=yyyy-mm-dd&to=yyyy-mm-dd"
		w := httptest.NewRecorder()
		req := testRequest(t, "GET", url, nil)
	})

	t.Run("request has query params other than to or from", func(t *testing.T) {
		url := "/clock?from=yyyy-mm-dd&foo=yyyy-mm-dd"
		w := httptest.NewRecorder()
		req := testRequest(t, "GET", url, nil)

		if w.Code != 200 {
		}
		record := data.Record{}
		err := json.Unmarshal(w.Body.Bytes(), &record)
	})

	t.Run("request has both from and to", func(t *testing.T) {
		url := "/clock?from=yyyy-mm-dd&to=yyyy-mm-dd"
		w := httptest.NewRecorder()
		req := testRequest(t, "GET", url, nil)
	})
}

func TestClock(t *testing.T) {
	t.Run("", func(t *testing.T) {
		w := *httptest.NewRecorder()
		req := testRequest(t, "POST", "/clock", nil)

		if w.Code != 200 {
		}
		record := data.Record{}
		err := json.Unmarshal(w.Body.Bytes(), &record)
	})
}

func testRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	t.Helper()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal("Test request couldn't be created")
	}
	return req
}
