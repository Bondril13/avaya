package CIUtilityWs

import (
	"context"
	"encoding/xml"
	"github.com/TheBookPeople/avaya/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// CustomerLogin -
type CustomerLogin struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CustomerLogin"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`
}

type CustomerLoginResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CustomerLoginResponse"`

	CustomerLoginResult string `xml:"CustomerLoginResult,omitempty"`
}

type GetAnonymousSessionKey struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAnonymousSessionKey"`
}

// AnonymousSessionKeyResult -
type AnonymousSessionKeyResult struct {
	SessionKey  string `xml:"SessionKey,omitempty"`
	AnonymousID string `xml:"AnonymousID,omitempty"`
}

// GetAnonymousSessionKeyResponse -
type GetAnonymousSessionKeyResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAnonymousSessionKeyResponse"`

	GetAnonymousSessionKeyResult *AnonymousSessionKeyResult `xml:"GetAnonymousSessionKeyResult,omitempty"`
}

type GetAnonymousCustomerID struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAnonymousCustomerID"`

	LoginResult *AnonymousLoginResult `xml:"LoginResult,omitempty"`

	EmailAddress string `xml:"EmailAddress,omitempty"`

	PhoneNumber string `xml:"PhoneNumber,omitempty"`
}

type GetAnonymousCustomerIDResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAnonymousCustomerIDResponse"`

	GetAnonymousCustomerIDResult int64 `xml:"GetAnonymousCustomerIDResult,omitempty"`
}

type GetAndUpdateAnonymousCustomerID struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAndUpdateAnonymousCustomerID"`

	LoginResult *AnonymousLoginResult //`xml:"LoginResult,omitempty"`

	EmailAddress string `xml:"EmailAddress,omitempty"`

	PhoneNumber string `xml:"PhoneNumber,omitempty"`

	ThisCustomer *CICustomerReadType `xml:"ThisCustomer"`
}

type GetAndUpdateAnonymousCustomerIDResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAndUpdateAnonymousCustomerIDResponse"`

	GetAndUpdateAnonymousCustomerIDResult int64 `xml:"GetAndUpdateAnonymousCustomerIDResult,omitempty"`
}

type ExtendedCustomerLogin struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ExtendedCustomerLogin"`

	Username string `xml:"username,omitempty"`

	ExLoginPassword string `xml:"exLoginPassword,omitempty"`
}

type ExtendedCustomerLoginResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ExtendedCustomerLoginResponse"`

	ExtendedCustomerLoginResult string `xml:"ExtendedCustomerLoginResult,omitempty"`
}

type CustomerLogoff struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CustomerLogoff"`

	Username string `xml:"username,omitempty"`
}

type CustomerLogoffResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CustomerLogoffResponse"`

	CustomerLogoffResult int64 `xml:"CustomerLogoffResult,omitempty"`
}

type CustomerEndSession struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CustomerEndSession"`

	Username string `xml:"username,omitempty"`

	SessionKey string `xml:"SessionKey,omitempty"`
}

type CustomerEndSessionResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CustomerEndSessionResponse"`

	CustomerEndSessionResult int64 `xml:"CustomerEndSessionResult,omitempty"`
}

type CustomerLogoffByContactID struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CustomerLogoffByContactID"`

	Username string `xml:"username,omitempty"`

	ContactID int64 `xml:"ContactId,omitempty"`

	SessionKey string `xml:"SessionKey,omitempty"`
}

type CustomerLogoffByContactIDResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com CustomerLogoffByContactIDResponse"`

	CustomerLogoffByContactIDResult int64 `xml:"CustomerLogoffByContactIDResult,omitempty"`
}

type AdministratorLogin struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AdministratorLogin"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`
}

type AdministratorLoginResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AdministratorLoginResponse"`

	AdministratorLoginResult string `xml:"AdministratorLoginResult,omitempty"`
}

type AdministratorLogoff struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AdministratorLogoff"`

	Username string `xml:"username,omitempty"`
}

type AdministratorLogoffResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AdministratorLogoffResponse"`

	AdministratorLogoffResult int64 `xml:"AdministratorLogoffResult,omitempty"`
}

type IsContactCentreClosed struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com IsContactCentreClosed"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type IsContactCentreClosedResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com IsContactCentreClosedResponse"`

	IsContactCentreClosedResult bool `xml:"IsContactCentreClosedResult,omitempty"`
}

type GetServerTime struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetServerTime"`
}

type GetServerTimeResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetServerTimeResponse"`

	GetServerTimeResult *soap.CIDateTime `xml:"GetServerTimeResult,omitempty"`
}

type GetServerUTCTime struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetServerUTCTime"`

	SessionKey string `xml:"SessionKey,omitempty"`
}

type GetServerUTCTimeResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetServerUTCTimeResponse"`

	GetServerUTCTimeResult *soap.CIDateTime `xml:"GetServerUTCTimeResult,omitempty"`
}

