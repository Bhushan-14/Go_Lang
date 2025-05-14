package main

import (
	"fmt"
	// "path"
)

// Go only supports for loop
func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// //While loop using for loop
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i == 5{
	// 		continue // skip the rest of the loop
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i == 5{
	// 		break // exit the loop
	// 	}
	// 	fmt.Println(i)
	// }

	for i := range 34 {
		if i == 7 {
			fmt.Println("i is 7")
		}
	}
}
