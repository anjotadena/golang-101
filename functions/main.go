package main

import "fmt"

// Functions - a functions takes input and returns an output.

// Function signature provides information
// takes 2 parameters and returns int value
func addUp(x int, y int) int {
	return x + y
}

// returns multiple value
func getRank() (int, string) {
	i := 2
	s := "Ging Freecs"

	return i, s
}

// variadic function
// - are defined as a function of indefinite arity
// - this means that they accept a variable number of arguments
// - syntax (...)
func sumNumbers(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// functions named return values
func sayHi() (x, y string, z int) {
	x = "hello"
	y = "world"
	z = 2
	return
}

// 5 + 4 + 3 + 2 + 1
// recursive functions - a functions that call itself until a particular condition is met
func sumOfNumbers(x int) int {
	fmt.Println(x)
	if x <= 0 {
		return x
	}

	return x + sumOfNumbers(x-1)
}

// Passing a function as an argument
func anotherFunction(f func() string) string {
	return f()
}

func main() {
	fmt.Println(addUp(1, 2)) // calling function

	rank, name := getRank()

	fmt.Printf("Hunter Rank: %v", rank)
	fmt.Println()
	fmt.Printf("Hunter Name: %v", name)

	fmt.Println()
	fmt.Println("Total of sum numbers: ", sumNumbers(1, 2, 3, 4, 5))

	// using function named return values
	x, y, z := sayHi()

	fmt.Println("Name x: ", x)
	fmt.Println("Name y: ", y)
	fmt.Println("Name z: ", z)
	r := sumOfNumbers(20)

	fmt.Println(r)

	// Function as a value
	fnHelloWorld := func() string {
		return "Hello, World!"
	}

	fmt.Println(fnHelloWorld())

	// Call passing a function as an argument
	fmt.Println(anotherFunction(fnHelloWorld))
}
