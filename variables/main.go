package main

import (
	"fmt"
	"reflect"
	"strconv"
)

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

	// CHECK TYPE OF VARIABLES
	var s string = "string"
	var i int = 10
	var flt float32 = 1.2

	fmt.Println(reflect.TypeOf(s))   // > string
	fmt.Println(reflect.TypeOf(i))   // > int
	fmt.Println(reflect.TypeOf(flt)) // > float32

	s = "true"

	pb, err := strconv.ParseBool(s)

	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Println(pb)                           // true
	fmt.Println("IS BOOLEAN: ", pb)           // true
	fmt.Println("Type: ", reflect.TypeOf(pb)) // bool

	// Pointers
	x := "Hello, World!"

	fmt.Println(x) // this prints a value
	// base 16format
	fmt.Println(&x) // this prints a variable address in which the location of memory stored

	y := 1
	fmt.Println(&y) // this will display memory address

	// using pointer
	showMemoryAddress(y)
	// the example above will display 2 different memory address

	// to avoid another memory to used
	showMemoryAddressPointer(&y)

}

func showMemoryAddress(x int) {
	fmt.Println(&x)
}

func showMemoryAddressPointer(x *int) {
	fmt.Println(x)  // display parameter memory address
	fmt.Println(*x) // this will display the value
}
