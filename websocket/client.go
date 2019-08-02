package websocket

import (
	"log"
	"net/http"

	"github.com/agungdwiprasetyo/go-utils/debug"
	"github.com/gorilla/websocket"
)

// Client model
type Client struct {
	ID     string
	Server *Server
	Conn   *websocket.Conn
	Data   chan []byte
}

// Usecase for ws client
type Usecase interface {
	Join(string, *Client) error
}

// Handler websocket chat service
type Handler struct {
	server *Server
	uc     Usecase
}

// NewWebsocketHandler constructor
func NewWebsocketHandler(s *Server, uc Usecase) *Handler {
	Handler := &Handler{server: s, uc: uc}
	return Handler
}

// ServeHTTP Socket handler
func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	sock, err := h.server.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := req.Header.Get("Sec-Websocket-Key")
	groupID := req.URL.Query().Get("groupId")

	client := new(Client)
	client.ID = id
	client.Conn = sock
	client.Data = make(chan []byte)
	client.Server = h.server
	h.server.Connect <- client

	err = h.uc.Join(groupID, client)
	if err != nil {
		debug.Println(err)
		h.server.Disconnect <- client
	}
}
