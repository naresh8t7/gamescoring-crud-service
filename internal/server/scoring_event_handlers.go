package server

import (
	"encoding/json"
	"fmt"
	"gamescoring/internal/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *HttpServer) ListScoringEvents(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("ListScoringEvents", time.Now())()
	events, err := h.repository.ListScoringEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	writeResponse(w, events)
}

func (h *HttpServer) CreateScoringEvent(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("CreateScoringEvent", time.Now())()
	req := model.ScoringEvent{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validateScoringEvent(req); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create scoring event, Error in fetching game: %v", err), http.StatusBadRequest)
		return
	}
	resp := Response{}
	_, err = h.repository.UpsertScoringEvent(&req)
	if err != nil {
		resp.Status = "Failed to Create scoring Event"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully created scoring event"
		w.WriteHeader(http.StatusOK)
	}
	writeResponse(w, resp)
}

func (h *HttpServer) UpdateScoringEvent(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("UpdateScoringEvent", time.Now())()
	vars := mux.Vars(r)
	eventID := vars["eventID"]
	req := model.ScoringEvent{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validateScoringEvent(req); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update scoring event, Error in fetching game: %v", err), http.StatusBadRequest)
		return
	}
	resp := Response{}
	_, err = h.repository.UpsertScoringEvent(&req)
	if err != nil {
		resp.Status = "Failed to update scoring event " + eventID
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully updated scoring event"
		w.WriteHeader(http.StatusOK)
	}
	writeResponse(w, resp)
}

func (h *HttpServer) GetScoringEvent(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("GetScoringEvent", time.Now())()
	vars := mux.Vars(r)
	eventID := vars["eventID"]
	if eventID == "" {
		http.Error(w, "Invalid eventID", http.StatusBadRequest)
		return
	}
	scoringEvent, err := h.repository.ScoringEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeResponse(w, scoringEvent)
}

func (h *HttpServer) DeleteScoringEvent(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("DeleteScoringEvent", time.Now())()
	vars := mux.Vars(r)
	eventID := vars["eventID"]
	if eventID == "" {
		http.Error(w, "Invalid eventID", http.StatusBadRequest)
		return
	}
	err := h.repository.DeleteScoringEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Scoring event with ID " + eventID + " has been deleted"))
}

func (h *HttpServer) validateScoringEvent(req model.ScoringEvent) error {
	_, err := h.repository.Game(req.GameID)
	return err
}
