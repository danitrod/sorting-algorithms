package main

import (
	"fmt"

	"github.com/danitrod/sorting-algorithms/internal/config"
)

func compareInt(a, b int) bool {
	return a > b
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic("error: " + err.Error())
	}

	array := []int{4, 3, 2, 1, 5, 6, 7, 8, 9, 10}
	cfg.Algorithm.Sort(array, compareInt)
	fmt.Println(array)
}
