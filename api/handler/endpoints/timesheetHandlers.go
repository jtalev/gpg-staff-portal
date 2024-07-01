package endpoints

import (
	"fmt"
	"net/http"

	"github.com/jtalev/gpg-staff-portal/pkg/mockdb"
)

type TimesheetHandler struct {
	Db		mockdb.Db
}

func(t *TimesheetHandler) CreateTimesheetHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
	}

	start := r.FormValue("start")
	end := r.FormValue("end")
	date := r.FormValue("date")

	// Todo: validate values

	// If no errors add to database

	

	fmt.Println(start, end, date)
}