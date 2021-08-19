package api

import (
	"fmt"
	"log"
	"os"

	"github.com/RaphaelNagato/goresume-api/api/controllers"
	"github.com/RaphaelNagato/goresume-api/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize("postgres", os.Getenv("DATABASE_URL"))

	seed.Load(server.DB)

	server.Run(os.Getenv("PORT"))

}
