// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// run by
// go run clock -port 8010 &
// go run clockwall NewYork=localhost:8010

// or
// TZ=US/Eastern go run clock -port 8010 &
// TZ=Asia/Tokyo go run clock -port 8020 &
// go run clockwall NewYork=localhost:8010 Tokyo=localhost:8020

// clockwall = client, from connection -> print
// question: how supplement src io.Reader with some information about location etc?
// 1. can fmt.Print to dst but i want this to happen for every line of io.Reader
// 2. maybe can change src, i.e. add to io.Reader

// In order to run printArgs as a goroutine need to safely print to os.Stdout, apparently the way to do this is with
// logger logger := log.New(os.Stdout, "", 0)

// Problem ; scanner.Scan() doesn't seem to work when i set mustCopy to a goroutine
// This doesn't work not sure what i'm doing wrong

func main() {
	// logger := log.New(os.Stdout, "", 0)
	outputStrings := [2]string{}
	for i, arg := range os.Args[1:] {
		printArg(arg, &outputStrings[i])
	}
	for _, s := range outputStrings {
		fmt.Println(len(s))
		fmt.Println(s)
	}
}

func printArg(arg string, outputString *string) {
	var location string
	var port int
	// fmt.Sscanf is greedy so cannot use "%s=localhost:%d"
	replacer := strings.NewReplacer("=", " ", ":", " ")
	argWhiteSpace := replacer.Replace(arg)
	_, err := fmt.Sscanf(argWhiteSpace, "%s localhost %d", &location, &port)
	if err != nil {
		fmt.Println(fmt.Errorf("invalid port number %s", arg))
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(location, outputString, conn)
}

func mustCopy(location string, dst *string, src io.Reader) {
	scanner := bufio.NewScanner(src)
	// dst.Write([]byte("must copy"))
	// (*dst) = location
	for scanner.Scan() {
		// dst.Write([]byte(fmt.Sprintf("\n%s: ", location)))
		// dst.Write(scanner.Bytes())
		// if err := scanner.Err(); err != nil {
		// 	log.Fatal(err)
		// }
		(*dst) = location + scanner.Text()
	}
}

//!-
