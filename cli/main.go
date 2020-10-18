package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Customizing help
	flag.Usage = func() {
		usageText := `Usage example04 [OPTION]
			An example of customizing usage output

			-s, --s		example string argument, default: String help text

			-i, --i		example integer argument, default: Int help text

			-b, --b		example boolean argument, default: Bool help text`

		fmt.Fprintf(os.Stderr, "%s\n", usageText)
	}

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
