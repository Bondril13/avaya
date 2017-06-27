package avaya

import (
	"context"
	"fmt"
	"log"
)

// DirectConversation - A conversation between a client and an advisor.
type DirectConversation struct {
	client       Client
	closed       bool
	answered     bool
	name         string
	sessionKey   string
	customerID   int64
	contactID    int64
	skillset     *Skillset
	lastReadTime int64
	verbose      bool
}

func (c *DirectConversation) CustomerID() int64 { return c.customerID }
func (c *DirectConversation) Name() string      { return c.name }

func (c *DirectConversation) Keepalive(ctx context.Context, isTyping bool) error {
	if c.IsClosed() {
		return fmt.Errorf("Conversation closed, not sending keep alive")
	}
	return c.client.KeepAlive(ctx, c.sessionKey, c.contactID, isTyping)
}

func (c *DirectConversation) WriteMessage(ctx context.Context, message string) error {
	if c.IsClosed() {
		return fmt.Errorf("Conversation closed, not writing message")
	}
	return c.client.WriteMessage(ctx, c.sessionKey, c.contactID, message, fromCustomer)
}

func (c *DirectConversation) ReadMessages(ctx context.Context) ([]Message, bool, error) {
	if c.closed {
		return nil, false, fmt.Errorf("Conversation closed, not reading messages")
	}

	soapMessages, err := c.client.ReadMessages(ctx, c.sessionKey, c.contactID, false, c.lastReadTime)
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

func (c *DirectConversation) Verbose(verbose bool) {
	c.verbose = verbose
	c.client.Verbose = verbose
}

func (c *DirectConversation) Close(ctx context.Context) {
	if c.closed {
		return
	}
	c.closed = true
	if c.answered {
		log.Println("Close() answered")
		if err := c.client.WriteMessage(ctx, c.sessionKey, c.contactID, "The customer has disconnected.", customerDisconnected); err != nil {
			log.Println("Close(): failed to send disconnect message:", err)
		}
	} else {
		log.Println("Close() abandoning queue")
		c.client.AbandonQue(ctx, c.sessionKey, c.contactID, "The conversation has ended.")
	}
	c.client.EndSession(ctx, c.sessionKey, c.contactID)
}

func (c *DirectConversation) IsAnswered() bool { return c.answered }

func (c *DirectConversation) IsClosed() bool { return c.closed }

// NewConversation - Creates a new Conversation
func NewDirectConversation(ctx context.Context, c Client, name, email, skillsetName string) (Conversation, error) {
	sessionID, anonymousID, err := c.AnonymousLogin(ctx)
	if err != nil {
		return nil, err
	}

	customerID, err := c.CustomerID(ctx, sessionID, int64(anonymousID), email)
	if err != nil {
		return nil, err
	}

	skillset, err := c.Skillset(ctx, sessionID, skillsetName)
	if err != nil {
		return nil, err
	}

	contactID, err := c.RequestChat(ctx, customerID, sessionID, skillset.ID)
	if err != nil {
		return nil, err
	}

	return &DirectConversation{
		c,
		false,
		false,
		name,
		sessionID,
		customerID,
		contactID,
		skillset,
		1,
		false,
	}, nil
}
