package heapsort_test

import (
	"testing"

	"github.com/danitrod/sorting-algorithms/internal/heapsort"
	"github.com/stretchr/testify/assert"
)

func compareInt(a, b int) bool {
	return a > b
}

func TestHeapSort(t *testing.T) {
	t.Run("Should sort array with heap sort", func(t *testing.T) {
		arr := []int{5, 1, 3, 90, 8, 2, 1, -5, 7, 9, 23}

		heapsort.New[int]().Sort(arr, compareInt)

		for i := 0; i < len(arr)-1; i++ {
			assert.LessOrEqual(t, arr[i], arr[i+1])
		}
	})
}
