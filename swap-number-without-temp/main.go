package main

import "fmt"

func main() {
	fmt.Print("Enter first number: ")
	var firstNum int
	fmt.Scan(&firstNum)

	fmt.Print("Enter second number: ")
	var secondNum int
	fmt.Scan(&secondNum)

	// SWAP Numbers
	firstNum -= secondNum //
	secondNum = firstNum + secondNum
	firstNum = secondNum - firstNum

	fmt.Println("First number: ", firstNum)
	fmt.Println("Second number: ", secondNum)
}
