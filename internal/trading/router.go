package trading

import (
	"github.com/gin-gonic/gin"

	"github.com/reddtsai/ExchangeBackend/internal/trading/handler"
)

// NewRouter create api router
func NewRouter(gin *gin.Engine) {
	order := gin.Group("/order")
	{
		order.POST("/buy", handler.Buy)
	}
}
