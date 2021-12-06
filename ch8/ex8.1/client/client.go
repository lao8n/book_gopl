package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// run with
// TZ=US/Eastern go run server -port 8010 &
// TZ=Asia/Tokyo go run server -port 8020 &
// go run client NewYork=localhost:8010 Tokyo=localhost:8020

func main() {
	for _, arg := range os.Args[1:] {
		location, portArg := parseArg(arg)
		conn, err := net.Dial("tcp", portArg)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go clockWall(os.Stdout, conn, location)
	}
	time.Sleep(10 * time.Second) // crucial otherwise just ends
}

func parseArg(arg string) (location string, portArg string) {
	replacer := strings.NewReplacer("=", " ", ":", " ")
	argWhiteSpace := replacer.Replace(arg)
	var port int
	fmt.Sscanf(argWhiteSpace, "%s localhost %d", &location, &port) //ignore error
	portArg = fmt.Sprintf("localhost:%d", port)
	return location, portArg
}

func clockWall(dst io.Writer, src net.Conn, location string) {
	input := bufio.NewScanner(src)
	for input.Scan() {
		io.WriteString(dst, location+" "+input.Text()+"\n")
	}
}
