package controllers
import (
	"net/http"
	"fmt"
	"encoding/json"
	. "github.com/davdwhyte87/travelfy/models"
	"gopkg.in/mgo.v2/bson"
	. "github.com/davdwhyte87/travelfy/dao"
	. "github.com/davdwhyte87/travelfy/utils"
	"github.com/dgrijalva/jwt-go"
	. "github.com/davdwhyte87/travelfy/config"
	"time"
)
var config = Config{}
var dao = UserDAO{}

var secreteKey string

func init() {
	config.Read()
	secreteKey = config.SecreteKey
}


// CreateUser ... This function crea
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	// fmt.Printf("%v\n", r.Body)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Printf("%v\n", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid request m payload")
		return
	}
	user.ID = bson.NewObjectId()
	// hash password 
	var hashError error
	user.Password, hashError = HashPassword(user.Password)
	if hashError != nil {
		RespondWithError(w, http.StatusBadRequest, "Error encrypting password")
	}

	// save user to database
	if err := dao.Insert(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, user)
} 

// LoginUser ... This function validates a users identity and then gives the user an auth token
func LoginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	// fmt.Printf("%v\n", r.Body)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Printf("%v\n", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid request m payload")
		return
	}
	// get the user from database
	userData, userDataError := dao.FindByEmail(user.Email)
	if userDataError !=nil {
		RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	// fmt.Printf("%+v\n", userData)
	if CheckPasswordHash(user.Password, userData.Password) {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 72).Unix(),
			"email": userData.Email,
			"id": userData.ID,
			"name": userData.Name,
		}
		signedString, tokenErr := token.SignedString([]byte(secreteKey))
		if tokenErr !=nil {
			RespondWithError(w, http.StatusInternalServerError, "Error generating token")
			return
		}
		returnData := map[string] string{"token": signedString, "message":"Siccessful"}
		RespondWithJson(w, http.StatusCreated, returnData)
	} else {
		RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
	}
	return
}