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
}
