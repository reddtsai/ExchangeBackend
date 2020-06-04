package main

import (
	"github.com/gin-gonic/gin"

	"github.com/reddtsai/ExchangeBackend/internal/trading"
)

func main() {
	engine := gin.Default()
	trading.NewRouter(engine)
	engine.Run(":5001")

	// middleware
	// r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// authorized := r.Group("/")
}
