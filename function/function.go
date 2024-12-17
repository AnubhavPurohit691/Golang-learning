package main

import "fmt"

func main() {
	s, i := myFunction("anubhav")
	fmt.Println(s, i)
}
func myFunction(p1 string) (s string, i int) {
	s = fmt.Sprintf("%s function", p1)
	i = 10

	return
}
