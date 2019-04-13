package domain

import "time"

// Message model
type Message struct {
	ID        int        `gorm:"column:id; primary_key:yes" json:"id,omitempty"`
	GroupID   int        `gorm:"column:group_id" json:"groupId"`
	ClientID  string     `gorm:"column:client_id" json:"clientId,omitempty"`
	Event     string     `gorm:"column:event" json:"event,omitempty"`
	Title     string     `gorm:"column:title" json:"title,omitempty"`
	Content   string     `gorm:"column:content" json:"content,omitempty"`
	Timestamp *time.Time `gorm:"column:timestamp" json:"timestamp,omitempty"`
}

func (m *Message) TableName() string {
	return "chat_messages"
}
