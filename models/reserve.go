package models

import (

	"gopkg.in/mgo.v2/bson"
)

// Reserve ... Data representation of a th system reserve
type Reserve struct {
	ID         bson.ObjectId
	LiquidCash int
	Balance    int
	Name       string
	CreatedAt  string
	UpdatedAt  string
}
