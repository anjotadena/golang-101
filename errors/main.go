package main

import (
	"fmt"
	"io/ioutil"
)

// Go offers an interesting approach to errors

func main() {
	//
	file, err := ioutil.ReadFile("notes.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s\n", file)

	// err := errors.New("Something went wrong")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Formatting errors
	name, role := "Richard Jupp", "Drummer"

	err = fmt.Errorf("The %v %v quit", role, name)

	if err != nil {
		fmt.Println(err)
	}

	n, errHalf := Half(20)

	if errHalf != nil {
		fmt.Println(errHalf)
		return
	}
	fmt.Println(n)

	panic("Oh no. I can do no more. Goodbye.")
}

// Errors return from a function
func Half(n int) (int, error) {
	if n%2 != 0 {
		return -1, fmt.Errorf("Cannot half %v", n)
	}
	return n / 2, nil
}
