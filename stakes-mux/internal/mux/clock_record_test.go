package mux

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"stakes/internal/data"
	"testing"
	"time"

	"github.com/google/uuid"
)

// fakeuuidgen tracks and outputs uuids deterministically for test purposes
type fakeuuidgen struct {
	bytes []byte
	pos   int
}

func (f fakeuuidgen) incID() {
	if f.bytes[f.pos] < 127 {
		f.bytes[f.pos]++
		return
	}
	f.bytes[f.pos-1]++
	for pos := f.pos; pos < 16 && f.bytes[pos] == 127; pos++ {
		f.bytes[pos] = 0
	}
	if f.pos--; f.pos < 0 {
		panic("how did you use up the entire UUID space...?")
	}
}

func (f fakeuuidgen) reset() {
	for i := range f.bytes {
		f.bytes[i] = 0
	}
	f.pos = 15
}

func (f fakeuuidgen) uuid() uuid.UUID {
	f.incID()
	return uuid.Must(uuid.FromBytes(f.bytes))
}

var (
	update = flag.Bool("u", false, "update .golden.json files")
	fake   = fakeuuidgen{
		bytes: make([]byte, 16),
		pos:   15,
	}
)

// Implements Record Table interface
type mockRecordTable struct {
	db []data.Record
}

func (table mockRecordTable) InsertRecord(email string, clockedAt time.Time) *data.Record {
	record := data.Record{
		ID:      fake.uuid(),
		Email:   email,
		ClockIn: time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
	}
	table.db = append(table.db, record)
	return &record
}

func (table mockRecordTable) FinishRecord(id uuid.UUID, clockedAt time.Time) *data.Record {
	for _, record := range table.db {
		if record.ID == id && !record.ClockIn.IsZero() && record.ClockOut.IsZero() {
			record.ClockOut = time.Date(2006, 1, 2, 8, 0, 0, 0, time.UTC)
			return &record
		}
	}
	return nil
}

func (table mockRecordTable) FindUnfinishedRecord(email string) uuid.UUID {
	for _, record := range table.db {
		if record.Email == email && !record.ClockIn.IsZero() && record.ClockOut.IsZero() {
			return record.ID
		}
	}
	return uuid.Nil
}

func (table mockRecordTable) FindRecordsInTimeFrame(email string, from, to time.Time) []data.Record {
	records := make([]data.Record, 0, 3)
	for _, record := range table.db {
		if record.Email == email {
			records = append(records, record)
		}
	}
	return records
}

// Unit tests

// verifies recorded response with golden file
// could extend this to errors too.
func verifySuccess(goldenFileName string) func(w *httptest.ResponseRecorder) error {
	return func(w *httptest.ResponseRecorder) error {
		if w.Code != http.StatusOK {
			msg := fmt.Sprintf("Expected code to be %d, got %d.", http.StatusOK, w.Code)
			return errors.New(msg)
		}
		goldenFilePath := "goldenfiles/" + goldenFileName
		got := w.Body.Bytes()
		if *update {
			ioutil.WriteFile(goldenFilePath, got, 0644)
		}
		expected, err := ioutil.ReadFile(goldenFilePath)
		if err != nil {
			return err
		}
		if !bytes.Equal(expected, got) {
			msg := fmt.Sprintf("Expected response to be %s, got %s", expected, got)
			return errors.New(msg)
		}
		return nil
	}
}

func Test_getRecords(t *testing.T) {
	fake.reset()
	stakesSrv := &StakesServer{
		Table: mockRecordTable{
			db: []data.Record{
				{
					ID:       fake.uuid(),
					Email:    "test@email.com",
					ClockIn:  time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
					ClockOut: time.Date(2006, 1, 2, 8, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	verifyError := func(w *httptest.ResponseRecorder) error {
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

	tests := []struct {
		Scenario       string
		URL            string
		VerifyResponse func(*httptest.ResponseRecorder) error
	}{
		{
			Scenario:       "url has no query params",
			URL:            "/clock",
			VerifyResponse: verifyError,
		},
		{
			Scenario:       "url missing to query param",
			URL:            "/clock?from=2006-01-02",
			VerifyResponse: verifyError,
		},
		{
			Scenario:       "url missing from query param",
			URL:            "/clock?to=2006-01-03",
			VerifyResponse: verifyError,
		},
		{
			Scenario:       "url has query params instead of to or from",
			URL:            "/clock?foo=2006-01-02&bar=2006-01-03",
			VerifyResponse: verifyError,
		},
		{
			Scenario:       "url uses wrong format for times in query params",
			URL:            "/clock?from=01-02-2016&to=01-03-2016",
			VerifyResponse: verifyError,
		},
		{
			Scenario:       "url ignores query params other than from and to",
			URL:            "/clock?from=2006-01-02&foo=2006-01-02&to=2006-01-03",
			VerifyResponse: verifySuccess("expectedRecordsOutput.json"),
		},
		{
			Scenario:       "request has both from and to",
			URL:            "/clock?from=2006-01-02&to=2006-01-03",
			VerifyResponse: verifySuccess("expectedRecordsOutput.json"),
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
	fake.reset()
	stakesSrv := &StakesServer{
		Table: mockRecordTable{
			db: []data.Record{
				{
					ID:       fake.uuid(),
					Email:    "test2@email.com",
					ClockIn:  time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
					ClockOut: time.Time{},
				},
			},
		},
	}

	// verifyError := func(w *httptest.ResponseRecorder) error {
	// 	return errors.New("do these come up in unit testing?")
	// }

	tests := []struct {
		Scenario       string
		Email          string
		VerifyResponse func(*httptest.ResponseRecorder) error
	}{
		{
			Scenario:       "clock in",
			Email:          "test1@email.com",
			VerifyResponse: verifySuccess("clockInResponse.json"),
		},
		{
			Scenario:       "clock out",
			Email:          "test2@email.com",
			VerifyResponse: verifySuccess("clockOutResponse.json"),
		},
	}

	for _, test := range tests {
		t.Run(test.Scenario, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/clock", nil)
			if err != nil {
				t.Fatal("test request could not be created")
			}
			req = req.WithContext(newContextWithUserID(req.Context(), test.Email))
			w := httptest.NewRecorder()
			stakesSrv.handleClock(w, req)

			if err = test.VerifyResponse(w); err != nil {
				t.Error(err)
			}
		})
	}
}
