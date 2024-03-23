package bubblesort

import "github.com/danitrod/sorting-algorithms/internal"

type BubbleSort[T any] struct{}

func New[T any]() internal.SortingAlgorithm[T] {
	return BubbleSort[T]{}
}

func (b BubbleSort[T]) String() string {
	return "Bubble Sort"
}

func (b BubbleSort[T]) Sort(arr []T, compare func(a, b T) bool) {
	n := len(arr)
	for n > 0 {
		newN := 0
		for i := 0; i < n-1; i++ {
			if compare(arr[i+1], arr[i]) {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				newN = i + 1
			}
		}
		n = newN
	}
}
