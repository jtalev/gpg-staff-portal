package handler

import (
	"net/http"
	"path/filepath"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "pages", "login.html")
	http.ServeFile(w, r, path)
}