package main

import (
	"testing"
)

var Json string = "test.json"

func TestInitAndSaveJSON(t *testing.T) {
	t.Run("TestInitJSON", func(t *testing.T) {
		err := InitJSON(Json, &Players)
		//if Players is empty, then the test fails
		if err != nil {
			t.Errorf("Error in InitJSON, %v", err)
		}
	})
	t.Run("TestSaveJSON", func(t *testing.T) {
		err := InitJSON(Json, &Players)
		if err != nil {
			t.Errorf("Error in InitJSON %v", err)
		}
		//add a new player to the list
		Players.Players = append(Players.Players, JPlayer{Name: "Test", Wins: 0})
		err = SaveJSON(Json)
		if err != nil {
			t.Errorf("Error in SaveJSON %v", err)
		}
		err = InitJSON(Json, &Players)
		if err != nil {
			t.Errorf("Error in second InitJSON %v", err)
		}
		//if Players does not contain Name : Test with Wins :0, then the test fails
		if Players.Players[len(Players.Players)-1].Name != "Test" || Players.Players[len(Players.Players)-1].Wins != 0 {
			t.Errorf("Players.JPlayers does not contain Name : Test with Wins :0")
		}
	})

}

func TestSetPlayer(t *testing.T) {
	t.Run("TestSetPlayer 'Test' wins", func(t *testing.T) {
		err := InitJSON(Json, &Players)
		if err != nil {
			t.Errorf("Error in InitJSON")
		}
		wins, err := SetPlayer("Test")
		if err != nil {
			t.Errorf("Error in SetPlayer")
		}
		if wins == -1 || err != nil {
			t.Errorf("Players.JPlayers does not contain Name : Test with Wins :4, or an error : %v", err)
		}
	})
	t.Run("TestSetPlayer 'Test123' wins and get error", func(t *testing.T) {
		err := InitJSON(Json, &Players)
		if err != nil {
			t.Errorf("Error in InitJSON")
		}
		_, err = SetPlayer("Test123")
		if err == nil {
			t.Errorf("no error in SetPlayer")
		}
	})
}

func TestGetPlayer(t *testing.T) {
	InitJSON(Json, &Players)
	//if GetPlayer returns an empty JPlayer, then the test fails
	player, err := GetPlayer("Test")
	if player == (JPlayer{}) || err != nil {
		t.Errorf("GetPlayer(\"Test\") returned an empty JPlayer, or an error : %v", err)
	}
}

func TestRemovePlayer(t *testing.T) {
	t.Run("TestRemoveJSON", func(t *testing.T) {
		err := InitJSON(Json, &Players)
		if err != nil {
			t.Errorf("Error in InitJSON")
		}
		err = RemovePlayer("Test")
		if err != nil {
			t.Errorf("Error in RemovePlayer")
		}
		err = SaveJSON(Json)
		if err != nil {
			t.Errorf("Error in SaveJSON")
		}
		err = InitJSON(Json, &Players)
		if err != nil {
			t.Errorf("Error in second InitJSON")
		}
		player, err := GetPlayer("Test")
		if player.Name == "Test" || err.Error() != "player not found in getplayer" {
			t.Errorf("Players.JPlayers does contain Name : Test with Wins :0, or an error : %v", err)
		}
	})
}
