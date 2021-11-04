// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Need to modify dup2 to print the names of all files in which each
// duplicated line occurs
// We haven't learnt structs yet, so maybe best solution is a double
// map?
// Rather than a double nested map i will make a separate map for
// storing the file name from the line (which we assume is unique anyway)

// note no testing
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s", filenames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, filenames map[string]string) {
	input := bufio.NewScanner(f)
	// read line by line
	for input.Scan() {
		filenames[input.Text()] = filename
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
