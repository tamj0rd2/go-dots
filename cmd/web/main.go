package main

import (
	"context"
	"github.com/tamj0rd2/go-dots/src/telemetry"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
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

	http.Handle("/", websocket.Handler(func(conn *websocket.Conn) {
		logger.Info("websocket reached")
		if _, err := conn.Write([]byte("bleh")); err != nil {
			logger.Error("error writing to websocket", zap.Error(err))
		}
	}))

	logger.Info("Server started", zap.String("port", port))
	return http.ListenAndServe(":"+port, nil)
}
