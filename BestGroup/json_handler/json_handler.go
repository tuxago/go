package json_handler

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"sync"
)

// Deserialize JPlayers
type JPlayers struct {
	Players []JPlayer
}

type JPlayer struct {
	Name string
	Wins int
}

var (
	muPlayers sync.RWMutex
	Players   JPlayers
)

func InitJSON(jsonI string) error {
	jsonFile, err := os.Open(jsonI)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	json.Unmarshal(byteValue, &Players)
	return nil
}

func SaveJSON(jsonI string) error {
	muPlayers.Lock()
	defer muPlayers.Unlock()
	// open output file
	jsonFile, err := os.OpenFile(jsonI, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// convert to json
	jsonData, err := json.Marshal(Players)
	if err != nil {
		return err
	}

	//write to file
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func GetPlayer(name string) (JPlayer, error) {
	muPlayers.RLock()
	defer muPlayers.RUnlock()
	for _, player := range Players.Players {
		if player.Name == name {
			return player, nil
		}
	}
	return JPlayer{}, errors.New("player not found in getplayer")
}

func SetPlayer(name string) (int, error) {
	muPlayers.Lock()
	defer muPlayers.Unlock()
	for i, player := range Players.Players {
		if player.Name == name {
			wins := Players.Players[i].Wins + 1
			Players.Players[i].Wins = wins
			return wins, nil
		}
	}
	return -1, errors.New("player not found in setplayer")
}

func RemovePlayer(name string) error {
	muPlayers.Lock()
	defer muPlayers.Unlock()
	for i, player := range Players.Players {
		if player.Name == name {
			Players.Players = append(Players.Players[:i], Players.Players[i+1:]...)
			return nil
		}
	}
	return errors.New("player not found in removeplayer")
}

func AddPlayer(name string) error {
	muPlayers.Lock()
	defer muPlayers.Unlock()
	for _, player := range Players.Players {
		if player.Name == name {
			return errors.New("player already exists")
		}
	}
	Players.Players = append(Players.Players, JPlayer{Name: name, Wins: 0})
	return nil
}

func FormatPlayers(format string) (string, error) {
	muPlayers.RLock()
	defer muPlayers.RUnlock()
	switch format {
	case "string":
		// Players to string in var str
		var str string
		for _, player := range Players.Players {
			str += player.Name + " " + strconv.Itoa(player.Wins) + "\n"
		}
		return str, nil
	case "csv":
		var str string
		str += "Name,Wins\n"
		// Players to string in var str
		for _, player := range Players.Players {
			str += player.Name + "," + strconv.Itoa(player.Wins) + "\n"
		}
		return str, nil
	case "html":
		var str string
		str += "<table>\n"
		str += "<tr><th>Name</th><th>Wins</th></tr>\n"
		// Players to string in var str
		for _, player := range Players.Players {
			str += "<tr><td>" + player.Name + "</td><td>" + strconv.Itoa(player.Wins) + "</td></tr>\n"
		}
		str += "</table>\n"
		return str, nil
	case "xml":
		var str string
		str += "<players>\n"
		// Players to string in var str
		for _, player := range Players.Players {
			str += "<player>\n"
			str += "<name>" + player.Name + "</name>\n"
			str += "<wins>" + strconv.Itoa(player.Wins) + "</wins>\n"
			str += "</player>\n"
		}
		str += "</players>\n"
		return str, nil
	default:
		// convert to json
		jsonData, err := json.Marshal(Players)
		if err != nil {
			return "", err
		}
		return string(jsonData), nil
	}
}
