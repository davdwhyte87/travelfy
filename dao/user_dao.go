package dao

import (
	. "github.com/davdwhyte87/travelfy/models"
	"gopkg.in/mgo.v2/bson"
)

type UserDAO struct {

}

// FindAll ... get list of users
func (m *UserDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// FindById ... get a user by its id
func (m *UserDAO) FindById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// FindByEmail ... get a user by the email
func (m *UserDAO) FindByEmail(email string) (User, error) {
	var user User

	err := db.C(COLLECTION).Find(bson.M{"email":email}).One(&user)
	return user, err
}

// Insert a user into database
func (m *UserDAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// Delete an existing user
func (m *UserDAO) Delete(user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func (m *UserDAO) Update(user User) error {
	err := db.C(COLLECTION).Update(bson.M{"_id":user.ID}, &user)
	return err
}