// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.
package values

import "fmt"

func main() {
	// concatination
	fmt.Println("Go" + "Lang")

	fmt.Println("1 + 1 = ", 1+1)

	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
