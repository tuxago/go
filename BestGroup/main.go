package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"sync"
	"os"
	jsonhandler "github.com/tuxago/go/BestGroup/json_handler"
)

var PlayerWins = map[string]int{
	"Pepper": 20,
	"Salt":   0,
}

var logfile string = "./server.log"
var logmutex sync.Mutex

func main() {
	//init the json_handler package
	jsonhandler.InitJSON("players.json")
	http.HandleFunc("/players/", PlayerServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	println("Server started on port 8080")
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	logrequest(r)
	//trim the /players/ from the request
	player := strings.TrimPrefix(r.URL.Path, "/players")
	// check if the option ?format is present and get the value
	format := r.URL.Query().Get("format")
	if player == "" || player == "/" {
		if r.Method == http.MethodGet {
			list, err := jsonhandler.FormatPlayers(format)
			if err != nil {
			} else {
				loganswer("List of Players")
				fmt.Fprint(w, list)
			}
		} else {
			loganswer("No player name called")
			fmt.Fprint(w, "No player name called")
		}
		return
	}
	player = player[1:]
	//get or post
	switch r.Method {
	case http.MethodPost:
		wins, err := jsonhandler.SetPlayer(player)
		if err != nil {
			loganswer("Player "+player+" doesn't exist")
			fmt.Fprint(w, "Player "+player+" doesn't exist")
		} else {
			loganswer(fmt.Sprint(wins))
			fmt.Fprint(w, wins)
		}

	case http.MethodGet:
		_player, err := jsonhandler.GetPlayer(player)
		if err != nil {
			loganswer("Player "+player+" doesn't exist")
			fmt.Fprint(w, "Player "+player+" doesn't exist")
		} else {
			loganswer(fmt.Sprint(_player.Wins))
			fmt.Fprint(w, _player.Wins)
		}
	}
}
//[yyyy-mm-dd:hh-mm-ss-mmss] Recieved $URL with $METHOD
func logrequest(r *http.Request) {
	logmutex.Lock()
	defer logmutex.Unlock()
	ctime := time.Now().Format(time.RFC850)
	logtext := "[" + ctime + "] Recieved " + r.URL.Path + " with " + r.Method + " method \n"
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer f.Close()
	if _, err = f.WriteString(logtext); err != nil {
		return
	}
}

func loganswer(answer string) {
	logmutex.Lock()
	defer logmutex.Unlock()
	ctime := time.Now().Format(time.RFC850)
	logtext := "[" + ctime + "] Answered with " + answer + "\n"
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer f.Close()
	if _, err = f.WriteString(logtext); err != nil {
		return
	}
}
