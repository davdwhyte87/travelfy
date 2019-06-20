package controllers
import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	. "github.com/davdwhyte87/travelfy/models"
	"gopkg.in/mgo.v2/bson"
	. "github.com/davdwhyte87/travelfy/dao"
	. "github.com/davdwhyte87/travelfy/utils"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"

)

var dao = UserDAO{}

var secreteKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	secreteKey, _ = os.LookupEnv("SECRETE_KEY")
}


// CreateUser ... This function crea
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	// fmt.Printf("%v\n", r.Body)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		//fmt.Printf("%v\n", err.Error())
		RespondWithError(w, http.StatusBadRequest, "Invalid request m payload")
		return
	}
	user.ID = bson.NewObjectId()
	// set confirmed to false 
	b := false
	user.Confirmed = &b
	user.IsDriver = false
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

	// remove the password for safety
	user.Password = "0"
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

func BecomeDriver(w http.ResponseWriter, r *http.Request) {
	// get email from request
	email := r.Context().Value("email")
	
	fmt.Printf("%+v\n", email)
}