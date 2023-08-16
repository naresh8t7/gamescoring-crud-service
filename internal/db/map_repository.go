package db

import (
	"errors"
	"gamescoring/internal/model"
	"sync"
)

type data map[string]interface{}

type mapRepsoitory struct {
	sync.RWMutex // this is not needed if we are using actual db
	memdb        map[string]data
}

func (repository *mapRepsoitory) GetStrikeoutsCountPerGame(gameID string) (int, error) {
	return 0, nil
}

func (r *mapRepsoitory) UpsertGame(game *model.Game) (*model.Game, error) {
	if game == nil {
		return game, errors.New("Game expected")
	}
	r.Lock()
	defer r.Unlock()
	if r.memdb[string(model.GameModel)] == nil {
		r.memdb[string(model.GameModel)] = make(data)
	}

	// Assumption: The Game ID will be unique for each particular Game. If we are using database it would throw error.
	// here we just upsert, this will override existing game.
	r.memdb[string(model.GameModel)][game.ID] = game
	return game, nil
}

func (r *mapRepsoitory) Game(id string) (*model.Game, error) {
	if game, ok := r.memdb[string(model.GameModel)][id]; ok {
		return game.(*model.Game), nil
	}
	return nil, errors.New("Game not found")
}

func (r *mapRepsoitory) DeleteGame(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	delete(r.memdb[string(model.GameModel)], id)
	return nil
}

func (r *mapRepsoitory) ScoringEvent(id string) (*model.ScoringEvent, error) {
	if se, ok := r.memdb[string(model.ScoringEventModel)][id]; ok {
		return se.(*model.ScoringEvent), nil
	}
	return nil, errors.New("scoring event not found")
}

func (r *mapRepsoitory) UpsertScoringEvent(ge *model.ScoringEvent) (*model.ScoringEvent, error) {
	if ge == nil {
		return ge, errors.New("scoring event expected")
	}
	r.Lock()
	defer r.Unlock()
	if r.memdb[string(model.ScoringEventModel)] == nil {
		r.memdb[string(model.ScoringEventModel)] = make(data)
	}
	// Assumption: The Game ID will be unique for each scoring event. If we are using database it would throw error.
	// here we just upsert, this will override existing scoring event.
	r.memdb[string(model.ScoringEventModel)][ge.ID] = ge
	return ge, nil
}

func (r *mapRepsoitory) DeleteScoringEvent(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	delete(r.memdb[string(model.ScoringEventModel)], id)
	return nil
}

func (r *mapRepsoitory) ListGames() ([]*model.Game, error) {
	// TODO to support pagination
	games := make([]*model.Game, 0)
	for _, game := range r.memdb[string(model.GameModel)] {
		games = append(games, game.(*model.Game))
	}
	return games, nil
}

func (r *mapRepsoitory) ListScoringEvents() ([]*model.ScoringEvent, error) {
	// TODO to support pagination
	events := make([]*model.ScoringEvent, 0)
	for _, event := range r.memdb[string(model.ScoringEventModel)] {
		events = append(events, event.(*model.ScoringEvent))
	}
	return events, nil
}

func (r *mapRepsoitory) initData() {
	r.Lock()
	defer r.Unlock()
	if r.memdb[string(model.GameModel)] == nil {
		r.memdb[string(model.GameModel)] = make(data)
	}
	if r.memdb[string(model.ScoringEventModel)] == nil {
		r.memdb[string(model.ScoringEventModel)] = make(data)
	}
	for _, g := range Games {
		r.memdb[string(model.GameModel)][g.ID] = g
	}

	for _, e := range Events {
		r.memdb[string(model.ScoringEventModel)][e.ID] = e
	}
}
