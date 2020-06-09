package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"github.com/reddtsai/ExchangeBackend/internal/trading/model"
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
	// order
	// nosql
	data := &model.Order{
		ID:     bson.NewObjectId(),
		Symbol: "BTC",
		Type:   "1",
		Price:  req.Price,
		Amount: req.Amount,
		Unix:   time.Now().Unix(),
	}
	err := model.Repository.PlaceOrder(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// err := model.Repository.DB.Insert(model.OrderCollection, data)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, resp)
	// 	return
	// }
	// // queue
	// err = model.Repository.MQ.Send(model.OrderTopic, data)
	// if err != nil {
	// 	// TODO
	// 	// remove db order
	// 	c.JSON(http.StatusBadRequest, resp)
	// 	return
	// }
	// response
	resp.Status = "ok"
	resp.ID = data.ID
	c.JSON(http.StatusOK, resp)
}
