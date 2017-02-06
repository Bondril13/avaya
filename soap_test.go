package avaya

import "testing"

func TestAnonymousLogin(t *testing.T) {
	sesh1, anonymousID, err := anonymousLogin()
	if err != nil {
		t.Error(err)
	}
	sesh2, anonmousID2, err := anonymousLogin()
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

	sessionID, anonymousID, err := anonymousLogin()
	if err != nil {
		t.Error(err)
	}

	_, err = customerID(sessionID, anonymousID, "test@test.co.uk")
	if err != nil {
		t.Error(err)
	}
}

//TODO test skillset()
