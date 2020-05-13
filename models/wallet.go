package models

import "gopkg.in/mgo.v2/bson"

// Wallet ... Data representation of a users wallet
type Wallet struct {
	ID bson.ObjectId 
	UserID bson.ObjectId
	// Riders []bson.ObjectId
	
	CreatedAt string
	UpdatedAt string
}