type GetServerRAWTime struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetServerRAWTime"`

	SessionKey string `xml:"SessionKey,omitempty"`
}

type GetServerRAWTimeResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetServerRAWTimeResponse"`

	GetServerRAWTimeResult *soap.CIDateTime `xml:"GetServerRAWTimeResult,omitempty"`
}

type TimeStampToMilliseconds struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com TimeStampToMilliseconds"`

	Timestamp *CITimeStamp `xml:"timestamp,omitempty"`
}

type TimeStampToMillisecondsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com TimeStampToMillisecondsResponse"`

	TimeStampToMillisecondsResult int64 `xml:"TimeStampToMillisecondsResult,omitempty"`
}

type MillisecondsToTimeStamp struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com MillisecondsToTimeStamp"`

	Milliseconds int64 `xml:"milliseconds,omitempty"`

	UTCOffsetMins int32 `xml:"UTCOffsetMins,omitempty"`
}

type MillisecondsToTimeStampResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com MillisecondsToTimeStampResponse"`

	MillisecondsToTimeStampResult *CITimeStamp `xml:"MillisecondsToTimeStampResult,omitempty"`
}

type GetTotalQueued struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetTotalQueued"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

type GetTotalQueuedResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetTotalQueuedResponse"`

	GetTotalQueuedResult int64 `xml:"GetTotalQueuedResult,omitempty"`
}

type GetTotalQueuedToSkillset struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetTotalQueuedToSkillset"`

	SessionKey string `xml:"sessionKey,omitempty"`

	SkillsetID int64 `xml:"skillsetID,omitempty"`
}

type GetTotalQueuedToSkillsetResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetTotalQueuedToSkillsetResponse"`

	GetTotalQueuedToSkillsetResult int64 `xml:"GetTotalQueuedToSkillsetResult,omitempty"`
}

type AnonymousLoginResult struct {
	//XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com AnonymousLoginResult"`
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com LoginResult"`
	//	XMLName xml.Name `xml:"AnonymousLoginResult"`

	SessionKey string `xml:"SessionKey,omitempty"`

	AnonymousID int64 `xml:"AnonymousID,omitempty"`
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

type CICustomerReadType struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CICustomerReadType"`
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ThisCustomer"`

	ID int64 `xml:"id,omitempty"`

	Title string `xml:"title,omitempty"`

	FirstName string `xml:"firstName,omitempty"`

	LastName string `xml:"lastName,omitempty"`

	Username string `xml:"username,omitempty"`

	RegisterDate *soap.CIDateTime `xml:"registerDate"`

	AddressList *ArrayOfCIAddressReadType `xml:"addressList"`

	PhoneNumberList *ArrayOfCIPhoneNumberReadType `xml:"phoneNumberList"`

	EmailAddressList *ArrayOfCIEmailAddressReadType `xml:"emailAddressList"`

	CustomFieldList *ArrayOfCICustomFieldReadType `xml:"customFieldList"`

	DefaultAddress *CIAddressReadType `xml:"defaultAddress"`

	DefaultPhoneNumber *CIPhoneNumberReadType `xml:"defaultPhoneNumber"`

	DefaultEmailAddress *CIEmailAddressReadType `xml:"defaultEmailAddress"`
}

type ArrayOfCIAddressReadType struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIAddressReadType"`

	CIAddressReadType []*CIAddressReadType `xml:"CIAddressReadType"`
}

type CIAddressReadType struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIAddressReadType"`

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
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIPhoneNumberReadType"`

	CIPhoneNumberReadType []*CIPhoneNumberReadType `xml:"CIPhoneNumberReadType,omitempty"`
}

type CIPhoneNumberReadType struct {
	//		XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIPhoneNumberReadType"`

	ID int64 `xml:"id,omitempty"`

	InternationalCode string `xml:"internationalCode,omitempty"`

	AreaCode string `xml:"areaCode,omitempty"`

	Number string `xml:"number,omitempty"`

	DoNotCall bool `xml:"doNotCall,omitempty"`

	DefaultPhoneNumber bool `xml:"defaultPhoneNumber,omitempty"`

	PhoneNumberType *CIPhoneNumberType `xml:"phoneNumberType,omitempty"`
}

type ArrayOfCIEmailAddressReadType struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCIEmailAddressReadType"`

	CIEmailAddressReadType []*CIEmailAddressReadType `xml:"CIEmailAddressReadType,omitempty"`
}

type CIEmailAddressReadType struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIEmailAddressReadType"`

	ID int64 `xml:"id,omitempty"`

	EmailAddress string `xml:"emailAddress,omitempty"`

	DefaultEmailAddress bool `xml:"defaultEmailAddress,omitempty"`
}

type ArrayOfCICustomFieldReadType struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCICustomFieldReadType"`

	CICustomFieldReadType []*CICustomFieldReadType `xml:"CICustomFieldReadType,omitempty"`
}

type CICustomFieldReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CICustomFieldReadType"`

	ID int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Text string `xml:"text,omitempty"`

	IsTextVisible bool `xml:"isTextVisible,omitempty"`
}

type CITimeStamp struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CITimeStamp"`

	Day int32 `xml:"day,omitempty"`

	Month int32 `xml:"month,omitempty"`

	Year int32 `xml:"year,omitempty"`

	Hour int32 `xml:"hour,omitempty"`

	Minute int32 `xml:"minute,omitempty"`

	Second int32 `xml:"second,omitempty"`

	UTCOffsetMins int32 `xml:"UTCOffsetMins,omitempty"`
}

type Soap struct {
	client *soap.SOAPClient
}

// NewSoap -
func NewSoap(baseURL string, tls bool, auth *soap.BasicAuth) *Soap {
	url := baseURL + "/ccmmwebservices/CIUtilityWs.asmx"
	client := soap.NewSOAPClient(url, tls, auth)

	return &Soap{
		client: client,
	}
}

func (service *Soap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

func (service *Soap) CustomerLogin(ctx context.Context, request *CustomerLogin) (*CustomerLoginResponse, error) {
	response := new(CustomerLoginResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/CustomerLogin", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) GetAnonymousSessionKey(ctx context.Context, request *GetAnonymousSessionKey) (*GetAnonymousSessionKeyResponse, error) {
	response := new(GetAnonymousSessionKeyResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetAnonymousSessionKey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) GetAnonymousCustomerID(ctx context.Context, request *GetAnonymousCustomerID) (*GetAnonymousCustomerIDResponse, error) {
	response := new(GetAnonymousCustomerIDResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetAnonymousCustomerID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) GetAndUpdateAnonymousCustomerID(ctx context.Context, request *GetAndUpdateAnonymousCustomerID) (*GetAndUpdateAnonymousCustomerIDResponse, error) {
	response := new(GetAndUpdateAnonymousCustomerIDResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetAndUpdateAnonymousCustomerID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) ExtendedCustomerLogin(ctx context.Context, request *ExtendedCustomerLogin) (*ExtendedCustomerLoginResponse, error) {
	response := new(ExtendedCustomerLoginResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/ExtendedCustomerLogin", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) CustomerLogoff(ctx context.Context, request *CustomerLogoff) (*CustomerLogoffResponse, error) {
	response := new(CustomerLogoffResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/CustomerLogoff", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) CustomerEndSession(ctx context.Context, request *CustomerEndSession) (*CustomerEndSessionResponse, error) {
	response := new(CustomerEndSessionResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/CustomerEndSession", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) CustomerLogoffByContactID(ctx context.Context, request *CustomerLogoffByContactID) (*CustomerLogoffByContactIDResponse, error) {
	response := new(CustomerLogoffByContactIDResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/CustomerLogoffByContactID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) AdministratorLogin(ctx context.Context, request *AdministratorLogin) (*AdministratorLoginResponse, error) {
	response := new(AdministratorLoginResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/AdministratorLogin", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) AdministratorLogoff(ctx context.Context, request *AdministratorLogoff) (*AdministratorLogoffResponse, error) {
	response := new(AdministratorLogoffResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/AdministratorLogoff", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) IsContactCentreClosed(ctx context.Context, request *IsContactCentreClosed) (*IsContactCentreClosedResponse, error) {
	response := new(IsContactCentreClosedResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/IsContactCentreClosed", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) GetServerTime(ctx context.Context, request *GetServerTime) (*GetServerTimeResponse, error) {
	response := new(GetServerTimeResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetServerTime", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) GetServerUTCTime(ctx context.Context, request *GetServerUTCTime) (*GetServerUTCTimeResponse, error) {
	response := new(GetServerUTCTimeResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetServerUTCTime", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) GetServerRAWTime(ctx context.Context, request *GetServerRAWTime) (*GetServerRAWTimeResponse, error) {
	response := new(GetServerRAWTimeResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetServerRAWTime", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) TimeStampToMilliseconds(ctx context.Context, request *TimeStampToMilliseconds) (*TimeStampToMillisecondsResponse, error) {
	response := new(TimeStampToMillisecondsResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/TimeStampToMilliseconds", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) MillisecondsToTimeStamp(ctx context.Context, request *MillisecondsToTimeStamp) (*MillisecondsToTimeStampResponse, error) {
	response := new(MillisecondsToTimeStampResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/MillisecondsToTimeStamp", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) GetTotalQueued(ctx context.Context, request *GetTotalQueued) (*GetTotalQueuedResponse, error) {
	response := new(GetTotalQueuedResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetTotalQueued", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *Soap) GetTotalQueuedToSkillset(ctx context.Context, request *GetTotalQueuedToSkillset) (*GetTotalQueuedToSkillsetResponse, error) {
	response := new(GetTotalQueuedToSkillsetResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetTotalQueuedToSkillset", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
