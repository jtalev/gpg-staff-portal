package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func ClockOnHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "pages", "clockOn.html")
	navPath := filepath.Join("..", "..", "web", "components", "nav.html")
	headPath := filepath.Join("..", "..", "web", "components", "header.html")

	tmpl, err := template.ParseFiles(path, headPath, navPath)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
