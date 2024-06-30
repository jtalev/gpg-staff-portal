package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homePath := filepath.Join("..", "..", "web", "pages", "home.html")
	navPath := filepath.Join("..", "..", "web", "components", "nav.html")
	headPath := filepath.Join("..", "..", "web", "components", "header.html")
	navMobPath := filepath.Join("..", "..", "web", "components", "navMob.html")

	tmpl, err := template.ParseFiles(homePath, headPath, navPath, navMobPath)
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
