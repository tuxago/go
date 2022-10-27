package format

import (
	"encoding/json"
	"sort"
	"strconv"

	"github.com/tuxago/go/BestGroup/store"
)

func FormatPlayers(format string, players []store.Player) (string, error) {
	sort.SliceStable(players, func(i, j int) bool {
		return players[i].Name < players[j].Name
	})

	switch format {
	case "string":
		// Players to string in var str
		var str string
		for _, player := range players {
			str += player.Name + " " + strconv.Itoa(player.Wins) + "\n"
		}
		return str, nil
	case "csv":
		var str string
		str += "Name,Wins\n"
		// Players to string in var str
		for _, player := range players {
			str += player.Name + "," + strconv.Itoa(player.Wins) + "\n"
		}
		return str, nil
	case "html":
		var str string
		str += "<table>\n"
		str += "<tr><th>Name</th><th>Wins</th></tr>\n"
		// Players to string in var str
		for _, player := range players {
			str += "<tr><td>" + player.Name + "</td><td>" + strconv.Itoa(player.Wins) + "</td></tr>\n"
		}
		str += "</table>\n"
		return str, nil
	case "xml":
		var str string
		str += "<players>\n"
		// Players to string in var str
		for _, player := range players {
			str += "<player>\n"
			str += "<name>" + player.Name + "</name>\n"
			str += "<wins>" + strconv.Itoa(player.Wins) + "</wins>\n"
			str += "</player>\n"
		}
		str += "</players>\n"
		return str, nil
	default:
		// convert to json
		jsonData, err := json.Marshal(players)
		if err != nil {
			return "", err
		}
		return string(jsonData), nil
	}
}
