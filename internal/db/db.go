package db

import (
	"gamescoring/internal/model"
	"log"
	"time"

	"github.com/hashicorp/go-memdb"
)

var (
	Games = []*model.Game{
		{
			ID:     "6690cf59-79de-445c-b9f7-04b7f1ee7991",
			Start:  time.Now().Add(3 * time.Hour),
			End:    time.Now().Add(6 * time.Hour),
			Arrive: time.Now().Add(2 * time.Hour),
		},
		{
			ID:     "6690cf59-79de-445c-b9f7-04b7f1ee7992",
			Start:  time.Now().Add(6 * time.Hour),
			End:    time.Now().Add(9 * time.Hour),
			Arrive: time.Now().Add(5 * time.Hour),
		},
		{
			ID:     "6690cf59-79de-445c-b9f7-04b7f1ee7993",
			Start:  time.Now().Add(9 * time.Hour),
			End:    time.Now().Add(12 * time.Hour),
			Arrive: time.Now().Add(8 * time.Hour),
		},
		{
			ID:     "6690cf59-79de-445c-b9f7-04b7f1ee7994",
			Start:  time.Now().Add(12 * time.Hour),
			End:    time.Now().Add(15 * time.Hour),
			Arrive: time.Now().Add(11 * time.Hour),
		},
	}
	Events = []*model.ScoringEvent{
		{
			ID:        "486585db-75f2-467d-a825-b37777c96530",
			GameID:    "6690cf59-79de-445c-b9f7-04b7f1ee7991",
			Timestamp: time.Now().Add(185 * time.Minute),
			Data: model.GameData{
				Code: "pitch",
				Attributes: model.Attributes{
					AdvancesCount: true,
					Result:        "ball_in_play",
				},
			},
		},
		{
			ID:        "486585db-75f2-467d-a825-b37777c96531",
			GameID:    "6690cf59-79de-445c-b9f7-04b7f1ee7992",
			Timestamp: time.Now().Add(365 * time.Minute),
			Data: model.GameData{
				Code: "pitch",
				Attributes: model.Attributes{
					AdvancesCount: true,
					Result:        "ball_in_play",
				},
			},
		},
		{
			ID:        "486585db-75f2-467d-a825-b37777c96532",
			GameID:    "6690cf59-79de-445c-b9f7-04b7f1ee7993",
			Timestamp: time.Now().Add(555 * time.Minute),
			Data: model.GameData{
				Code: "pitch",
				Attributes: model.Attributes{
					AdvancesCount: true,
					Result:        "ball_in_play",
				},
			},
		},
		{
			ID:        "486585db-75f2-467d-a825-b37777c96533",
			GameID:    "6690cf59-79de-445c-b9f7-04b7f1ee7994",
			Timestamp: time.Now().Add(725 * time.Minute),
			Data: model.GameData{
				Code: "pitch",
				Attributes: model.Attributes{
					AdvancesCount: true,
					Result:        "ball_in_play",
				},
			},
		},
	}
)

type Repository interface {
	UpsertGame(game *model.Game) (*model.Game, error)
	Game(id string) (*model.Game, error)
	DeleteGame(id string) error
	ListGames() ([]*model.Game, error)

	UpsertScoringEvent(ge *model.ScoringEvent) (*model.ScoringEvent, error)
	ScoringEvent(id string) (*model.ScoringEvent, error)
	DeleteScoringEvent(id string) error
	ListScoringEvents() ([]*model.ScoringEvent, error)
}

func NewRepository() Repository {
	r := &mapRepsoitory{
		memdb: make(map[string]data),
	}
	r.initData()
	return r
}

func NewMemDBRepository() Repository {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			string(model.GameModel): &memdb.TableSchema{
				Name: string(model.GameModel),
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			string(model.ScoringEventModel): &memdb.TableSchema{
				Name: string(model.ScoringEventModel),
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalf("Unable to create db schema %v", err)
	}
	repo := &memDBRepsository{
		memdb: db,
	}
	repo.initData()
	return repo
}
