package acceptance

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type GameDriver interface {
	HealthCheck() error
	CreateGame() (string, error)
	JoinGame(string) error
}

func TestPlayingAGameWith2Players(t *testing.T) {
	topT := t
	driver := NewOTIGameDriver()

	var (
		gameID string
		err    error
	)

	t.Run("The website is healthy", func(t *testing.T) {
		err = driver.HealthCheck()
		assert.NoError(t, err, "the application should be healthy")
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
