package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemDBOperations(t *testing.T) {
	memdb := NewMemDBRepository()

	games, err := memdb.ListGames()

	assert.NoError(t, err, "No error expected")

	assert.Equal(t, 4, len(games), "should be equal to initial backfill")

	game, err := memdb.Game(games[0].ID)
	assert.NoError(t, err, "No error expected")
	assert.Equal(t, games[0], game, "Game expected")

	err = memdb.DeleteGame(games[0].ID)
	assert.NoError(t, err, "No error expected")

	games, err = memdb.ListGames()

	assert.NoError(t, err, "No error expected")

	assert.Equal(t, 3, len(games), "3 games expected as we deleted one above")

	game, err = memdb.UpsertGame(games[0])
	assert.NoError(t, err, "No error expected")
	assert.Equal(t, games[0], game, "Game expected")

}
