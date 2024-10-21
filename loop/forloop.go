package main

import "fmt"

// there is only for keyword is there in go lang
func main() {
	// to work like while loop this is the syntax
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// infinite loop
	// for {
	// 	fmt.Println("hello")
	// } this is infinite loop

	for i := 0; i <= 5; i++ {
		fmt.Println("anubhav", i)
	}
	for range 5 {
		fmt.Println("hello")
	}
	for i := range 7 {
		fmt.Println(i)
	}

	if age := 15; age > 12 {
		fmt.Println("age is more than 12")
	}
}
