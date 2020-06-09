package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Storage trading storage
type Storage struct {
	mongo *mongoDB
	mq    *messageQueue
	sql   *mysqlDB
}

// Repository trading storage instance
var Repository *Storage

// NewRepository create trading storage instance
func NewRepository() error {
	mongo, err := newMongoDB()
	if err != nil {
		return err
	}
	mq, err := newNSQ()
	if err != nil {
		return err
	}
	mysql, err := newMySQL()
	if err != nil {
		return err
	}
	Repository = &Storage{
		mongo: mongo,
		mq:    mq,
		sql:   mysql,
	}
	return nil
}

// PlaceOrder place an order
func (s *Storage) PlaceOrder(order *Order) error {
	// mongo
	err := s.mongo.insert("order", order)
	if err != nil {
		return err
	}
	// nsq
	buf, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = s.mq.push("order", buf)
	if err != nil {
		return err
	}
	// mysql
	var sb strings.Builder
	sb.WriteString(`INSERT INTO EXCHANGE.ORDER VALUES (`)
	sb.WriteString(`"` + order.ID.Hex() + `",`)
	sb.WriteString(`"` + order.Symbol + `",`)
	sb.WriteString(`"` + order.Type + `",`)
	sb.WriteString(`"` + order.Amount + `",`)
	sb.WriteString(`"` + order.Price + `",`)
	sb.WriteString(strconv.FormatInt(order.Unix, 10) + `)`)
	fmt.Println(sb.String())
	err = s.sql.insert(sb.String())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Close close trading repository
func (s Storage) Close() {
	s.mongo.close()
	s.sql.close()
	s.mq.close()
	s.mongo = nil
	s.mq = nil
}
