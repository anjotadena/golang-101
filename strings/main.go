package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// \a U+0007 alert or bell
// \b U+0008 backspace
// \f U+000C form feed
// \n U+000A line feed or newline
// \r U+000D carriage return
// \t U+0009 horizontal tab
// \v U+000b vertical tab
// \\ U+005c backslash
// \' U+0027 single quote (valid escape only with rune literals)
// \" U+0022 double quote (valid escape only with string literals)

// only type string can be concatenated
func main() {
	s := "Can you hear me?"
	s += "\nHear me screamin'?"

	fmt.Println(s)

	i := 1
	e := "egg"

	intToString := strconv.Itoa(i)
	breakfast := intToString + " " + e
	fmt.Println(breakfast)

	// cancatenating strings with a buffer
	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}

	fmt.Println(buffer.String())

	// Searching for substring
	fmt.Println(strings.Index("surface", "face"))
	fmt.Println(strings.Index("moon", "aer"))

	// Trimming and Trailing Whitespace
	fmt.Println(strings.TrimSpace("          I don't need all this space          "))
	fmt.Println("TEST")
}
