package CIWebCommsWs

import (
	"encoding/xml"
	"tbp/avaya/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// CreateWebCommsSession -
type CreateWebCommsSession struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CreateWebCommsSession"`

	ContactID int64 `xml:"contactID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// CreateWebCommsSessionResponse -
type CreateWebCommsSessionResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CreateWebCommsSessionResponse"`

	CreateWebCommsSessionResult int64 `xml:"CreateWebCommsSessionResult,omitempty"`
}

// WriteChatMessage -
type WriteChatMessage struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com WriteChatMessage"`

	ContactID int64 `xml:"contactID,omitempty"`

	Message string `xml:"message,omitempty"`

	HiddenMessage string `xml:"hiddenMessage,omitempty"`

	ChatMessageType *CIChatMessageType `xml:"chatMessageType,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// WriteChatMessageResponse -
type WriteChatMessageResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com WriteChatMessageResponse"`

	WriteChatMessageResult int64 `xml:"WriteChatMessageResult,omitempty"`
}

// ReadChatMessage -
type ReadChatMessage struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadChatMessage"`

	ContactID int64 `xml:"contactID,omitempty"`

	LastReadTime *CIDateTime `xml:"lastReadTime,omitempty"`

	IsWriting bool `xml:"isWriting"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// ReadChatMessageResponse -
type ReadChatMessageResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadChatMessageResponse"`

	ReadChatMessageResult *CIMultipleChatMessageReadType //`xml:"ReadChatMessageResult,omitempty"`
}

// RequestChatHistory -
type RequestChatHistory struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestChatHistory"`

	ContactID int64 `xml:"contactID,omitempty"`

	Email string `xml:"email,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// RequestChatHistoryResponse -
type RequestChatHistoryResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestChatHistoryResponse"`

	RequestChatHistoryResult int64 `xml:"RequestChatHistoryResult,omitempty"`
}

// GetWebOnHoldURLs -
type GetWebOnHoldURLs struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetWebOnHoldURLs"`

	TagName string `xml:"tagName,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetWebOnHoldURLsResponse -
type GetWebOnHoldURLsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetWebOnHoldURLsResponse"`

	GetWebOnHoldURLsResult *CIMultipleOnHoldURLReadType `xml:"GetWebOnHoldURLsResult,omitempty"`
}

// UpdateAliveTime -
type UpdateAliveTime struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateAliveTime"`

	ContactID int64 `xml:"contactID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// UpdateAliveTimeResponse -
type UpdateAliveTimeResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateAliveTimeResponse"`

	UpdateAliveTimeResult *soap.CIDateTime `xml:"UpdateAliveTimeResult,omitempty"`
}

// UpdateAliveTimeAndUpdateIsTyping -
type UpdateAliveTimeAndUpdateIsTyping struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateAliveTimeAndUpdateIsTyping"`

	ContactID int64 `xml:"contactID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`

	IsTyping bool `xml:"isTyping"`
}

// UpdateAliveTimeAndUpdateIsTypingResponse -
type UpdateAliveTimeAndUpdateIsTypingResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateAliveTimeAndUpdateIsTypingResponse"`

	UpdateAliveTimeAndUpdateIsTypingResult *soap.CIDateTime `xml:"UpdateAliveTimeAndUpdateIsTypingResult,omitempty"`
}

// AbandonQueuingWebCommsContact -
type AbandonQueuingWebCommsContact struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AbandonQueuingWebCommsContact"`

	ContactID int64 `xml:"contactID,omitempty"`

	ClosureComment string `xml:"closureComment,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// AbandonQueuingWebCommsContactResponse -
type AbandonQueuingWebCommsContactResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AbandonQueuingWebCommsContactResponse"`

	AbandonQueuingWebCommsContactResult int64 `xml:"AbandonQueuingWebCommsContactResult,omitempty"`
}

// AbandonQueuingWCContact -
type AbandonQueuingWCContact struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AbandonQueuingWCContact"`

	ContactID int64 `xml:"contactID,omitempty"`

	ClosureComment string `xml:"closureComment,omitempty"`

	ClosedReasonCodeValue int64 `xml:"closedReasonCodeValue,omitempty"`

	ClosedReasonCodeSpecified bool `xml:"closedReasonCodeSpecified"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// AbandonQueuingWCContactResponse -
type AbandonQueuingWCContactResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AbandonQueuingWCContactResponse"`

	AbandonQueuingWCContactResult int64 `xml:"AbandonQueuingWCContactResult,omitempty"`
}

// GetCustomerTextChatLabel -
type GetCustomerTextChatLabel struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerTextChatLabel"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetCustomerTextChatLabelResponse -
type GetCustomerTextChatLabelResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerTextChatLabelResponse"`

	GetCustomerTextChatLabelResult string `xml:"GetCustomerTextChatLabelResult,omitempty"`
}

// GetCustomerTextChatName -
type GetCustomerTextChatName struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerTextChatName"`

	ContactID int64 `xml:"contactID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetCustomerTextChatNameResponse -
type GetCustomerTextChatNameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerTextChatNameResponse"`

	GetCustomerTextChatNameResult *CICustomerChatNameReadType `xml:"GetCustomerTextChatNameResult,omitempty"`
}

