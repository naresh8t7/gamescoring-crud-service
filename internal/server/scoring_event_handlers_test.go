package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gamescoring/internal/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListScoringEventHandler(t *testing.T) {
	httpServer := GetHttpServer()
	req, _ := http.NewRequest("GET", "/ScoringEvent", nil)
	rec := httptest.NewRecorder()

	httpServer.ListScoringEvents(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}

func TestCreateScoringEventHandler(t *testing.T) {
	httpServer := GetHttpServer()
	event := db.Events[0]
	ScoringEventJSON, _ := json.Marshal(event)
	req, _ := http.NewRequest("POST", "/scoringEvent", bytes.NewBuffer(ScoringEventJSON))
	rec := httptest.NewRecorder()

	httpServer.CreateScoringEvent(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}

func TestUpdateScoringEventHandler(t *testing.T) {
	httpServer := GetHttpServer()
	event := db.Events[0]
	ScoringEventID := event.ID

	eventJSON, _ := json.Marshal(event)
	req, _ := http.NewRequest("PUT", "/ScoringEvent/"+ScoringEventID, bytes.NewBuffer(eventJSON))
	rec := httptest.NewRecorder()

	httpServer.UpdateScoringEvent(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}

func TestGetScoringEventHandler(t *testing.T) {
	httpServer := GetHttpServer()
	httpServer.Router.Methods("GET").Path("/scoringEvent/{eventID}").HandlerFunc(httpServer.GetScoringEvent)
	event := db.Events[0]
	url := fmt.Sprintf("/scoringEvent/%s", event.ID)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err, "error not expected")
	rec := httptest.NewRecorder()
	httpServer.Router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}

func TestCreateScoringEventHandler_Failure(t *testing.T) {
	httpServer := GetHttpServer()
	event := db.Events[0]
	event.GameID = "dummy-id"
	ScoringEventJSON, _ := json.Marshal(event)
	req, _ := http.NewRequest("POST", "/scoringEvent", bytes.NewBuffer(ScoringEventJSON))
	rec := httptest.NewRecorder()

	httpServer.CreateScoringEvent(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "Response status should not be OK")
}

func TestDeleteScoringEventHandler(t *testing.T) {
	httpServer := GetHttpServer()
	httpServer.Router.Methods("DELETE").Path("/scoringEvent/{eventID}").HandlerFunc(httpServer.DeleteScoringEvent)
	event := db.Events[0]
	eventID := event.ID
	req, _ := http.NewRequest("DELETE", "/scoringEvent/"+eventID, nil)
	rec := httptest.NewRecorder()

	httpServer.Router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}
