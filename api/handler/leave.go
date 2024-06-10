package handler

import (
	"net/http"
	"path/filepath"
)

func LeaveHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "static", "leave.html")
	http.ServeFile(w, r, path)
}