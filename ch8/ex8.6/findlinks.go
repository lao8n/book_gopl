// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 240.

// Crawl1 crawls web links starting with the command-line arguments.
//
// This version quickly exhausts available file descriptors
// due to excessive concurrent calls to links.Extract.
//
// Also, it never terminates because the worklist is never closed.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

// to implement depth we have two options
// 1. have associated with each url a depth int and don't call crawl once depth==3
// 2. depth is a property of a list, so don't add to worklist unless depth is correct -> but this is crawl as well

// rather than sotoring depth with worklist instead we make seen do two jobs 1. track seen and 2. track depth

// key decision rule: only add children links (i.e. crawl) if your depth is less than < maxDepth
// question: how know your depth? either 1. store in worklist 2. store in a separate dict, but only nkow depth of children when know parent's depth
// once add children to worklist the depth is lost.

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

type LinkDepth struct {
	Url   string
	Depth int
}

//!+main
func main() {
	maxDepth := *(flag.Int("depth", 3, "depth of url"))
	flag.Parse()
	fmt.Println(maxDepth)
	worklist := make(chan []LinkDepth)
	// Start with the command-line arguments.
	go func() {
		var initialUrls []LinkDepth
		for _, url := range os.Args[1:] {
			initialUrls = append(initialUrls, LinkDepth{url, 0})
		}
		worklist <- initialUrls
	}() // needed to avoid deadlock in which both the main goroutine and crawler goroutine attempt to send to each other whilst neither is receiving

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, linkStruct := range list {
			if !seen[linkStruct.Url] && linkStruct.Depth < maxDepth {
				seen[linkStruct.Url] = true
				go func(link string, depth int) { // take link as explicit parameter to avoid variable capture
					var crawledUrls []LinkDepth
					for _, url := range crawl(link) {
						crawledUrls = append(crawledUrls, LinkDepth{url, depth + 1})
					}
					worklist <- crawledUrls
				}(linkStruct.Url, linkStruct.Depth)
			}
		}
	}
}

//!-main

/*
//!+output
$ go build gopl.io/ch8/crawl1
$ ./crawl1 http://gopl.io/
http://gopl.io/
https://golang.org/help/

https://golang.org/doc/
https://golang.org/blog/
...
2015/07/15 18:22:12 Get ...: dial tcp: lookup blog.golang.org: no such host
2015/07/15 18:22:12 Get ...: dial tcp 23.21.222.120:443: socket:
                                                        too many open files
...
//!-output
*/
