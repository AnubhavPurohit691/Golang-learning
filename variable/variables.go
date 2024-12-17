package main

import "fmt"

func main() {

	fmt.Println("helloworld")
	var hello int
	hello = 56
	// we can also write it like
	anubhav := 78
	fmt.Println(anubhav)
	println(hello)
	fmt.Printf("%v anubhav is ", anubhav)

	// printf is used to print type of value printf(%T)%T is imp to print type of value

	// sprintf is used to return formatted string and printf is used to print formattedstring
	// %v is used for default for any value like in c and %d is used for int , %f is for float ,%s string
	fmt.Printf("%.1f is float \n", 5.6)
	// .1 is used to make float visible upto tenth digit

	i := 42
	f := float64(i)
	u := uint(f)

	fmt.Printf("%T %T", f, u)
}

//   unsigned int is only positive integer
// we use go build to generate exe and to run do go run variable.go
