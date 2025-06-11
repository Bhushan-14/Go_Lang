package main

import "fmt"

// Slices are dynamic arrays
// Slices are more flexible than arrays
// Slices are reference types
// provide useful built-in functions

func main() {

	//only decalre slice, not initialized
	//uninitialized slice has nil values
	// var nums []int

	// fmt.Println("Slice: ", nums)

	// //length of slice
	// fmt.Println("Length of slice: ", len(nums))

	// fmt.Println(nums == nil) //check if slice is nil
	//create slice using make(types, 0) function

	var nums = make([]int, 2, 5)
	fmt.Println(cap(nums)) //capacity of maximum elements can fit
	fmt.Println("Slice: ", nums)
	fmt.Println(nums == nil) //check if slice is nil

	nums = append(nums, 34)
	fmt.Println("Slice after appending: ", nums)
}
