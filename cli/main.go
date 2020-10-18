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

	flag.Parse()
	fmt.Println("value of s:", *s)
}
