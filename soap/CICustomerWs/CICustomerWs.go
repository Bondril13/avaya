package CICustomerWs

import (
	"encoding/xml"
	"tbp/avaya/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// RegisterNewCustomer -
type RegisterNewCustomer struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RegisterNewCustomer"`

	NewCustomer *CICustomerWriteType `xml:"newCustomer,omitempty"`

	NewPhoneNumber *CIPhoneNumberWriteType `xml:"newPhoneNumber,omitempty"`

	NewAddress *CIAddressWriteType `xml:"newAddress,omitempty"`

	NewEmailAddress *CIEmailAddressWriteType `xml:"newEmailAddress,omitempty"`
}

type RegisterNewCustomerResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RegisterNewCustomerResponse"`

	RegisterNewCustomerResult int64 `xml:"RegisterNewCustomerResult,omitempty"`
}

type RegisterAnonymousCustomer struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RegisterAnonymousCustomer"`

	NewCustomer *CICustomerWriteType `xml:"newCustomer,omitempty"`

	NewPhoneNumber *CIPhoneNumberWriteType `xml:"newPhoneNumber,omitempty"`
}

type RegisterAnonymousCustomerResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RegisterAnonymousCustomerResponse"`

	RegisterAnonymousCustomerResult int64 `xml:"RegisterAnonymousCustomerResult,omitempty"`
}

type RequestTextChat struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestTextChat"`

	CustID int64 `xml:"custID,omitempty"`

	CreateAsClosed bool `xml:"createAsClosed,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`

	NewContact *CIContactWriteType `xml:",omitempty"`
}

type RequestTextChatResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestTextChatResponse"`

	RequestTextChatResult int64 `xml:"RequestTextChatResult,omitempty"`
}

type RequestScheduledCallback struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestScheduledCallback"`

	CustID int64 `xml:"custID,omitempty"`

	NewContact *CIContactWriteType `xml:",omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type RequestScheduledCallbackResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestScheduledCallbackResponse"`

	RequestScheduledCallbackResult int64 `xml:"RequestScheduledCallbackResult,omitempty"`
}

type RequestScheduledCallbackUTC struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestScheduledCallbackUTC"`

	CustID int64 `xml:"custID,omitempty"`

	NewContact *CIContactWriteType `xml:",omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type RequestScheduledCallbackUTCResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestScheduledCallbackUTCResponse"`

	RequestScheduledCallbackUTCResult int64 `xml:"RequestScheduledCallbackUTCResult,omitempty"`
}

type RequestImmediateCallback struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestImmediateCallback"`

	CustID int64 `xml:"custID,omitempty"`

	NewContact *CIContactWriteType `xml:",omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type RequestImmediateCallbackResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RequestImmediateCallbackResponse"`

	RequestImmediateCallbackResult int64 `xml:"RequestImmediateCallbackResult,omitempty"`
}

type SendPasswordReminder struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com SendPasswordReminder"`

	EmailAddress string `xml:"emailAddress,omitempty"`

	SkillsetName string `xml:"skillsetName,omitempty"`
}

type SendPasswordReminderResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com SendPasswordReminderResponse"`

	SendPasswordReminderResult int64 `xml:"SendPasswordReminderResult,omitempty"`
}

type ReadFirstBlockOfContacts struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadFirstBlockOfContacts"`

	CustomerID int64 `xml:"customerID,omitempty"`

	NumOfContacts int64 `xml:"numOfContacts,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type ReadFirstBlockOfContactsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadFirstBlockOfContactsResponse"`

	ReadFirstBlockOfContactsResult *CIContactBlockReadType `xml:"ReadFirstBlockOfContactsResult,omitempty"`
}

type ReadLastBlockOfContacts struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadLastBlockOfContacts"`

	CustomerID int64 `xml:"customerID,omitempty"`

	NumOfContacts int64 `xml:"numOfContacts,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type ReadLastBlockOfContactsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadLastBlockOfContactsResponse"`

	ReadLastBlockOfContactsResult *CIContactBlockReadType `xml:"ReadLastBlockOfContactsResult,omitempty"`
}

type ReadPreviousBlockOfContacts struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadPreviousBlockOfContacts"`

	CustomerID int64 `xml:"customerID,omitempty"`

	NumOfContacts int64 `xml:"numOfContacts,omitempty"`

	StartContactID int64 `xml:"startContactID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type ReadPreviousBlockOfContactsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadPreviousBlockOfContactsResponse"`

	ReadPreviousBlockOfContactsResult *CIContactBlockReadType `xml:"ReadPreviousBlockOfContactsResult,omitempty"`
}

