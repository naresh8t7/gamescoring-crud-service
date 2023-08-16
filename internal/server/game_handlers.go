package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"gamescoring/internal/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *HttpServer) ListGames(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("ListGames", time.Now())()
	games, err := h.repository.ListGames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	writeResponse(w, games)
}

func (h *HttpServer) CreateGame(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("CreateGame", time.Now())()
	req := model.Game{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateGame(req); err != nil {
		http.Error(w, fmt.Sprintf("Game cannot be created: %v", err), http.StatusBadRequest)
		return
	}
	resp := Response{}
	_, err = h.repository.UpsertGame(&req)
	if err != nil {
		resp.Status = "Failed to create game"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully created Game"
		w.WriteHeader(http.StatusOK)
	}
	writeResponse(w, resp)
}

func (h *HttpServer) UpdateGame(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("UpdateGame", time.Now())()
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	req := model.Game{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validateGame(req); err != nil {
		http.Error(w, fmt.Sprintf("Game cannot be updated: %v", err), http.StatusBadRequest)
		return
	}
	resp := Response{}
	_, err = h.repository.UpsertGame(&req)
	if err != nil {
		resp.Status = "Failed to update game " + gameID
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully updated game"
		w.WriteHeader(http.StatusOK)
	}
	writeResponse(w, resp)
}

func (h *HttpServer) GetGame(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("GetGame", time.Now())()
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	if gameID == "" {
		http.Error(w, "Invalid gameID", http.StatusBadRequest)
		return
	}
	game, err := h.repository.Game(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeResponse(w, game)
}

func (h *HttpServer) DeleteGame(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("DeleteGame", time.Now())()
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	if gameID == "" {
		http.Error(w, "Invalid gameID", http.StatusBadRequest)
		return
	}
	err := h.repository.DeleteGame(gameID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Game with ID " + gameID + " has been deleted"))

}

func validateGame(req model.Game) error {
	var err error
	if req.End.Before(req.Start) {
		err = errors.New("end Time is less than the start time")
	}
	if req.Start.Before(req.Arrive) {
		err = errors.New("start Time is less than the arrive time")
	}
	return err
}
