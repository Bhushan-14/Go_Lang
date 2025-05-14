package main

import "fmt"


// names := "Sagar" // Short variable declaration out of function is not allowed
// var names = "Sagar" // This is a valid declaration but not recommended
func main() {
	// var f_name string = "Sagar"
	f_name := "Sagar" // Short variable declaration

	// f_name := "Mayur" // This will cause an error because f_name is already declared

	var l_name = "Patil"
	var age = 21
	l_name = "Patel"
	age = 23
	// age = 23.4 // This will cause an error because age is declared as an integer
	fmt.Println(f_name)
	fmt.Println(l_name)
	fmt.Println(age)
}
