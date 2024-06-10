package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	h1 := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello Geelong Paint Group employee!")
	}

	r.HandleFunc("/", h1)

	

	log.Println("Starting server on port 8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("Server startup failure: ", err)
	}
}
