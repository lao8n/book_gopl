// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import "fmt"

//!+
func max(vals ...int) int {
	max := 0
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

//!-

func main() {
	//!+main
	fmt.Println(max())           //  "0"
	fmt.Println(max(3))          //  "3"
	fmt.Println(max(3, 2, 1, 4)) //  "4"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(max(values...)) // "4"
	//!-slice
}