type ReadNextBlockOfContacts struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadNextBlockOfContacts"`

	CustomerID int64 `xml:"customerID,omitempty"`

	NumOfContacts int64 `xml:"numOfContacts,omitempty"`

	StartContactID int64 `xml:"startContactID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type ReadNextBlockOfContactsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadNextBlockOfContactsResponse"`

	ReadNextBlockOfContactsResult *CIContactBlockReadType `xml:"ReadNextBlockOfContactsResult,omitempty"`
}

type GetNumberOfContactsByTime struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetNumberOfContactsByTime"`

	CustomerID int64 `xml:"customerID,omitempty"`

	Timestamp *CIDateTime `xml:"timestamp,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetNumberOfContactsByTimeResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetNumberOfContactsByTimeResponse"`

	GetNumberOfContactsByTimeResult int64 `xml:"GetNumberOfContactsByTimeResult,omitempty"`
}

type ReadCustomer struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadCustomer"`

	ID int64 `xml:"id,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type ReadCustomerResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadCustomerResponse"`

	ReadCustomerResult *CICustomerReadType `xml:"ReadCustomerResult,omitempty"`
}

type GetDefaultAddress struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetDefaultAddress"`

	ID int64 `xml:"id,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetDefaultAddressResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetDefaultAddressResponse"`

	GetDefaultAddressResult *CIAddressReadType `xml:"GetDefaultAddressResult,omitempty"`
}

type GetDefaultPhoneNumber struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetDefaultPhoneNumber"`

	ID int64 `xml:"id,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetDefaultPhoneNumberResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetDefaultPhoneNumberResponse"`

	GetDefaultPhoneNumberResult *CIPhoneNumberReadType `xml:"GetDefaultPhoneNumberResult,omitempty"`
}

type GetDefaultEmailAddress struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetDefaultEmailAddress"`

	ID int64 `xml:"id,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetDefaultEmailAddressResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetDefaultEmailAddressResponse"`

	GetDefaultEmailAddressResult *CIEmailAddressReadType `xml:"GetDefaultEmailAddressResult,omitempty"`
}

type GetCustomerByUsername struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByUsername"`

	Username string `xml:"username,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetCustomerByUsernameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByUsernameResponse"`

	GetCustomerByUsernameResult *CICustomerReadType `xml:"GetCustomerByUsernameResult,omitempty"`
}

type GetCustomerByEmailAddress struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByEmailAddress"`

	EmailAddress string `xml:"emailAddress,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetCustomerByEmailAddressResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByEmailAddressResponse"`

	GetCustomerByEmailAddressResult *CICustomerReadType `xml:"GetCustomerByEmailAddressResult,omitempty"`
}

type GetCustomerByPhoneNumber struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByPhoneNumber"`

	IntCode string `xml:"intCode,omitempty"`

	AreaCode string `xml:"areaCode,omitempty"`

	Number string `xml:"number,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetCustomerByPhoneNumberResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByPhoneNumberResponse"`

	GetCustomerByPhoneNumberResult *CIMultipleCustomerIDReadType `xml:"GetCustomerByPhoneNumberResult,omitempty"`
}

type GetCustomerByFirstName struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByFirstName"`

	FirstName string `xml:"firstName,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetCustomerByFirstNameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByFirstNameResponse"`

	GetCustomerByFirstNameResult *CIMultipleCustomerIDReadType `xml:"GetCustomerByFirstNameResult,omitempty"`
}

type GetCustomerByLastName struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByLastName"`

	LastName string `xml:"lastName,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetCustomerByLastNameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByLastNameResponse"`

	GetCustomerByLastNameResult *CIMultipleCustomerIDReadType `xml:"GetCustomerByLastNameResult,omitempty"`
}

type GetCustomerByName struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByName"`

	FirstName string `xml:"firstName,omitempty"`

	LastName string `xml:"lastName,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetCustomerByNameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetCustomerByNameResponse"`

	GetCustomerByNameResult *CIMultipleCustomerIDReadType `xml:"GetCustomerByNameResult,omitempty"`
}

type UpdateTitle struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateTitle"`

	CustID int64 `xml:"custID,omitempty"`

	NewTitle string `xml:"newTitle,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type UpdateTitleResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateTitleResponse"`

	UpdateTitleResult int64 `xml:"UpdateTitleResult,omitempty"`
}

type UpdateFirstName struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateFirstName"`

	CustID int64 `xml:"custID,omitempty"`

	NewFirstName string `xml:"newFirstName,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type UpdateFirstNameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateFirstNameResponse"`

	UpdateFirstNameResult int64 `xml:"UpdateFirstNameResult,omitempty"`
}

type UpdateLastName struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateLastName"`

	CustID int64 `xml:"custID,omitempty"`

	NewLastName string `xml:"newLastName,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type UpdateLastNameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateLastNameResponse"`

	UpdateLastNameResult int64 `xml:"UpdateLastNameResult,omitempty"`
}

