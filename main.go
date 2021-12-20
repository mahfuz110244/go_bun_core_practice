package main

import (
	"fmt"
	"go-echo-bun-crud-service/schema"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	toml "github.com/pelletier/go-toml"
	"github.com/uptrace/bun"
)

const API_PATH = "/api/v1"
const SetMaxOpenConns = 20
const LIMIT = 10

var config *toml.Tree

//var db *sql.DB
var db *bun.DB

//var db *gorm.DB
var database *schema.Queries

// server is used to implement helloworld.GreeterServer.

func main() {
	//Load config file
	var err error
	config, err = toml.LoadFile("config.ini")
	if err != nil {
		log.Fatalf("Toml Error: %v", err)
	}

	//DB connection
	dbconfig := config.Get("database").(*toml.Tree)
	db = dbconn(dbconfig)
	defer db.Close()
	//database = schema.New(db)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", status)

	e.GET("/status", getStatus)
	e.GET("/status/:id", getStatusById)
	e.DELETE("/status/:id", deleteStatusById)
	e.POST("/status", SaveStatus)
	e.PUT("/status/:id", updateStatus)
	e.PUT("/status", bulkUpdateStatus)

	// Start server
	address := fmt.Sprintf("%s:%s", config.Get("host").(string), config.Get("port").(string))
	e.Logger.Fatal(e.Start(address))

}
