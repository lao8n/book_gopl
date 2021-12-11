// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//go:build ignore
// +build ignore

// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
//
// The "+build ignore" tag (see p.295) excludes this file from the
// thumbnail package, but it can be compiled as a command and run like
// this:
//
// Run with:
//   $ go run $GOPATH/src/gopl.io/ch8/thumbnail/main.go
//   foo.jpeg
//   ^D
//
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopl.io/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}

// 1. add goroutines
// 2. but problem is need to wait for them so add channel to communicate over
// 3. but want errors as well but need to make sure use buffered channel (or drain)
// 4. but want unknown number of file names so communicate via channel
// 5. but need to keep track of loop operations so use a sync.WaitGroup
