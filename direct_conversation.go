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

// CustomerID - returns the customer ID
func (c *DirectConversation) CustomerID() int64 { return c.customerID }

// Name - returns the customer name
func (c *DirectConversation) Name() string { return c.name }

// Keepalive - tells the AACC server that the client is still active, and if they are typing
func (c *DirectConversation) Keepalive(ctx context.Context, isTyping bool) error {
	if c.IsClosed() {
		return fmt.Errorf("Conversation closed, not sending keep alive")
	}
	return c.client.KeepAlive(ctx, c.sessionKey, c.contactID, isTyping)
}

// WriteMessage - set a message to the AACC server
func (c *DirectConversation) WriteMessage(ctx context.Context, message string) error {
	if c.IsClosed() {
		return fmt.Errorf("Conversation closed, not writing message")
	}
	return c.client.WriteMessage(ctx, c.sessionKey, c.contactID, message, fromCustomer)
}

// ReadMessages - gets any unread message from the AACC server
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

// SetVerbose - enable more detailed logging
func (c *DirectConversation) SetVerbose(verbose bool) {
	c.verbose = verbose
	c.client.Verbose = verbose
}

// Close - end the conversation
func (c *DirectConversation) Close(ctx context.Context) error {
	if c.closed {
		return nil
	}
	c.closed = true
	if c.answered {
		log.Println("Close() answered")
		if err := c.client.WriteMessage(ctx, c.sessionKey, c.contactID, "The customer has disconnected.", customerDisconnected); err != nil {
			log.Println("Close(): failed to send disconnect message:", err)
		}
	} else {
		log.Println("Close() abandoning queue")
		if err := c.client.AbandonQueue(ctx, c.sessionKey, c.contactID, "The conversation has ended."); err != nil {
			log.Printf("Error abandoning queue: %v", err)
		}
	}
	return c.client.EndSession(ctx, c.sessionKey, c.contactID)
}

// IsAnswered - check if an advisor has answered this conversation
func (c *DirectConversation) IsAnswered() bool { return c.answered }

// IsClosed - check if the conversation has already been ended
func (c *DirectConversation) IsClosed() bool { return c.closed }

// NewDirectConversation - Creates a new Conversation
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
