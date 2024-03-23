package bencher

import (
	"fmt"
	"log/slog"
	"math"
	"os"
	"sort"
	"time"

	"github.com/gocarina/gocsv"
)

const outputFile = "data/results.csv"

type Bencher struct {
	algorithm  string
	iterations int
}

func New(algorithm string, iterations int) *Bencher {
	return &Bencher{
		algorithm:  algorithm,
		iterations: iterations,
	}
}

type BenchResults struct {
	Algorithm         string `csv:"algorithm"`
	Average           int64  `csv:"average"`
	Fastest           int64  `csv:"fastest"`
	Iterations        int    `csv:"iterations"`
	Median            int64  `csv:"median"`
	Slowest           int64  `csv:"slowest"`
	StandardDeviation int64  `csv:"standard_deviation"`
}

func (b Bencher) Bench(fn func()) (BenchResults, error) {
	slog.Info(fmt.Sprintf("Running experiment with algorithm %s and %d iterations", b.algorithm, b.iterations))

	results := make([]int64, 0, b.iterations)
	for i := 0; i < b.iterations; i++ {
		start := time.Now()
		fn()
		results = append(results, time.Since(start).Microseconds())
	}

	avg := int64(0)
	for _, r := range results {
		avg += r
	}
	avg /= int64(b.iterations)

	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	stdDev := b.getStandardDeviation(results, avg)
	benchResults := BenchResults{
		Algorithm:         b.algorithm,
		Average:           avg,
		Fastest:           results[0],
		Iterations:        b.iterations,
		Median:            results[b.iterations/2],
		Slowest:           results[b.iterations-1],
		StandardDeviation: stdDev,
	}

	return benchResults, b.saveToFile(benchResults)
}

func (b Bencher) getStandardDeviation(results []int64, avg int64) int64 {
	sumOfDistanceSquares := int64(0)
	for _, r := range results {
		sumOfDistanceSquares += (r - avg) * (r - avg)
	}

	avgSquareDistance := sumOfDistanceSquares / int64(len(results))
	sd := math.Sqrt(float64(avgSquareDistance))

	return int64(sd)
}

func (b Bencher) saveToFile(results BenchResults) error {
	resultsFile, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer resultsFile.Close()

	fileResults := []*BenchResults{}
	err = gocsv.UnmarshalFile(resultsFile, &fileResults)
	if err != nil {
		return err
	}

	fileResults = append(fileResults, &results)
	csvContent, err := gocsv.MarshalString(&fileResults)
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, []byte(csvContent), os.ModePerm)
}
