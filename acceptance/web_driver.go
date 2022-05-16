package acceptance

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
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
		healthURL := fmt.Sprintf("%s/health", d.cfg.HTTPBaseURL.String())

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
	url := d.cfg.WSBaseURL
	url.Path = fmt.Sprintf("/games/%s", id)

	ws, err := websocket.Dial(url.String(), "", d.cfg.HTTPBaseURL.String())
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

type WebGameDriverConfig struct {
	HTTPBaseURL url.URL
	WSBaseURL   url.URL
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

	cfg = WebGameDriverConfig{
		HTTPBaseURL: *parsedURL,
		WSBaseURL:   *parsedURL,
	}
	cfg.WSBaseURL.Scheme = "ws"

	return cfg, nil
}
