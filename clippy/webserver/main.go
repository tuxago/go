// basic go api server
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
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

func PlayerListToMessage(players Players) string {
	message := "["
	for i, player := range players.Players {
		message += "{\"Name:\":\"" + player.Name + "\", \"Wins:\":" + strconv.Itoa(player.Wins) + ", \"TotalGames:\":" + strconv.Itoa(player.TotalGames) + ", \"Team:\":\"" + player.Team + "\"}"
		if i != len(players.Players)-1 {
			message += ", "
		}
	}
	message += "]"
	return message
}

func sortPlayersByWins(players Players) Players {
	sort.Slice(players.Players, func(i, j int) bool {
		return players.Players[i].Wins < players.Players[j].Wins
	})
	return players
}

func sortPlayersByName(players Players) Players {
	sort.Slice(players.Players, func(i, j int) bool {
		return players.Players[i].Name < players.Players[j].Name
	})
	return players
}

func sortPlayersByTeam(players Players) Players {
	sort.Slice(players.Players, func(i, j int) bool {
		return players.Players[i].Team < players.Players[j].Team
	})
	return players
}

func sortPlayersByTotalGames(players Players) Players {
	sort.Slice(players.Players, func(i, j int) bool {
		return players.Players[i].TotalGames < players.Players[j].TotalGames
	})
	return players
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

func getURLSortType(url string) string {
	info := URLSplit(url)
	if len(info) < 4 {
		return "/"
	}
	return info[3]
}

func urlHandler(url string) string {
	////base := getURLbase(url)
	player := getURLPlayer(url)
	category := strings.ToLower(getURLCategory(url))

	message := ""

	switch player {
	case "sort":
		sorttype := strings.ToLower(getURLSortType(url))
		switch sorttype {
		case "wins":
			message = PlayerListToMessage(sortPlayersByWins(parsePlayers()))
		case "name":
			message = PlayerListToMessage(sortPlayersByName(parsePlayers()))
		case "team":
			message = PlayerListToMessage(sortPlayersByTeam(parsePlayers()))
		case "totalgames":
			message = PlayerListToMessage(sortPlayersByTotalGames(parsePlayers()))
		default:
			message = "Invalid sort type"
		}
	default:
		switch category {
		case "wins":
			message = strconv.Itoa(getPlayerScore(player))
		case "totalgames":
			message = strconv.Itoa(getPlayerTotalGames(player))
		case "team":
			message = getPlayerTeam(player)
		case "/":
			message = PlayerListToMessage(parsePlayers())
		default:
			message = "{\"Name:\":" + player + ", \"Wins:\":" + strconv.Itoa(getPlayerScore(player)) + ", \"TotalGames:\":" + strconv.Itoa(getPlayerTotalGames(player)) + ", \"Team:\":\"" + getPlayerTeam(player) + "\"}"
		}
	}
	return message
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		url := r.URL.Path
		message := urlHandler(url)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		//increase the player's score

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
