package mux

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"stakes/internal/data"
	"testing"
	"time"

	"github.com/google/uuid"
)

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

func TestClockRecord(t *testing.T) {
	if *integration {
		t.Skip("Unit tests skipped - running integration tests instead")
	}

	stakesSrv := &StakesServer{
		Table:  MockRecordTable{},
		Router: http.NewServeMux(),
		Logger: log.New(ioutil.Discard, "", 0),
	}
	stakesSrv.MapRoutes()

}

// func TestGetRecords(t *testing.T) {
// 	url := "/clock?from=yyyy-mm-dd&to=yyyy-mm-dd"
// 	request, _ := http.NewRequest("GET", url, nil)
// 	response := *httptest.NewRecorder()
// 	ServeHTTP(&response, request)
// 	// if writer.Code != 200 {}
// 	record := data.Record{}
// 	err := json.Unmarshal(response.Body.Bytes(), &record)
// 	if err != nil {
// 	}
// 	// check err and fields
// }

// func TestClock(t *testing.T) {
// 	request, _ := http.NewRequest("POST", "/clock", nil)
// 	response := *httptest.NewRecorder()
// 	mux.ServeHTTP(&response, request)
// }

func testRequest(t *testing.T) {
	t.Helper()
}