type UpdateUsername struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateUsername"`

	CustID int64 `xml:"custID,omitempty"`

	NewUsername string `xml:"newUsername,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type UpdateUsernameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateUsernameResponse"`

	UpdateUsernameResult int64 `xml:"UpdateUsernameResult,omitempty"`
}

type UpdatePassword struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdatePassword"`

	CustID int64 `xml:"custID,omitempty"`

	OldPassword string `xml:"oldPassword,omitempty"`

	NewPassword string `xml:"newPassword,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type UpdatePasswordResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdatePasswordResponse"`

	UpdatePasswordResult int64 `xml:"UpdatePasswordResult,omitempty"`
}

type UpdateCustomer struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateCustomer"`

	ID int64 `xml:"id,omitempty"`

	NewCustomerData *CICustomerWriteType `xml:"newCustomerData,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type UpdateCustomerResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com UpdateCustomerResponse"`

	UpdateCustomerResult int64 `xml:"UpdateCustomerResult,omitempty"`
}

type AddAddress struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AddAddress"`

	ID int64 `xml:"id,omitempty"`

	NewAddress *CIAddressWriteType `xml:"newAddress,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type AddAddressResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AddAddressResponse"`

	AddAddressResult int64 `xml:"AddAddressResult,omitempty"`
}

type AddPhoneNumber struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AddPhoneNumber"`

	ID int64 `xml:"id,omitempty"`

	NewPhoneNumber *CIPhoneNumberWriteType `xml:"newPhoneNumber,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type AddPhoneNumberResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AddPhoneNumberResponse"`

	AddPhoneNumberResult int64 `xml:"AddPhoneNumberResult,omitempty"`
}

type AddEmailAddress struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AddEmailAddress"`

	ID int64 `xml:"id,omitempty"`

	NewEmailAddress *CIEmailAddressWriteType `xml:"newEmailAddress,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type AddEmailAddressResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AddEmailAddressResponse"`

	AddEmailAddressResult int64 `xml:"AddEmailAddressResult,omitempty"`
}

type AddCustomField struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AddCustomField"`

	ID int64 `xml:"id,omitempty"`

	NewCustomField *CICustomFieldWriteType `xml:"newCustomField,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type AddCustomFieldResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AddCustomFieldResponse"`

	AddCustomFieldResult int64 `xml:"AddCustomFieldResult,omitempty"`
}

type RemoveAddress struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RemoveAddress"`

	CustID int64 `xml:"custID,omitempty"`

	AddressID int64 `xml:"addressID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type RemoveAddressResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RemoveAddressResponse"`

	RemoveAddressResult int64 `xml:"RemoveAddressResult,omitempty"`
}

type RemovePhoneNumber struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RemovePhoneNumber"`

	CustID int64 `xml:"custID,omitempty"`

	PhoneNumberID int64 `xml:"phoneNumberID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type RemovePhoneNumberResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RemovePhoneNumberResponse"`

	RemovePhoneNumberResult int64 `xml:"RemovePhoneNumberResult,omitempty"`
}

type RemoveEmailAddress struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RemoveEmailAddress"`

	CustID int64 `xml:"custID,omitempty"`

	EmailAddressID int64 `xml:"emailAddressID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type RemoveEmailAddressResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RemoveEmailAddressResponse"`

	RemoveEmailAddressResult int64 `xml:"RemoveEmailAddressResult,omitempty"`
}

type RemoveCustomField struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RemoveCustomField"`

	CustID int64 `xml:"custID,omitempty"`

	CustomFieldID int64 `xml:"customFieldID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type RemoveCustomFieldResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com RemoveCustomFieldResponse"`

	RemoveCustomFieldResult int64 `xml:"RemoveCustomFieldResult,omitempty"`
}

type CIPhoneNumberType string

const (
	CIPhoneNumberTypeUnknown CIPhoneNumberType = "Unknown"

	CIPhoneNumberTypeHome CIPhoneNumberType = "Home"

	CIPhoneNumberTypeBusiness CIPhoneNumberType = "Business"

	CIPhoneNumberTypeFax CIPhoneNumberType = "Fax"

	CIPhoneNumberTypeMobile CIPhoneNumberType = "Mobile"

	CIPhoneNumberTypeOther CIPhoneNumberType = "Other"

	CIPhoneNumberTypeUnspecified CIPhoneNumberType = "Unspecified"
)

type CIContactPriority string

const (
	CIContactPriorityPriority1Highest CIContactPriority = "Priority1Highest"

	CIContactPriorityPriority2High CIContactPriority = "Priority2High"

	CIContactPriorityPriority3MediumHigh CIContactPriority = "Priority3MediumHigh"

	CIContactPriorityPriority4MediumLow CIContactPriority = "Priority4MediumLow"

	CIContactPriorityPriority5Low CIContactPriority = "Priority5Low"

	CIContactPriorityPriority6Lowest CIContactPriority = "Priority6Lowest"

	CIContactPriorityUnspecified CIContactPriority = "Unspecified"
)

