package main

import "fmt"

func main() {

	// ARRAYS
	var chesses [2]string // array of length 2

	chesses[0] = "Mariolles"
	chesses[1] = "Epoisses"

	// chesses[2] = "Test" // this will cause out of bound index

	// SLICES
	// - defined as a contigous segment of an underlying array

	sChesses := make([]string, 0)

	// sChesses[0] = "A"
	// sChesses[1] = "B"

	// sChesses[2] = "C"
	sChesses = append(sChesses, "A")
	sChesses = append(sChesses, "B")
	sChesses = append(sChesses, "C")
	sChesses = append(sChesses, "D")

	// or
	sChesses = append(sChesses, "E", "F", "G")

	fmt.Println(sChesses)

	// Delete element
	sChesses = append(sChesses[:1], sChesses[1+1:]...)
	fmt.Println(sChesses)

	// copy one slice to another
	smellyChesses := make([]string, 6)
	copy(smellyChesses, sChesses)
	fmt.Println(smellyChesses)

	// Working with MAPS
	// - is an unordered group of elements that is accessed by a key rather than an index
	var players = make(map[string]int)

	players["anjo"] = 100
	players["mayflor"] = 10
	players["anya"] = 95
	players["intruder"] = 50
	players["intruder"] = 40

	fmt.Println(players)
	fmt.Println(players["anjo"])

	// delete map element
	delete(players, "intruder")
	fmt.Println(players)
}
