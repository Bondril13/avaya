package CIUtilityWs

import (
	_ "fmt"
	"tbp/avaya/soap"
	"testing"
)

func TestFoo(t *testing.T) {
	ciUtil := NewSoap("", false, &soap.BasicAuth{})
	resp, err := ciUtil.GetAnonymousSessionKey(&GetAnonymousSessionKey{})
	if err != nil {
		t.Errorf("Failed to get anonymous session key: %v", err)
	}
	t.Log(resp.GetAnonymousSessionKeyResult.AnonymousID)
	t.Log(resp.GetAnonymousSessionKeyResult.SessionKey)
}
