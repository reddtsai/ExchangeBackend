package notification

import (
	"github.com/gin-gonic/gin"

	"github.com/reddtsai/go-ringing"
)

// New create notification service
func New() (*gin.Engine, error) {
	// ringing
	n := &ringing.NSQSetting{
		NSQAddr:      "192.168.1.73:4161",
		Topics:       []string{"order"},
		TopicChannel: "ch1",
	}
	c := ringing.NewConfig(n)
	r, err := ringing.New(c)
	if err != nil {
		return nil, err
	}

	// gin
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		r.HandleRequest(c.Writer, c.Request)
	})

	return g, nil
}