// GetTrunkAccessCode -
type GetTrunkAccessCode struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetTrunkAccessCode"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetTrunkAccessCodeResponse -
type GetTrunkAccessCodeResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetTrunkAccessCodeResponse"`

	GetTrunkAccessCodeResult string `xml:"GetTrunkAccessCodeResult,omitempty"`
}

// GetContactOnHoldMessages -
type GetContactOnHoldMessages struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetContactOnHoldMessages"`

	Contactid int64 `xml:"contact_id,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetContactOnHoldMessagesResponse -
type GetContactOnHoldMessagesResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetContactOnHoldMessagesResponse"`

	GetContactOnHoldMessagesResult *CIMultipleOnHoldMessages `xml:"GetContactOnHoldMessagesResult,omitempty"`
}

// CIChatMessageType -
type CIChatMessageType string

const (
	// CIChatMessageTypeChatMessagefromCustomer -
	CIChatMessageTypeChatMessagefromCustomer CIChatMessageType = "Chat_Message_from_Customer"

	//CIChatMessageTypeChatMessagefromAgent -
	CIChatMessageTypeChatMessagefromAgent CIChatMessageType = "Chat_Message_from_Agent"

	//CIChatMessageTypePagePushedbyCustomer -
	CIChatMessageTypePagePushedbyCustomer CIChatMessageType = "Page_Pushed_by_Customer"

	//CIChatMessageTypePagePushedbyAgent -
	CIChatMessageTypePagePushedbyAgent CIChatMessageType = "Page_Pushed_by_Agent"

	//CIChatMessageTypeFormSharedbyCustomer -
	CIChatMessageTypeFormSharedbyCustomer CIChatMessageType = "Form_Shared_by_Customer"

	//CIChatMessageTypeFormSharedbyAgent -
	CIChatMessageTypeFormSharedbyAgent CIChatMessageType = "Form_Shared_by_Agent"

	//CIChatMessageTypeCallMeRequestfromCustomer -
	CIChatMessageTypeCallMeRequestfromCustomer CIChatMessageType = "Call_Me_Request_from_Customer"

	//CIChatMessageTypePagePreviewedbyCustomer -
	CIChatMessageTypePagePreviewedbyCustomer CIChatMessageType = "Page_Previewed_by_Customer"

	//CIChatMessageTypePagePreviewedbyAgent -
	CIChatMessageTypePagePreviewedbyAgent CIChatMessageType = "Page_Previewed_by_Agent"

	// CIChatMessageTypeSessionDisconnectedbyCustomer -
	CIChatMessageTypeSessionDisconnectedbyCustomer CIChatMessageType = "Session_Disconnected_by_Customer"

	//CIChatMessageTypeSessionDisconnectedbyAgent -
	CIChatMessageTypeSessionDisconnectedbyAgent CIChatMessageType = "Session_Disconnected_by_Agent"

	//CIChatMessageTypePrivateChatMessagebetweenAgents -
	CIChatMessageTypePrivateChatMessagebetweenAgents CIChatMessageType = "Private_Chat_Message_between_Agents"

	//CIChatMessageTypeComfortMessage -
	CIChatMessageTypeComfortMessage CIChatMessageType = "Comfort_Message"
)

// CIDateTime -
type CIDateTime struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIDateTime"`

	Milliseconds int64 `xml:"http://datatypes.ci.ccmm.applications.nortel.com milliseconds,omitempty"`
}

// CIMultipleChatMessageReadType -
type CIMultipleChatMessageReadType struct {
	//XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ReadChatMessageResult"` // CIMultipleChatMessageReadType"`
	XMLName xml.Name `xml:"ReadChatMessageResult"` // CIMultipleChatMessageReadType"`

	ListOfChatMessages *ArrayOfCIChatMessageReadType `xml:"listOfChatMessages"`

	IsWriting bool `xml:"isWriting"`

	LastReadTime *soap.CIDateTime `xml:"lastReadTime"`
}

// ArrayOfCIChatMessageReadType -
type ArrayOfCIChatMessageReadType struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com listOfArrayOfCIChatMessageReadType"`

	CIChatMessageReadType []*CIChatMessageReadType `xml:"CIChatMessageReadType"`
}

// CIChatMessageReadType -
type CIChatMessageReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIChatMessageReadType"`

	ChatMessage string `xml:"chatMessage,omitempty"`

	HiddenMessage string `xml:"hiddenMessage,omitempty"`

	WriteTime *soap.CIDateTime `xml:"writeTime,omitempty"`

	ChatMessageType *CIChatMessageType `xml:"chatMessageType,omitempty"`
}

// CIMultipleOnHoldURLReadType -
type CIMultipleOnHoldURLReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIMultipleOnHoldURLReadType"`

	URLList *ArrayOfCIOnHoldURLReadType `xml:"URLList,omitempty"`
}

// ArrayOfCIOnHoldURLReadType -
type ArrayOfCIOnHoldURLReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIOnHoldURLReadType"`

	CIOnHoldURLReadType []*CIOnHoldURLReadType `xml:"CIOnHoldURLReadType,omitempty"`
}

