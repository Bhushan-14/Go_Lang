package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("Your age is:", age)

	if age < 18 {
		fmt.Println("You are a minor.")
	}else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")	
	}else {
		fmt.Println("You are a senior citizen.")
	}
}
