package controllers

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	DAO "github.com/davdwhyte87/travelfy/dao"
// 	Models "github.com/davdwhyte87/travelfy/models"
// 	"github.com/joho/godotenv"
// 	"gopkg.in/mgo.v2/bson"
// )

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 	}

// }

// var userDao = DAO.UserDAO{}

// // // CreateUser ... This function crea
// // func CreateUser() {
// // 	adminName, _ := os.LookupEnv("ADMIN_NAME")
// // 	adminEmail, _ := os.LookupEnv("ADMIN_EMAIL")
// // 	adminPass, _ := os.LookupEnv("ADMIN_PASS")
// // 	var user Models.User
// // 	user.ID = bson.NewObjectId()
// // 	user.Email = adminEmail
// // 	user.Password = adminPass
// // 	user.Name = adminName
// // 	// create user on db
// // 	userErr := userDao.Insert(user)
// // 	if userErr != nil {
// // 		fmt.Print("Error seeding user data")
// // 	}
// // }
