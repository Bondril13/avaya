package avaya

import (
	"fmt"
	"log"
)

// Conversation - A conversation between a client and an advisor.
type Conversation struct {
	closed       bool
	answered     bool
	Name         string
	Email        string
	SessionKey   string
	CustomerID   int64
	ContactID    int64
	Skillset     *Skillset
	lastReadTime int64
}

// Message - a message from Avaya
type Message struct {
	Text   string
	Type   string
	Hidden string
	Time   int64
}

func (c *Conversation) Keepalive(isTyping bool) error {
	return keepAlive(c.SessionKey, c.ContactID, isTyping)
}

func (c *Conversation) WriteMessage(message string) error {
	return writeMessage(c.SessionKey, c.ContactID, message)
}

func (c *Conversation) ReadMessages() ([]Message, bool, error) {
	if c.closed {
		return nil, false, fmt.Errorf("Cannot read messages - conversation is closed")
	}

	soapMessages, err := readMessages(c.SessionKey, c.ContactID, false, c.lastReadTime)
	if err != nil {
		return nil, false, err
	}

	//	c.lastReadTime = time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))

	messages := []Message{}

	advisorTyping := false
	if soapMessages != nil {
		c.lastReadTime = soapMessages.LastReadTime.Milliseconds
		advisorTyping = soapMessages.IsWriting

		if soapMessages.ListOfChatMessages != nil {
			for _, m := range soapMessages.ListOfChatMessages.CIChatMessageReadType {
				c.answered = true
				messages = append(messages, Message{
					Text:   m.ChatMessage,
					Hidden: m.HiddenMessage,
					Type:   string(*m.ChatMessageType),
					Time:   m.WriteTime.Milliseconds,
				})
			}
		}
	}

	return messages, advisorTyping, nil
}

func (c *Conversation) Close() {
	if c.closed {
		return
	}
	c.closed = true
	if c.answered {
		log.Println("Close() answered")
	} else {
		log.Println("Close() abandoning queue")
		AbandonQue(c.SessionKey, c.ContactID, "The conversation has ended.")
	}
	EndSession(c.SessionKey, c.ContactID)
}

// NewConversation - Creates a new Conversation
func NewConversation(name, email, skillsetName string) (*Conversation, error) {
	sessionID, anonymousID, err := anonymousLogin()
	if err != nil {
		return nil, err
	}

	customerID, err := customerID(sessionID, int64(anonymousID), email)
	if err != nil {
		return nil, err
	}

	skillset, err := skillset(sessionID, skillsetName)
	if err != nil {
		return nil, err
	}

	contactID, err := requestChat(customerID, sessionID, skillset.ID)
	if err != nil {
		return nil, err
	}

	return &Conversation{
		false,
		false,
		name,
		email,
		sessionID,
		customerID,
		contactID,
		skillset,
		1,
	}, nil
}
