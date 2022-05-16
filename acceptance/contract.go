package acceptance

type GameDriver interface {
	HealthCheck() error
	JoinGame(string) error
	//LeaveGame(id string) error
	//IsConnectedToGame(id string) (bool, error)
}
