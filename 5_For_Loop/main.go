package main

import "fmt"

//Go only supports for loop
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//While loop using for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}
