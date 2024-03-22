package quicksort_test

import (
	"testing"

	"github.com/danitrod/sorting-algorithms/internal/quicksort"
	"github.com/stretchr/testify/assert"
)

func compareInt(a, b int) bool {
	return a > b
}

func TestQuickSort(t *testing.T) {
	t.Run("Should sort array with quick sort, sequentially", func(t *testing.T) {
		arr := []int{5, 1, 3, 90, 8, 2, 1, -5, 7, 9, 23}

		quicksort.New[int](0).Sort(arr, compareInt)

		for i := 0; i < len(arr)-1; i++ {
			assert.GreaterOrEqual(t, arr[i], arr[i+1])
		}
	})

	t.Run("Should sort array with quick sort, with 8 threads", func(t *testing.T) {
		arr := []int{5, 1, 3, 90, 8, 2, 1, -5, 7, 9, 23}
		for i := 0; i < 1000; i++ {
			arr = append(arr, i)
		}

		quicksort.New[int](8).Sort(arr, compareInt)

		for i := 0; i < len(arr)-1; i++ {
			assert.GreaterOrEqual(t, arr[i], arr[i+1])
		}
	})
}
