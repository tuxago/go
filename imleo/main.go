package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
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
	http.HandleFunc("/", RootServer)
	http.ListenAndServe(":8080", nil)
	fmt.Println("starting the server on port 8080")
}

func RootServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "Hello")
		return
	}

	sub := strings.Split(r.URL.Path, "/")[1]
	switch sub {
	case "players":
		PlayerServer(w, r)
	case "games":
		GameServer(w, r)
	default:
		fmt.Fprint(w, "404: Not found")
	}
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	name := ""
	if len(split) >= 3 {
		name = split[2]
	}

	switch r.Method {
	case http.MethodGet:
		if name == "" {
			names := GetPlayerList()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(names)
		} else {
			score := GetScore(name)
			fmt.Fprint(w, score)
		}
	case http.MethodPost:
		IncreaseScore(name)
	default:
		fmt.Fprint(w, "unkown method")
	}
}

func GameServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "unkown method")
		return
	}

	split := strings.Split(r.URL.Path, "/")
	id := -1
	if len(split) >= 3 && split[2] != "" {
		var err error
		id, err = strconv.Atoi(split[2])
		if err != nil {
			fmt.Fprintf(w, "invalid id %q", split[2])
			return
		}
	}

	if id == -1 {
		games := GetGameList()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(games)
	} else {
		for _, game := range games {
			if game.Id == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(game)
				return
			}
		}
		fmt.Fprint(w, "game does not exist")
	}
}

func GetPlayerList() []string {
	names := []string{}
	for _, player := range players {
		names = append(names, player.Name)
	}
	sort.Strings(names)
	return names
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
