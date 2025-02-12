package tilastokeskus

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-tilastokeskus/utils"
)

func (c *Client) NewSendStatisticsRequest() SendStatisticsRequest {
	r := SendStatisticsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	r.requestHeader = r.NewRequestHeader()
	return r
}

type SendStatisticsRequest struct {
	client        *Client
	queryParams   *SendStatisticsQueryParams
	pathParams    *SendStatisticsPathParams
	method        string
	headers       http.Header
	requestBody   SendStatisticsRequestBody
	requestHeader SendStatisticsRequestHeader
}

func (r SendStatisticsRequest) NewQueryParams() *SendStatisticsQueryParams {
	return &SendStatisticsQueryParams{}
}

type SendStatisticsQueryParams struct {
}

func (p SendStatisticsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SendStatisticsRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r SendStatisticsRequest) NewPathParams() *SendStatisticsPathParams {
	return &SendStatisticsPathParams{}
}

type SendStatisticsPathParams struct {
}

func (p *SendStatisticsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SendStatisticsRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *SendStatisticsRequest) SetMethod(method string) {
	r.method = method
}

func (r *SendStatisticsRequest) Method() string {
	return r.method
}

func (r SendStatisticsRequest) NewRequestHeader() SendStatisticsRequestHeader {
	return SendStatisticsRequestHeader{}
}

func (r *SendStatisticsRequest) RequestHeader() *SendStatisticsRequestHeader {
	return &r.requestHeader
}

func (r *SendStatisticsRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type SendStatisticsRequestHeader struct{}

func (r SendStatisticsRequest) NewRequestBody() SendStatisticsRequestBody {
	return SendStatisticsRequestBody{
		DataTransfer: DataTransfer{
			UserID:   r.client.Username(),
			Password: r.client.Password(),
			Flags:    "",
			DataID:   8313, // fixed value
			Data: Data{
				Majoitustilasto{
					Xmlns: "http://tempuri.org/majoitustilasto.xsd",
				},
			},
		},
	}
}

type SendStatisticsRequestBody struct {
	DataTransfer DataTransfer `xml:"stat:DataTransfer"`
}

func (r *SendStatisticsRequest) RequestBody() *SendStatisticsRequestBody {
	return &r.requestBody
}

func (r *SendStatisticsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SendStatisticsRequest) SetRequestBody(body SendStatisticsRequestBody) {
	r.requestBody = body
}

func (r *SendStatisticsRequest) NewResponseBody() *SendStatisticsResponseBody {
	return &SendStatisticsResponseBody{}
}

type SendStatisticsResponseBody struct {
	XMLName              xml.Name             `xml:"Body"`
	DataTransferResponse DataTransferResponse `xml:"DataTransferResponse"`
}

func (r *SendStatisticsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("StatTrans.asmx", r.PathParams())
	return &u
}

func (r *SendStatisticsRequest) Do() (SendStatisticsResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)

	return *responseBody, err
}

type FolioDetailsRequest struct {
	XMLName      xml.Name `xml:"HcsSendStatistics"`
	Xmlns        string   `xml:"xmlns,attr"`
	PropertyCode string   `xml:"PropertyCode"`
	UniqueID     string   `xml:"UniqueID"`
	PMSNumber    string   `xml:"PMSNumber"`
	Language     string   `xml:"Language"`
}
