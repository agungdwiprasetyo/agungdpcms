package delivery

import (
	"log"
	"net/http"

	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/usecase"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

// Handler websocket chat service
type Handler struct {
	server *chat.Server
	uc     usecase.Chat
}

// NewWebsocketHandler constructor
func NewWebsocketHandler(s *chat.Server) *Handler {
	return &Handler{server: s, uc: usecase.New()}
}

// Socket handler
func (h *Handler) Socket(res http.ResponseWriter, req *http.Request) {
	sock, err := h.server.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := req.Header.Get("Sec-Websocket-Key")
	roomID := req.URL.Query().Get("roomId")
	debug.Println(req.URL.Query().Get("token"))

	var client chat.Client
	client.ID = id
	client.Conn = sock
	client.Data = make(chan []byte)
	client.Server = h.server
	h.server.Connect <- &client

	h.uc.Join(roomID, &client)
}
