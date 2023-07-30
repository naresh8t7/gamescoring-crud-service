package server

import (
	"encoding/json"
	"gamescoring/internal/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *HttpServer) ListGames(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("ListGames", time.Now())()
	resp := Response{}
	games, err := h.repository.ListGames()
	if err != nil {
		resp.Status = "Failed to list Games"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully fetched games"
		w.WriteHeader(http.StatusOK)
	}
	j, _ := json.Marshal(&games)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}

func (h *HttpServer) CreateGame(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("CreateGame", time.Now())()
	req := model.Game{}
	err := json.NewDecoder(r.Body).Decode(&req)
	resp := Response{}
	if err != nil {
		resp.Status = "Invalid request body"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.repository.UpsertGame(&req)
	if err != nil {
		resp.Status = "Failed to create game"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully created Game"
		w.WriteHeader(http.StatusOK)
	}
	j, _ := json.Marshal(&resp)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}

func (h *HttpServer) UpdateGame(w http.ResponseWriter, r *http.Request) {
	defer h.statsCollection("UpdateGame", time.Now())()
	vars := mux.Vars(r)
	gameID := vars["gameID"]
	req := model.Game{}
	err := json.NewDecoder(r.Body).Decode(&req)
	resp := Response{}
	if err != nil {
		resp.Status = "Invalid request body"
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.repository.UpsertGame(&req)
	if err != nil {
		resp.Status = "Failed to update game " + gameID
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Status = "Succesfully updated game"
		w.WriteHeader(http.StatusOK)
	}
	j, _ := json.Marshal(&resp)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
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
	j, _ := json.Marshal(&game)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
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
