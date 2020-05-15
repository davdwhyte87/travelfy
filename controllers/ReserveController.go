package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	DAO "github.com/davdwhyte87/travelfy/dao"
	Models "github.com/davdwhyte87/travelfy/models"
	Utils "github.com/davdwhyte87/travelfy/utils"
	"gopkg.in/mgo.v2/bson"
)

var reserveDao = DAO.ReserveDAO{}


// InfuseLids ...
func InfuseLids(w http.ResponseWriter, r *http.Request) {
	type InfuseLidsRequest struct {
		Amount int
	}
	var reqData InfuseLidsRequest
	// populate the user object with data from requests
	err := Utils.DecodeReq(r, &reqData)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}

	// Validate input data
	valOk, errInput := Utils.InfuseLidsValidator(r)
	if valOk == false {

		Utils.RespondWithJSON(w, http.StatusBadRequest, errInput)
		return
	}
	// get the reserve
	reserveName, _ := os.LookupEnv("RESERVE_NAME")
	// var reserve Models.Reserve
	reserve, findErr := reserveDao.FindByName(reserveName)
	// fmt.Printf("%v\n", reserve)
	if reserve.ID == "" {
		// fmt.Print("No found bro!")
		createNew()
		return
	}
	if findErr != nil {

		Utils.RespondWithError(w, http.StatusNotFound, "Reserve error not found")
		return
	}
	if reserve.CreatedAt == "" {
		// create new reserve
		st := createNew()
		if st {
			Utils.RespondWithError(w, http.StatusNotFound, "not found but created")
			return
		}
		Utils.RespondWithError(w, http.StatusNotFound, "not found ")
		return
	}

	reserve.Balance = reserve.Balance + (reqData.Amount)
	// reserve.Balance = 0
	print(reserve.Balance)
	updateOk := reserveDao.Update(reserve)
	if updateOk != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, "could not add Lids")
		fmt.Printf("%v\n", updateOk)
		return
	}

	Utils.RespondWithOk(w, "Lids have been added to the system")
	return
}

func createNew() bool {
	reserveName, _ := os.LookupEnv("RESERVE_NAME")
	var reserve Models.Reserve
	dt := time.Now()
	reserve.CreatedAt = dt.Format("01-02-2006 15:04:05")
	reserve.UpdatedAt = reserve.CreatedAt
	reserve.Balance = 0
	reserve.ID = bson.NewObjectId()
	reserve.LiquidCash = 0.0
	reserve.Name = reserveName
	insertErr := reserveDao.Insert(reserve)
	if insertErr != nil {
		return false
	}

	return true
}

// Take ... takes cash from the bank reserve
func Take(amount int) bool{
	reserveName, _ := os.LookupEnv("RESERVE_NAME")
	reserve, findErr := reserveDao.FindByName(reserveName)
	if reserve.ID == "" {
		// fmt.Print("No found bro!")
		
		return false
	}

	if findErr != nil{
		return false
	}

	dt := time.Now()
	reserve.UpdatedAt = dt.Format("01-02-2006 15:04:05")
	reserve.Balance = reserve.Balance - amount
	reserve.LiquidCash = reserve.LiquidCash + amount
	insertErr := reserveDao.Update(reserve)

	if insertErr != nil {
		return false
	}

	return true	
}
