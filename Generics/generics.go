package main

import "fmt"

func status[T any](status T) T {
	return status
}

type Order[T comparable] struct {
	gateway []T
}

func main() {
	fmt.Println(status("hello"))

	gate := Order[string]{
		gateway: []string{"1", "2"},
	}

	fmt.Println(gate)
}
