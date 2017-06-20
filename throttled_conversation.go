package avaya

import (
	"context"
	"log"
	"time"
)

type requestChan (chan func())

// ThrottledConversation - A conversation between a client and an advisor in
// which all requests are performed sequentially on a wrapped conversation.
// Not thread safe.
type ThrottledConversation struct {
	wrapped       Conversation
	requests      requestChan
	lastKeepAlive time.Time
}

func (tc *ThrottledConversation) queueRequest(ctx context.Context, f func()) error {
	done := make(chan bool)
	go func() { // Done in goroutine so to honour ctx.
		defer close(done)
		tc.requests <- f
		done <- true
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Name - of customer
func (tc *ThrottledConversation) Name() string { return tc.wrapped.Name() }

// IsClosed - is wrapped conversation closed
func (tc *ThrottledConversation) IsClosed() bool { return tc.wrapped.IsClosed() }

// Keepalive - Keep chat active, and send "is typing" messages
func (tc *ThrottledConversation) Keepalive(ctx context.Context, isTyping bool) error {
	if sinceLast := time.Now().Sub(tc.lastKeepAlive).Seconds(); sinceLast < 10 {
		log.Printf("Keepalive: Not bothering as last one was only %f seconds ago", sinceLast)
		return nil
	}

	resp := make(chan error)
	err := tc.queueRequest(ctx, func() {
		defer close(resp)
		resp <- tc.wrapped.Keepalive(ctx, isTyping)
	})
	if err != nil {
		return err
	}
	err = <-resp
	tc.lastKeepAlive = time.Now()
	return err
}

// WriteMessage - Send chat message
func (tc *ThrottledConversation) WriteMessage(ctx context.Context, message string) error {
	resp := make(chan error)
	err := tc.queueRequest(ctx, func() {
		defer close(resp)
		resp <- tc.wrapped.WriteMessage(ctx, message)
	})
	if err != nil {
		return err
	}
	return <-resp
}

type readMessageResponse struct {
	messages      []Message
	advisorTyping bool
	err           error
}

// ReadMessages - Non-blocking read to check for messages from the advisor
func (tc *ThrottledConversation) ReadMessages(ctx context.Context) ([]Message, bool, error) {
	resp := make(chan readMessageResponse)

	err := tc.queueRequest(ctx, func() {
		defer close(resp)
		m, typing, e := tc.wrapped.ReadMessages(ctx)
		resp <- readMessageResponse{m, typing, e}
	})
	if err != nil {
		return []Message{}, false, err
	}
	r := <-resp
	return r.messages, r.advisorTyping, r.err
}

// Close - close the wrapped conversation
func (tc *ThrottledConversation) Close(ctx context.Context) {
	if tc.wrapped.IsClosed() {
		return
	}

	done := make(chan bool)

	err := tc.queueRequest(ctx, func() {
		defer close(done)
		tc.wrapped.Close(ctx)
		done <- true
	})
	if err != nil {
		log.Printf("Problem closing conversation:", err)
	}
	<-done
}

// Throttle - Executes queued requests sequentially to stop the server being overloaded
// Can be pooled if you wish.
type Throttle struct {
	requests (requestChan)
}

// ThrottleConversation - Wraps a conversation with a throttled version
func (t *Throttle) ThrottleConversation(c Conversation) Conversation {
	return &ThrottledConversation{
		wrapped:  c,
		requests: t.requests,
	}
}

// Close - closes the request queue channel
func (t *Throttle) Close() {
	close(t.requests)
}

// NewThrottle - Creates a Throttle
func NewThrottle() *Throttle {
	requests := make(requestChan)
	go func() {
		for {
			req, ok := <-requests
			if !ok {
				log.Println("Throttle closed.")
				return
			}
			req()
			time.Sleep(time.Second)
		}
	}()
	return &Throttle{requests}
}
