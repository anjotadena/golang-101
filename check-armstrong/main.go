package main

import "fmt"

// A positive integer is called an armstrong number of order n if
// the sum of cubes of each digits is equal to the number itself.
func main() {
	number := 0
	remainder := 0
	result := 0

	for number == 0 {
		fmt.Print("Enter any three digit number: ")
		fmt.Scan(&number)
	}

	tempNumber := number

	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder

		// cut number
		tempNumber /= 10

		if tempNumber == 0 {
			break // exit loop
		}
	}

	if result == number {
		fmt.Printf("%d is an armstrong number.", number)
	} else {
		fmt.Printf("%d is not an armstrong number.", number)
	}
}
