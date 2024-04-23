package database

import (
	"os"
	"strconv"
	"time"

	databasewrapper "github.com/irvankadhafi/go-boilerplate/database/wrapper"
)

var (
	dbSql *DbSql
)

func StartDBConnection() {

	initDBMain()
}

func closeDBConnections() {
	closeDBMain()
}

func initDBMain() {
	var err error
	maxRetry, convErr := strconv.Atoi(os.Getenv("MAX_RETRY"))
	if convErr != nil {
		maxRetry = 3
	}

	dbTimeout, convErr := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if convErr != nil {
		dbTimeout = 120
	}

	dbSql = InitConnectionDB("postgres", Config{
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		SslMode:      os.Getenv("DB_SSL_MODE"),
		TimeZone:     os.Getenv("DB_TIMEZONE"),
		MaxRetry:     maxRetry,
		Timeout:      time.Duration(dbTimeout) * time.Second,
	}, &databasewrapper.DatabaseWrapper{})

	err = dbSql.Connect()

	if err != nil {
		os.Exit(1)
		return
	}

	dbSql.SetMaxIdleConns(0)
	dbSql.SetMaxOpenConns(100)

}

func closeDBMain() {
	if err := dbSql.ClosePmConnection(); err != nil {
		return
	}
}
