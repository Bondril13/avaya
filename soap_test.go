package avaya

import (
	"context"
	"testing"
)

func TestAnonymousLogin(t *testing.T) {
	client := NewClient("http", "host")
	ctx := context.Background()
	sesh1, anonymousID, err := client.AnonymousLogin(ctx)
	if err != nil {
		t.Error(err)
	}
	sesh2, anonmousID2, err := client.AnonymousLogin(ctx)
	if err != nil {
		t.Error(err)
	}

	if anonymousID == 0 || anonmousID2 == 0 {
		t.Errorf("anonymousID should be non-zero")
	}
	if sesh1 == sesh2 {
		t.Errorf("sessionKeys should be different: %q & %q were the same.", sesh1, sesh2)
	}
}

func TestCustomerID(t *testing.T) {

	client := NewClient("http", "host")
	ctx := context.Background()
	sessionID, anonymousID, err := client.AnonymousLogin(ctx)
	if err != nil {
		t.Error(err)
	}

	_, err = client.CustomerID(ctx, sessionID, anonymousID, "test@test.co.uk")
	if err != nil {
		t.Error(err)
	}
}

//TODO test skillset()
