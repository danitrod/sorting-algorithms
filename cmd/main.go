package main

import (
	"fmt"
	"log/slog"
	"sort"
	"time"

	"github.com/danitrod/sorting-algorithms/internal/config"
	"github.com/danitrod/sorting-algorithms/internal/players"
)

func compareInt(a, b int) bool {
	return a > b
}

func main() {
	playersFromDataset := players.FetchFromDataset()
	cfg, err := config.NewConfig[players.Player]()
	if err != nil {
		panic("error: " + err.Error())
	}

	bench(500, func() {
		unsortedPlayers := make([]players.Player, len(playersFromDataset))
		copy(unsortedPlayers, playersFromDataset)
		cfg.Algorithm.Sort(unsortedPlayers, players.Compare)
	})
}

func bench(iterations int, fn func()) {
	results := make([]time.Duration, 0, iterations)
	for i := 0; i < iterations; i++ {
		start := time.Now()
		fn()
		results = append(results, time.Since(start))
	}

	avg := int64(0)
	for _, r := range results {
		avg += r.Nanoseconds()
	}
	avg /= int64(iterations)

	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	slog.Info(fmt.Sprintf("Fastest time: %d ns", results[0]))
	slog.Info(fmt.Sprintf("Slowest time: %d ns", results[iterations-1]))
	slog.Info(fmt.Sprintf("Median time: %d ns", results[iterations/2]))
	slog.Info(fmt.Sprintf("Average time: %d ns", avg))
}
