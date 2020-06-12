package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"github.com/reddtsai/ExchangeBackend/internal/trading/model"
)

const (
	limitBuy  = "1"
	limitSell = "2"
)

const (
	btcusdt = "BTCUSDT"
)

// Buy buy service
func Buy(c *gin.Context) {
	// request
	var req model.OrderReq
	resp := &model.OrderResp{
		Status: "err",
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// buy limite order
	data := &model.Order{
		ID:     bson.NewObjectId(),
		Symbol: btcusdt,
		Type:   limitBuy,
		Price:  req.Price,
		Amount: req.Amount,
		Unix:   time.Now().Unix(),
	}
	err := model.Repository.PlaceOrder(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Status = "ok"
	resp.ID = data.ID
	c.JSON(http.StatusOK, resp)
}

// Sell sell service
func Sell(c *gin.Context) {
	// request
	var req model.OrderReq
	resp := &model.OrderResp{
		Status: "err",
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// sell limite order
	data := &model.Order{
		ID:     bson.NewObjectId(),
		Symbol: btcusdt,
		Type:   limitSell,
		Price:  req.Price,
		Amount: req.Amount,
		Unix:   time.Now().Unix(),
	}
	err := model.Repository.PlaceOrder(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Status = "ok"
	resp.ID = data.ID
	c.JSON(http.StatusOK, resp)
}
