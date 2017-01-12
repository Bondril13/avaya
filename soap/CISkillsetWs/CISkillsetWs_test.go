package CISkillsetWs

import (
	"tbp/avaya/soap"
	"tbp/avaya/soap/CIUtilityWs"
	"testing"
)

func TestReadSkillsetByName(t *testing.T) {
	ciSkillset := NewSoap("", false, &soap.BasicAuth{})
	request := ReadSkillsetByName{
		SessionKey:   sessionKey(t),
		SkillsetName: "WC_Default_Skillset",
	}
	resp, err := ciSkillset.ReadSkillsetByName(&request)
	if err != nil {
		t.Errorf("Failed to read skillset: %v", err)
	}

	t.Log(resp.ReadSkillsetByNameResult.ID)
	t.Log(resp.ReadSkillsetByNameResult.Name)
	t.Log(resp.ReadSkillsetByNameResult.Tag)

}

func sessionKey(t *testing.T) string {
	ciUtil := CIUtilityWs.NewSoap("", false, &soap.BasicAuth{})
	resp, err := ciUtil.GetAnonymousSessionKey(&CIUtilityWs.GetAnonymousSessionKey{})
	if err != nil {
		t.Errorf("Failed to get anonymous session key: %v", err)
	}
	return resp.GetAnonymousSessionKeyResult.SessionKey
}
