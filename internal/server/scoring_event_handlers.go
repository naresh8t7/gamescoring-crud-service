package server

import (
	"encoding/json"
	"gamescoring/internal/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *HttpServer) ListScoringEvents(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("ListScoringEvents", time.Now())()
	resp := Response{}
	events, err := h.repository.ListScoringEvents()
	if err != nil {
		resp.Status = "Failed to list scoring events"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully scoring events"
		w.WriteHeader(http.StatusOK)
	}
	j, _ := json.Marshal(&events)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}

func (h *HttpServer) CreateScoringEvent(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("CreateScoringEvent", time.Now())()
	req := model.ScoringEvent{}
	err := json.NewDecoder(r.Body).Decode(&req)
	resp := Response{}
	if err != nil {
		resp.Status = "Invalid request body"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = h.repository.UpsertScoringEvent(&req)
	if err != nil {
		resp.Status = "Failed to Create scoring Event"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully created scoring event"
		w.WriteHeader(http.StatusOK)
	}
	j, _ := json.Marshal(&resp)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}

func (h *HttpServer) UpdateScoringEvent(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("UpdateScoringEvent", time.Now())()
	vars := mux.Vars(r)
	eventID := vars["eventID"]
	req := model.ScoringEvent{}
	err := json.NewDecoder(r.Body).Decode(&req)
	resp := Response{}
	if err != nil {
		resp.Status = "Invalid request body"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.repository.UpsertScoringEvent(&req)
	if err != nil {
		resp.Status = "Failed to update scoring event " + eventID
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully updated scoring event"
		w.WriteHeader(http.StatusOK)
	}
	j, _ := json.Marshal(&resp)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
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
	j, _ := json.Marshal(&scoringEvent)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
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
