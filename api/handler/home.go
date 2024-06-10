package handler

import (
	"log"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "static", "home.html")
	log.Println(path)

	http.ServeFile(w, r, path)
}
