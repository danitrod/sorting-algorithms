package main

import (
	"log/slog"

	"github.com/danitrod/sorting-algorithms/internal/bencher"
	"github.com/danitrod/sorting-algorithms/internal/config"
	"github.com/danitrod/sorting-algorithms/internal/players"
)

const numIterations = 500

func main() {
	playersFromDataset := players.FetchFromDataset()
	cfg, err := config.NewConfig[players.Player]()
	if err != nil {
		panic("error: " + err.Error())
	}

	b := bencher.New(cfg.Algorithm.String(), numIterations)

	results, err := b.Bench(func() {
		unsortedPlayers := make([]players.Player, len(playersFromDataset))
		copy(unsortedPlayers, playersFromDataset)
		cfg.Algorithm.Sort(unsortedPlayers, players.Compare)
	})
	if err != nil {
		panic("error: " + err.Error())
	}

	slog.Info("Successfully executed experiment", slog.Attr{Key: "results", Value: slog.AnyValue(results)})
}
