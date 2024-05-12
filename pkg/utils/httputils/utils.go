package httputils

import (
	"net"
	"net/http"
	"sync"
	"time"
)

// HTTPClientPool manages a pool of http.Client instances
type HTTPClientPool struct {
	clients []*http.Client
	lock    sync.Mutex
	index   int
}

// NewHTTPClientPool creates a new pool of http.Client instances
func NewHTTPClientPool(maxClients int, timeout time.Duration) *HTTPClientPool {
	pool := &HTTPClientPool{
		clients: make([]*http.Client, maxClients),
	}
	for i := range pool.clients {
		transport := &http.Transport{
			MaxIdleConns:       10000,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: timeout,
			}).DialContext,
		}
		pool.clients[i] = &http.Client{
			Transport: transport,
			Timeout:   timeout,
		}
	}
	return pool
}

// Get retrieves a client from the pool, rotating in a round-robin fashion
func (p *HTTPClientPool) Get() *http.Client {
	p.lock.Lock()
	defer p.lock.Unlock()
	client := p.clients[p.index]
	p.index = (p.index + 1) % len(p.clients)
	return client
}
