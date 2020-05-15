package controllers

import (
	// "fmt"
	"net/http"
	"os"
	// "gopkg.in/mgo.v2/bson"
	DAO "github.com/davdwhyte87/travelfy/dao"
	Models "github.com/davdwhyte87/travelfy/models"
	Utils "github.com/davdwhyte87/travelfy/utils"
)

var userDao = DAO.UserDAO{}

// SetSuperAdmin ... This function gets a users wallet with the ID
func SetSuperAdmin(w http.ResponseWriter, r *http.Request) {
	var user Models.User

	// get user id from request
	requestUserID := r.Context().Value("user_id")
	requestUserIDString := requestUserID.(string)
	// fmt.Print("UserId", requestUserIDString)
	if requestUserID == "" {
		Utils.RespondWithError(w, http.StatusNotFound, "An error occured")
	}

	// get wallet from database
	var getWalletError error
	user, getWalletError = userDao.FindById(requestUserIDString)
	// fmt.Printf("WalletData : %v\n", wallet)
	if getWalletError != nil {
		Utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// check if the user is the super 
	adminEmail, _ := os.LookupEnv("ADMIN_EMAIL")
	if user.Email == adminEmail {
		// update user identity
		user.Type = 1
		updateErr := userDao.Update(user)
		if updateErr != nil{
			Utils.RespondWithError(w, http.StatusBadRequest, "UnAuthorized")
			return
		}
	} else{
		Utils.RespondWithError(w, http.StatusUnauthorized, "UnAuthorized")
		return
	}

	Utils.RespondWithOk(w, "You are now the SuperAdmin")
	return
}


