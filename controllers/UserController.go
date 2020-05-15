package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	DAO "github.com/davdwhyte87/travelfy/dao"
	Models "github.com/davdwhyte87/travelfy/models"
	Utils "github.com/davdwhyte87/travelfy/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"
)

var dao = DAO.UserDAO{}
var walletDao = DAO.WalletDAO{}
var secreteKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	secreteKey, _ = os.LookupEnv("SECRETE_KEY")
}

// CreateUser ... This function crea
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close()

	var user Models.User
	// populate the user object with data from requests
	err := Utils.DecodeReq(r, &user)
	// fmt.Printf("%+v\n", user)
	// fmt.Printf("%+v\n", err)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}

	// Validate input data
	ok, errInput := Utils.CreateUserValidator(r)
	if ok == false {

		Utils.RespondWithJSON(w, http.StatusBadRequest, errInput)
		return
	}

	user.ID = bson.NewObjectId()
	user.Confirmed = false
	user.IsDriver = false
	user.Type = 0
	// hash password
	var hashError error
	user.Password, hashError = Utils.HashPassword(user.Password)
	if hashError != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "Error encrypting password")
		return
	}

	// save user to database
	if err := dao.Insert(user); err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// create a users wallet
	var wallet Models.Wallet
	wallet.ID = bson.NewObjectId()

	dt := time.Now()

	wallet.CreatedAt = dt.Format("01-02-2006 15:04:05")
	wallet.UpdatedAt = wallet.CreatedAt
	wallet.Balance = 0
	wallet.UserID = user.ID
	if creatWalletErr := walletDao.Insert(wallet); err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, creatWalletErr.Error())
		return
	}

	// send the user a welcome email
	var emialData Utils.EmailData
	emialData.EmailTo = user.Email
	emialData.ContentData = map[string]interface{}{"Name": user.Name}
	emialData.Template = "hello.html"
	emialData.Title = "Welcome!"
	mailSent := Utils.SendEmail(emialData)
	if mailSent {
		print("mail sent")
	} else {
		print("email not sent")
	}
	// remove the password for safety
	user.Password = "0"
	// return response
	Utils.RespondWithJSON(w, http.StatusCreated, user)
	return
}

// LoginUser ... This function validates a users identity and then gives the user an auth token
func LoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user Models.User
	// populate the user object with data from requests
	err := Utils.DecodeReq(r, &user)
	// fmt.Printf("%+v\n", user)
	// fmt.Printf("%+v\n", err)
	if err != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "This is an invalid request object. Cannot decode on server")
		return
	}

	// Validate input data
	ok, errInput := Utils.LoginUserValidator(r)
	if ok == false {
		Utils.RespondWithJSON(w, http.StatusBadRequest, errInput)
		return
	}

	// get the user from database
	userData, userDataError := dao.FindByEmail(user.Email)
	if userDataError != nil {
		Utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	// fmt.Printf("%+v\n", userData)
	if Utils.CheckPasswordHash(user.Password, userData.Password) {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.MapClaims{
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
			"email": userData.Email,
			"id":    userData.ID,
			"type":  userData.Type,
			"name":  userData.Name,
		}
		signedString, tokenErr := token.SignedString([]byte(secreteKey))
		if tokenErr != nil {
			Utils.RespondWithError(w, http.StatusInternalServerError, "Error generating token")
			return
		}
		returnData := map[string]string{"token": signedString, "message": "Successful"}
		Utils.RespondWithJSON(w, http.StatusOK, returnData)
	} else {
		Utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
	}
	return
}

// BecomeDriver ...
func BecomeDriver(w http.ResponseWriter, r *http.Request) {
	// get email from request
	requestEmail := r.Context().Value("email")
	if requestEmail == nil {
		Utils.RespondWithError(w, http.StatusNotFound, "An error occured")
	}

	//email := fmt.Sprintf("%+v\n", requestEmail)
	//print(reflect.TypeOf(email).String())
	//get user with the email

	userData, userDataError := dao.FindByEmail(requestEmail.(string))
	if userDataError != nil {
		fmt.Printf("%+v\n", userDataError)
		Utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	//fmt.Printf("%+v\n", userData)
	userData.IsDriver = true
	updateErr := dao.Update(userData)
	if updateErr != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, "Error updating data")
		return
	}

	//returnDatau := map[] {"user":userData, "message":"Successful"}

	userData.Password = ""
	var users = []interface{}{map[string]Models.User{"user": userData}, map[string]string{"name": "sjsklkldnk"}}

	//users[0] = map[string]User {"user":userData}
	var returnData = Utils.ReturnData{Status: http.StatusOK, Data: users}
	Utils.RespondWithJSON(w, http.StatusCreated, returnData)
	return
}



