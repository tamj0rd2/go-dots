package acceptance

import "github.com/tamj0rd2/go-dots/src/domain"

type GameDriver interface {
	HealthCheck() error
	JoinGame(gameID string) error
	GetBoard(gameID string) (domain.Board, error)
}
