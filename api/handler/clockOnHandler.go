package handler

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

type ClockOnData struct {
	Timesheets []types.Timesheet
}

func ClockOnHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("..", "..", "web", "pages", "clockOn.html")
	navPath := filepath.Join("..", "..", "web", "components", "nav.html")
	navmobPath := filepath.Join("..", "..", "web", "components", "navMob.html")
	headPath := filepath.Join("..", "..", "web", "components", "header.html")
	tsPath := filepath.Join("..", "..", "web", "components", "timesheet.html")
	udPath := filepath.Join("..", "..", "web", "components", "userdetails.html")

	data := ClockOnData{
		Timesheets: make([]types.Timesheet, 7),
	}

	tmpl, err := template.ParseFiles(
		path, 
		headPath, 
		navPath, 
		navmobPath, 
		tsPath, 
		udPath, 
	)

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
