package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/binus-thesis-team/product-service/internal/model"
	"github.com/binus-thesis-team/product-service/pkg/utils/httputils"
	"io/ioutil"
	"net/http"
	"time"
)

type httpClient struct {
	clients *httputils.HTTPClientPool
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
	httpClientPool := httputils.NewHTTPClientPool(200, 10*time.Second)

	return &httpClient{
		baseURL: baseURL,
		clients: httpClientPool,
	}
}

func (h *httpClient) FindByProductID(ctx context.Context, id int64) (*model.Product, error) {
	client := h.clients.Get()
	defer h.clients.Put(client)

	url := fmt.Sprintf("%s/products/%d", h.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
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
	client := h.clients.Get()
	defer h.clients.Put(client)

	// Construct the URL with query parameters
	url := fmt.Sprintf("%s/products?query=%s", h.baseURL, query)

	// Create the request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := client.Do(req)
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
