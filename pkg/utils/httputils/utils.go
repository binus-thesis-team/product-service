package httputils

import (
	"net"
	"net/http"
	"time"
)

// HTTPClientPool manages a pool of http.Client instances
type HTTPClientPool struct {
	pool chan *http.Client
}

// NewHTTPClientPool creates a new pool of http.Client instances
func NewHTTPClientPool(maxClients int, timeout time.Duration) *HTTPClientPool {
	pool := make(chan *http.Client, maxClients)
	for i := 0; i < maxClients; i++ {
		client := &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
			},
			Timeout: timeout,
		}
		pool <- client
	}
	return &HTTPClientPool{pool: pool}
}

func (p *HTTPClientPool) Get() *http.Client {
	return <-p.pool
}

func (p *HTTPClientPool) Put(client *http.Client) {
	p.pool <- client
}
