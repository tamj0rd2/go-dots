package acceptance

import (
	"github.com/alecthomas/assert/v2"
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
		assert.NoError(t, err, "the application should be healthy")
	})

	t.Run("Both players can join the game", func(t *testing.T) {
		err = firstPlayerDriver.JoinGame(gameID)
		assert.NoError(topT, err)

		err := secondPlayerDriver.JoinGame(gameID)
		assert.NoError(topT, err)
	})

	//t.Run("When a player leaves the game, the other also gets disconnected", func(t *testing.T) {
	//	WIP(t)
	//	err = firstPlayerDriver.LeaveGame(gameID)
	//	assert.NoError(topT, err)
	//
	//	WIP(t)
	//	isConnected, err := firstPlayerDriver.IsConnectedToGame(gameID)
	//	assert.NoError(topT, err)
	//	assert.False(topT, isConnected)
	//})
}

func WIP(t *testing.T) {
	t.Helper()

	if shouldRunWIPTests, _ := os.LookupEnv("WIP_ACCEPTANCE_ENABLED"); shouldRunWIPTests != "true" {
		t.Skip("skipping WIP acceptance test")
	}
}
