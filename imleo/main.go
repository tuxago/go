package main

import (
	"fmt"
	"net/http"
	"strings"
)

var scores = map[string]int{
	"Pepper":  20,
	"Salt":    10,
	"Paprika": 30,
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

func GetScore(name string) int {
	return scores[name]
}

func IncreaseScore(name string) int {
	score := GetScore(name)
	score += 1
	scores[name] = score
	return score
}

func resetScores() {
	// for name := range scores {
	// 	scores[name] = 0
	// }
	scores = map[string]int{
		"Pepper":  20,
		"Salt":    10,
		"Paprika": 30,
	}
}
