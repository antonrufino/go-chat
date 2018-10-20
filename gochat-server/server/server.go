package server

import (
	"bufio"
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

		errorChannel := make(chan error)
		messageChannel := make(chan string, 10)

		go acceptThread(listener, messageChannel, errorChannel)
		go messageHandlerThread(messageChannel)

		for {
			select {
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

func acceptThread(listener net.Listener, messageChannel chan string, errorChannel chan error) {
	for {
		conn, err := listener.Accept()
		if (err != nil) {
			errorChannel <- err
			continue
		}

		AddToRoom("default", conn)

		go recieveMessages(conn, messageChannel)
	}
}

func recieveMessages(conn net.Conn, messageChannel chan string) {
	fmt.Println("Reciever thread created for connection from", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Closing connection.")
			conn.Close()
			return
		} else {
			messageChannel <- msg
		}
	}
}

func messageHandlerThread(messageChannel <-chan string) {
	fmt.Println("Message thread created")

	for {
		select {
		case msg := <-messageChannel:
			room, _ := GetRoom("default")

			for _, conn := range room {
				fmt.Println("Sending to", conn.RemoteAddr())
				conn.Write([]byte(msg))
			}
		}
	}
}