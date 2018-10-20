package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {
	ip, port, err := parseFlags();
	if (err != nil) {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("IP:", ip)
	fmt.Println("Port:", port)
}

func parseFlags() (string, uint16, error) {
	ip := flag.String("ip", "0.0.0.0", "IP address of server.")
	port := flag.Uint("port", 0, "Port to listen on.")

	flag.Parse()

	if *port > 0xffff {
		return "", 0, errors.New("Port must be in the range of 0-65535")
	}

	return *ip, uint16(*port), nil
}