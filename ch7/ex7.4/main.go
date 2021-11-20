package main

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

type StringReader string

// copies bytes to string
func (s *StringReader) Read(p []byte) (n int, err error) {
	copy(p, *s)
	return len(*s), io.EOF
}

// return concrete implementation that satisfies io.Reader interface
func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr // cannot do &StringReader(s)
}

func main() {
	_, err := html.Parse(NewReader("<p>Hello, world!</p>"))
	if err != nil {
		log.Fatalf("NewReader failed")
	}
	// need html printer
}

/*
type Reader interface {
	Read(p []byte)(n int, err error)
}
*/
