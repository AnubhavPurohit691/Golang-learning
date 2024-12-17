package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice => dynamic array

	var nums []int // this is now nil like in other programming laguage it is null
	println(len(nums))
	println(nums == nil)

	slice := make([]int, 2, 5) // here we make a slice using make and 2 is its length of starting array and 5 is  initial capacity
	fmt.Println(slice)         // this is not nil

	fmt.Println(append(slice, 5))
	fmt.Println(cap(slice)) // this will give capacity of slice

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2[0:2]) //will exclude 3 it will print 0,1 index only

	// slices pakage
	fmt.Println(slices.Equal(slice, slice2))
}
