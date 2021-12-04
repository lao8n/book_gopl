// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		var location string
		var port int
		// fmt.Sscanf is greedy so cannot use "%s=localhost:%d"
		replacer := strings.NewReplacer("=", " ", ":", " ")
		argWhiteSpace := replacer.Replace(arg)
		fmt.Printf("replacer %s\n", argWhiteSpace)
		_, err := fmt.Sscanf(argWhiteSpace, "%s localhost %d", &location, &port)
		if err != nil {
			fmt.Println(fmt.Errorf("invalid port number %s", arg))
		}
		fmt.Printf("%s: ", location)
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		mustCopy(os.Stdout, conn)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//!-
