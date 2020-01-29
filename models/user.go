package models

import "gopkg.in/mgo.v2/bson"

// User ... this is a representation of a user on the server
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name" validate:"required,min=2,max=100"`
	Email string  `bson:"email" json:"email" validate:"required,email"`
	Password string  `bson:"password" json:"password" validate:"required,alpha"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Bio string        `bson:"bio" json:"bio"`
	Location string `bson:"location" json:"location"`
	Confirmed bool `bson:"confirmed" json:"confirmed"`
	IsDriver bool `bson:"is_driver" json:"is_driver"`
}


