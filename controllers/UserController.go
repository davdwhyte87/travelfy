package controllers
import (
	"net/http"
	"fmt"
	"encoding/json"
	. "github.com/davdwhyte87/travelfy/models"
	"gopkg.in/mgo.v2/bson"
	. "github.com/davdwhyte87/travelfy/dao"
	. "github.com/davdwhyte87/travelfy/utils"
)

var dao = UserDAO{}




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
	if err := dao.Insert(user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, user)
} 