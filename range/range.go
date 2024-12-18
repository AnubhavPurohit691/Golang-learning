package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"age": 3, "age2": 4}

	for key, value := range m {
		fmt.Println(value, key)
	}

	for k, v := range "golang" {
		fmt.Println(k, v)
	}

}
