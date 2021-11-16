// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main
// ./fetch https://golang.org | go run main.go

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
// func visit(links []string, n *html.Node) []string {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				links = append(links, a.Val)
// 			}
// 		}
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		links = visit(links, c)
// 	}
// 	return links
// }

func visit(n *html.Node) []string {
	var self []string
	var firstChildRelatives []string
	var siblingRelatives []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				self = append(self, a.Val)
			}
		}
	}
	// two recursive cases 1. first child 2. sibling
	if n.FirstChild != nil {
		firstChildRelatives = visit(n.FirstChild)
	}
	if n.NextSibling != nil {
		siblingRelatives = visit(n.NextSibling)
	}
	return append(self, append(firstChildRelatives, siblingRelatives...)...)
	// cleaner implementation found online
	// func visit(links []string, n *html.Node) []string {
	// links = visit(links, n.FirstChild)
	// links = visit(links, n.NextSibling)
	// return links
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
