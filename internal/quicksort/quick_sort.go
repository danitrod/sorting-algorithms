package quicksort

import (
	"fmt"
	"sync"

	"github.com/danitrod/sorting-algorithms/internal"
)

type QuickSort[T any] struct {
	maxThreads  int
	usedThreads int
}

func New[T any](maxThreads int) internal.SortingAlgorithm[T] {
	if maxThreads%2 != 0 {
		panic("maxThreads must be an even number")
	}

	return QuickSort[T]{
		maxThreads:  maxThreads,
		usedThreads: 0,
	}
}

func (q QuickSort[T]) String() string {
	if q.maxThreads == 0 {
		return "Quick Sort"
	}

	return fmt.Sprintf("Parallel Quick Sort (%d threads)", q.maxThreads)
}

func (q QuickSort[T]) Sort(arr []T, compare func(a, b T) bool) {
	q.quicksort(arr, 0, len(arr)-1, compare)
}

func (q QuickSort[T]) quicksort(arr []T, lo, hi int, compare func(a, b T) bool) {
	if lo >= hi || lo < 0 {
		return
	}

	p := partition(arr, lo, hi, compare)

	if q.usedThreads >= q.maxThreads {
		q.quicksort(arr, lo, p-1, compare)
		q.quicksort(arr, p+1, hi, compare)
		return
	}

	q.usedThreads += 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		q.quicksort(arr, lo, p-1, compare)
		wg.Done()
	}()
	go func() {
		q.quicksort(arr, p+1, hi, compare)
		wg.Done()
	}()

	wg.Wait()
}

func partition[T any](arr []T, lo, hi int, compare func(a, b T) bool) int {
	pivot := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if compare(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[hi], arr[i] = arr[i], arr[hi]
	return i
}
