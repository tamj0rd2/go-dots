package main

import (
	"context"
	"github.com/tamj0rd2/go-dots/src/telemetry"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {
	logger, ctx := telemetry.NewApplicationLogger(context.Background())
	defer logger.Sync()

	if err := startServer(ctx); err != nil {
		logger.Fatal("Server exited", zap.Error(err))
	}
}

func startServer(ctx context.Context) error {
	port, isPortInEnv := os.LookupEnv("PORT")
	if !isPortInEnv {
		port = "8000"
	}

	logger := telemetry.GetLogger(ctx)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("App is healthy")
		w.WriteHeader(http.StatusOK)
	})

	logger.Info("Server started", zap.String("port", port))
	return http.ListenAndServe(":"+port, nil)
}
