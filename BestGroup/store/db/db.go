package playerdb

import (
	"database/sql"

	"github.com/tuxago/go/BestGroup/store"
)

type PlayerStore struct {
	db *sql.DB
}

func NewPlayerStore(db *sql.DB) *PlayerStore {
	return &PlayerStore{
		db: db,
	}
}

func (ps *PlayerStore) GetPlayer(name string) (store.Player, error) {
	playerName, wins, err := GetPlayer(ps.db, name)
	if err != nil {
		return store.Player{}, err
	}
	return store.Player{Name: playerName, Wins: wins}, nil
}

func (ps *PlayerStore) IncWins(name string) (int, error) {
	wins, err := IncWins(ps.db, name)
	if err != nil {
		return -1, err
	}
	return wins, nil
}

func (ps *PlayerStore) RemovePlayer(name string) error {
	err := RemovePlayer(ps.db, name)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PlayerStore) AddPlayer(name string) error {
	err := AddPlayer(ps.db, name)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PlayerStore) GetAllPlayers() ([]store.Player, error) {
	players, err := GetAllPlayers(ps.db)
	if err != nil {
		return []store.Player{}, err
	}
	return players, nil
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		panic(err)
	}
	return db
}

func AddPlayer(db *sql.DB, name string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into players(name, wins) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, 0)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func GetPlayer(db *sql.DB, name string) (string, int, error) {
	rows, err := db.Query("select name, wins from players where name = ?", name)
	if err != nil {
		return "", -1, err
	}
	defer rows.Close()
	var playerName string
	var wins int
	for rows.Next() {
		err = rows.Scan(&playerName, &wins)
		if err != nil {
			return "", -1, err
		}
	}
	err = rows.Err()
	if err != nil {
		return "", -1, err
	}
	return playerName, wins, nil
}

func IncWins(db *sql.DB, name string) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return -1, err
	}
	stmt, err := tx.Prepare("update players set wins = wins + 1 where name = ?")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	if err != nil {
		return -1, err
	}
	tx.Commit()
	return 42, nil
}

func RemovePlayer(db *sql.DB, name string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("delete from players where name = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func GetAllPlayers(db *sql.DB) ([]store.Player, error) {
	rows, err := db.Query("select name, wins from players")
	if err != nil {
		return []store.Player{}, err
	}
	defer rows.Close()
	var playerName string
	var wins int
	var players []store.Player
	for rows.Next() {
		err = rows.Scan(&playerName, &wins)
		if err != nil {
			return []store.Player{}, err
		}
		player := store.Player{
			Name: playerName,
			Wins: wins,
		}
		players = append(players, player)
	}
	err = rows.Err()
	if err != nil {
		return []store.Player{}, err
	}
	return players, nil
}
