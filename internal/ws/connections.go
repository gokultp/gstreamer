package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var (
	// Connections is a map of streamerId and ws connections,
	Connections map[string][]*websocket.Conn = make(map[string][]*websocket.Conn)
)

func StartEventListener(streamerId uint64, conn *websocket.Conn) {
	key := fmt.Sprintf("%d", streamerId)
	Connections[key] = append(Connections[key], conn)
}

func WriteEvent(streamer string, event []byte) {
	for _, conn := range Connections[streamer] {
		conn.WriteMessage(1, event)
	}
}
