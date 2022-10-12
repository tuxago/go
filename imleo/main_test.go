package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGETRoot(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	RootServer(response, request)

	got := response.Body.String()
	want := "Hello"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGETPepper(t *testing.T) {
	SetScore("Pepper", 20)

	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	RootServer(response, request)

	got := response.Body.String()
	want := "20"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGETSalt(t *testing.T) {
	SetScore("Salt", 10)

	request, _ := http.NewRequest(http.MethodGet, "/players/Salt", nil)
	response := httptest.NewRecorder()

	RootServer(response, request)

	got := response.Body.String()
	want := "10"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGETPlayerList(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/", nil)
	response := httptest.NewRecorder()

	RootServer(response, request)

	got := response.Body.String()
	want := "[\"Paprika\",\"Pepper\",\"Salt\"]\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	request, _ = http.NewRequest(http.MethodGet, "/players", nil)
	response = httptest.NewRecorder()

	RootServer(response, request)

	got = response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPOSTPepper(t *testing.T) {
	SetScore("Pepper", 20)
	want := 21

	request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	RootServer(response, request)

	got := GetScore("Pepper")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestIncreasePepper(t *testing.T) {
	SetScore("Pepper", 20)
	want := 21

	got := IncreaseScore("Pepper")
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	got = GetScore("Pepper")
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPlayers(t *testing.T) {
	SetScore("Pepper", 20)
	SetScore("Salt", 10)
	SetScore("Paprika", 30)

	gotPepper := GetScore("Pepper")
	wantPepper := 20
	if gotPepper != wantPepper {
		t.Errorf("got %d, want %d", gotPepper, wantPepper)
	}

	gotSalt := GetScore("Salt")
	wantSalt := 10
	if gotSalt != wantSalt {
		t.Errorf("got %d, want %d", gotSalt, wantSalt)
	}

	gotPaprika := GetScore("Paprika")
	wantPaprika := 30
	if gotPaprika != wantPaprika {
		t.Errorf("got %d, want %d", gotPaprika, wantPaprika)
	}
	gotCurry := GetScore("Curry")
	wantCurry := -1
	if gotCurry != wantCurry {
		t.Errorf("got %d, want %d", gotCurry, wantCurry)
	}
}

func TestPlayerList(t *testing.T) {
	got := GetPlayerList()
	want := []string{
		"Paprika",
		"Pepper",
		"Salt",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("go %v, want %v", got, want)
	}
}

func TestGamesList(t *testing.T) {
	setGames([]Game{
		{0, time.Date(0, 5, 14, 0, 0, 0, 0, time.UTC), "Pepper", "Salt"},
		{1, time.Date(0, 3, 22, 0, 0, 0, 0, time.UTC), "Pepper", "Paprika"},
		{2, time.Date(0, 9, 1, 0, 0, 0, 0, time.UTC), "Salt", "Paprika"},
	})

	got := GetGameList()
	want := []Game{
		{1, time.Date(0, 3, 22, 0, 0, 0, 0, time.UTC), "Pepper", "Paprika"},
		{0, time.Date(0, 5, 14, 0, 0, 0, 0, time.UTC), "Pepper", "Salt"},
		{2, time.Date(0, 9, 1, 0, 0, 0, 0, time.UTC), "Salt", "Paprika"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
