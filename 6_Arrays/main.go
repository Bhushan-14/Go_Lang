package main

import "fmt"

func main() {
	var nums [4]int
	fmt.Println("Length of array: ", len(nums))

	// nums[0] = 1
	// nums[1] = 2
	//if we not add any value to the array, it will be initialized with 0
	// nums[3] = 4

	// fmt.Println("Enter 4 integers:")
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Printf("Element %d: ", i)
	// 	fmt.Scanln(&nums[i])
	// }

	// fmt.Println("Array: ", nums)

	// fmt.Println("Array Elements:")
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// var vals [5]string = [5]string{"Sagar", "Mayur", "Hitesh", "Hrishi", "Ronit"}
	// fmt.Println("Array with values: ", vals)

	// default values
	// 	int -> 0
	// 	string -> ""
	// 	bool -> false

	// names := [4]string{"Sagar", "Mayur", "Hitesh", "Hrishi"}
	// fmt.Println("Array with values: ", names)

	//2D array
	var matrix [2][2]int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Enter element [%d][%d]: ", i, j)
			fmt.Scanln(&matrix[i][j])
		}
	}
}


/*
	- fixed size, that is predictable
	- memory optimization
	- Constant time access
*/