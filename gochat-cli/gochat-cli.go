package main

import (
	"bufio"
	"fmt"
	"flag"
	"net"
	"os"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address of GoChat server")
	port := flag.Uint("port", 0, "Port used by GoChat server")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Username is required")
	}
	username := flag.Args()[0]

	if *port <= 0 || *port > 0xffff {
		fmt.Println("Error: Port must be between 1 and 65535.")
		return
	}

	conn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", *ip, uint16(*port)))
	if (err != nil) {
		fmt.Println(err)
		return
	}

	go messageThread(conn)

	connWriter := bufio.NewWriter(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if (err != nil) {
			fmt.Println(err)
		}

		msg = fmt.Sprintf("(%s) %s", username, msg)
		_, err = connWriter.WriteString(msg)
		if err != nil {
			fmt.Println(err)
		}

		connWriter.Flush()
	}
}

func messageThread(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if (err != nil) {
			fmt.Println("Lost connection to server.")
			return
		}

		fmt.Print(msg)
	}
}