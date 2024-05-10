package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/binus-thesis-team/product-service/internal/model"
	"io/ioutil"
	"net/http"
)

type httpClient struct {
	baseURL string
}

type successResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

func NewHTTPRestClient(baseURL string) ProductServiceClient {
	return &httpClient{baseURL: baseURL}
}
func (h *httpClient) FindByProductID(ctx context.Context, id int64) (*model.Product, error) {
	url := fmt.Sprintf("%s/products/%d", h.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
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

func (h *httpClient) SearchAllProducts(ctx context.Context, query string) (ids []int64, count int64, err error) {
	//TODO implement me
	panic("implement me")
}
