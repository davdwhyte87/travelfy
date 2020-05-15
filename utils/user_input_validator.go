package utils

// import "fmt"
import (
	"fmt"
	"net/http"

	// "reflect"

	DAO "github.com/davdwhyte87/travelfy/dao"
	"github.com/thedevsaddam/govalidator"
)

var dao = DAO.UserDAO{}

func init() {
	govalidator.AddCustomRule("user_exists", func(field string, rule string, message string, value interface{}) error {
		valSlice := value.(string)
		println(valSlice)
		user, _ := dao.FindByEmail(valSlice)
		if user.Email != "" {
			return fmt.Errorf("This user email exists")
		}
		// if err != nil {
		// 	return fmt.Errorf(err.Error())

		// }
		return nil
	})
}

// CreateUserValidator ...
func CreateUserValidator(r *http.Request) (bool, interface{}) {
	rules := govalidator.MapData{
		"Name":     []string{"required", "between:3,50"},
		"Email":    []string{"required", "min:4", "max:100", "email", "user_exists"},
		"Password": []string{"required", "min:4", "max:20"},
	}
	// var user Models.User
	data := make(map[string]interface{}, 0)
	// messages := govalidator.MapData{
	// 	"Name": []string{"Name field is required", "Name should be between 3 to 50 charachers"},
	// 	"Email":    []string{"Email field is required", "", "", "A valid email is required"},
	// 	"Password": []string {"Password is required", "", ""},
	// }

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    &data,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}

	if len(e) == 0 {
		return true, err
	}
	return false, err
}

// LoginUserValidator ...
func LoginUserValidator(r *http.Request) (bool, interface{}) {
	rules := govalidator.MapData{
		"Name":  []string{"required", "between:3,50"},
		"Email": []string{"required", "min:4", "max:100", "email"},
	}
	// var user Models.User
	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    &data,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}

	if len(e) == 0 {
		return true, err
	}
	return false, err
}

// InfuseLidsValidator ...
func InfuseLidsValidator(r *http.Request) (bool, interface{}) {
	rules := govalidator.MapData{
		"Amount":  []string{"required", "float"},
	}
	// var user Models.User
	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    &data,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}

	if len(e) == 0 {
		return true, err
	}
	return false, err
}

