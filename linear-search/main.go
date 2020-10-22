package main

import "fmt"

type Result struct {
	Index int
	Value int
}

// This technique pass over the list of elements,
// by using the index to move from the beginning of the list to the end.
// Each element is examined and if it does not match the search item, the next item is examined.
// By hopping from one item to its next, the list is passed over sequentially.

func main() {
	items := []int{5, 3, 2, 5, 1, 3, 4}
	s := 1
	var result Result
	fmt.Println(result)
	for i, item := range items {
		if item == s {
			result = Result{i, item}

			break
		}
	}

	if result.Value != s {
		fmt.Printf("%d doesn't exist in %d\n", s, items)
	} else {
		fmt.Printf("%d exist in %d. Found in index %d\n", s, items, result.Index)
	}
}