type CIContactSource string

const (
	CIContactSourceNotAvailable CIContactSource = "NotAvailable"

	CIContactSourceEMail CIContactSource = "EMail"

	CIContactSourceAgentCreated CIContactSource = "AgentCreated"

	CIContactSourceWeb CIContactSource = "Web"

	CIContactSourceUnspecified CIContactSource = "Unspecified"
)

type CIContactStatus string

const (
	CIContactStatusNew CIContactStatus = "New"

	CIContactStatusOpen CIContactStatus = "Open"

	CIContactStatusClosed CIContactStatus = "Closed"

	CIContactStatusWaiting CIContactStatus = "Waiting"

	CIContactStatusUnspecified CIContactStatus = "Unspecified"
)

type CIContactType string

const (
	CIContactTypeFax CIContactType = "Fax"

	CIContactTypeSMS CIContactType = "SMS"

	CIContactTypeVoiceMail CIContactType = "VoiceMail"

	CIContactTypeWhiteMail CIContactType = "WhiteMail"

	CIContactTypeOther CIContactType = "Other"

	CIContactTypeScheduledCallback CIContactType = "ScheduledCallback"

	CIContactTypeVoice CIContactType = "Voice"

	CIContactTypeEmail CIContactType = "Email"

	CIContactTypeWebCommunications CIContactType = "WebCommunications"

	CIContactTypeOutbound CIContactType = "Outbound"

	CIContactTypeVideo CIContactType = "Video"

	CIContactTypeUnspecified CIContactType = "Unspecified"
)

type CICallbackStatus string

const (
	CICallbackStatusUnknown CICallbackStatus = "Unknown"

	CICallbackStatusNoCallback CICallbackStatus = "NoCallback"

	CICallbackStatusCallCompleted CICallbackStatus = "CallCompleted"

	CICallbackStatusCallNotAnswered CICallbackStatus = "CallNotAnswered"

	CICallbackStatusWrongTelephoneNumber CICallbackStatus = "WrongTelephoneNumber"

	CICallbackStatusNoNumberSelected CICallbackStatus = "NoNumberSelected"

	CICallbackStatusBusy CICallbackStatus = "Busy"

	CICallbackStatusOther CICallbackStatus = "Other"

	CICallbackStatusEMailInQueue CICallbackStatus = "EMailInQueue"

	CICallbackStatusEMailSent CICallbackStatus = "EMailSent"

	CICallbackStatusEMailUndeliverable CICallbackStatus = "EMailUndeliverable"

	CICallbackStatusEMailNotSent CICallbackStatus = "EMailNotSent"

	CICallbackStatusEMailAddressNotProvided CICallbackStatus = "EMailAddressNotProvided"

	CICallbackStatusEMailReceived CICallbackStatus = "EMailReceived"

	CICallbackStatusEMailSendInProgress CICallbackStatus = "EMailSendInProgress"

	CICallbackStatusFaxInQueue CICallbackStatus = "FaxInQueue"

	CICallbackStatusFaxSent CICallbackStatus = "FaxSent"

	CICallbackStatusWrongFaxNumber CICallbackStatus = "WrongFaxNumber"

	CICallbackStatusFaxNotSent CICallbackStatus = "FaxNotSent"

	CICallbackStatusMailPosted CICallbackStatus = "MailPosted"

	CICallbackStatusResponseCancelled CICallbackStatus = "ResponseCancelled"

	CICallbackStatusWebreply CICallbackStatus = "Webreply"

	CICallbackStatusContactClosed CICallbackStatus = "ContactClosed"

	CICallbackStatusContactReOpened CICallbackStatus = "ContactReOpened"

	CICallbackStatusContactTransferred CICallbackStatus = "ContactTransferred"

	CICallbackStatusContactBarred CICallbackStatus = "ContactBarred"

	CICallbackStatusAgentNote CICallbackStatus = "AgentNote"

	CICallbackStatusChatResponse CICallbackStatus = "ChatResponse"

	CICallbackStatusDone CICallbackStatus = "Done"

	CICallbackStatusCreated CICallbackStatus = "Created"

	CICallbackStatusReceived CICallbackStatus = "Received"

	CICallbackStatusUnspecified CICallbackStatus = "Unspecified"
)

type CIActionSource string

