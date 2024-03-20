package config

import (
	"errors"
	"flag"

	"github.com/danitrod/sorting-algorithms/internal"
	"github.com/danitrod/sorting-algorithms/internal/bubblesort"
	"github.com/danitrod/sorting-algorithms/internal/heapsort"
)

const (
	flagBubbleSort = "bubblesort"
	flagHeapSort   = "heapsort"
)

var errUnknownAlgorithm = errors.New("unknown algorithm")

type config struct {
	Algorithm internal.SortingAlgorithm[int]
}

func NewConfig() (config, error) {
	flag.Parse()

	var algo internal.SortingAlgorithm[int]
	switch flag.Arg(0) {
	case string(flagBubbleSort):
		algo = bubblesort.New[int]()
	case string(flagHeapSort):
		algo = heapsort.New[int]()
	default:
		return config{}, errUnknownAlgorithm
	}

	return config{Algorithm: algo}, nil
}
