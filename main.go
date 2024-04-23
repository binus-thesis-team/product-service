package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irvankadhafi/go-boilerplate/database"
	databasewrapper "github.com/irvankadhafi/go-boilerplate/database/wrapper"
	"github.com/irvankadhafi/go-boilerplate/internal/router"
	logrushook "github.com/irvankadhafi/go-boilerplate/pkg/logrus_hook"
	"github.com/irvankadhafi/go-boilerplate/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.AddHook(&logrushook.Trace{})
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	var err error
	maxRetry, convErr := strconv.Atoi(os.Getenv("MAX_RETRY"))
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
		SslMode:      os.Getenv("SSL_MODE"),
		TimeZone:     os.Getenv("TZ"),
		MaxRetry:     maxRetry,
		Timeout:      time.Duration(dbTimeout) * time.Second,
	}, &databasewrapper.DatabaseWrapper{})

	err = dbSql.Connect()
	fmt.Println(err)

	if err != nil {
		os.Exit(1)
		return
	}

	dbSql.SetMaxIdleConns(0)
	dbSql.SetMaxOpenConns(100)

	location, locErr := utils.SetTimeLocation("Asia/Jakarta")
	if locErr != nil {
		panic(locErr)
	}

	time.Local = location

	ginEngine := gin.Default()

	ginEngine.GET("/ping", ping)
	router.Add(ginEngine, dbSql)

	if err := ginEngine.Run(); err != nil {
		logrus.Fatal(err)
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"service_name": os.Getenv("SERVICE_NAME"),
		"mode":         os.Getenv("MODE"),
	})
}
