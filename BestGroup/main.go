package main

import (
	"fmt"
	"net/http"
	"strings"
)

var PlayerWins = map[string]int{
	"Pepper": 20,
	"Salt":   0,
}

func main() {
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
	if player == "" || player == "/" {
		fmt.Fprint(w, "No player name called")
		return
	}
	player = player[1:]
	//get or post
	switch r.Method {
	case http.MethodPost:
		SETPlayerWins(player)
	case http.MethodGet:
		wins:= GETPlayerWins(player)
		if wins == -1 {
			fmt.Fprint(w, "Player " + player + " doesn't exist")
		} else {
			fmt.Fprint(w, wins)
		}
	}
}
func GETPlayerWins(name string) int {
	wins, ok := PlayerWins[name]
	if ok {
		return wins
	}

	return -1
}

func SETPlayerWins(name string) {
	PlayerWins[name]++
}
