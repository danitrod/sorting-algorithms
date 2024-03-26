package main

import (
	"fmt"
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

	sizes := []int{
		len(playersFromDataset) / 20,
		len(playersFromDataset) / 15,
		len(playersFromDataset) / 10,
		len(playersFromDataset) / 5,
		len(playersFromDataset) / 2,
		int(float32(len(playersFromDataset)) / 1.5),
		len(playersFromDataset),
	}

	b := bencher.New(cfg.Algorithm.String(), numIterations)

	for _, size := range sizes {
		results, err := b.Bench(func() {
			unsortedPlayers := make([]players.Player, size)
			copy(unsortedPlayers, playersFromDataset[:size])
			cfg.Algorithm.Sort(unsortedPlayers, players.Compare)
		}, size)
		if err != nil {
			panic("error: " + err.Error())
		}
		slog.Info(fmt.Sprintf("Successfully executed experiment with size %d", size), slog.Attr{Key: "results", Value: slog.AnyValue(results)})
	}
}
