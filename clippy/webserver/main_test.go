package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var expectedScores = map[string]int{
	"Pepper":  0,
	"Jhon":    1,
	"Brendon": -1,
	"Louis":   2}

func TestGETPlayers(t *testing.T) {
	for name, score := range expectedScores {
		t.Run(name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/players/"+name+"/wins", nil)
			response := httptest.NewRecorder()

			PlayerServer(response, request)

			got := response.Body.String()
			want := strconv.Itoa(score)

			if got != want {
				if !(got == "This player does not exist" && want == "-1") {
					t.Errorf("got %q, want %q", got, want)
				}
			}
		})
	}
}

func TestGetPlayerScore(t *testing.T) {
	for name, want := range expectedScores {
		t.Run(name, func(t *testing.T) {
			got := getPlayerScore(name)
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
func TestGetUnknownPlayerScore(t *testing.T) {
	want := -1
	got := getPlayerScore("UnkownPlayerNotInDatabase")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestParsePlayer(t *testing.T) {
	for name, want := range expectedScores {
		t.Run(name, func(t *testing.T) {
			got := getPlayerScore(name)
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestUrlHandler(t *testing.T) {
	for request, want := range map[string]string{
		"/players/Pepper/wins":  "0",
		"/players/Jhon/wins":    "1",
		"/players/Brendon/wins": "-1",
		"/players/Louis/wins":   "2",
		"/players/sort/wins":    "[{\"Name:\":\"Pepper\", \"Wins:\":0, \"TotalGames:\":9, \"Team:\":\"Red\"}, {\"Name:\":\"Jhon\", \"Wins:\":1, \"TotalGames:\":8, \"Team:\":\"Green\"}, {\"Name:\":\"Louis\", \"Wins:\":2, \"TotalGames:\":7, \"Team:\":\"Purpel\"}, {\"Name:\":\"Paul\", \"Wins:\":3, \"TotalGames:\":6, \"Team:\":\"Green\"}, {\"Name:\":\"Mia\", \"Wins:\":4, \"TotalGames:\":5, \"Team:\":\"White\"}, {\"Name:\":\"Jesse\", \"Wins:\":5, \"TotalGames:\":4, \"Team:\":\"Blue\"}, {\"Name:\":\"Walter\", \"Wins:\":6, \"TotalGames:\":3, \"Team:\":\"White\"}]",
		"/players/sort/name":    "[{\"Name:\":\"Jesse\", \"Wins:\":5, \"TotalGames:\":4, \"Team:\":\"Blue\"}, {\"Name:\":\"Jhon\", \"Wins:\":1, \"TotalGames:\":8, \"Team:\":\"Green\"}, {\"Name:\":\"Louis\", \"Wins:\":2, \"TotalGames:\":7, \"Team:\":\"Purpel\"}, {\"Name:\":\"Mia\", \"Wins:\":4, \"TotalGames:\":5, \"Team:\":\"White\"}, {\"Name:\":\"Paul\", \"Wins:\":3, \"TotalGames:\":6, \"Team:\":\"Green\"}, {\"Name:\":\"Pepper\", \"Wins:\":0, \"TotalGames:\":9, \"Team:\":\"Red\"}, {\"Name:\":\"Walter\", \"Wins:\":6, \"TotalGames:\":3, \"Team:\":\"White\"}]"} {
		t.Run(request, func(t *testing.T) {
			got := urlHandler(request)
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