// CIOnHoldURLReadType -
type CIOnHoldURLReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIOnHoldURLReadType"`

	URL string `xml:"URL,omitempty"`

	Description string `xml:"description,omitempty"`

	Tag string `xml:"tag,omitempty"`

	HoldTime int64 `xml:"holdTime,omitempty"`

	Sequence int64 `xml:"sequence,omitempty"`
}

// CICustomerChatNameReadType -
type CICustomerChatNameReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CICustomerChatNameReadType"`

	FirstName string `xml:"firstName,omitempty"`

	LastName string `xml:"lastName,omitempty"`
}

// CIMultipleOnHoldMessages -
type CIMultipleOnHoldMessages struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIMultipleOnHoldMessages"`

	ListOfOnHoldMessages *ArrayOfCICOnHoldMessages `xml:"listOfOnHoldMessages,omitempty"`
}

// ArrayOfCICOnHoldMessages -
type ArrayOfCICOnHoldMessages struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCICOnHoldMessages"`

	CICOnHoldMessages []*CICOnHoldMessages `xml:"CICOnHoldMessages,omitempty"`
}

// CICOnHoldMessages -
type CICOnHoldMessages struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CICOnHoldMessages"`

	Message string `xml:"Message,omitempty"`
}

// Soap -
type Soap struct {
	client *soap.SOAPClient
}

// NewSoap -
func NewSoap(url string, tls bool, auth *soap.BasicAuth) *Soap {
	if url == "" {
		url = "http://***REMOVED***/ccmmwebservices/CIWebCommsWs.asmx"
	}
	client := soap.NewSOAPClient(url, tls, auth)

	return &Soap{
		client: client,
	}
}

// SetHeader -
func (service *Soap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

// CreateWebCommsSession -
func (service *Soap) CreateWebCommsSession(request *CreateWebCommsSession) (*CreateWebCommsSessionResponse, error) {
	response := new(CreateWebCommsSessionResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/CreateWebCommsSession", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// WriteChatMessage -
func (service *Soap) WriteChatMessage(request *WriteChatMessage) (*WriteChatMessageResponse, error) {
	response := new(WriteChatMessageResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/WriteChatMessage", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ReadChatMessage -
func (service *Soap) ReadChatMessage(request *ReadChatMessage) (*ReadChatMessageResponse, error) {
	response := new(ReadChatMessageResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/ReadChatMessage", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RequestChatHistory -
func (service *Soap) RequestChatHistory(request *RequestChatHistory) (*RequestChatHistoryResponse, error) {
	response := new(RequestChatHistoryResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RequestChatHistory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetWebOnHoldURLs -
func (service *Soap) GetWebOnHoldURLs(request *GetWebOnHoldURLs) (*GetWebOnHoldURLsResponse, error) {
	response := new(GetWebOnHoldURLsResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetWebOnHoldURLs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateAliveTime -
func (service *Soap) UpdateAliveTime(request *UpdateAliveTime) (*UpdateAliveTimeResponse, error) {
	response := new(UpdateAliveTimeResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/UpdateAliveTime", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateAliveTimeAndUpdateIsTyping -
func (service *Soap) UpdateAliveTimeAndUpdateIsTyping(request *UpdateAliveTimeAndUpdateIsTyping) (*UpdateAliveTimeAndUpdateIsTypingResponse, error) {
	response := new(UpdateAliveTimeAndUpdateIsTypingResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/UpdateAliveTimeAndUpdateIsTyping", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AbandonQueuingWebCommsContact -
func (service *Soap) AbandonQueuingWebCommsContact(request *AbandonQueuingWebCommsContact) (*AbandonQueuingWebCommsContactResponse, error) {
	response := new(AbandonQueuingWebCommsContactResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/AbandonQueuingWebCommsContact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// AbandonQueuingWCContact -
func (service *Soap) AbandonQueuingWCContact(request *AbandonQueuingWCContact) (*AbandonQueuingWCContactResponse, error) {
	response := new(AbandonQueuingWCContactResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/AbandonQueuingWCContact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCustomerTextChatLabel -
func (service *Soap) GetCustomerTextChatLabel(request *GetCustomerTextChatLabel) (*GetCustomerTextChatLabelResponse, error) {
	response := new(GetCustomerTextChatLabelResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetCustomerTextChatLabel", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCustomerTextChatName -
func (service *Soap) GetCustomerTextChatName(request *GetCustomerTextChatName) (*GetCustomerTextChatNameResponse, error) {
	response := new(GetCustomerTextChatNameResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetCustomerTextChatName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetTrunkAccessCode -
func (service *Soap) GetTrunkAccessCode(request *GetTrunkAccessCode) (*GetTrunkAccessCodeResponse, error) {
	response := new(GetTrunkAccessCodeResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetTrunkAccessCode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetContactOnHoldMessages -
func (service *Soap) GetContactOnHoldMessages(request *GetContactOnHoldMessages) (*GetContactOnHoldMessagesResponse, error) {
	response := new(GetContactOnHoldMessagesResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetContactOnHoldMessages", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
