package chat

import (
	"github.com/gorilla/websocket"
)

// Client model
type Client struct {
	ID     string
	Server *Server
	Conn   *websocket.Conn
	Data   chan []byte
}
