package avaya

// Conversation - an avaya conversation
type Conversation interface {
	Keepalive(isTyping bool) error
	WriteMessage(text string) error
	ReadMessages() ([]Message, bool, error)
	Close()
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
