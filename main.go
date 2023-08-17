package pinecone

import (
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

func NewClient(index, environment, apiKey string, debug bool) *Client {
	return &Client{
		Index:       index,
		Environment: environment,
		APIKey:      apiKey,
		Debug:       debug,
		Resty:       resty.New(),
	}
}

func (c *Client) GetHeaders() map[string]string {
	return map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"Api-Key":      c.APIKey,
	}
}

func (c *Client) BaseURL() string {
	return "https://" + c.Index + ".svc." + c.Environment + ".pinecone.io"
}

func (c *Client) DescribeIndexStats(request *DescribeIndexStatsRequest) (*DescribeIndexStatsResponse, error) {
	result := &DescribeIndexStatsResponse{}

	_, err := c.makePostRequest("/describe_index_stats", request, result)

	return result, err
}

func (c *Client) Query(request *QueryRequest) (*QueryResponse, error) {
	result := &QueryResponse{}

	_, err := c.makePostRequest("/query", request, result)

	return result, err
}

func (c *Client) Delete(request *DeleteRequest) (*DeleteResponse, error) {
	result := &DeleteResponse{}

	_, err := c.makePostRequest("/vectors/delete", request, result)

	return result, err
}

func (c *Client) Fetch(request *FetchRequest) (*FetchResponse, error) {
	result := &FetchResponse{}

	var queryParams = map[string]string{}

	// Convert ID slice to comma separated string
	if len(request.IDs) > 0 {
		queryParams["ids"] = strings.Join(request.IDs, ",")
	}

	// Check if Namespace is not nil before dereferencing
	if request.Namespace != nil {
		queryParams["namespace"] = *request.Namespace
	}

	log.Println(queryParams)
	_, err := c.makeGetRequest("/vectors/fetch", queryParams, result)

	return result, err
}

func (c *Client) Update(request *UpdateRequest) (*UpdateResponse, error) {
	result := &UpdateResponse{}

	_, err := c.makePostRequest("/vectors/update", request, result)

	return result, err
}

func (c *Client) Upsert(request *UpsertRequest) (*UpsertResponse, error) {
	result := &UpsertResponse{}

	_, err := c.makePostRequest("/vectors/upsert", request, result)

	return result, err
}

func (c *Client) makePostRequest(endpoint string, body interface{}, result interface{}) (*resty.Response, error) {
	var errResp *ErrorResponse

	request := c.Resty.SetDebug(c.Debug).
		R().
		EnableTrace().
		SetHeaders(c.GetHeaders()).
		SetBody(body).
		SetError(&errResp).
		SetResult(&result)

	resp, err := request.Post(c.BaseURL() + endpoint)

	if errResp != nil {
		return nil, errors.New(errResp.Message)
	}

	return resp, err
}

func (c *Client) makeGetRequest(endpoint string, queryParams map[string]string, result interface{}) (*resty.Response, error) {
	var errResp *ErrorResponse
	request := c.Resty.SetDebug(c.Debug).
		R().
		EnableTrace().
		SetHeaders(c.GetHeaders()).
		SetQueryParams(queryParams).
		SetError(&errResp).
		SetResult(&result)

	resp, err := request.Get(c.BaseURL() + endpoint)

	if errResp != nil {
		return nil, errors.New(errResp.Message)
	}

	return resp, err
}
