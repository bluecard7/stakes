package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
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
	mux      *http.ServeMux
	response httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	mux = http.NewServeMux()
	response = *httptest.NewRecorder()
	mux.HandleFunc("/clock", ClockHandler(MockRecordTable{}))
	os.Exit(m.Run())
}

func TestGetRecords(t *testing.T) {
	url := "/clock?from=yyyy-mm-dd&to=yyyy-mm-dd"
	request, _ := http.NewRequest("GET", url, nil)
	mux.ServeHTTP(&response, request)
	// if writer.Code != 200 {}
	record := data.Record{}
	err := json.Unmarshal(response.Body.Bytes(), &record)
	if err != nil {
	}
	// check err and fields
}

func TestClock(t *testing.T) {
	request, _ := http.NewRequest("POST", "/clock", nil)
	mux.ServeHTTP(&response, request)
}
