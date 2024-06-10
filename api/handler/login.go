package handler

import (
	"log"
	"net/http"
	"path/filepath"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "static")
	log.Println(path)

	http.ServeFile(w, r, path)
}
