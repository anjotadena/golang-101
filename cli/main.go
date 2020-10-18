package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func flagUsage() {
	usageText := `SubCommands is an example of CLI tool.
		Usage:
			SubCommand command [arguments]
			The commands are:
				uppercase uppercase a string
				lowercase lowercase a string
			Use "SubCommand [command] --help" for more information about a command.
	`
	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func SubCommands() {
	flag.Usage = flagUsage
	uppercase := flag.NewFlagSet("uppercase", flag.ExitOnError)
	lowercase := flag.NewFlagSet("lowercase", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "uppercase":
		s := uppercase.String("s", "", "A string of text to be uppercased")
		uppercase.Parse(os.Args[2:])
		fmt.Println(strings.ToUpper(*s))
	case "lowercase":
		s := lowercase.String("s", "", "A string of text to be lowercased")
		lowercase.Parse(os.Args[2:])
		fmt.Println(strings.ToLower(*s))
	default:
		flag.Usage()
	}
}

// POSIX Compliance
// https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html

func main() {

	if true {
		SubCommands()
		return
	}

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
