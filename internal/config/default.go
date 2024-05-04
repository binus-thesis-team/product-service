package config

import "time"

const (
	DefaultDatabaseMaxIdleConns    = 3
	DefaultDatabaseMaxOpenConns    = 5
	DefaultDatabaseConnMaxLifetime = 1 * time.Hour
	DefaultDatabasePingInterval    = 1 * time.Second
	DefaultDatabaseRetryAttempts   = 3
	DefaultDatabaseTimeout         = 120

	DefaultRedisCacheTTL = 15 * time.Minute
)
