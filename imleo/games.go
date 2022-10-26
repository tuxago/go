package main

import (
	"sort"
	"time"
)

type Game struct {
	Id      int
	Date    time.Time
	Player1 string
	Player2 string
}

var games = []Game{}

func GetGameList() []Game {
	copy := games
	sort.Slice(copy, func(i, j int) bool {
		return copy[i].Date.Before(copy[j].Date)
	})
	return copy
}

func setGames(new []Game) {
	games = new
}
