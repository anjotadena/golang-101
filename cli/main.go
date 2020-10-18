package main

import (
	"fmt"
	"os"
)

func main() {
	// Accessing cli arguments
	for i, arg := range os.Args {
		fmt.Println("argument", i, "is", arg)
	}
}
