package middleware

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var SecreteKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	SecreteKey, _ = os.LookupEnv("SECRETE_KEY")
}

