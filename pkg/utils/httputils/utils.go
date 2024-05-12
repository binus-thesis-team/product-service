package httputils

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

type HTTPClientOptions struct {
	Timeout             time.Duration
	MaxIdleConns        int  // Maximum number of idle connections across all hosts
	MaxIdleConnsPerHost int  // Maximum number of idle connections per host
	UseCircuitBreaker   bool // Configurable use of circuit breaker

	// RetryCount retry the operation if found error.
	// When set to <= 1, then it means no retry
	RetryCount int

	// RetryInterval next interval for retry.
	RetryInterval time.Duration
}

type PooledHTTPClient struct {
	client  *http.Client
	options *HTTPClientOptions
}

func NewPooledHTTPClient(options *HTTPClientOptions) *PooledHTTPClient {
	transport := &http.Transport{
		MaxIdleConns:        options.MaxIdleConns,
		MaxIdleConnsPerHost: options.MaxIdleConnsPerHost,
		IdleConnTimeout:     90 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   options.Timeout,
	}

	return &PooledHTTPClient{
		client:  client,
		options: options,
	}
}

// HTTPRequestInvoker function type for retrying HTTP requests
type HTTPRequestInvoker func() (*http.Response, error)

// retryableInvoke that handles retries for HTTP requests.
func (h *PooledHTTPClient) retryableInvoke(req *http.Request, retryCount int, retryInterval time.Duration) (*http.Response, error) {
	var response *http.Response
	err := Retry(retryCount, retryInterval, func() error {
		var err error
		response, err = h.client.Do(req)
		if err != nil {
			return err // Network errors etc
		}
		if response.StatusCode >= 500 { // Retry on server errors
			return NewRetryStopper(fmt.Errorf("server error: %d %s", response.StatusCode, response.Status))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}

// Do makes an HTTP request with retry logic
func (h *PooledHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if h.options.UseCircuitBreaker {
		response, err := h.doWithCircuitBreaker(req)
		if err != nil {
			logrus.Warnf("Request failed even after retries and circuit breaker: %v", err)
			return nil, err
		}
		return response, nil
	}
	return h.retryableInvoke(req, h.options.RetryCount, h.options.RetryInterval)
}

// doWithCircuitBreaker handles requests using the Hystrix circuit breaker
func (h *PooledHTTPClient) doWithCircuitBreaker(req *http.Request) (*http.Response, error) {
	commandName := "http_request_" + req.Method + "_" + req.URL.Path
	var response *http.Response
	var err error

	errChan := hystrix.GoC(req.Context(), commandName, func(ctx context.Context) error {
		response, err = h.client.Do(req)
		if err != nil {
			return err
		}
		if response.StatusCode >= 400 { // Treat 4xx and 5xx as reasons to trip the circuit
			return &HTTPError{StatusCode: response.StatusCode}
		}
		return nil
	}, nil)

	// Wait for the command to complete
	err = <-errChan
	if err != nil {
		return nil, err
	}
	return response, nil
}

type HTTPError struct {
	StatusCode int
}

func (e *HTTPError) Error() string {
	return http.StatusText(e.StatusCode)
}

// RetryStopper allows controlling the retry mechanism by stopping it when necessary.
type RetryStopper struct {
	error
}

// NewRetryStopper wraps an error that should stop retries.
func NewRetryStopper(err error) RetryStopper {
	return RetryStopper{err}
}

// Retry function to handle retries with exponential backoff.
func Retry(attempts int, sleep time.Duration, fn func() error) error {
	err := fn()
	if err != nil {
		if _, ok := err.(RetryStopper); ok {
			return err // Stop retrying if it's a RetryStopper error.
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return Retry(attempts, sleep*2, fn)
		}
		return err
	}
	return nil
}
