package acceptance

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type GameDriver interface {
	CreateGame() (GameID, error)
	JoinGame(GameID) error
}

type GameID string

func TestPlayingAGameWith2Players(t *testing.T) {
	topT := t

	var (
		driver GameDriver
		gameID GameID
		err    error
	)

	t.Run("this test actually runs", func(t *testing.T) {
		assert.Equal(topT, 1, 1)
	})

	t.Run("Player 1 creates a game", func(t *testing.T) {
		WIP(t)
		gameID, err = driver.CreateGame()
		assert.NoError(topT, err)
	})

	t.Run("Player 2 can join the game", func(t *testing.T) {
		WIP(t)
		err := driver.JoinGame(gameID)
		assert.NoError(topT, err)
	})
}
