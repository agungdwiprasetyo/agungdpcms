package usecase

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/src/shared"
	"github.com/gorilla/websocket"
)

type chatImpl struct {
	repo repository.Chat
}

// New chat usecase
func New(conf *config.Config) Chat {
	return &chatImpl{repo: repository.NewChatRepo(conf.DB)}
}

func (uc *chatImpl) Join(roomID string, client *chat.Client) error {
	// TODO: add client join to roomId in database (repository)
	groupID, _ := strconv.Atoi(roomID)
	res := uc.repo.FindGroupByID(groupID)
	if res.Error != nil {
		return res.Error
	}

	now := time.Now()

	// Read client message
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v\n", r)
			}
			message, _ := json.Marshal(domain.Message{ClientID: client.ID, Event: "disconnect"})
			client.Server.Message <- message

			client.Server.Disconnect <- client
			client.Conn.Close()
		}()

		for {
			_, message, err := client.Conn.ReadMessage()
			if err != nil {
				break
			}
			var m domain.Message
			json.Unmarshal(message, &m)
			m.ClientID = client.ID
			m.Timestamp = &now
			uc.repo.SaveMessage(&m)

			message, _ = json.Marshal(m)
			client.Server.Message <- message
		}
	}()

	// Write client message
	go func() {
		defer func() {
			client.Conn.Close()
		}()

		for {
			select {
			case msg, ok := <-client.Data:
				if !ok {
					client.Conn.WriteMessage(websocket.CloseMessage, nil)
					break
				}

				err := client.Conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					client.Conn.Close()
				}
			}
		}
	}()

	return nil
}

func (uc *chatImpl) FindAllMessagesByGroupID(groupID int32) (res shared.Result) {
	res = uc.repo.FindAllMessageByGroupID(int(groupID))
	if res.Error != nil {
		return
	}

	messages := res.Data.([]*domain.Message)
	data := new(serializer.MessageListSchema)
	for _, m := range messages {
		data.Data = append(data.Data, &serializer.MessageSchema{Message: m})
	}

	res.Data = data
	return res
}