package main

import (
	"fmt"
	"maps"
)

// map are like obj , dict, hash
func main() {

	m := make(map[string]string) //key value pair

	m["name"] = "anubhav"
	map4 := m
	fmt.Println(m["name"])
	fmt.Println(map4)
	fmt.Println(len(m["name"]))
	// delete(m, "name")
	// clear(m)
	mapf := map[string]int{"age1": 40, "age2": 30}
	map3 := map[string]int{"age1": 40, "age2": 30}

	value, ok := mapf["age1"]

	if ok {
		fmt.Println("exist", value)

	} else {
		fmt.Println("not found")
	}

	fmt.Println(maps.Equal(map3, mapf))
}
