// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 96.

// Dedup prints only one instance of each line; duplicates are removed.
package main

import "fmt"

func main() {
	s := []string{"A", "A", "B", "B", "A", "C"}
	fmt.Println(removeDuplicates(s))
}

//!+
func removeDuplicates(input []string) []string {
	var cLast string
	iDedup := 0
	for i := 0; i < len(input); {
		c := input[i]
		// no duplicate
		if c != cLast {
			input[iDedup], cLast = c, c
			iDedup++
		}
		i++
	}
	return input[:iDedup]
}

//!-
