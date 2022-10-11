package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	Name  string
	Score int
}

var players = []Player{
	{"Pepper", 20},
	{"Salt", 10},
	{"Paprika", 30},
}

func main() {
	http.HandleFunc("/players/", PlayerServer)
	http.ListenAndServe(":8080", nil)
	fmt.Println("starting the server on port 8080")
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.Path, "/")[2]

	if r.Method == http.MethodGet {
		score := GetScore(name)
		fmt.Fprint(w, score)
	} else {
		IncreaseScore(name)
	}
}

func GetPlayer(name string) *Player {
	for _, player := range players {
		if player.Name == name {
			return &player
		}
	}
	return nil
}

func GetScore(name string) int {
	player := GetPlayer(name)
	if player == nil {
		return -1
	}
	return player.Score
}

func SetScore(name string, score int) int {
	for i, player := range players {
		if player.Name == name {
			players[i].Score = score
			return players[i].Score
		}
	}
	return -1
}

func IncreaseScore(name string) int {
	return SetScore(name, GetScore(name)+1)
}

func resetScores() {
	for _, player := range players {
		player.Score = 0
	}
}
