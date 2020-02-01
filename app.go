package main

import (

	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/davdwhyte87/travelfy/dao"
	. "github.com/davdwhyte87/travelfy/controllers"
	. "github.com/davdwhyte87/travelfy/middlewares"
	"github.com/joho/godotenv"
	"os"
)


var dao = DatabaseDAO{}



// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	Server, _ := os.LookupEnv("SERVER")
	database, _ := os.LookupEnv("DATABASE")
	dao.Server = Server
	dao.Database = database
	dao.Connect()
}


func main() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()
	userRouter := v1.PathPrefix("/user").Subrouter()
	vehicleRouter := v1.PathPrefix("/vehicle").Subrouter()
	becomeAdriverRouter := r.PathPrefix("/api/v1/").Subrouter()

	userRouter.HandleFunc("/signup", CreateUser).Methods("POST")
	// signUpRouter.Use(InputValidationMiddleware)

	userRouter.HandleFunc("/signin", LoginUser).Methods("POST")

	vehicleRouter.HandleFunc("/create",
	 MultipleMiddleware(CreateVehicle, AuthenticationMiddleware)).Methods("POST")


	becomeAdriverRouter.HandleFunc("/user/become_driver",
	 MultipleMiddleware(BecomeDriver, AuthenticationMiddleware)).Methods("GET")
	// becomeAdriverRouter.Use(AuthenticationMiddleware)


	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}