package grpcutils

import (
	"context"
	"encoding/base64"
	"runtime"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/imdario/mergo"
	grpcpool "github.com/processout/grpc-go-pool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// NewGRPCConnection establish a new grpc connection (based on pool)
func NewGRPCConnection(target string, poolSetting *GRPCConnectionPoolSetting, dialOptions ...grpc.DialOption) (*grpcpool.Pool, error) {
	poolSetting = applyGRPCConnectionPoolSetting(poolSetting)
	pool, err := grpcpool.New(func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(target, dialOptions...)
		if err != nil {
			logrus.Errorf("Error : %v", err)
			return nil, err
		}

		return conn, err
	}, poolSetting.MaxIdle, poolSetting.MaxActive, poolSetting.IdleTimeout, poolSetting.MaxConnLifetime)
	if err != nil {
		logrus.Errorf("Error : %v", err)
		return nil, err
	}
	return pool, nil
}

// GRPCConnectionPool wrapper type for pooled grpc connection
type GRPCConnectionPool struct {
	Conn *grpcpool.Pool
}

// GRPCConnectionPoolSetting if set, then treat as pooled connection
type GRPCConnectionPoolSetting struct {
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration
}

// defaultGRPCConnectionPoolSetting is a single connection setting
var defaultGRPCConnectionPoolSetting = &GRPCConnectionPoolSetting{
	MaxIdle:         10,
	MaxActive:       20,
	IdleTimeout:     0,
	MaxConnLifetime: 0,
}
var defaultPooledIdleTimeout = 1 * time.Second
var defaultPooledMaxConnLifetime = 60 * time.Minute

func applyGRPCConnectionPoolSetting(opts *GRPCConnectionPoolSetting) *GRPCConnectionPoolSetting {
	if opts == nil {
		return defaultGRPCConnectionPoolSetting
	}
	// if error occurs, also return options from input
	_ = mergo.Merge(opts, *defaultGRPCConnectionPoolSetting)

	// give default for wrong setting on pooled
	if opts.MaxActive > 1 && opts.MaxIdle > 1 {
		if opts.MaxConnLifetime <= 0 {
			opts.MaxConnLifetime = defaultPooledMaxConnLifetime
		}
		if opts.IdleTimeout <= 0 {
			opts.IdleTimeout = defaultPooledIdleTimeout
		}
	}
	return opts
}

// GRPCUnaryInterceptorOptions wrapper options for the grpc connection
type GRPCUnaryInterceptorOptions struct {
	// UseCircuitBreaker flag if the connection will implement a circuit breaker
	UseCircuitBreaker bool

	// RetryCount retry the operation if found error.
	// When set to <= 1, then it means no retry
	RetryCount int

	// RetryInterval next interval for retry.
	RetryInterval time.Duration

	// Timeout value, will return context deadline exceeded when the operation exceeds the duration
	Timeout time.Duration
}

var defaultGRPCUnaryInterceptorOptions = &GRPCUnaryInterceptorOptions{
	UseCircuitBreaker: false,
	RetryCount:        0,
	RetryInterval:     20 * time.Millisecond,
	Timeout:           1 * time.Second,
}

// UnaryClientInterceptor wrapper with circuit breaker, retry, timeout and metadata logging
func UnaryClientInterceptor(opts *GRPCUnaryInterceptorOptions) grpc.UnaryClientInterceptor {
	o := applyGRPCUnaryInterceptorOptions(opts)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := context.WithTimeout(ctx, o.Timeout)
		defer cancel()

		ctx = metadata.AppendToOutgoingContext(ctx, "caller", getCallerName(5))

		if o.UseCircuitBreaker {
			success := make(chan bool, 1)
			errC := hystrix.GoC(ctx, method, func(ctx context.Context) error {
				err := o.retryableInvoke(ctx, method, req, reply, cc, invoker, opts...)
				if err == nil {
					success <- true
				}
				return err
			}, nil)

			select {
			case out := <-success:
				logrus.Debugf("success %v", out)
				return nil
			case err := <-errC:
				logrus.Warnf("failed %s", err)
				return err
			}
		}

		return o.retryableInvoke(ctx, method, req, reply, cc, invoker, opts...)
	}
}

func UnaryClientInterceptorWithBasicAuth(opts *GRPCUnaryInterceptorOptions, username, password string) grpc.UnaryClientInterceptor {
	o := applyGRPCUnaryInterceptorOptions(opts)
	basicAuth := basicAuthHeader(username, password)

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := context.WithTimeout(ctx, o.Timeout)
		defer cancel()

		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", basicAuth)
		ctx = metadata.AppendToOutgoingContext(ctx, "caller", getCallerName(5))

		if o.UseCircuitBreaker {
			success := make(chan bool, 1)
			errC := hystrix.GoC(ctx, method, func(ctx context.Context) error {
				err := o.retryableInvoke(ctx, method, req, reply, cc, invoker, opts...)
				if err == nil {
					success <- true
				}
				return err
			}, nil)

			select {
			case out := <-success:
				logrus.Debugf("success %v", out)
				return nil
			case err := <-errC:
				logrus.Warnf("failed %s", err)
				return err
			}
		}

		return o.retryableInvoke(ctx, method, req, reply, cc, invoker, opts...)
	}
}

func applyGRPCUnaryInterceptorOptions(opts *GRPCUnaryInterceptorOptions) *GRPCUnaryInterceptorOptions {
	if opts == nil {
		return defaultGRPCUnaryInterceptorOptions
	}
	// if error occurs, also return options from input
	_ = mergo.Merge(opts, *defaultGRPCUnaryInterceptorOptions)
	return opts
}

func (o *GRPCUnaryInterceptorOptions) retryableInvoke(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return Retry(o.RetryCount, o.RetryInterval, func() error {
		err := invoker(ctx, method, req, reply, cc, opts...)

		if status.Code(err) != codes.Unavailable { // stop retrying unless Unavailable
			return NewRetryStopper(err)
		}
		return err
	})
}

// RetryStopper :nodoc:
type RetryStopper struct {
	error
}

// Retry :nodoc:
func Retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(RetryStopper); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

func NewRetryStopper(err error) RetryStopper {
	return RetryStopper{err}
}

func getCallerName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}

	return "failed to identify method caller"
}

func basicAuthHeader(username, password string) string {
	auth := username + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + encodedAuth
}
