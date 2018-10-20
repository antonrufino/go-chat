package main

import (
	"fmt"
	"flag"
	"net"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address of GoChat server")
	port := flag.Uint("port", 0, "Port used by GoChat server")

	flag.Parse()

	if *port <= 0 || *port > 0xffff {
		fmt.Println("Error: Port must be between 1 and 65535.")
		return
	}

	conn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", *ip, uint16(*port)))
	if (err != nil) {
		fmt.Println(err)
	}

	conn.Close()
}