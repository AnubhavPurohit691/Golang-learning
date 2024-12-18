package main

import (
	"fmt"
	"time"
)

// structs are like classes
type order struct {
	id        int
	name      string
	createdAt time.Time // nano secound precision
	status    string
}

// constructor is use to initialize
func New(id int, name string, status string) *order {

	Order := order{
		id:     id,
		name:   name,
		status: status,
	}
	return &Order
}

// create method in structs

// pointer is used to modify value but if you want to get only value don't use pointer
func (o *order) changestatus(status string) {
	o.status = status
}

func main() {
	var Orders order = order{
		id:     1,
		name:   "books",
		status: "received",
	}
	Orders.changestatus("confirm")
	Orders.createdAt = time.Now()

	fmt.Println(Orders.name, Orders.status)

	fmt.Println(New(1, "clock", "received"))

	// one time stuct
	language := struct {
		id   string
		name string
	}{
		id:   "anubhav",
		name: "yo",
	}

	fmt.Println(language.id, language.name)
}
