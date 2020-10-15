package main

import "fmt"

// control flow evaluates condition statements

func main() {
	b := true

	if b {
		fmt.Println("b is true")
	}

	a := false

	// IF ELSE
	if a {
		fmt.Println("a is true")
	} else {
		fmt.Println("a is not true")
	}

	// IF ELSE IF

	c := 3

	if c == 3 {
		fmt.Println("c is 3")
	} else if c == 2 {
		fmt.Println("c is 2")
	} else {
		fmt.Println("c is ", c)
	}

	// comparison operators
	// == equal toa
	// != not equal ta
	// < less thaa
	// <= less than or equal to
	// > greater than
	// >= greater than or equal to

	// Artihmentic operators
	// + Sum (addition)
	// - Difference (substraction)
	// * Product (multiplication)
	// / Quotient (division)
	// % Remainder (modulus)

	// Logical operators
	// && AND
	// || OR
	// ! NOT

	// SWITCH
	d := 1

	switch d {
	case 1:
		//
		fmt.Println("d is 1")
	case 2:
		//
		fmt.Println("d is 2")
	default:
		//
		fmt.Println("d is ", d)
	}

	// LOOP
	i := 0

	for i < 10 {
		i++

		fmt.Println("i is ", i)
	}

	// Init statement
	for i := 0; i < 10; i++ {
		fmt.Println("i is ", i)
	}
}
