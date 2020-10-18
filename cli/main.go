package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Accessing cli arguments
	for i, arg := range os.Args {
		fmt.Println("argument", i, "is", arg)
	}

	// Parsing command line flags
	s := flag.String("s", "Hello, World!", "String help text")

	i := flag.Int("i", 1, "Int help text")
	b := flag.Bool("b", false, "Bool help text")

	flag.Parse()
	fmt.Println("value of s:", *s)
	fmt.Println("value of i:", *i)
	fmt.Println("value of b:", *b)
}
