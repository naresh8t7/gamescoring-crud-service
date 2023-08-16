package db

import (
	"errors"
	"fmt"
	"gamescoring/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type PostgresDbProperties struct {
	DbHost     string
	DbName     string
	DbPort     string
	DbUser     string
	DbPassword string
}

type dbRepository struct {
	db *gorm.DB
}

func NewDBRepository(properties PostgresDbProperties) (Repository, error) {
	err := validatePostgresDbProperties(properties)
	if err != nil {
		return nil, err
	}
	db, err := setupDBConn(properties.DbHost, properties.DbPort, properties.DbName, properties.DbUser, properties.DbPassword, "GMT",
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "gamescoring.", // schema name
				SingularTable: false,
			},
		})
	if err != nil {
		return nil, err
	}
	repo := &dbRepository{db: db}
	// backfill with dummy data
	repo.initData()
	return repo, nil
}

func (repository *dbRepository) GetStrikeoutsCountPerGame(gameID string) (int, error) {
	if gameID == "" {
		return 0, errors.New("id cannot be empty")
	}
	_, err := repository.Game(gameID)
	if err != nil {
		return 0, err
	}
	var count int
	res := repository.db.Raw("select count(*) from scoring_events where game_id = ? and data->>attributes->>result = ? ", gameID, string(model.Strikeout)).Scan(&count)
	if res.Error != nil {
		return 0, res.Error
	}

	return count, nil

}

func (repository *dbRepository) UpsertGame(game *model.Game) (*model.Game, error) {
	if game == nil {
		return game, errors.New("Game expected")
	}
	_, err := repository.Game(game.ID)
	if err != nil {
		err := repository.db.Create(game).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := repository.db.Save(game).Error
		if err != nil {
			return nil, err
		}
	}

	return game, nil
}

func (repository *dbRepository) Game(id string) (*model.Game, error) {

	var game *model.Game
	err := repository.db.First(&game, "id = ?", id).Error
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return game, nil
}

func (repository *dbRepository) DeleteGame(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}

	return repository.db.Delete(&model.Game{}, "id = ?", id).Error
}

func (repository *dbRepository) ScoringEvent(id string) (*model.ScoringEvent, error) {

	se := &model.ScoringEvent{}
	err := repository.db.First(&se, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return se, nil
}

func (repository *dbRepository) UpsertScoringEvent(se *model.ScoringEvent) (*model.ScoringEvent, error) {
	if se == nil {
		return se, errors.New("scoring event expected")
	}

	_, err := repository.ScoringEvent(se.ID)
	if err != nil {
		err := repository.db.Create(se).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := repository.db.Save(se).Error
		if err != nil {
			return nil, err
		}
	}

	return se, nil
}

func (repository *dbRepository) DeleteScoringEvent(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}

	return repository.db.Delete(&model.ScoringEvent{}, "id = ?", id).Error
}

func (repository *dbRepository) ListGames() ([]*model.Game, error) {
	// TODO to support pagination

	var games []*model.Game
	err := repository.db.Find(&games).Error
	return games, err
}

func (repository *dbRepository) ListScoringEvents() ([]*model.ScoringEvent, error) {
	// TODO to support pagination

	var events []*model.ScoringEvent
	err := repository.db.Find(&events).Error
	return events, err
}

func validatePostgresDbProperties(properties PostgresDbProperties) error {
	if properties.DbHost == "" {
		return fmt.Errorf("PostgresDbProperties.DbHost must not be empty")
	}
	if properties.DbName == "" {
		return fmt.Errorf("PostgresDbProperties.DbName must not be empty")
	}
	if properties.DbPort == "" {
		return fmt.Errorf("PostgresDbProperties.DbPort must not be empty")
	}
	if properties.DbUser == "" {
		return fmt.Errorf("PostgresDbProperties.DbUser must not be empty")
	}
	if properties.DbPassword == "" {
		return fmt.Errorf("PostgresDbProperties.DbPassword must not be empty")
	}
	return nil
}

func setupDBConn(host, port, dbName, user, password,
	dbTimeZone string, gormConfig *gorm.Config) (*gorm.DB, error) {

	dbConnURL := fmt.Sprintf("host=%s  port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		host, port, user, password, dbName, dbTimeZone)

	dbClient, err := gorm.Open(postgres.Open(dbConnURL), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to datbase: %v", err)
	}

	return dbClient, nil
}

func (repository *dbRepository) initData() {
	for _, g := range Games {
		err := repository.db.Create(g).Error
		if err != nil {
			fmt.Printf(" error creating game %s", g.ID)
		}
	}
	for _, e := range Events {
		err := repository.db.Create(e).Error
		if err != nil {
			fmt.Printf(" error creating scoring event %s", e.ID)
		}
	}
}
