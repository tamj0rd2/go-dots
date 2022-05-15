package acceptance

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"net/http"
	"os"
)

type OTIGameDriver struct {
	httpClient *http.Client
	baseURL    string
}

func NewOTIGameDriver() *OTIGameDriver {
	baseURL, ok := os.LookupEnv("BASE_URL")
	if !ok {
		panic("BASE_URL not set")
	}

	return &OTIGameDriver{
		httpClient: http.DefaultClient,
		baseURL:    baseURL,
	}
}

func (i OTIGameDriver) HealthCheck() error {
	return backoff.Retry(func() error {
		healthURL := fmt.Sprintf("%s/health", i.baseURL)

		res, err := i.httpClient.Get(healthURL)
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

func (i OTIGameDriver) CreateGame() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (i OTIGameDriver) JoinGame(id string) error {
	//TODO implement me
	panic("implement me")
}
