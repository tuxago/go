package store

// Deserialize JPlayers
type Players struct {
	Players []Player
}

type Player struct {
	Name string
	Wins int
}

type PlayerStorage interface {
	GetPlayer(name string) (Player, error)
	IncWins(name string) (int, error)
	RemovePlayer(name string) error
	AddPlayer(name string) error
}
