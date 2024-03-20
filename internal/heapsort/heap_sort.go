package heapsort

import "github.com/danitrod/sorting-algorithms/internal"

type HeapSort[T any] struct{}

func New[T any]() internal.SortingAlgorithm[T] {
	return HeapSort[T]{}
}

func (h HeapSort[T]) Sort(arr []T, compare func(a, b T) bool) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, len(arr), i, compare)
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0, compare)
	}
}

func heapify[T any](arr []T, n, i int, compare func(a, b T) bool) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && compare(arr[left], arr[largest]) {
		largest = left
	}

	if right < n && compare(arr[right], arr[largest]) {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest, compare)
	}
}
