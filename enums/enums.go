package main

import "fmt"

type Orderstatus int

const (
	received Orderstatus = iota
	confirm              // it will increment bcoz of iota
)

func status(stat Orderstatus) {
	fmt.Println(stat)
}

func main() {
	status(received)
}
