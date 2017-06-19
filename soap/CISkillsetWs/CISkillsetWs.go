package CISkillsetWs

import (
	"context"
	"encoding/xml"
	"github.com/TheBookPeople/avaya/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// GetAllSkillsets -
type GetAllSkillsets struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAllSkillsets"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetAllSkillsetsResponse -
type GetAllSkillsetsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAllSkillsetsResponse"`

	GetAllSkillsetsResult *CIMultipleSkillsetsReadType `xml:"GetAllSkillsetsResult,omitempty"`
}

// GetAllWebSkillsets -
type GetAllWebSkillsets struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAllWebSkillsets"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetAllWebSkillsetsResponse -
type GetAllWebSkillsetsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAllWebSkillsetsResponse"`

	GetAllWebSkillsetsResult *CIMultipleSkillsetsReadType `xml:"GetAllWebSkillsetsResult,omitempty"`
}

// GetAllOutboundSkillsets -
type GetAllOutboundSkillsets struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAllOutboundSkillsets"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetAllOutboundSkillsetsResponse -
type GetAllOutboundSkillsetsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAllOutboundSkillsetsResponse"`

	GetAllOutboundSkillsetsResult *CIMultipleSkillsetsReadType `xml:"GetAllOutboundSkillsetsResult,omitempty"`
}

// GetAllEmailSkillsets -
type GetAllEmailSkillsets struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAllEmailSkillsets"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetAllEmailSkillsetsResponse -
type GetAllEmailSkillsetsResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetAllEmailSkillsetsResponse"`

	GetAllEmailSkillsetsResult *CIMultipleSkillsetsReadType `xml:"GetAllEmailSkillsetsResult,omitempty"`
}

// GetSkillsetByID -
type GetSkillsetByID struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetSkillsetByID"`

	ID int64 `xml:"id,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetSkillsetByIDResponse -
type GetSkillsetByIDResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetSkillsetByIDResponse"`

	GetSkillsetByIDResult *CISkillsetReadType `xml:"GetSkillsetByIDResult,omitempty"`
}

// GetSkillsetByName -
type GetSkillsetByName struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetSkillsetByName"`

	SkillsetName string `xml:"skillsetName,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// GetSkillsetByNameResponse -
type GetSkillsetByNameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com GetSkillsetByNameResponse"`

	GetSkillsetByNameResult *CISkillsetReadType `xml:"GetSkillsetByNameResult,omitempty"`
}

// ReadSkillsetByName -
type ReadSkillsetByName struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadSkillsetByName"`

	SkillsetName string `xml:"skillsetName,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// ReadSkillsetByNameResponse -
type ReadSkillsetByNameResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com ReadSkillsetByNameResponse"`

	ReadSkillsetByNameResult *CISkillsetReadType `xml:"ReadSkillsetByNameResult,omitempty"`
}

// IsSkillsetInService -
type IsSkillsetInService struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com IsSkillsetInService"`

	SkillsetID int64 `xml:"skillsetID,omitempty"`

	SessionKey string `xml:"sessionKey,omitempty"`
}

// IsSkillsetInServiceResponse -
type IsSkillsetInServiceResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com IsSkillsetInServiceResponse"`

	IsSkillsetInServiceResult bool `xml:"IsSkillsetInServiceResult,omitempty"`
}

// IsSkillsetNameInService -
type IsSkillsetNameInService struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com IsSkillsetNameInService"`

	SkillsetName string `xml:"skillsetName,omitempty"`
}

// IsSkillsetNameInServiceResponse -
type IsSkillsetNameInServiceResponse struct {
	XMLName xml.Name `xml:"http://webservices.ci.ccmm.applications.nortel.com IsSkillsetNameInServiceResponse"`

	IsSkillsetNameInServiceResult bool `xml:"IsSkillsetNameInServiceResult,omitempty"`
}

// CIMultipleSkillsetsReadType -
type CIMultipleSkillsetsReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CIMultipleSkillsetsReadType"`

	SkillsetList *ArrayOfCISkillsetReadType `xml:"skillsetList,omitempty"`
}

// ArrayOfCISkillsetReadType -
type ArrayOfCISkillsetReadType struct {
	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com ArrayOfCISkillsetReadType"`

	CISkillsetReadType []*CISkillsetReadType `xml:"CISkillsetReadType,omitempty"`
}

// CISkillsetReadType -
type CISkillsetReadType struct {
	//	XMLName xml.Name `xml:"http://datatypes.ci.ccmm.applications.nortel.com CISkillsetReadType"`

	ID int64 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Tag string `xml:"tag,omitempty"`
}

// Soap -
type Soap struct {
	client *soap.SOAPClient
}

// NewSoap -
func NewSoap(baseURL string, tls bool, auth *soap.BasicAuth) *Soap {
	url := baseURL + "/ccmmwebservices/CISkillsetWs.asmx"
	client := soap.NewSOAPClient(url, tls, auth)

	return &Soap{
		client: client,
	}
}

// SetHeader -
func (service *Soap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

// GetAllSkillsets -
func (service *Soap) GetAllSkillsets(ctx context.Context, request *GetAllSkillsets) (*GetAllSkillsetsResponse, error) {
	response := new(GetAllSkillsetsResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetAllSkillsets", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllWebSkillsets -
func (service *Soap) GetAllWebSkillsets(ctx context.Context, request *GetAllWebSkillsets) (*GetAllWebSkillsetsResponse, error) {
	response := new(GetAllWebSkillsetsResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetAllWebSkillsets", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllOutboundSkillsets -
func (service *Soap) GetAllOutboundSkillsets(ctx context.Context, request *GetAllOutboundSkillsets) (*GetAllOutboundSkillsetsResponse, error) {
	response := new(GetAllOutboundSkillsetsResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetAllOutboundSkillsets", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllEmailSkillsets -
func (service *Soap) GetAllEmailSkillsets(ctx context.Context, request *GetAllEmailSkillsets) (*GetAllEmailSkillsetsResponse, error) {
	response := new(GetAllEmailSkillsetsResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetAllEmailSkillsets", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetSkillsetByID -
func (service *Soap) GetSkillsetByID(ctx context.Context, request *GetSkillsetByID) (*GetSkillsetByIDResponse, error) {
	response := new(GetSkillsetByIDResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetSkillsetByID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetSkillsetByName -
func (service *Soap) GetSkillsetByName(ctx context.Context, request *GetSkillsetByName) (*GetSkillsetByNameResponse, error) {
	response := new(GetSkillsetByNameResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/GetSkillsetByName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ReadSkillsetByName -
func (service *Soap) ReadSkillsetByName(ctx context.Context, request *ReadSkillsetByName) (*ReadSkillsetByNameResponse, error) {
	response := new(ReadSkillsetByNameResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/ReadSkillsetByName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// IsSkillsetInService -
func (service *Soap) IsSkillsetInService(ctx context.Context, request *IsSkillsetInService) (*IsSkillsetInServiceResponse, error) {
	response := new(IsSkillsetInServiceResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/IsSkillsetInService", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// IsSkillsetNameInService -
func (service *Soap) IsSkillsetNameInService(ctx context.Context, request *IsSkillsetNameInService) (*IsSkillsetNameInServiceResponse, error) {
	response := new(IsSkillsetNameInServiceResponse)
	err := service.client.Call(ctx, "http://webservices.ci.ccmm.applications.nortel.com/IsSkillsetNameInService", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
