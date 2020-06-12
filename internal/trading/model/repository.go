package model

import (
	"encoding/json"
	"fmt"
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
	sql := `INSERT INTO EXCHANGE.ORDER VALUES ("%s","%s","%s","%s","%s","%d")`
	sql = fmt.Sprintf(sql, order.ID.Hex(), order.Symbol, order.Type, order.Amount, order.Price, order.Unix)
	err = s.sql.insert(sql)
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
	s.sql = nil
}
