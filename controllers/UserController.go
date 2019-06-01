package controllers
import (
	"strings"
	"net/http"
	"encoding/json"
	. "github.com/davdwhyte87/travelfy/models"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	. "github.com/davdwhyte87/travelfy/dao"
)

var dao = UserDAO{}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


// CreateUser ... This function crea
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	validate := validator.New()
	err := validate.Struct(user)
	var valerrors = []string{}
	if err != nil {
		if strings.Contains(err.Error(), "Description") {
			valerrors=  append(valerrors, "Description is invalid")
		}
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	user.ID = bson.NewObjectId()
	if err := dao.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
} 