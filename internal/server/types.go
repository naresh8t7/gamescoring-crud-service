package server

import (
	"gamescoring/internal/db"
	"gamescoring/internal/metrics"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

type HttpServer struct {
	repository db.Repository
	Router     *mux.Router
	Registry   *prometheus.Registry
	apiMetrics *metrics.ApiMetrics
}

type IGameScoringService interface {
	CreateGame(w http.ResponseWriter, r *http.Request)
	GetGame(w http.ResponseWriter, r *http.Request)
	ListGames(w http.ResponseWriter, r *http.Request)
	UpdateGame(w http.ResponseWriter, r *http.Request)
	DeleteGame(w http.ResponseWriter, r *http.Request)

	CreateScoringEvent(w http.ResponseWriter, r *http.Request)
	GetScoringEvent(w http.ResponseWriter, r *http.Request)
	ListScoringEvents(w http.ResponseWriter, r *http.Request)
	UpdateScoringEvent(w http.ResponseWriter, r *http.Request)
	DeleteScoringEvent(w http.ResponseWriter, r *http.Request)

	Home(w http.ResponseWriter, r *http.Request)
	Health(w http.ResponseWriter, r *http.Request)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
