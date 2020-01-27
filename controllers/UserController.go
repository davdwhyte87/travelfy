package controllers

import (
	"encoding/json"
	"fmt"
    "io/ioutil"
	DAO "github.com/davdwhyte87/travelfy/dao"
	Models "github.com/davdwhyte87/travelfy/models"
	Utils  "github.com/davdwhyte87/travelfy/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"time"
	"bytes"
)

var dao = DAO.UserDAO{}

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
	// validate user
	b, _ := ioutil.ReadAll(r.Body)
	var user Models.User
	// fmt.Printf("%v\n", r.Body)
	// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	// 	fmt.Printf("%v\n", err.Error())
	// 	Utils.RespondWithError(w, http.StatusBadRequest, "Invalid request m payload")
	// 	return
	// }
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	ok, errInput := Utils.CreateUserValidator(w,r)
	if ok == false {
		print("joll")
		Utils.RespondWithJson(w, http.StatusBadRequest, errInput)
		return
	}
	user.ID = bson.NewObjectId()
	// set confirmed to false 
	// b := false
	// user.Confirmed = &b
	user.IsDriver = false
	// hash password 
	var hashError error
	user.Password, hashError = Utils.HashPassword(user.Password)
	if hashError != nil {
		Utils.RespondWithError(w, http.StatusBadRequest, "Error encrypting password")
	}

	// save user to database
	if err := dao.Insert(user); err != nil {
		Utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// remove the password for safety
	user.Password = "0"
	Utils.RespondWithJson(w, http.StatusCreated, user)
} 


// LoginUser ... This function validates a users identity and then gives the user an auth token
func LoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user Models.User
	// fmt.Printf("%v\n", r.Body)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Printf("%v\n", err.Error())
		Utils.RespondWithError(w, http.StatusBadRequest, "Invalid request m payload")
		return
	}
	// get the user from database
	userData, userDataError := dao.FindByEmail(user.Email)
	if userDataError !=nil {
		Utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	// fmt.Printf("%+v\n", userData)
	if Utils.CheckPasswordHash(user.Password, userData.Password) {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 72).Unix(),
			"email": userData.Email,
			"id": userData.ID,
			"name": userData.Name,
		}
		signedString, tokenErr := token.SignedString([]byte(secreteKey))
		if tokenErr !=nil {
			Utils.RespondWithError(w, http.StatusInternalServerError, "Error generating token")
			return
		}
		returnData := map[string] string{"token": signedString, "message":"Siccessful"}
		Utils.RespondWithJson(w, http.StatusCreated, returnData)
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
	if userDataError !=nil {
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
	var users = []interface{}{map[string]Models.User {"user":userData}, map[string]string {"name":"sjsklkldnk"}}

	//users[0] = map[string]User {"user":userData}
	var returnData = Utils.ReturnData{ Status:http.StatusOK, Data:users }
	Utils.RespondWithJson(w, http.StatusCreated, returnData)
	return
}