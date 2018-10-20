package server

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port uint16
}

func (server Server) Listen(quitChannel chan bool) {
	address := fmt.Sprintf("%s:%d", server.Ip, server.Port)
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		fmt.Println("Error:", err)
		quitChannel <- true
	} else {
		fmt.Println("Server listening on", listener.Addr())

		connChannel := make(chan net.Conn)
		errorChannel := make(chan error)

		go acceptThread(listener, connChannel, errorChannel)

		for {
			select {
			case conn := <-connChannel:
				fmt.Println("Connection from", conn.RemoteAddr())
			case err := <-errorChannel:
				fmt.Println("Error:", err)
			case <-quitChannel:
				fmt.Print("Closing listener... ")
				listener.Close()
				fmt.Print("Done")
				return
			}
		}
	}
}

func acceptThread(listener net.Listener, connChannel chan net.Conn, errorChannel chan error) {
	for {
		conn, err := listener.Accept()
		if (err != nil) {
			errorChannel <- err
			return
		}

		connChannel <- conn
	}
}
