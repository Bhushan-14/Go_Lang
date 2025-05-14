package main

import (
	"fmt"
	"time"
)

func switch_stmt() {
	// var i int
	// fmt.Print("Enter day number (1 = Sunday, 7 = Saturday): ")
	// fmt.Scanln(&i)

	// switch i {
	// case 1:
	// 	fmt.Println("Sunday")
	// case 2:
	// 	fmt.Println("Monday")
	// case 3:
	// 	fmt.Println("Tuesday")
	// case 4:
	// 	fmt.Println("Wednesday")
	// case 5:
	// 	fmt.Println("Thursday")
	// case 6:
	// 	fmt.Println("Friday")
	// case 7:
	// 	fmt.Println("Saturday")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// Check actual day today from system date
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
}
