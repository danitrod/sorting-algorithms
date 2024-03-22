package config

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/danitrod/sorting-algorithms/internal"
	"github.com/danitrod/sorting-algorithms/internal/bubblesort"
	"github.com/danitrod/sorting-algorithms/internal/heapsort"
	"github.com/danitrod/sorting-algorithms/internal/quicksort"
)

const (
	flagBubbleSort = "bubblesort"
	flagHeapSort   = "heapsort"
	flagQuickSort  = "quicksort"
)

var errUnknownAlgorithm = errors.New("unknown algorithm")

type config[T any] struct {
	Algorithm internal.SortingAlgorithm[T]
}

func NewConfig[T any]() (config[T], error) {
	flag.Parse()

	var algo internal.SortingAlgorithm[T]
	switch flag.Arg(0) {
	case string(flagBubbleSort):
		algo = bubblesort.New[T]()
	case string(flagHeapSort):
		algo = heapsort.New[T]()
	case string(flagQuickSort):
		numThreads := flag.Arg(1)
		n, _ := strconv.Atoi(numThreads)
		slog.Info(fmt.Sprintf("Using number of threads: %d", n))
		algo = quicksort.New[T](n)
	default:
		return config[T]{}, errUnknownAlgorithm
	}

	slog.Info(fmt.Sprintf("Chosen algorithm: %s", flag.Arg(0)))
	return config[T]{Algorithm: algo}, nil
}
