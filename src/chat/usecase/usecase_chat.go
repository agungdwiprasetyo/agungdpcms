package usecase

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/agungdwiprasetyo/agungdpcms/shared/meta"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/serializer"
	"github.com/agungdwiprasetyo/agungdpcms/websocket"
	"github.com/agungdwiprasetyo/go-utils/debug"
	ws "github.com/gorilla/websocket"
)

type chatImpl struct {
	repo *repository.Repository
}

// New chat usecase
func New(repo *repository.Repository) Chat {
	return &chatImpl{
		repo: repo,
	}
}

func (uc *chatImpl) Join(roomID string, client *websocket.Client) error {
	groupID, _ := strconv.Atoi(roomID)
	res := uc.repo.Chat.FindGroupByID(groupID)
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
			tp, message, err := client.Conn.ReadMessage()
			debug.Println(tp)
			if err != nil {
				break
			}
			var m domain.Message
			json.Unmarshal(message, &m)
			m.ID = int(time.Now().Unix())
			m.ClientID = client.ID
			m.Timestamp = &now

			go uc.repo.Chat.SaveMessage(&m)

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
					client.Conn.WriteMessage(ws.CloseMessage, nil)
					break
				}

				err := client.Conn.WriteMessage(ws.TextMessage, msg)
				if err != nil {
					client.Conn.Close()
				}
			case dis := <-client.Server.Disconnect:
				client.Conn.WriteMessage(ws.CloseMessage, []byte(`{"message": "closed"}`))
				client.Server.RemoveClient(dis)
				debug.Println(len(dis.Server.Clients))
			}
		}
	}()

	return nil
}

func (uc *chatImpl) FindAllMessages(args *domain.Param) (res shared.Result) {
	filter := filter.Filter{Page: args.Page, Limit: args.Limit}
	filter.CalculateOffset()
	mt := &meta.Meta{Page: int(args.Page), Limit: int(args.Limit)}

	res = uc.repo.Chat.FindAllMessage(int(args.GroupID), int(filter.Offset), mt.Limit)
	if res.Error != nil {
		return
	}

	messages := res.Data.([]*domain.Message)
	data := new(serializer.MessageListSchema)
	for _, m := range messages {
		data.Data = append(data.Data, &serializer.MessageSchema{Message: m})
	}

	res = uc.repo.Chat.Count(&domain.Message{GroupID: int(args.GroupID)})
	if res.Error != nil {
		return
	}
	mt.TotalRecords = res.Data.(int)
	mt.CalculatePages()

	data.M = &meta.MetaSchema{Meta: mt}

	res.Data = data
	return res
}
