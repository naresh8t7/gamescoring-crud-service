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
	CreateScoringEvent(w http.ResponseWriter, r *http.Request)
	GetScoringEvent(w http.ResponseWriter, r *http.Request)
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
