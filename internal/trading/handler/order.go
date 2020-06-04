package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Order order model
type Order struct {
	Name string `binding:"required"`
}

// Buy create buy order
func Buy(c *gin.Context) {
	var o Order
	if err := c.ShouldBindWith(&o, binding.FormPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "err",
			"err":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   "redd",
	})
}
