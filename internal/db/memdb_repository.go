package db

import (
	"errors"
	"gamescoring/internal/model"
	"log"

	"github.com/hashicorp/go-memdb"
)

type memDBRepsository struct {
	memdb *memdb.MemDB
}

func (db *memDBRepsository) UpsertGame(game *model.Game) (*model.Game, error) {
	if game == nil {
		return game, errors.New("Game expected")
	}

	txn := db.memdb.Txn(true)
	defer txn.Abort()

	if err := txn.Insert(string(model.GameModel), game); err != nil {
		return game, err
	}
	txn.Commit()

	return game, nil
}

func (db *memDBRepsository) Game(id string) (*model.Game, error) {
	txn := db.memdb.Txn(false)
	defer txn.Abort()
	game, err := txn.First(string(model.GameModel), "id", id)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return &model.Game{}, memdb.ErrNotFound
	}
	return game.(*model.Game), nil
}

func (db *memDBRepsository) DeleteGame(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	txn := db.memdb.Txn(true)
	defer txn.Abort()

	err := txn.Delete(string(model.GameModel), &model.Game{ID: id})
	if err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (db *memDBRepsository) ScoringEvent(id string) (*model.ScoringEvent, error) {

	txn := db.memdb.Txn(false)
	defer txn.Abort()
	se, err := txn.First(string(model.ScoringEventModel), "id", id)
	if err != nil {
		return nil, err
	}
	if se == nil {
		return &model.ScoringEvent{}, memdb.ErrNotFound
	}
	return se.(*model.ScoringEvent), nil
}

func (db *memDBRepsository) UpsertScoringEvent(ge *model.ScoringEvent) (*model.ScoringEvent, error) {
	if ge == nil {
		return ge, errors.New("scoring event expected")
	}

	txn := db.memdb.Txn(true)
	defer txn.Abort()

	if err := txn.Insert(string(model.ScoringEventModel), ge); err != nil {
		return ge, err
	}
	txn.Commit()
	return ge, nil
}

func (db *memDBRepsository) DeleteScoringEvent(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	txn := db.memdb.Txn(true)
	defer txn.Abort()

	err := txn.Delete(string(model.ScoringEventModel), &model.ScoringEvent{ID: id})
	if err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (db *memDBRepsository) ListGames() ([]*model.Game, error) {
	// TODO to support pagination
	txn := db.memdb.Txn(false)
	defer txn.Abort()
	rows, err := txn.Get(string(model.GameModel), "id")
	if err != nil {
		return nil, err
	}
	games := make([]*model.Game, 0)
	for game := rows.Next(); game != nil; game = rows.Next() {
		games = append(games, game.(*model.Game))
	}
	return games, nil
}

func (db *memDBRepsository) ListScoringEvents() ([]*model.ScoringEvent, error) {
	// TODO to support pagination
	txn := db.memdb.Txn(false)
	defer txn.Abort()
	rows, err := txn.Get(string(model.ScoringEventModel), "id")
	if err != nil {
		return nil, err
	}
	events := make([]*model.ScoringEvent, 0)
	for event := rows.Next(); event != nil; event = rows.Next() {
		events = append(events, event.(*model.ScoringEvent))
	}
	return events, nil
}

func (db *memDBRepsository) initData() {
	txn := db.memdb.Txn(true)
	defer txn.Abort()

	for _, g := range Games {
		if err := txn.Insert(string(model.GameModel), g); err != nil {
			log.Printf(" Error %v", err)
		}
	}
	for _, e := range Events {
		if err := txn.Insert(string(model.ScoringEventModel), e); err != nil {
			log.Printf(" Error %v", err)
		}
	}
	// Commit the transaction
	txn.Commit()
}
