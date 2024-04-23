package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/irvankadhafi/go-boilerplate/database"
	databasewrapper "github.com/irvankadhafi/go-boilerplate/database/wrapper"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	maxRetry, convErr := strconv.Atoi(os.Getenv("MAX_RETRU"))
	if convErr != nil {
		maxRetry = 3
	}

	dbTimeout, convErr := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if convErr != nil {
		dbTimeout = 120
	}

	dbSql := database.InitConnectionDB("postgres", database.Config{
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

	dbSql.Connect()

	if err := migrationUp(dbSql); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("migrate up success")
}

func migrationUp(db *database.DbSql) error {
	migrate.SetTable("migrations")

	_, err := migrate.Exec(db.SqlDb, "postgres", getFileMigrationSource(), migrate.Up)
	if err != nil {
		return fmt.Errorf("fail execute migrations: %v", err)
	}

	return nil
}

func getFileMigrationSource() *migrate.FileMigrationSource {
	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join("database", "schema_migration"),
	}
	return migrations
}
