package trading

import (
	"github.com/gin-gonic/gin"

	"github.com/reddtsai/ExchangeBackend/internal/trading/handler"
	"github.com/reddtsai/ExchangeBackend/internal/trading/model"
)

// New create trading service
func New() (*gin.Engine, error) {
	engine := gin.Default()
	// router
	order := engine.Group("/order")
	{
		order.POST("/buy", handler.Buy)
		order.POST("/sell", handler.Sell)
	}

	// model
	err := model.NewRepository()
	if err != nil {
		return nil, err
	}

	return engine, nil
}

// Close close trading service
func Close() {
	model.Repository.Close()
}
