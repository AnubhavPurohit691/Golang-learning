package main

import "fmt"

type Object struct {
	id   string
	name string
}
type user struct {
	yo          string
	newcustomer Object
}

func main() {

	customer := Object{
		id:   "1",
		name: "anubhav",
	}

	userId := user{
		yo:          "anubhav",
		newcustomer: customer,
	}

	fmt.Println(userId)
}
