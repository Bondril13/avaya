package avaya

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/TheBookPeople/avaya/soap"
	"github.com/TheBookPeople/avaya/soap/CICustomerWs"
	"github.com/TheBookPeople/avaya/soap/CISkillsetWs"
	"github.com/TheBookPeople/avaya/soap/CIUtilityWs"
	"github.com/TheBookPeople/avaya/soap/CIWebCommsWs"
)

// Client - Avaya chat SOAP client
type Client struct {
	baseURL string // e.g. http://avaya-server
	Verbose bool
}

// NewClient - creates a new Client
func NewClient(proto, host string) Client {
	return Client{proto + "://" + host, false}
}

func (c Client) newCIUtilityWs() *CIUtilityWs.Soap {
	return CIUtilityWs.NewSoap(c.baseURL, false, &soap.BasicAuth{}, c.Verbose)
}

func (c Client) newCISkillset() *CISkillsetWs.Soap {
	return CISkillsetWs.NewSoap(c.baseURL, false, &soap.BasicAuth{}, c.Verbose)
}

func (c Client) newCIWebComms() *CIWebCommsWs.Soap {
	return CIWebCommsWs.NewSoap(c.baseURL, false, &soap.BasicAuth{}, c.Verbose)
}

func (c Client) newCICustomerWs() *CICustomerWs.Soap {
	return CICustomerWs.NewSoap(c.baseURL, false, &soap.BasicAuth{}, c.Verbose)
}

// AnonymousLogin - login to the AACC server
func (c Client) AnonymousLogin(ctx context.Context) (string, int64, error) {
	ciUtil := c.newCIUtilityWs()
	resp, err := ciUtil.GetAnonymousSessionKey(ctx, &CIUtilityWs.GetAnonymousSessionKey{})
	if err != nil {
		return "", 0, fmt.Errorf("Failed to get anonymous session key: %v", err)
	}
	anonymousID, err := strconv.Atoi(resp.GetAnonymousSessionKeyResult.AnonymousID)
	if err != nil {
		return "", 0, fmt.Errorf("Failed to convert anonymousID to int64: %v %v", resp.GetAnonymousSessionKeyResult.AnonymousID, err)
	}
	return resp.GetAnonymousSessionKeyResult.SessionKey, int64(anonymousID), nil
}

// KeepAlive - send a keepalive messsage to the AACC server
func (c Client) KeepAlive(ctx context.Context, sessionID string, contactID int64, isTyping bool) error {
	ciWebComms := c.newCIWebComms()
	_, err := ciWebComms.UpdateAliveTimeAndUpdateIsTyping(ctx, &CIWebCommsWs.UpdateAliveTimeAndUpdateIsTyping{
		ContactID:  contactID,
		SessionKey: sessionID,
		IsTyping:   isTyping,
	})

	return err
}

// CustomerID - Get a customer ID for a given session
func (c Client) CustomerID(ctx context.Context, sessionID string, anonymousID int64, email string) (int64, error) {

	ciUtil := c.newCIUtilityWs()
	resp, err := ciUtil.GetAndUpdateAnonymousCustomerID(ctx, &CIUtilityWs.GetAndUpdateAnonymousCustomerID{
		LoginResult: &CIUtilityWs.AnonymousLoginResult{
			AnonymousID: anonymousID,
			SessionKey:  sessionID,
		},
		EmailAddress: email,
		ThisCustomer: &CIUtilityWs.CICustomerReadType{
			AddressList: &CIUtilityWs.ArrayOfCIAddressReadType{
				CIAddressReadType: []*CIUtilityWs.CIAddressReadType{
					&CIUtilityWs.CIAddressReadType{},
				},
			},
		},
	})
	if err != nil {
		return 0, err
	}
	return resp.GetAndUpdateAnonymousCustomerIDResult, nil
}

// IsSkillsetInService - establish if a give n skillset is in use, e.g. an advisor is logged in
func (c Client) IsSkillsetInService(ctx context.Context, skillsetName string) (bool, error) {
	resp, err := c.newCISkillset().IsSkillsetNameInService(ctx, &CISkillsetWs.IsSkillsetNameInService{
		SkillsetName: skillsetName,
	})
	if err != nil {
		return false, err
	}
	return resp.IsSkillsetNameInServiceResult, nil
}

// Skillset - an AACC skillset (e.g. can handle orders, complaints etc..)
type Skillset struct {
	ID   int64
	Name string
}

