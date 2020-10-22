package main

import "fmt"

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := 33
	low := 0
	high := len(items) - 1

	for low <= high { // 0 - 1
		mid := (low + high) / 2 // 2

		if items[mid] < s { //
			low = mid + 1 // 1 + 1 = 2
		} else {
			high = mid - 1 // 2 - 1 = 1
		}
	}

	if low == len(items) || items[low] != s {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found")
	}

}
