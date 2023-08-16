package main

import (
	"github.com/go-resty/resty/v2"
)

const (
)

var restyClient = resty.New().R().EnableTrace()

type Client struct {
	Index string
	Environment string
    APIKey      string
}

func (c *Client) GetHeaders() map[string]string {
    return map[string]string{
        "Accept": "application/json",
        "Content-Type": "application/json",
        "Api-Key": c.APIKey,
    }
}

func (c *Client) BaseURL() string {
	return "https://" + c.Index + ".svc." + c.Environment + ".pinecone.io"
}

func (c *Client) DescribeIndexStats (body *DescribeIndexStatsRequest) (*DescribeIndexStatsResponse, error) {
	result := &DescribeIndexStatsResponse{}

	err := c.makePostRequest("/describe_index_stats", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Query(body *QueryRequest) (*QueryResponse, error) {
	result := &QueryResponse{}

    err := c.makePostRequest("/query", body, result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func (c *Client) Delete(body *DeleteRequest) (*DeleteResponse, error) {
	result := &DeleteResponse{}

	err := c.makePostRequest("/delete", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Fetch(body *FetchRequest) (*FetchResponse, error) {
	result := &FetchResponse{}

    err := c.makeGetRequest("/fetch", result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func (c *Client) Update(body *UpdateRequest) (*UpdateResponse, error) {
	result := &UpdateResponse{}

	err := c.makePostRequest("/update", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Upsert(body *UpsertRequest) (*UpsertResponse, error) {

	result := &UpsertResponse{}

    err := c.makePostRequest("/upsert", body, result)
    if err != nil {
        return nil, err
    }
    return result, nil
}


func (c *Client) makePostRequest(endpoint string, body, result interface{}) error {
    request := restyClient.
        SetHeaders(c.GetHeaders()).
        SetBody(body).	
        SetResult(result)

    _, err := request.Post(c.BaseURL() + endpoint)
    
    return err
}

func (c *Client) makeGetRequest(endpoint string, result interface{}) error {
    request := restyClient.
        SetHeaders(c.GetHeaders()).
        SetResult(result)

    _, err := request.Get(c.BaseURL() + endpoint)
    
    return err
}
