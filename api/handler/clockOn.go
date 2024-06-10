package handler

import (
	"net/http"
	"path/filepath"
)

func ClockOnHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "static", "clockOn.html")
	http.ServeFile(w, r, path)
}
