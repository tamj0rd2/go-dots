package src_test

import (
	"encoding/json"
	"github.com/alecthomas/assert/v2"
	"github.com/tamj0rd2/go-dots/src"
	"github.com/tamj0rd2/go-dots/src/domain"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUnitServer(t *testing.T) {
	t.Run("Checking health", func(t *testing.T) {
		handler := src.NewHandler(zap.L())

		server := httptest.NewServer(handler)
		defer server.Close()

		res, err := http.DefaultClient.Get(server.URL + "/health")
		assert.NoError(t, err)
		defer res.Body.Close()
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Joining a game opens a websocket connection", func(t *testing.T) {
		var (
			gameID = "some-game-id"
		)

		handler := src.NewHandler(zap.L())
		server := httptest.NewServer(handler)
		defer server.Close()

		wsURL, err := url.Parse(server.URL + "/ws/games/" + gameID)
		assert.NoError(t, err)
		wsURL.Scheme = "ws"
		t.Logf(wsURL.String())

		ws, err := websocket.Dial(wsURL.String(), "", server.URL)
		assert.NoError(t, err)
		defer ws.Close()
	})

	t.Run("Getting a game returns an empty game", func(t *testing.T) {
		var (
			gameID = "some-game-id"
		)

		handler := src.NewHandler(zap.L())
		server := httptest.NewServer(handler)
		defer server.Close()

		res, err := http.DefaultClient.Get(server.URL + "/games/" + gameID)
		assert.NoError(t, err)
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
		var game domain.Game
		assert.NoError(t, json.NewDecoder(res.Body).Decode(&game))
		assert.Equal(t, domain.Game{}, game)
	})
}
