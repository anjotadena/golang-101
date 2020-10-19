package main

import (
	parent "family/father"
	child "family/father/son"

	"fmt"
)

func main() {
	f := new(parent.Father)

	fmt.Println(f.Data("Anjo Tadena"))

	c := new(child.Son)
	fmt.Println(c.Data("Anya Xyra"))
}
