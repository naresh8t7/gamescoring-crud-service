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
	DBStore    RepositoryStore = "db"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(" Error loading %v", err)
	}
	var repo db.Repository
	if config.Store == string(MemDBStore) {
		repo = db.NewMemDBRepository()
	} else if config.Store == string(MemDBStore) {
		repo = db.NewRepository()
	} else {
		props := db.PostgresDbProperties{
			DbHost:     config.DbHost,
			DbPort:     config.DbPort,
			DbUser:     config.DbUser,
			DbPassword: config.DbPassword,
			DbName:     config.DbName,
		}
		repo, err = db.NewDBRepository(props)
		if err != nil {
			log.Fatalf(" Error connecting db %v", err)
		}
	}
	reg := metrics.New()
	s := server.NewHttpServer(repo, reg)
	s.AddRoutes()
	log.Println("starting HTTP sever")
	http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort), s.Router)
}
