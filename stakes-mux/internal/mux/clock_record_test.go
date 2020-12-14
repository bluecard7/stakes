package mux

import (
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

func (table MockRecordTable) FindRecordsInTimeFrame(email string, from, to time.Time) []data.Record {
	return nil
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
		// Router: http.NewServeMux(),
		// Logger: log.New(ioutil.Discard, "", 0),
	}

	requestAndRespond := func(url string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		req := testRequest(t, "GET", url, nil)
		stakesSrv.handleClock(w, req)
		return w
	}

	expectError := func(w *httptest.ResponseRecorder) {
		if w.Code != http.StatusBadRequest {
			t.Fatalf("Expected code to be %d, got %d", http.StatusBadRequest, w.Code)
		}
		errMsg := string(w.Body.Bytes())
		expectedMsg := "Need to specify from and to dates in yyyy-mm-dd format as query params.\n"
		if errMsg != expectedMsg {
			t.Errorf("Expected response to be \"%s\", got \"%s\"", expectedMsg, errMsg)
		}
	}

	// expectRecords := func(w *httptest.ResponseRecorder) {
	// }

	t.Run("request has no query params", func(t *testing.T) {
		url := "/clock"
		expectError(requestAndRespond(url))
	})

	t.Run("request missing to query param", func(t *testing.T) {
		url := "/clock?from=2006-01-02"
		expectError(requestAndRespond(url))
	})

	t.Run("request missing from query param", func(t *testing.T) {
		url := "/clock?to=2006-01-03"
		expectError(requestAndRespond(url))
	})

	t.Run("request has query params instead of to or from", func(t *testing.T) {
		url := "/clock?foo=2006-01-02&bar=2006-01-03"
		expectError(requestAndRespond(url))
	})

	// t.Run("request ignores query params other than from and to", func(t *testing.T) {
	// 	url := "/clock?from=2006-01-02&foo=2006-01-02&to=2006-01-03"
	// 	w := httptest.NewRecorder()
	// 	req := testRequest(t, "GET", url, nil)
	// 	stakesSrv.handleClock(w, req)
	// })

	// // how to vet from and to are in the yyyy-mm-dd format?

	// t.Run("request has both from and to", func(t *testing.T) {
	// 	url := "/clock?from=2006-01-02&to=2006-01-03"
	// 	w := httptest.NewRecorder()
	// 	req := testRequest(t, "GET", url, nil)
	// 	stakesSrv.handleClock(w, req)
	// })
}

func Test_clock(t *testing.T) {
	// t.Run("", func(t *testing.T) {
	// 	w := *httptest.NewRecorder()
	// 	req := testRequest(t, "POST", "/clock", nil)

	// 	if w.Code != 200 {
	// 	}
	// 	record := data.Record{}
	// 	err := json.Unmarshal(w.Body.Bytes(), &record)
	// })
}

func testRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	t.Helper()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal("Test request couldn't be created")
	}
	return req
}
