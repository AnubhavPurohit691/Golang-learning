package main

import "fmt"

func main() {
	num := 1
	fmt.Println(num, &num) // this is &num address of num
	fmt.Println(pointer(&num))
}
func pointer(n *int) int {
	*n = 5
	return *n
}
