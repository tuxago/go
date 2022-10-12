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

func getPlayerList() string {
	//	var playerList string
	//	players := parsePlayers()
	//	for _, player := range players.Players {
	//		playerList += "{" + player.Name + ", " + strconv.Itoa(player.Wins) + ", " + strconv.Itoa(player.TotalGames) + ", " + player.Team + "},\n"
	//	}
	//	return playerList

	message := ""
	var players Players
	jsonFile, err := os.Open("./players.json")
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &players)
	message = "["
	for i, player := range players.Players {
		message += player.Name
		if i != len(players.Players)-1 {
			message += ", "
		}
	}
	message += "]"
	return message
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

func getPlayerTotalGames(name string) int {
	players := parsePlayers()
	for _, player := range players.Players {
		if player.Name == name {
			return player.TotalGames
		}
	}
	return -1
}

func getPlayerTeam(name string) string {
	players := parsePlayers()
	for _, player := range players.Players {
		if player.Name == name {
			return player.Team
		}
	}
	return "No team"
}

func URLSplit(url string) []string {
	if url[len(url)-1] != '/' {
		url += "/"
	}
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
	category := strings.ToLower(getURLCategory(r.URL.Path))

	message := ""

	switch category {
	case "wins":
		message = strconv.Itoa(getPlayerScore(player))
	case "totalgames":
		message = strconv.Itoa(getPlayerTotalGames(player))
	case "team":
		message = getPlayerTeam(player)
	case "/":
		message = getPlayerList()
	default:
		message = "{\"Name:\":" + player + ", \"Wins:\":" + strconv.Itoa(getPlayerScore(player)) + ", \"TotalGames:\":" + strconv.Itoa(getPlayerTotalGames(player)) + ", \"Team:\":\"" + getPlayerTeam(player) + "\"}"
	}

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		//increase the player's score

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
