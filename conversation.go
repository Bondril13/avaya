package avaya

import "context"

// Conversation - an avaya conversation
type Conversation interface {
	Keepalive(ctx context.Context, isTyping bool) error
	WriteMessage(context.Context, string) error
	ReadMessages(context.Context) ([]Message, bool, error)
	Close(context.Context)
	IsClosed() bool
	Name() string
}

// Message - a message from Avaya
type Message struct {
	Text   string
	Type   string
	Hidden string
	Time   int64
}

// NewConversation - Creates a new Conversation
func NewConversation(name, email, skillsetName string) (Conversation, error) {
	return newDirectConversation(name, email, skillsetName)
}
