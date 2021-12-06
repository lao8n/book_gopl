// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

type PortNumber int

func (pf portFlag) String() string   { return fmt.Sprintf("%d", pf.portNumber) }
func (pn PortNumber) String() string { return fmt.Sprintf("%d", pn) }

type portFlag struct{ portNumber PortNumber }

func (f *portFlag) Set(s string) error {
	var port PortNumber
	_, err := fmt.Sscanf(s, "%d", &port)
	if err != nil {
		return fmt.Errorf("invalid port number %s", s)
	}
	f.portNumber = port
	return nil
}

func main() {

	pf := portFlag{0}
	flag.CommandLine.Var(&pf, "port", "")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", pf.portNumber))
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently // don't need to this as single client for multiple servers
	}
	//!-
}
