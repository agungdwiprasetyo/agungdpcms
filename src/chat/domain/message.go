package domain

import "time"

// Message model
type Message struct {
	ID        string     `json:"id,omitempty"`
	ClientID  string     `json:"clientID,omitempty"`
	Content   string     `json:"content,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}