// Skillset - create a Skillset for a given name & session
func (c Client) Skillset(ctx context.Context, sessionID, name string) (*Skillset, error) {
	ciSkillset := c.newCISkillset()
	resp, err := ciSkillset.GetSkillsetByName(ctx, &CISkillsetWs.GetSkillsetByName{
		SessionKey:   sessionID,
		SkillsetName: name,
	})

	if err != nil {
		return nil, err
	}

	return &Skillset{
		ID:   resp.GetSkillsetByNameResult.ID,
		Name: resp.GetSkillsetByNameResult.Name,
	}, nil
}

// RequestChat - start a conversation
func (c Client) RequestChat(ctx context.Context, customerID int64, sessionID string, skillsetID int64) (int64, error) {
	ciCustomerWs := c.newCICustomerWs()

	resp, err := ciCustomerWs.RequestTextChat(ctx, &CICustomerWs.RequestTextChat{
		CustID:     customerID,
		SessionKey: sessionID,
		NewContact: &CICustomerWs.CIContactWriteType{
			SkillsetID: skillsetID,
		},
	})

	if err != nil {
		return 0, err
	}

	return resp.RequestTextChatResult, nil
}

// ReadMessages - read unread messages from the AACC server
func (c Client) ReadMessages(ctx context.Context, sessionKey string, contactID int64, isWriting bool, lastReadTime int64) (*CIWebCommsWs.CIMultipleChatMessageReadType, error) {
	ciWebComms := CIWebCommsWs.NewSoap(c.baseURL, false, &soap.BasicAuth{}, c.Verbose)
	resp, err := ciWebComms.ReadChatMessage(ctx, &CIWebCommsWs.ReadChatMessage{
		ContactID: contactID,
		IsWriting: isWriting,
		LastReadTime: &CIWebCommsWs.CIDateTime{
			Milliseconds: lastReadTime,
		},
		SessionKey: sessionKey,
	})

	if err != nil {
		return nil, err
	}

	result := resp.ReadChatMessageResult
	if c.Verbose {
		log.Println(resp)
	}
	return result, nil
}

const (
	fromCustomer         = CIWebCommsWs.CIChatMessageTypeChatMessagefromCustomer
	customerDisconnected = CIWebCommsWs.CIChatMessageTypeSessionDisconnectedbyCustomer
	agentDisconnected    = CIWebCommsWs.CIChatMessageTypeSessionDisconnectedbyAgent
)

// WriteMessage - send a message to the AACC server
func (c Client) WriteMessage(ctx context.Context, sessionKey string, contactID int64, message string, msgType CIWebCommsWs.CIChatMessageType) error {

	ciWebComms := CIWebCommsWs.NewSoap(c.baseURL, false, &soap.BasicAuth{}, c.Verbose)
	resp, err := ciWebComms.WriteChatMessage(ctx, &CIWebCommsWs.WriteChatMessage{
		ContactID:       contactID,
		Message:         message,
		SessionKey:      sessionKey,
		ChatMessageType: &msgType,
	})

	if err != nil {
		return err
	}

	fmt.Printf("###\nchat message result:%d\n###", resp.WriteChatMessageResult)
	return nil
}

// AbandonQueue - remove a session from the AACC chat queue
func (c Client) AbandonQueue(ctx context.Context, sessionKey string, contactID int64, reason string) error {
	ciWebComms := c.newCIWebComms()
	_, err := ciWebComms.AbandonQueuingWebCommsContact(ctx, &CIWebCommsWs.AbandonQueuingWebCommsContact{
		SessionKey:     sessionKey,
		ContactID:      contactID,
		ClosureComment: reason,
	})
	return err
}

// EndSession - stop the conversation
func (c Client) EndSession(ctx context.Context, sessionKey string, contactID int64) error {
	ciUtil := c.newCIUtilityWs()
	_, err := ciUtil.CustomerLogoffByContactID(ctx, &CIUtilityWs.CustomerLogoffByContactID{
		SessionKey: sessionKey,
		ContactID:  contactID,
		// FIXME: Does username do anything?
	})
	return err
}

/*
web_1          | <env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:tns="http://webservices.ci.ccmm.applications.nortel.com" xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ins0="http://datatypes.ci.ccmm.applications.nortel.com">
web_1          |   <env:Body>
web_1          |     <tns:RequestTextChat>
web_1          |       <tns:custID>198853</tns:custID>
web_1          |       <tns:sessionKey>4145hiDT00</tns:sessionKey>
web_1          |       <tns:newContact>
web_1          |         <ins0:skillsetID>11</ins0:skillsetID>
web_1          |       </tns:newContact>
web_1          |     </tns:RequestTextChat>
web_1          |   </env:Body>
web_1          | </env:Envelope>
*/
