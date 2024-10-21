package main

import (
	"fmt"
)

func main() {
	// switch time.Now().Weekday() {
	// case time.Monday:
	// 	fmt.Println("it's monday")
	// case time.Saturday:
	// 	fmt.Println("it's saturday")
	// default:
	// 	fmt.Println("it")
	// }
	whoamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("interger")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("default", t)
		}
	}
	whoamI("hello")
}
