package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return n
	}

	return n * factorial(n-1)
}

func main() {
	var n int

	fmt.Print("Enter a positive integer (0 - 50): ")
	fmt.Scan(&n)

	if n <= 0 || n >= 50 {
		fmt.Println("Invalid number selection. Please provide 0 - 50 integers.")

		return
	}

	fmt.Printf("Factorial of %d is %d\n", n, factorial(n))
}
