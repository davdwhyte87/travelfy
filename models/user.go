package models

import "gopkg.in/mgo.v2/bson"

// User ... this is a representation of a user on the server
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name" validate:"required,alpha"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Bio string        `bson:"description" json:"description" validate:"required,alpha"`
}