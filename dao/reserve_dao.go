package dao

import (
	"github.com/davdwhyte87/travelfy/models"
	"gopkg.in/mgo.v2/bson"
)

// ReserveDAO ...
type ReserveDAO struct {
}

// FindByName ... get a user by the email
func (m *ReserveDAO) FindByName(name string) (models.Reserve, error) {
	var reserve models.Reserve
	err := db.C(COLLECTION).Find(bson.M{"name": name}).One(&reserve)
	return reserve, err
}

// Update an existing user
func (m *ReserveDAO) Update(reserve models.Reserve) error {
	err := db.C(COLLECTION).Update(bson.M{"id":reserve.ID}, &reserve)
	return err
}

// Insert a user into database
func (m *ReserveDAO) Insert(reserve models.Reserve) error {
	err := db.C(COLLECTION).Insert(&reserve)
	return err
}

// // FindAll ... get list of users
// func (m *WalletDAO) FindAll() ([]models.Wallet, error) {
// 	var wallets []models.Wallet
// 	err := db.C(COLLECTION).Find(bson.M{}).All(&wallets)
// 	return wallets, err
// }
