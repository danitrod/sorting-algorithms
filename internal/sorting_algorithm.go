package internal

import "fmt"

type SortingAlgorithm[T any] interface {
	fmt.Stringer
	Sort(arr []T, compare func(a, b T) bool)
}
