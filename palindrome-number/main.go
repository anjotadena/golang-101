package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	result := 0
	var remainder int

	temp := num

	for {
		remainder = num % 10
		result = result*10 + remainder
		num /= 10

		if num == 0 {
			break
		}
	}

	if temp == result {
		fmt.Printf("%d is a palindrome\n", temp)
		return
	}

	fmt.Printf("%d is not a palindrome\n", temp)
}