const (
	CIActionSourceNotAvailable CIActionSource = "NotAvailable"

	CIActionSourceAutomatedResponse CIActionSource = "AutomatedResponse"

	CIActionSourceEMailfromCustomer CIActionSource = "EMailfromCustomer"

	CIActionSourceEMailfromAgenttoCustomer CIActionSource = "EMailfromAgenttoCustomer"

	CIActionSourcePhonefromCustomer CIActionSource = "PhonefromCustomer"

	CIActionSourcePhonefromAgenttoCustomer CIActionSource = "PhonefromAgenttoCustomer"

	CIActionSourceTextChatfromCustomer CIActionSource = "TextChatfromCustomer"

	CIActionSourceTextChatfromAgenttoCustomer CIActionSource = "TextChatfromAgenttoCustomer"

	CIActionSourceOtherfromCustomer CIActionSource = "OtherfromCustomer"

	CIActionSourceOtherfromAgenttoCustomer CIActionSource = "OtherfromAgenttoCustomer"

	CIActionSourceTransfertoAgent CIActionSource = "TransfertoAgent"

	CIActionSourceTransfertoSkillset CIActionSource = "TransfertoSkillset"

	CIActionSourceTransfertoExternalAgent CIActionSource = "TransfertoExternalAgent"

	CIActionSourceAgentCreatedContact CIActionSource = "AgentCreatedContact"

	CIActionSourceAgentCreated CIActionSource = "AgentCreated"

	CIActionSourceCallbackRequestfromCustomer CIActionSource = "CallbackRequestfromCustomer"

	CIActionSourcePasswordRemindertoCustomer CIActionSource = "PasswordRemindertoCustomer"

	CIActionSourceOutboundVoiceCallNotMade CIActionSource = "OutboundVoiceCallNotMade"

	CIActionSourceSupervisorClosed CIActionSource = "SupervisorClosed"

	CIActionSourceCustomerAbandonedWebOnHold CIActionSource = "CustomerAbandonedWebOnHold"

	CIActionSourceClosedByCustomApplication CIActionSource = "ClosedByCustomApplication"

	CIActionSourceChatTextHistory CIActionSource = "ChatTextHistory"

	CIActionSourceUnspecified CIActionSource = "Unspecified"
)

type CICustomerWriteType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CICustomerWriteType"`

	Title string `xml:"title,omitempty"`

	FirstName string `xml:"firstName,omitempty"`

	LastName string `xml:"lastName,omitempty"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`
}

type CIPhoneNumberWriteType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIPhoneNumberWriteType"`

	InternationalCode string `xml:"internationalCode,omitempty"`

	AreaCode string `xml:"areaCode,omitempty"`

	Number string `xml:"number,omitempty"`

	PhoneNumberType *CIPhoneNumberType `xml:"phoneNumberType,omitempty"`

	DoNotCall bool `xml:"doNotCall,omitempty"`

	DoNotCallSpecified bool `xml:"doNotCallSpecified,omitempty"`

	DefaultPhoneNumber bool `xml:"defaultPhoneNumber,omitempty"`

	DefaultPhoneNumberSpecified bool `xml:"defaultPhoneNumberSpecified,omitempty"`
}

type CIAddressWriteType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIAddressWriteType"`

	Line1 string `xml:"line1,omitempty"`

	Line2 string `xml:"line2,omitempty"`

	Line3 string `xml:"line3,omitempty"`

	Line4 string `xml:"line4,omitempty"`

	Line5 string `xml:"line5,omitempty"`

	Zipcode string `xml:"zipcode,omitempty"`

	Country string `xml:"country,omitempty"`

	DefaultAddress bool `xml:"defaultAddress,omitempty"`

	DefaultAddressSpecified bool `xml:"defaultAddressSpecified,omitempty"`
}

type CIEmailAddressWriteType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIEmailAddressWriteType"`

	EmailAddress string `xml:"emailAddress,omitempty"`

	DefaultEmailAddress bool `xml:"defaultEmailAddress,omitempty"`

	DefaultEmailAddressSpecified bool `xml:"defaultEmailAddressSpecified,omitempty"`
}

type CIContactWriteType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com" newContact` // CIContactWriteType"`

	SkillsetID int64 `xml:"skillsetID,omitempty"`

	Priority *CIContactPriority `xml:"priority,omitempty"`

	Timezone int16 `xml:"timezone,omitempty"`

	Text string `xml:"text,omitempty"`

	TextHTML string `xml:"textHTML,omitempty"`

	Subject string `xml:"subject,omitempty"`

	CallbackTime *CIDateTime `xml:"callbackTime,omitempty"`

	WebOnHoldTag string `xml:"webOnHoldTag,omitempty"`

	CustomFields []*CICustomFieldWriteType `xml:"customFields,omitempty"`
}

type CIDateTime struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIDateTime"`

	Milliseconds int64 `xml:"milliseconds,omitempty"`
}

type CICustomFieldWriteType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CICustomFieldWriteType"`

	Name string `xml:"name,omitempty"`

	Text string `xml:"text,omitempty"`

	IsTextVisible bool `xml:"isTextVisible,omitempty"`
}

type CIContactBlockReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIContactBlockReadType"`

	CustomerID int64 `xml:"customerID,omitempty"`

	ContactsList *ArrayOfCIContactReadType `xml:"contactsList,omitempty"`

	ContactsRemainingBefore int64 `xml:"contactsRemainingBefore,omitempty"`

	ContactsRemainingAfter int64 `xml:"contactsRemainingAfter,omitempty"`
}

type ArrayOfCIContactReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIContactReadType"`

	CIContactReadType []*CIContactReadType `xml:"CIContactReadType,omitempty"`
}

type CIContactReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIContactReadType"`

	ID int64 `xml:"id,omitempty"`

	CustomerID int64 `xml:"customerID,omitempty"`

	OriginalSubject string `xml:"originalSubject,omitempty"`

	Source *CIContactSource `xml:"source,omitempty"`

	Status *CIContactStatus `xml:"status,omitempty"`

	Skillset *CISkillsetReadType `xml:"skillset,omitempty"`

	Priority *CIContactPriority `xml:"priority,omitempty"`

	Timezone int64 `xml:"timezone,omitempty"`

	WebOnHoldTag string `xml:"webOnHoldTag,omitempty"`

	ArrivalTime *CIDateTime `xml:"arrivalTime,omitempty"`

	ClosedTime *CIDateTime `xml:"closedTime,omitempty"`

	OpenTime *CIDateTime `xml:"openTime,omitempty"`

	OpenDuration int64 `xml:"openDuration,omitempty"`

	MailTo string `xml:"MailTo,omitempty"`

	MailFrom string `xml:"MailFrom,omitempty"`

	MailCC string `xml:"MailCC,omitempty"`

	ContactType *CIContactType `xml:"contactType,omitempty"`

	Agent *CIAgentReadType `xml:"agent,omitempty"`

	ActionList *ArrayOfCIActionReadType `xml:"actionList,omitempty"`

	CustomFieldList *ArrayOfCICustomFieldReadType `xml:"customFieldList,omitempty"`
}

type CISkillsetReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CISkillsetReadType"`

	ID int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Tag string `xml:"tag,omitempty"`
}

type CIAgentReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIAgentReadType"`

	ID int64 `xml:"id,omitempty"`

	FirstName string `xml:"firstName,omitempty"`

	LastName string `xml:"lastName,omitempty"`
}

type ArrayOfCIActionReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIActionReadType"`

	CIActionReadType []*CIActionReadType `xml:"CIActionReadType,omitempty"`
}

type CIActionReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIActionReadType"`

	ID int64 `xml:"id,omitempty"`

	ContactID int64 `xml:"contactID,omitempty"`

	Agent *CIAgentReadType `xml:"agent,omitempty"`

	Subject string `xml:"subject,omitempty"`

	Text string `xml:"text,omitempty"`

	TextHTML string `xml:"textHTML,omitempty"`

	CallbackTime *CIDateTime `xml:"callbackTime,omitempty"`

	CallbackStatus *CICallbackStatus `xml:"callbackStatus,omitempty"`

	CreationTime *CIDateTime `xml:"creationTime,omitempty"`

	Source *CIActionSource `xml:"source,omitempty"`

	MailTo string `xml:"mailTo,omitempty"`

	MailCC string `xml:"mailCC,omitempty"`

	MailBCC string `xml:"mailBCC,omitempty"`

	Comment string `xml:"comment,omitempty"`

	TimeAllocated int64 `xml:"timeAllocated,omitempty"`

	NumberUsed string `xml:"numberUsed,omitempty"`

	OutboundTalkTime int64 `xml:"outboundTalkTime,omitempty"`

	OutboundDispositionCode string `xml:"outboundDispositionCode,omitempty"`

	ActionType *CIContactType `xml:"actionType,omitempty"`

	CustomFieldList *ArrayOfCICustomFieldReadType `xml:"customFieldList,omitempty"`
}

type ArrayOfCICustomFieldReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCICustomFieldReadType"`

	CICustomFieldReadType []*CICustomFieldReadType `xml:"CICustomFieldReadType,omitempty"`
}

type CICustomFieldReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CICustomFieldReadType"`

	ID int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Text string `xml:"text,omitempty"`

	IsTextVisible bool `xml:"isTextVisible,omitempty"`
}

type CICustomerReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CICustomerReadType"`

	ID int64 `xml:"id,omitempty"`

	Title string `xml:"title,omitempty"`

	FirstName string `xml:"firstName,omitempty"`

	LastName string `xml:"lastName,omitempty"`

	Username string `xml:"username,omitempty"`

	RegisterDate *CIDateTime `xml:"registerDate,omitempty"`

	AddressList *ArrayOfCIAddressReadType `xml:"addressList,omitempty"`

	PhoneNumberList *ArrayOfCIPhoneNumberReadType `xml:"phoneNumberList,omitempty"`

	EmailAddressList *ArrayOfCIEmailAddressReadType `xml:"emailAddressList,omitempty"`

	CustomFieldList *ArrayOfCICustomFieldReadType `xml:"customFieldList,omitempty"`

	DefaultAddress *CIAddressReadType `xml:"defaultAddress,omitempty"`

	DefaultPhoneNumber *CIPhoneNumberReadType `xml:"defaultPhoneNumber,omitempty"`

	DefaultEmailAddress *CIEmailAddressReadType `xml:"defaultEmailAddress,omitempty"`
}

type ArrayOfCIAddressReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIAddressReadType"`

	CIAddressReadType []*CIAddressReadType `xml:"CIAddressReadType,omitempty"`
}

type CIAddressReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIAddressReadType"`

	ID int64 `xml:"id,omitempty"`

	Line1 string `xml:"line1,omitempty"`

	Line2 string `xml:"line2,omitempty"`

	Line3 string `xml:"line3,omitempty"`

	Line4 string `xml:"line4,omitempty"`

	Line5 string `xml:"line5,omitempty"`

	Zipcode string `xml:"zipcode,omitempty"`

	Country string `xml:"country,omitempty"`

	DefaultAddress bool `xml:"defaultAddress,omitempty"`
}

type ArrayOfCIPhoneNumberReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIPhoneNumberReadType"`

	CIPhoneNumberReadType []*CIPhoneNumberReadType `xml:"CIPhoneNumberReadType,omitempty"`
}

type CIPhoneNumberReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIPhoneNumberReadType"`

	ID int64 `xml:"id,omitempty"`

	InternationalCode string `xml:"internationalCode,omitempty"`

	AreaCode string `xml:"areaCode,omitempty"`

	Number string `xml:"number,omitempty"`

	DoNotCall bool `xml:"doNotCall,omitempty"`

	DefaultPhoneNumber bool `xml:"defaultPhoneNumber,omitempty"`

	PhoneNumberType *CIPhoneNumberType `xml:"phoneNumberType,omitempty"`
}

type ArrayOfCIEmailAddressReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIEmailAddressReadType"`

	CIEmailAddressReadType []*CIEmailAddressReadType `xml:"CIEmailAddressReadType,omitempty"`
}

type CIEmailAddressReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIEmailAddressReadType"`

	ID int64 `xml:"id,omitempty"`

	EmailAddress string `xml:"emailAddress,omitempty"`

	DefaultEmailAddress bool `xml:"defaultEmailAddress,omitempty"`
}

type CIMultipleCustomerIDReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIMultipleCustomerIDReadType"`

	CustomerIDs *ArrayOfLong `xml:"customerIDs,omitempty"`
}

type ArrayOfLong struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfLong"`

	Long []int64 `xml:"long,omitempty"`
}

type CICustomerWsSoap struct {
	client *soap.SOAPClient
}

func NewCICustomerWsSoap(url string, tls bool, auth *soap.BasicAuth) *CICustomerWsSoap {
	if url == "" {
		url = "http://***REMOVED***/ccmmwebservices/CICustomerWs.asmx"
	}
	client := soap.NewSOAPClient(url, tls, auth)

	return &CICustomerWsSoap{
		client: client,
	}
}

