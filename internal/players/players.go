package players

import (
	"os"

	"github.com/gocarina/gocsv"
)

const fileLocation = "./data/players.csv"

type Player struct {
	Name          string `csv:"Player"`
	Goals         int    `csv:"Performance_Gls"`
	Assists       int    `csv:"Performance.1_Ast"`
	PlayingTime   int    `csv:"Playing Time.1_Min"`
	MatchesPlayed int    `csv:"MP"`
}

func Compare(p1, p2 Player) bool {
	if p1.Goals != p2.Goals {
		return p1.Goals > p2.Goals
	}

	if p1.Assists != p2.Assists {
		return p1.Assists > p2.Assists
	}

	if p1.MatchesPlayed != p2.MatchesPlayed {
		return p1.MatchesPlayed > p2.MatchesPlayed
	}

	if p1.PlayingTime != p2.PlayingTime {
		return p1.PlayingTime > p2.PlayingTime
	}

	return p1.Name < p2.Name
}

func FetchFromDataset() []Player {
	csvFile, err := os.Open(fileLocation)
	if err != nil {
		panic("error opening file: " + err.Error())
	}
	defer csvFile.Close()

	players := []Player{}
	if err := gocsv.UnmarshalFile(csvFile, &players); err != nil {
		panic(err)
	}

	// Sum accross player seasons
	playerSums := map[string]Player{}
	for _, player := range players {
		if _, ok := playerSums[player.Name]; !ok {
			playerSums[player.Name] = player
			continue
		}

		playerSums[player.Name] = Player{
			Name:          player.Name,
			Goals:         playerSums[player.Name].Goals + player.Goals,
			Assists:       playerSums[player.Name].Assists + player.Assists,
			PlayingTime:   playerSums[player.Name].PlayingTime + player.PlayingTime,
			MatchesPlayed: playerSums[player.Name].MatchesPlayed + player.MatchesPlayed,
		}
	}

	players = make([]Player, 0, len(playerSums))
	for _, player := range playerSums {
		players = append(players, player)
	}

	return players
}
