// basic go api server
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Players struct {
	Players []Player
}

type Player struct {
	Name       string
	Wins       int
	TotalGames int
	Team       string
}

var players = map[string]int{
	"Pepper":      20,
	"Jhon":        20,
	"Brendon":     10,
	"Louis":       30,
	"Harry":       40,
	"Zayn":        50,
	"Niall":       60,
	"Mia":         70, //Théo's sister
	"Taylor":      30,
	"Mary":        10,
	"Maris":       20,
	"Maryoumaris": 30,
	"Jhonny":      30,
	"Sins":        40,
}

func main() {
	http.HandleFunc("/players/", PlayerServer)
	http.ListenAndServe(":8080", nil)
}

func parsePlayers() Players {
	var players Players
	jsonFile, err := os.Open("./players.json")
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &players)
	return players
}

func getPlayerScore(name string) int {
	players := parsePlayers()
	for _, player := range players.Players {
		if player.Name == name {
			return player.Wins
		}
	}
	return -1
}

func URLSplit(url string) []string {
	return strings.Split(url, "/")
}

func getURLbase(url string) string {
	info := URLSplit(url)
	if len(info) < 2 {
		return "/"
	}
	return info[1]
}

func getURLPlayer(url string) string {
	info := URLSplit(url)
	if len(info) < 3 {
		return "/"
	}
	return info[2]
}

func getURLCategory(url string) string {
	info := URLSplit(url)
	if len(info) < 4 {
		return "/"
	}
	return info[3]
}

// returns the wins of a player on GET /players/{name}/wins
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	//base := getURLbase(r.URL.Path)
	player := getURLPlayer(r.URL.Path)
	//category := getURLCategory(r.URL.Path)

	score := getPlayerScore(player)

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		if score != -1 {
			w.Write([]byte(strconv.Itoa(score)))
		} else {
			w.Write([]byte("This player does not exist"))
		}
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		//increase the player's score

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
