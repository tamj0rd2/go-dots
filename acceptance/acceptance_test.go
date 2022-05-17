package acceptance

import (
	"github.com/alecthomas/assert/v2"
	"github.com/tamj0rd2/go-dots/src/domain"
	"os"
	"testing"
)

func TestPlayingAGameWith2Players(t *testing.T) {
	cfg, err := NewWebGameDriverConfig()
	assert.NoError(t, err)
	firstPlayerDriver := NewWebGameDriver(cfg)
	secondPlayerDriver := NewWebGameDriver(cfg)

	var (
		gameID = "some-game-id"
		topT   = t
	)

	t.Run("The website is healthy", func(t *testing.T) {
		err = firstPlayerDriver.HealthCheck()
		assert.NoError(topT, err, "the application should be healthy")
	})

	t.Run("Both players can join the game", func(t *testing.T) {
		err = firstPlayerDriver.JoinGame(gameID)
		assert.NoError(topT, err)

		err := secondPlayerDriver.JoinGame(gameID)
		assert.NoError(topT, err)
	})

	t.Run("When both players have joined the game, they both see a blank board", func(t *testing.T) {
		firstPlayerBoard, err := firstPlayerDriver.GetBoard(gameID)
		assert.NoError(topT, err)
		assert.Equal(topT, domain.EmptyBoard, firstPlayerBoard)

		secondPlayerBoard, err := secondPlayerDriver.GetBoard(gameID)
		assert.NoError(topT, err)
		assert.Equal(topT, domain.EmptyBoard, secondPlayerBoard)
	})
}

func WIP(t *testing.T) {
	t.Helper()

	if shouldRunWIPTests, _ := os.LookupEnv("WIP_ACCEPTANCE_ENABLED"); shouldRunWIPTests != "true" {
		t.Skip("skipping WIP acceptance test")
	}
}
