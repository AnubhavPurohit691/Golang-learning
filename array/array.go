package main

import "fmt"

func main() {
	var array [3]int
	array[0] = 4
	fmt.Println(array)

	array1 := [2][2]int{{1, 2}, {2, 3}}
	fmt.Println(array1)

	wow := [2]int{2, 3}
	fmt.Println(wow)
	fmt.Println(len(wow))

	// this is used for fixed and memory optimation and constant time access
}
