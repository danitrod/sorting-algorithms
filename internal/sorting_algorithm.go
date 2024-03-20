package internal

type SortingAlgorithm[T any] interface {
	Sort(arr []T, compare func(a, b T) bool)
}
