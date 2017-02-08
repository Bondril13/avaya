package avaya

import (
	"context"
	"fmt"
	"log"
)

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

func (c *DirectConversation) CustomerID() int64 { return c.customerID }
func (c *DirectConversation) Name() string      { return c.name }

func (c *DirectConversation) Keepalive(ctx context.Context, isTyping bool) error {
	if c.IsClosed() {
		return fmt.Errorf("Conversation closed, not sending keep alive")
	}
	return keepAlive(ctx, c.sessionKey, c.contactID, isTyping)
}

func (c *DirectConversation) WriteMessage(ctx context.Context, message string) error {
	if c.IsClosed() {
		return fmt.Errorf("Conversation closed, not writing message")
	}
	return writeMessage(ctx, c.sessionKey, c.contactID, message, fromCustomer)
}

func (c *DirectConversation) ReadMessages(ctx context.Context) ([]Message, bool, error) {
	if c.closed {
		return nil, false, fmt.Errorf("Conversation closed, not reading messages")
	}

	soapMessages, err := readMessages(ctx, c.sessionKey, c.contactID, false, c.lastReadTime)
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
				if *m.ChatMessageType == agentDisconnected {
					c.closed = true
				}
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

func (c *DirectConversation) Close(ctx context.Context) {
	if c.closed {
		return
	}
	c.closed = true
	if c.answered {
		log.Println("Close() answered")
		if err := writeMessage(ctx, c.sessionKey, c.contactID, "The customer has disconnected.", customerDisconnected); err != nil {
			log.Println("Close(): failed to send disconnect message:", err)
		}
	} else {
		log.Println("Close() abandoning queue")
		AbandonQue(ctx, c.sessionKey, c.contactID, "The conversation has ended.")
	}
	EndSession(ctx, c.sessionKey, c.contactID)
}

func (c *DirectConversation) IsClosed() bool { return c.closed }

// NewConversation - Creates a new Conversation
func NewDirectConversation(ctx context.Context, name, email, skillsetName string) (Conversation, error) {
	sessionID, anonymousID, err := anonymousLogin(ctx)
	if err != nil {
		return nil, err
	}

	customerID, err := customerID(ctx, sessionID, int64(anonymousID), email)
	if err != nil {
		return nil, err
	}

	skillset, err := skillset(ctx, sessionID, skillsetName)
	if err != nil {
		return nil, err
	}

	contactID, err := requestChat(ctx, customerID, sessionID, skillset.ID)
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
