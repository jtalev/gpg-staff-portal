package handler

import (
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "static", "home.html")
	http.ServeFile(w, r, path)
}
