package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/binus-thesis-team/product-service/internal/model"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type httpClient struct {
	client  *http.Client
	baseURL string
}

type successResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type searchResponse struct {
	Count int64   `json:"count"`
	Ids   []int64 `json:"ids"`
}

func NewHTTPRestClient(baseURL string) ProductServiceClient {
	transport := &http.Transport{
		MaxIdleConns:        100,              // Increase max idle connections if necessary
		MaxIdleConnsPerHost: 100,              // Set the max idle connections per host
		IdleConnTimeout:     15 * time.Second, // Keep idle connections alive for longer
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // Timeout for TCP connections
			KeepAlive: 30 * time.Second, // Keep-alive period for active network connections
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
	}

	return &httpClient{
		client:  client,
		baseURL: baseURL,
	}
}

func (h *httpClient) FindByProductID(ctx context.Context, id int64) (*model.Product, error) {
	url := fmt.Sprintf("%s/products/%d", h.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server error: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response successResponse[*model.Product]
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, fmt.Errorf("API call failed with response: %s", string(body))
	}

	return response.Data, nil
}

func (h *httpClient) FindProductIDsByQuery(ctx context.Context, query string) (ids []int64, count int64, err error) {
	// Construct the URL with query parameters
	url := fmt.Sprintf("%s/products?query=%s", h.baseURL, query)

	// Create the request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	// Check if the status code is what we expect
	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("server error: %s", resp.Status)
	}

	// Read and parse the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	var result searchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, 0, err
	}

	return result.Ids, result.Count, nil
}
