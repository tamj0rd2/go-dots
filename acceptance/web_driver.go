package acceptance

import (
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/tamj0rd2/go-dots/src/domain"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"os"
)

type WebGameDriver struct {
	httpClient *http.Client
	cfg        WebGameDriverConfig
}

func NewWebGameDriver(cfg WebGameDriverConfig) GameDriver {
	return &WebGameDriver{
		httpClient: &http.Client{},
		cfg:        cfg,
	}
}

func (d WebGameDriver) HealthCheck() error {
	return backoff.Retry(func() error {
		healthURL := fmt.Sprintf("%s/health", d.cfg.HTTPBaseURL)

		res, err := d.httpClient.Get(healthURL)
		if err != nil {
			return fmt.Errorf("failed to make request to %s: %w", healthURL, err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("health check failed with status code %d", res.StatusCode)
		}

		return nil
	}, backoff.NewExponentialBackOff())
}

func (d WebGameDriver) JoinGame(id string) error {
	url := fmt.Sprintf("%s/games/%s", d.cfg.WSBaseURL, id)

	ws, err := websocket.Dial(url, "", d.cfg.HTTPBaseURL)
	if err != nil {
		return err
	}
	defer ws.Close()

	//if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
	//	return nil
	//}
	//
	//var msg = make([]byte, 512)
	//var n int
	//if n, err = ws.Read(msg); err != nil {
	//	return nil
	//}
	//fmt.Printf("Received: %s.\n", msg[:n])
	return nil
}

func (d WebGameDriver) GetBoard(gameID string) (domain.Board, error) {
	url := fmt.Sprintf("%s/games/%s", d.cfg.HTTPBaseURL, gameID)

	res, err := d.httpClient.Get(url)
	if err != nil {
		return domain.EmptyBoard, fmt.Errorf("failed to make request to %s: %w", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return domain.EmptyBoard, fmt.Errorf("failed to get board with status code %d - %s", res.StatusCode, url)
	}

	var game domain.Game
	if err := json.NewDecoder(res.Body).Decode(&game); err != nil {
		return domain.EmptyBoard, fmt.Errorf("failed to decode board: %w", err)
	}

	return game.Board, nil
}

type WebGameDriverConfig struct {
	HTTPBaseURL string
	WSBaseURL   string
}

func NewWebGameDriverConfig() (cfg WebGameDriverConfig, err error) {
	rawBaseURL, ok := os.LookupEnv("BASE_URL")
	if !ok {
		return
	}

	parsedURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}

	wsBaseURL := *parsedURL
	wsBaseURL.Scheme = "ws"

	return WebGameDriverConfig{
		HTTPBaseURL: parsedURL.String(),
		WSBaseURL:   wsBaseURL.String() + "/ws",
	}, nil
}
