package avaya

import (
	"fmt"
	"log"
)

type Conversation interface {
	Keepalive(isTyping bool) error
	WriteMessage(text string) error
	ReadMessages() ([]Message, bool, error)
	Close()
	Name() string
}

// DirectConversation - A conversation between a client and an advisor.
type DirectConversation struct {
	closed       bool
	answered     bool
	name         string
	sessionKey   string
	customerID   int64
	contactID    int64
	skillset     *Skillset
	lastReadTime int64
}

// Message - a message from Avaya
type Message struct {
	Text   string
	Type   string
	Hidden string
	Time   int64
}

func (c *DirectConversation) CustomerID() int64 { return c.customerID }
func (c *DirectConversation) Name() string      { return c.name }

func (c *DirectConversation) Keepalive(isTyping bool) error {
	return keepAlive(c.sessionKey, c.contactID, isTyping)
}

func (c *DirectConversation) WriteMessage(message string) error {
	return writeMessage(c.sessionKey, c.contactID, message, fromCustomer)
}

func (c *DirectConversation) ReadMessages() ([]Message, bool, error) {
	if c.closed {
		return nil, false, fmt.Errorf("Cannot read messages - conversation is closed")
	}

	soapMessages, err := readMessages(c.sessionKey, c.contactID, false, c.lastReadTime)
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

func (c *DirectConversation) Close() {
	if c.closed {
		return
	}
	c.closed = true
	if c.answered {
		log.Println("Close() answered")
		if err := writeMessage(c.sessionKey, c.contactID, "The customer has disconnected.", customerDisconnected); err != nil {
			log.Println("Close(): failed to send disconnect message:", err)
		}
	} else {
		log.Println("Close() abandoning queue")
		AbandonQue(c.sessionKey, c.contactID, "The conversation has ended.")
	}
	EndSession(c.sessionKey, c.contactID)
}

// NewConversation - Creates a new Conversation
func NewConversation(name, email, skillsetName string) (Conversation, error) {
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

	return &DirectConversation{
		false,
		false,
		name,
		sessionID,
		customerID,
		contactID,
		skillset,
		1,
	}, nil
}
