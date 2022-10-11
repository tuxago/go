package main

import (
	"TeamGo/json_handler"
	"fmt"
	"net/http"
	"strings"
)

var PlayerWins = map[string]int{
	"Pepper": 20,
	"Salt":   0,
}

func main() {
	//init the json_handler package
	json_handler.InitJSON("players.json")
	http.HandleFunc("/players/", PlayerServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	println("Server started on port 8080")
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	//trim the /players/ from the request
	player := strings.TrimPrefix(r.URL.Path, "/players")
	// check if the option ?format is present and get the value
	format := r.URL.Query().Get("format")
	if player == "" || player == "/" {
		if r.Method == http.MethodGet {
			list, err := json_handler.FormatPlayers(format)
			if err != nil {
			} else {
				fmt.Fprint(w, list)
			}
		} else {
			fmt.Fprint(w, "No player name called")
		}
		return
	}
	player = player[1:]
	//get or post
	switch r.Method {
	case http.MethodPost:
		wins, err := json_handler.SetPlayer(player)
		if err != nil {
			fmt.Fprint(w, "Player "+player+" doesn't exist")
		} else {
			fmt.Fprint(w, wins)
		}

	case http.MethodGet:
		_player, err := json_handler.GetPlayer(player)
		if err != nil {
			fmt.Fprint(w, "Player "+player+" doesn't exist")
		} else {
			fmt.Fprint(w, _player.Wins)
		}
	}
}
