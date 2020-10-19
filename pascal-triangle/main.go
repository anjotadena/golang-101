package main

import (
	"fmt"
)

func main() {
	var rows int
	var temp int = 1

	fmt.Println(rows)
	// _, err = fmt.Scan(&rows)

	for rows <= 0 {
		fmt.Print("Enter number of rows: ")
		fmt.Scan(&rows)
		fmt.Println()
	}

	// 5
	for i := 0; i < rows; i++ {
		//
		for j := 1; j <= rows-i; j++ { // 4
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1 // default
			} else {
				temp = temp * (i - k + 1) / k
			}

			fmt.Printf(" %d", temp)
		}

		fmt.Println("")
	}
}
