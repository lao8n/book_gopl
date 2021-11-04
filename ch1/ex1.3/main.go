package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("efficient")
	efficient(os.Args)
	fmt.Println("inefficient")
	efficient(os.Args)
}

func efficient(args []string) {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.100fs elapsed\n", time.Since(start).Seconds())
}

func inefficient(args []string) {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.100fs elapsed\n", time.Since(start).Seconds())
}
