package utils

import (

	// "fmt"
	"net/http"
	Models "github.com/davdwhyte87/travelfy/models"
	"github.com/thedevsaddam/govalidator"
)

// CreateUserValidator ...
func CreateUserValidator(w http.ResponseWriter, r *http.Request) (bool, interface{}){
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,8"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"phone":    []string{"digits:11"},
		"agree":    []string{"bool"},
		"dob":      []string{"date"},
	}
	var user Models.User
	messages := govalidator.MapData{
		"username": []string{"required:আপনাকে অবশ্যই ইউজারনেম দিতে হবে", "between:ইউজারনেম অবশ্যই ৩-৮ অক্ষর হতে হবে"},
		"phone":    []string{"digits:ফোন নাম্বার অবশ্যই ১১ নম্বারের হতে হবে"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Data:            &user,
		Messages:        messages, // custom message map (Optional)
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}

	println(err)
	if e == nil {
		return true, err
	}
	return false, err
	// RespondWithJson(w, http.StatusBadRequest, err)
}