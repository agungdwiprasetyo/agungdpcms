package delivery

import (
	"log"
	"net/http"

	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/usecase"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

// WsHandler websocket chat service
type WsHandler struct {
	server *chat.Server
	uc     usecase.Chat
}

// NewWebsocketHandler constructor
func NewWebsocketHandler(s *chat.Server, uc usecase.Chat) *WsHandler {
	return &WsHandler{server: s, uc: uc}
}

// Socket handler
func (h *WsHandler) Socket(res http.ResponseWriter, req *http.Request) {
	sock, err := h.server.Upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := req.Header.Get("Sec-Websocket-Key")
	groupID := req.URL.Query().Get("groupId")
	debug.Println(req.URL.Query().Get("token"))

	var client chat.Client
	client.ID = id
	client.Conn = sock
	client.Data = make(chan []byte)
	client.Server = h.server
	h.server.Connect <- &client

	err = h.uc.Join(groupID, &client)
	if err != nil {
		debug.Println(err)
		h.server.Disconnect <- &client
	}
}
