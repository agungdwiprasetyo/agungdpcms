package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/gorilla/websocket"
)

// Server for websocket
type Server struct {
	Clients             map[string]*Client
	Message             chan []byte
	Connect, Disconnect chan *Client
	Upgrader            websocket.Upgrader
	sync.RWMutex
}

// NewServer init new websocket server
func NewServer() *Server {
	server := new(Server)
	server.Clients = make(map[string]*Client)
	server.Connect = make(chan *Client)
	server.Disconnect = make(chan *Client)
	server.Message = make(chan []byte)
	server.Upgrader = websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool {
			return true
		},
	}

	return server
}

// ListenAndServe start serve client
func (s *Server) ListenAndServe() {
	if !config.GlobalEnv.Websocket {
		return
	}
	log.Println("Websocket server up and running in path /ws")

	for {
		select {
		case client := <-s.Connect:
			log.Printf("%s joined\n", client.ID)
			s.addClient(client)

		case client := <-s.Disconnect:
			log.Printf("%s leave\n", client.ID)
			s.RemoveClient(client)

		case message := <-s.Message:
			var m domain.Message
			json.Unmarshal(message, &m)
			for id, client := range s.Clients {
				if m.ClientID != id {
					// write to all client connection
					s.RLock()
					client.Data <- message
					s.RUnlock()
				}
			}
		}
	}
}

func (s *Server) addClient(c *Client) {
	s.Lock()
	defer s.Unlock()

	s.Clients[c.ID] = c
}

func (s *Server) RemoveClient(c *Client) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Clients[c.ID]; ok {
		delete(s.Clients, c.ID)
	}
}
