package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Order order model
type Order struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Symbol string        `bson:"symbol" json:"symbol"`
	Type   string        `bson:"type" json:"type"`
	Amount string        `bson:"amount" json:"amount"`
	Price  string        `bson:"price" json:"price"`
	Unix   int64         `bson:"ts" json:"ts"`
}

// OrderReq order request
type OrderReq struct {
	Symbol string `form:"symbol" json:"symbol" binding:"required"`
	Amount string `form:"amount" json:"amount" binding:"required"`
	Price  string `form:"price" json:"price" binding:"required"`
}

// OrderResp order response
type OrderResp struct {
	Status string        `json:"status"`
	ID     bson.ObjectId `json:"id,omitempty"`
}
