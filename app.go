package main

import (
	"fmt"
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

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}



func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}





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
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")

	signUpRouter := r.PathPrefix("/api/v1/").Subrouter()
	signInRouter := r.PathPrefix("/api/v1/").Subrouter()
	becomeAdriverRouter := r.PathPrefix("/api/v1/").Subrouter()

	signUpRouter.HandleFunc("/user/signup", CreateUser).Methods("POST")
	// signUpRouter.Use(InputValidationMiddleware)

	signInRouter.HandleFunc("/user/signin", LoginUser).Methods("POST")
	signInRouter.Use(InputValidationMiddleware)


	becomeAdriverRouter.HandleFunc("/user/become_driver",
	 MultipleMiddleware(BecomeDriver, AuthenticationMiddleware)).Methods("GET")
	// becomeAdriverRouter.Use(AuthenticationMiddleware)

	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}