func (service *CICustomerWsSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

func (service *CICustomerWsSoap) RegisterNewCustomer(request *RegisterNewCustomer) (*RegisterNewCustomerResponse, error) {
	response := new(RegisterNewCustomerResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RegisterNewCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RegisterAnonymousCustomer(request *RegisterAnonymousCustomer) (*RegisterAnonymousCustomerResponse, error) {
	response := new(RegisterAnonymousCustomerResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RegisterAnonymousCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RequestTextChat(request *RequestTextChat) (*RequestTextChatResponse, error) {
	response := new(RequestTextChatResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RequestTextChat", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RequestScheduledCallback(request *RequestScheduledCallback) (*RequestScheduledCallbackResponse, error) {
	response := new(RequestScheduledCallbackResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RequestScheduledCallback", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RequestScheduledCallbackUTC(request *RequestScheduledCallbackUTC) (*RequestScheduledCallbackUTCResponse, error) {
	response := new(RequestScheduledCallbackUTCResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RequestScheduledCallbackUTC", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RequestImmediateCallback(request *RequestImmediateCallback) (*RequestImmediateCallbackResponse, error) {
	response := new(RequestImmediateCallbackResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RequestImmediateCallback", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) SendPasswordReminder(request *SendPasswordReminder) (*SendPasswordReminderResponse, error) {
	response := new(SendPasswordReminderResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/SendPasswordReminder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) ReadFirstBlockOfContacts(request *ReadFirstBlockOfContacts) (*ReadFirstBlockOfContactsResponse, error) {
	response := new(ReadFirstBlockOfContactsResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/ReadFirstBlockOfContacts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) ReadLastBlockOfContacts(request *ReadLastBlockOfContacts) (*ReadLastBlockOfContactsResponse, error) {
	response := new(ReadLastBlockOfContactsResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/ReadLastBlockOfContacts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) ReadPreviousBlockOfContacts(request *ReadPreviousBlockOfContacts) (*ReadPreviousBlockOfContactsResponse, error) {
	response := new(ReadPreviousBlockOfContactsResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/ReadPreviousBlockOfContacts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) ReadNextBlockOfContacts(request *ReadNextBlockOfContacts) (*ReadNextBlockOfContactsResponse, error) {
	response := new(ReadNextBlockOfContactsResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/ReadNextBlockOfContacts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetNumberOfContactsByTime(request *GetNumberOfContactsByTime) (*GetNumberOfContactsByTimeResponse, error) {
	response := new(GetNumberOfContactsByTimeResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetNumberOfContactsByTime", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) ReadCustomer(request *ReadCustomer) (*ReadCustomerResponse, error) {
	response := new(ReadCustomerResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/ReadCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetDefaultAddress(request *GetDefaultAddress) (*GetDefaultAddressResponse, error) {
	response := new(GetDefaultAddressResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetDefaultAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetDefaultPhoneNumber(request *GetDefaultPhoneNumber) (*GetDefaultPhoneNumberResponse, error) {
	response := new(GetDefaultPhoneNumberResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetDefaultPhoneNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetDefaultEmailAddress(request *GetDefaultEmailAddress) (*GetDefaultEmailAddressResponse, error) {
	response := new(GetDefaultEmailAddressResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetDefaultEmailAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetCustomerByUsername(request *GetCustomerByUsername) (*GetCustomerByUsernameResponse, error) {
	response := new(GetCustomerByUsernameResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetCustomerByUsername", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetCustomerByEmailAddress(request *GetCustomerByEmailAddress) (*GetCustomerByEmailAddressResponse, error) {
	response := new(GetCustomerByEmailAddressResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetCustomerByEmailAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetCustomerByPhoneNumber(request *GetCustomerByPhoneNumber) (*GetCustomerByPhoneNumberResponse, error) {
	response := new(GetCustomerByPhoneNumberResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetCustomerByPhoneNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetCustomerByFirstName(request *GetCustomerByFirstName) (*GetCustomerByFirstNameResponse, error) {
	response := new(GetCustomerByFirstNameResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetCustomerByFirstName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetCustomerByLastName(request *GetCustomerByLastName) (*GetCustomerByLastNameResponse, error) {
	response := new(GetCustomerByLastNameResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetCustomerByLastName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) GetCustomerByName(request *GetCustomerByName) (*GetCustomerByNameResponse, error) {
	response := new(GetCustomerByNameResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/GetCustomerByName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) UpdateTitle(request *UpdateTitle) (*UpdateTitleResponse, error) {
	response := new(UpdateTitleResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/UpdateTitle", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) UpdateFirstName(request *UpdateFirstName) (*UpdateFirstNameResponse, error) {
	response := new(UpdateFirstNameResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/UpdateFirstName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) UpdateLastName(request *UpdateLastName) (*UpdateLastNameResponse, error) {
	response := new(UpdateLastNameResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/UpdateLastName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) UpdateUsername(request *UpdateUsername) (*UpdateUsernameResponse, error) {
	response := new(UpdateUsernameResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/UpdateUsername", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) UpdatePassword(request *UpdatePassword) (*UpdatePasswordResponse, error) {
	response := new(UpdatePasswordResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/UpdatePassword", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) UpdateCustomer(request *UpdateCustomer) (*UpdateCustomerResponse, error) {
	response := new(UpdateCustomerResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/UpdateCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) AddAddress(request *AddAddress) (*AddAddressResponse, error) {
	response := new(AddAddressResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/AddAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) AddPhoneNumber(request *AddPhoneNumber) (*AddPhoneNumberResponse, error) {
	response := new(AddPhoneNumberResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/AddPhoneNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) AddEmailAddress(request *AddEmailAddress) (*AddEmailAddressResponse, error) {
	response := new(AddEmailAddressResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/AddEmailAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) AddCustomField(request *AddCustomField) (*AddCustomFieldResponse, error) {
	response := new(AddCustomFieldResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/AddCustomField", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RemoveAddress(request *RemoveAddress) (*RemoveAddressResponse, error) {
	response := new(RemoveAddressResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RemoveAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RemovePhoneNumber(request *RemovePhoneNumber) (*RemovePhoneNumberResponse, error) {
	response := new(RemovePhoneNumberResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RemovePhoneNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RemoveEmailAddress(request *RemoveEmailAddress) (*RemoveEmailAddressResponse, error) {
	response := new(RemoveEmailAddressResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RemoveEmailAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *CICustomerWsSoap) RemoveCustomField(request *RemoveCustomField) (*RemoveCustomFieldResponse, error) {
	response := new(RemoveCustomFieldResponse)
	err := service.client.Call("http://webservices.ci.ccmm.applications.nortel.com/RemoveCustomField", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
