package server

import (
	"errors"
	"net"
	"sync"
)

var roomDB map[string][]net.Conn = make(map[string][]net.Conn)
var mutex *sync.Mutex = &sync.Mutex{}

func AddToRoom(roomName string, conn net.Conn) {
	mutex.Lock()

	_, ok := roomDB[roomName]
	if !ok {
		roomDB[roomName] = make([]net.Conn, 0)
	}

	roomDB[roomName] = append(roomDB[roomName], conn)

	mutex.Unlock()
}

func GetRoom(roomName string) ([]net.Conn, error) {
	mutex.Lock()

	room, ok := roomDB[roomName]
	if !ok {
		mutex.Unlock()
		return nil, errors.New("Room does not exist")
	}

	mutex.Unlock()
	return room, nil
}