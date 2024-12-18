package main

import "fmt"

// normal function//=============================

// func main() {
// 	s, i := myFunction("anubhav")
// 	fmt.Println(s, i)
// }
// func myFunction(p1 string) (s string, i int) {
// 	s = fmt.Sprintf("%s function", p1)
// 	i = 10

// 	return
// }

//closer//=========================================

// package main

// import "fmt"

// func main() {

// 	a := closer()
// 	fmt.Println(a(5))
// }

// func closer() func(int) int {

// 	return func(c int) int {
// 		return c
// 	}
// }

// variadic function//=============================================

func main() {

	fmt.Println(varia(1, 2, 3, 4))

	// for slice

	slice := []int{1, 2, 3}
	fmt.Println(varia(slice...)) // dot is in right in slice

}

// ...interface{} is any type in go
func varia(nums ...int) int {
	total := 0
	for k, num := range nums {
		fmt.Println(k)
		total = total + num
	}
	return total

}
