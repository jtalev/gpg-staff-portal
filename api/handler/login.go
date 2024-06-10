package handler

import (
	"log"
	"net/http"
	"path/filepath"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "static", "login.html")
	log.Println(path)

	http.ServeFile(w, r, path)
}
