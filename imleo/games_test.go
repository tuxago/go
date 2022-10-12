package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGETGameList(t *testing.T) {
	setGames([]Game{
		{0, time.Date(0, 5, 14, 0, 0, 0, 0, time.UTC), "Pepper", "Salt"},
		{1, time.Date(0, 3, 22, 0, 0, 0, 0, time.UTC), "Pepper", "Paprika"},
		{2, time.Date(0, 9, 1, 0, 0, 0, 0, time.UTC), "Salt", "Paprika"},
	})

	request, _ := http.NewRequest(http.MethodGet, "/games/", nil)
	response := httptest.NewRecorder()

	RootServer(response, request)

	got := response.Body.String()
	want := "[{\"Id\":1,\"Date\":\"0000-03-22T00:00:00Z\",\"Player1\":\"Pepper\",\"Player2\":\"Paprika\"},{\"Id\":0,\"Date\":\"0000-05-14T00:00:00Z\",\"Player1\":\"Pepper\",\"Player2\":\"Salt\"},{\"Id\":2,\"Date\":\"0000-09-01T00:00:00Z\",\"Player1\":\"Salt\",\"Player2\":\"Paprika\"}]\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	request, _ = http.NewRequest(http.MethodGet, "/games", nil)
	response = httptest.NewRecorder()

	RootServer(response, request)

	got = response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
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
