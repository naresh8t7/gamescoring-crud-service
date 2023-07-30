package main

import (
	"fmt"
	"gamescoring/internal/config"
	"gamescoring/internal/db"
	"gamescoring/internal/metrics"
	"gamescoring/internal/server"
	"log"
	"net/http"
)

type RepositoryStore string

const (
	MapStore   RepositoryStore = "map"
	MemDBStore RepositoryStore = "memdb"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(" Error loading %v", err)
	}
	var repo db.Repository
	if config.Store == string(MemDBStore) {
		repo = db.NewMemDBRepository()
	} else {
		repo = db.NewRepository()
	}
	reg := metrics.New()
	s := server.NewHttpServer(repo, reg)
	s.AddRoutes()
	log.Println("starting HTTP sever")
	http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort), s.Router)
}
