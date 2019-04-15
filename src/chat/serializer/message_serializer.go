package serializer

import (
	"github.com/agungdwiprasetyo/agungdpcms/shared/meta"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/domain"
)

type MessageSchema struct {
	Message *domain.Message
}

func (r *MessageSchema) ID() int32 {
	return int32(r.Message.ID)
}

func (r *MessageSchema) GroupID() int32 {
	return int32(r.Message.GroupID)
}

func (r *MessageSchema) ClientID() string {
	return r.Message.ClientID
}

func (r *MessageSchema) Event() string {
	return r.Message.Event
}

func (r *MessageSchema) Title() string {
	return r.Message.Title
}

func (r *MessageSchema) Content() string {
	return r.Message.Content
}

type MessageListSchema struct {
	M    *meta.MetaSchema
	Data []*MessageSchema
}

func (r *MessageListSchema) Meta() *meta.MetaSchema {
	return r.M
}

// Messages method
func (r *MessageListSchema) Messages() []*MessageSchema {
	return r.Data
}
