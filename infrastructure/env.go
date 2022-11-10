package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

//To load env variables
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to LOAD ENV Variables")
	}
}
