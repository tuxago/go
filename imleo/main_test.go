package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPepper(t *testing.T) {
	SetScore("Pepper", 20)

	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

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

	PlayerServer(response, request)

	got := response.Body.String()
	want := "10"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
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
