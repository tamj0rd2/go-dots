package src

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tamj0rd2/go-dots/src/domain"
	"github.com/tamj0rd2/go-dots/src/telemetry"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"net/http"
)

type Handler struct {
	r *mux.Router
}

func NewHandler(logger *zap.Logger) *Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(telemetry.AddLoggerToContext(r.Context(), logger))
			logger.Info("request received", zap.String("url", r.URL.String()), zap.String("method", r.Method))
			handler.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/games/{gameID}", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(domain.Game{})
	})

	r.Handle("/ws/games/{gameID}", websocket.Handler(func(conn *websocket.Conn) {
		logger := telemetry.GetLoggerFromContext(conn.Request().Context())
		logger.Info("reached websocket")
	}))

	return &Handler{r}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}
