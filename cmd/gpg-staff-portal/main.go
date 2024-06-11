package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jtalev/gpg-staff-portal/api"
)

func getEnvVariable(key string) string {
	err := godotenv.Load("../../config/.env")
	if err != nil {
		log.Println("Failed to load .env file")
	}
	return os.Getenv(key)
}

func main() {

	var PORT = getEnvVariable("PORT")
	r := mux.NewRouter()

	api.SetupRoutes(r)

	log.Println("Starting server on port:", PORT)
	if err := http.ListenAndServe(":" + PORT, r); err != nil {
		log.Fatal("Server startup failure: ", err)
	}
}
