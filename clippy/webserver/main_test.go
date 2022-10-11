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
