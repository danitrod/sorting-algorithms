bubble:
	go run cmd/main.go bubblesort

heap:
	go run cmd/main.go heapsort

quick:
	go run cmd/main.go quicksort

quick-%:
	go run cmd/main.go quicksort $(subst quick-,,$@)

test:
	go test ./...
