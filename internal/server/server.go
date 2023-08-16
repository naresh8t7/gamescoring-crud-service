package server

import (
	"encoding/json"
	"fmt"
	"gamescoring/internal/db"
	"gamescoring/internal/metrics"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHttpServer(r db.Repository, reg *prometheus.Registry) *HttpServer {
	return &HttpServer{
		Router:     mux.NewRouter(),
		repository: r,
		Registry:   reg,
		apiMetrics: metrics.NewAPIMetrics(),
	}
}

func (h *HttpServer) AddRoutes() {
	// Games Endpoints
	h.Router.Methods("POST").Path("/games").HandlerFunc(h.CreateGame)
	h.Router.Methods("GET").Path("/games").HandlerFunc(h.ListGames)
	h.Router.Methods("GET").Path("/games/{gameID}").HandlerFunc(h.GetGame)
	h.Router.Methods("PUT").Path("/games/{gameID}").HandlerFunc(h.UpdateGame)
	h.Router.Methods("DELETE").Path("/games/{gameID}").HandlerFunc(h.DeleteGame)

	// Scoring Events Endpoints
	h.Router.Methods("POST").Path("/scoringEvents").HandlerFunc(h.CreateScoringEvent)
	h.Router.Methods("GET").Path("/scoringEvents/{eventID}").HandlerFunc(h.GetScoringEvent)
	h.Router.Methods("PUT").Path("/scoringEvents/{eventID}").HandlerFunc(h.UpdateScoringEvent)
	h.Router.Methods("GET").Path("/scoringEvents").HandlerFunc(h.ListScoringEvents)
	h.Router.Methods("DELETE").Path("/scoringEvents/{eventID}").HandlerFunc(h.DeleteScoringEvent)

	// Misc Endpoints about info, health and metrics
	h.Router.Methods("GET").Path("/").HandlerFunc(h.Home)
	h.Router.Methods("GET").Path("/health").HandlerFunc(h.Health)
	h.Registry.MustRegister(h.apiMetrics.ApiRequestCounts)
	h.Registry.MustRegister(h.apiMetrics.APiRequestProcessingTime)
	h.Router.Methods("GET").Path("/metrics").Handler(promhttp.HandlerFor(h.Registry, promhttp.HandlerOpts{}))
}

func (h *HttpServer) statsCollection(path string, startTime time.Time) func() {
	return func() {
		h.apiMetrics.ApiRequestCounts.WithLabelValues(path).Inc()
		h.apiMetrics.APiRequestProcessingTime.WithLabelValues(path).Observe(time.Since(startTime).Seconds())
	}
}

func (h *HttpServer) Home(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("Home", time.Now())()
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<pre>")
	fmt.Fprintf(w, "%s<br><br>", time.Now().Format(time.RFC1123Z))
	fmt.Fprintf(w, "%s<br>", "Game Scoring Service....")
	fmt.Fprintf(w, "</pre>")
}

func (h *HttpServer) Health(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("Health", time.Now())()
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<pre>")
	fmt.Fprintf(w, "%s<br><br>", time.Now().Format(time.RFC1123Z))
	fmt.Fprintf(w, "%s<br>", "Game Scoring Service. Status : OK")
	fmt.Fprintf(w, "</pre>")
}

func writeResponse(w http.ResponseWriter, resp interface{}) {
	j, _ := json.Marshal(&resp)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}
