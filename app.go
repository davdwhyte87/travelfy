package main

import (
	"log"
	"net/http"
	"os"

	"github.com/davdwhyte87/travelfy/controllers"
	. "github.com/davdwhyte87/travelfy/dao"
	. "github.com/davdwhyte87/travelfy/middlewares"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	walletRouter := v1.PathPrefix("/wallet").Subrouter()
	adminRouter := v1.PathPrefix("/admin").Subrouter()
	becomeAdriverRouter := r.PathPrefix("/api/v1/").Subrouter()

	userRouter.HandleFunc("/signup", controllers.CreateUser).Methods("POST")
	// signUpRouter.Use(InputValidationMiddleware)

	userRouter.HandleFunc("/signin", controllers.LoginUser).Methods("POST")

	userRouter.HandleFunc("/me_admin",
		MultipleMiddleware(controllers.SetSuperAdmin, AuthenticationMiddleware)).Methods("GET")

	vehicleRouter.HandleFunc("/create",
		MultipleMiddleware(controllers.CreateVehicle, AuthenticationMiddleware)).Methods("POST")

	becomeAdriverRouter.HandleFunc("/user/become_driver",
		MultipleMiddleware(controllers.BecomeDriver, AuthenticationMiddleware)).Methods("GET")
	// becomeAdriverRouter.Use(AuthenticationMiddleware)

	walletRouter.HandleFunc("/my_wallet",
		MultipleMiddleware(controllers.GetWallet, AuthenticationMiddleware)).Methods("GET")

	walletRouter.HandleFunc("/all",
		MultipleMiddleware(controllers.GetAllWallets, AdminAuthMiddleware)).Methods("GET")

	// admin routes
	adminRouter.HandleFunc("/add_lids", MultipleMiddleware(controllers.InfuseLids)).Methods("POST")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
