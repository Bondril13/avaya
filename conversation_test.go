package avaya

import "testing"

func TestNewConversation(t *testing.T) {
	name := "Testy McTestface"
	email := "test@test.io"

	c, err := NewConversation(name, email, "WC_Default_Skillset")
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name, expected, actual interface{}
	}{
		{"name", name, c.Name},
		{"email", email, c.Email},
		{"customerID", false, c.CustomerID == 0},
		{"sessionKey", len("412f3iMi00"), len(c.SessionKey)},
		{"skillset name", "WC_Default_Skillset", c.Skillset.Name},
		// TODO:		{"contact ID", 12, c.contactID},
		// TODO: skillset ID not tested
	}

	for _, test := range tests {
		if test.expected != test.actual {
			t.Errorf("Expected %q to be %q but was %q", test.name, test.expected, test.actual)
		}
	}
}
