package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/reddtsai/ExchangeBackend/internal/trading"
)

func main() {
	engine, err := trading.New()
	if err != nil {
		log.Println(err.Error())
		return
	}
	engine.Run(":5001")

	// middleware
	// r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// authorized := r.Group("/")
}
