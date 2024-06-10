package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jtalev/gpg-staff-portal/api"
)

func main() {

	r := mux.NewRouter()

	api.SetupRoutes(r)

	log.Println("Starting server on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("Server startup failure: ", err)
	}
}
