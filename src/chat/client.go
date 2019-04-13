package chat

import (
	"encoding/json"
	"log"

	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/gorilla/websocket"
)

// Client model
type Client struct {
	ID     string
	Server *Server
	Conn   *websocket.Conn
	Data   chan []byte
}

func (c *Client) Read() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v\n", r)
			}
			c.Server.Disconnect <- c
			c.Conn.Close()
		}()

		for {
			_, message, err := c.Conn.ReadMessage()
			if err != nil {
				break
			}
			var m domain.Message
			json.Unmarshal(message, &m)
			m.ClientID = c.ID

			message, _ = json.Marshal(m)
			c.Server.Message <- message
		}
	}()
}

// Write function will write message to connected client
func (c *Client) Write() {
	go func() {
		defer func() {
			c.Conn.Close()
		}()

		for {
			select {
			case msg, ok := <-c.Data:
				if !ok {
					c.Conn.WriteMessage(websocket.CloseMessage, nil)
					break
				}

				err := c.Conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					c.Conn.Close()
				}
			}
		}
	}()
}
