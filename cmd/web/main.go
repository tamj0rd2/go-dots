package main

import (
	"context"
	"github.com/tamj0rd2/go-dots/src"
	"github.com/tamj0rd2/go-dots/src/telemetry"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {
	logger, _ := telemetry.NewApplicationLogger(context.Background())
	defer logger.Sync()

	port, isPortInEnv := os.LookupEnv("PORT")
	if !isPortInEnv {
		port = "8000"
	}

	logger.Info("Starting server", zap.String("port", port))
	if err := http.ListenAndServe(":"+port, src.NewHandler(logger)); err != nil {
		logger.Fatal("Server exited", zap.Error(err))
	}
}
