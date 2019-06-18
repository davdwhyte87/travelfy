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

	server, _ := os.LookupEnv("SERVER")
	database, _ := os.LookupEnv("DATABASE")
	dao.Server = server
	dao.Database = database
	dao.Connect()
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")

	r.HandleFunc("/user/signup", CreateUser).Methods("POST")
	r.Use(InputValidationMiddleware)

	r.HandleFunc("/user/signin", LoginUser).Methods("POST")
	r.Use(InputValidationMiddleware)
	 
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}