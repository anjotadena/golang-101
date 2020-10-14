package main

import "fmt"

func main() {
	var a = "Initial"

	fmt.Println(a)

	var b, c int = 1, 2

	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var arr [4]string
	fmt.Println("ARRAY")
	fmt.Println(arr[0])

	// fmt.Println(arr[11]) # out of boulds for 4-element array
}
