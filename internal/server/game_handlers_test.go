package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gamescoring/internal/db"
	"gamescoring/internal/metrics"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetHttpServer() *HttpServer {
	repo := db.NewMemDBRepository()
	reg := metrics.New()
	server := NewHttpServer(repo, reg)

	return server
}

func TestListGamesHandler(t *testing.T) {
	httpServer := GetHttpServer()
	req, _ := http.NewRequest("GET", "/games", nil)
	rec := httptest.NewRecorder()

	httpServer.ListGames(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}

func TestCreateGameHandler(t *testing.T) {
	httpServer := GetHttpServer()
	game := db.Games[0]
	gameJSON, _ := json.Marshal(game)
	req, _ := http.NewRequest("POST", "/games", bytes.NewBuffer(gameJSON))
	rec := httptest.NewRecorder()

	httpServer.CreateGame(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}

func TestUpdateGameHandler(t *testing.T) {
	httpServer := GetHttpServer()
	game := db.Games[0]
	gameID := game.ID

	gameJSON, _ := json.Marshal(game)
	req, _ := http.NewRequest("PUT", "/games/"+gameID, bytes.NewBuffer(gameJSON))
	rec := httptest.NewRecorder()

	httpServer.UpdateGame(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}

func TestGetGameHandler(t *testing.T) {
	httpServer := GetHttpServer()
	httpServer.Router.Methods("GET").Path("/games/{gameID}").HandlerFunc(httpServer.GetGame)
	game := db.Games[0]
	url := fmt.Sprintf("/games/%s", game.ID)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err, "error not expected")
	rec := httptest.NewRecorder()
	httpServer.Router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}

func TestDeleteGameHandler(t *testing.T) {
	httpServer := GetHttpServer()
	httpServer.Router.Methods("DELETE").Path("/games/{gameID}").HandlerFunc(httpServer.DeleteGame)
	game := db.Games[0]
	gameID := game.ID
	req, _ := http.NewRequest("DELETE", "/games/"+gameID, nil)
	rec := httptest.NewRecorder()

	httpServer.Router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Response status should be OK")
}
