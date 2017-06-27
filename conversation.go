package avaya

import (
	"context"
)

// Conversation - an avaya conversation
type Conversation interface {
	Keepalive(ctx context.Context, isTyping bool) error
	WriteMessage(context.Context, string) error
	ReadMessages(context.Context) (messages []Message, advisorTyping bool, err error)
	Close(context.Context)
	IsAnswered() bool
	IsClosed() bool
	Name() string
	Verbose(bool)
}

// Message - a message from Avaya
type Message struct {
	Text   string
	Type   string
	Hidden string
	Time   int64
}
