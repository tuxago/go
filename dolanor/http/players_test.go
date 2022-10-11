package main

import (
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestPlayersService(t *testing.T) {
	client := http.DefaultClient

	serviceHost := "http://192.168.20.16:8080"
	testHost := os.Getenv("TEST_HOST")
	if testHost != "" {
		serviceHost = testHost
	}

	req, err := http.NewRequest(http.MethodGet, serviceHost+"/players/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	type player struct {
		Name string
	}

	var got []player
	want := []player{
		{Name: "Pepper"},
		{Name: "Jason"},
		{Name: "Alix"},
	}

	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(got, want) {
		t.Fatal("wrong list of players")
	}

}
