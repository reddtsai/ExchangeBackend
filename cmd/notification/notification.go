package main

import (
	"log"

	"github.com/reddtsai/ExchangeBackend/internal/notification"
)

func main() {
	engine, err := notification.New()
	if err != nil {
		log.Println(err.Error())
		return
	}
	engine.Run(":5000")
}
